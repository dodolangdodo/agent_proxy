package backend

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Status represents the health status of a backend.
type Status string

const (
	StatusActive    Status = "active"
	StatusDegraded  Status = "degraded"
	StatusDisabled  Status = "disabled"
)

// Backend represents a downstream LLM provider endpoint.
type Backend struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	Provider      string            `json:"provider"` // "openai", "anthropic", "custom"
	BaseURL       string            `json:"base_url"`
	APIKey        string            `json:"-"` // encrypted at rest
	APIKeyEncrypted string          `json:"api_key_encrypted,omitempty"`
	Models        []string          `json:"models"`
	ModelMapping  map[string]string `json:"model_mapping,omitempty"` // proxy model → real backend model
	Weight           int               `json:"weight"`
	CostTier         string            `json:"cost_tier"`          // "prepaid" or "pay_per_token" (default if empty)
	MaxContextTokens int               `json:"max_context_tokens"` // 0 = unlimited
	SkipContextFilter bool             `json:"skip_context_filter"` // if true, bypass context-length filter
	MaxRPM           int               `json:"max_rpm"`
	MaxConcurrent int               `json:"max_concurrent"`
	Timeout       int               `json:"timeout_seconds"`
	Headers       map[string]string `json:"headers,omitempty"`
	Status        Status            `json:"status"`
	TokenBalance  int64             `json:"token_balance"` // total token budget, 0 = unlimited
	TokenUsed     int64             `json:"token_used"`    // running counter of tokens consumed
	// Auto-discovered quota from rate-limit response headers
	QuotaTokensTotal     int64     `json:"quota_tokens_total"`
	QuotaTokensRemaining int64     `json:"quota_tokens_remaining"`
	QuotaRequestsTotal     int64   `json:"quota_requests_total"`
	QuotaRequestsRemaining int64   `json:"quota_requests_remaining"`
	ConsecutiveFails int            `json:"-"`
	LastHealthCheck  time.Time      `json:"last_health_check"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
}

// RemainingTokens returns how many tokens are left for this backend.
// Prefers auto-discovered quota from rate-limit headers when available.
// Returns -1 if neither quota nor balance is set (unlimited).
func (b *Backend) RemainingTokens() int64 {
	if b.QuotaTokensTotal > 0 {
		return b.QuotaTokensRemaining
	}
	if b.TokenBalance <= 0 {
		return -1 // unlimited
	}
	remaining := b.TokenBalance - b.TokenUsed
	if remaining < 0 {
		return 0
	}
	return remaining
}

// UsageRatio returns 0.0-1.0 representing how much of the budget is consumed.
func (b *Backend) UsageRatio() float64 {
	if b.QuotaTokensTotal > 0 {
		if b.QuotaTokensTotal == 0 {
			return 0
		}
		return float64(b.QuotaTokensTotal-b.QuotaTokensRemaining) / float64(b.QuotaTokensTotal)
	}
	if b.TokenBalance <= 0 {
		return 0
	}
	return float64(b.TokenUsed) / float64(b.TokenBalance)
}

// CaptureRateLimitHeaders parses rate limit headers from a backend HTTP response.
// Supports OpenAI-style, Anthropic-style, and DeepSeek-style headers.
func (b *Backend) CaptureRateLimitHeaders(resp *http.Response) {
	// DeepSeek-style (X-RateLimit-Limit / X-RateLimit-Remaining) — request-level
	if v := resp.Header.Get("x-ratelimit-limit"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			b.QuotaRequestsTotal = n
		}
	}
	if v := resp.Header.Get("x-ratelimit-remaining"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			b.QuotaRequestsRemaining = n
		}
	}

	// OpenAI-style headers (token-level + request-level)
	if v := resp.Header.Get("x-ratelimit-limit-tokens"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			b.QuotaTokensTotal = n
		}
	}
	if v := resp.Header.Get("x-ratelimit-remaining-tokens"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			b.QuotaTokensRemaining = n
		}
	}
	if v := resp.Header.Get("x-ratelimit-limit-requests"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			b.QuotaRequestsTotal = n
		}
	}
	if v := resp.Header.Get("x-ratelimit-remaining-requests"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			b.QuotaRequestsRemaining = n
		}
	}

	// Anthropic-style headers (take precedence if present)
	if v := resp.Header.Get("anthropic-ratelimit-tokens-limit"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			b.QuotaTokensTotal = n
		}
	}
	if v := resp.Header.Get("anthropic-ratelimit-tokens-remaining"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			b.QuotaTokensRemaining = n
		}
	}
	if v := resp.Header.Get("anthropic-ratelimit-requests-limit"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			b.QuotaRequestsTotal = n
		}
	}
	if v := resp.Header.Get("anthropic-ratelimit-requests-remaining"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			b.QuotaRequestsRemaining = n
		}
	}
}

// ResolveModel returns the real backend model name for a proxy model.
// If a mapping exists, returns the mapped name; otherwise returns the original model.
func (b *Backend) ResolveModel(proxyModel string) string {
	if b.ModelMapping != nil {
		if realModel, ok := b.ModelMapping[proxyModel]; ok && realModel != "" {
			return realModel
		}
	}
	return proxyModel
}

// HealthCheckConfig defines health probing parameters.
type HealthCheckConfig struct {
	Interval       time.Duration
	Timeout        time.Duration
	MaxFails       int
	SuccessThreshold int
}

// Runtime contains mutable per-backend state (latency, rate limiting).
type Runtime struct {
	mu             sync.RWMutex
	EWMALatency    time.Duration // exponentially weighted moving average
	ErrorRate      float64       // 0.0 - 1.0
	ActiveRequests int32
	LastResult     time.Time
	TokensUsed     int64 // local counter, synced to Backend.TokenUsed periodically
}

func (r *Runtime) RecordResult(latency time.Duration, isError bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	const alpha = 0.3 // EWMA smoothing factor
	r.EWMALatency = time.Duration(float64(r.EWMALatency)*(1-alpha) + float64(latency)*alpha)
	r.LastResult = time.Now()

	if isError {
		r.ErrorRate = r.ErrorRate*(1-alpha) + alpha*1.0
	} else {
		r.ErrorRate = r.ErrorRate * (1 - alpha)
	}
}

func (r *Runtime) RecordTokens(count int64) int64 {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.TokensUsed += count
	return r.TokensUsed
}

func (r *Runtime) GetLatency() time.Duration {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.EWMALatency
}

func (r *Runtime) GetErrorRate() float64 {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.ErrorRate
}

func (r *Runtime) Score(weight int, minLatency time.Duration) float64 {
	lat := r.GetLatency()
	err := r.GetErrorRate()
	if lat == 0 {
		lat = minLatency
	}
	latScore := 1.0
	if minLatency > 0 && lat > 0 {
		latScore = float64(minLatency) / float64(lat)
	}
	return float64(weight) * latScore * (1.0 - err)
}

// encryptAPIKey encrypts an API key using AES-GCM with a static key.
// In production, the encryption key should come from a secure key management service.
func EncryptAPIKey(plaintext, encKey string) (string, error) {
	if plaintext == "" {
		return "", nil
	}
	block, err := aes.NewCipher([]byte(encKey)[:32])
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// decryptAPIKey decrypts an API key.
func decryptAPIKey(encoded, encKey string) (string, error) {
	if encoded == "" {
		return "", nil
	}
	ciphertext, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(encKey)[:32])
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", io.ErrUnexpectedEOF
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// ToJSON serializes a backend to JSON (exposing encrypted key).
func (b *Backend) ToJSON(encKey string) ([]byte, error) {
	enc, err := EncryptAPIKey(b.APIKey, encKey)
	if err != nil {
		return nil, err
	}
	clone := *b
	clone.APIKeyEncrypted = enc
	clone.APIKey = ""
	return json.Marshal(clone)
}
