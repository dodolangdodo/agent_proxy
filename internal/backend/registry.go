package backend

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

// Registry manages all backends with atomic snapshot swaps.
type Registry struct {
	mu         sync.RWMutex
	backends   map[string]*Backend
	runtimes   map[string]*Runtime
	byModel    map[string][]string // model name -> backend IDs
	encKey     string
	httpClient *http.Client

	// Snapshot for lock-free reads
	snapshot atomic.Value // *registrySnapshot
}

type registrySnapshot struct {
	backends map[string]*Backend
	runtimes map[string]*Runtime
	byModel  map[string][]string
}

// NewRegistry creates a backend registry.
func NewRegistry(encKey string) *Registry {
	r := &Registry{
		backends: make(map[string]*Backend),
		runtimes: make(map[string]*Runtime),
		byModel:  make(map[string][]string),
		encKey:   encKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
	r.updateSnapshot()
	return r
}

// Add adds or updates a backend and atomically swaps the snapshot.
func (r *Registry) Add(b *Backend) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	existing, exists := r.backends[b.ID]
	r.backends[b.ID] = b

	if !exists {
		r.runtimes[b.ID] = &Runtime{}
	}

	// Update model index: remove old associations if updating
	if existing != nil {
		r.removeFromModelIndex(b.ID, existing.Models)
	}
	for _, m := range b.Models {
		r.byModel[m] = appendIfMissing(r.byModel[m], b.ID)
	}

	r.updateSnapshot()
	return nil
}

// Remove removes a backend by ID.
func (r *Registry) Remove(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if b, ok := r.backends[id]; ok {
		r.removeFromModelIndex(id, b.Models)
		delete(r.backends, id)
		delete(r.runtimes, id)
		r.updateSnapshot()
	}
}

// Get returns a backend by ID (from the lock-free snapshot — read only).
func (r *Registry) Get(id string) *Backend {
	snap := r.snapshot.Load().(*registrySnapshot)
	return snap.backends[id]
}

// GetLive returns the live, mutable backend by ID. Must be called while
// holding r.mu (or for single-threaded mutation).
func (r *Registry) GetLive(id string) *Backend {
	return r.backends[id]
}

// List returns all backends.
func (r *Registry) List() []*Backend {
	snap := r.snapshot.Load().(*registrySnapshot)
	result := make([]*Backend, 0, len(snap.backends))
	for _, b := range snap.backends {
		result = append(result, b)
	}
	return result
}

// GetForModel returns all backends that support a given model.
func (r *Registry) GetForModel(model string) []*Backend {
	snap := r.snapshot.Load().(*registrySnapshot)
	ids := snap.byModel[model]
	result := make([]*Backend, 0, len(ids))
	for _, id := range ids {
		if b, ok := snap.backends[id]; ok && b.Status == StatusActive {
			result = append(result, b)
		}
	}
	return result
}

// RecordResult records the outcome of a request for adaptive strategies.
func (r *Registry) RecordResult(backendID string, latency time.Duration, isError bool) {
	r.mu.RLock()
	rt, ok := r.runtimes[backendID]
	r.mu.RUnlock()
	if ok {
		rt.RecordResult(latency, isError)
	}
}

// RecordTokens records token consumption for a backend.
// Snapshot is refreshed asynchronously by StartSnapshotRefresh to avoid
// deep-copy overhead on the hot path.
func (r *Registry) RecordTokens(backendID string, count int64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if b, ok := r.backends[backendID]; ok {
		b.TokenUsed += count
	}
	if rt, ok := r.runtimes[backendID]; ok {
		rt.TokensUsed += count
	}
}

// CaptureRateLimits extracts rate-limit headers from a backend response and
// updates the live backend's auto-discovered quota fields.
// Snapshot is refreshed asynchronously by StartSnapshotRefresh.
func (r *Registry) CaptureRateLimits(backendID string, resp *http.Response) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if b, ok := r.backends[backendID]; ok {
		b.CaptureRateLimitHeaders(resp)
	}
}

// StartSnapshotRefresh runs updateSnapshot periodically in a background goroutine.
// This avoids the deep-copy overhead on every request in RecordTokens/CaptureRateLimits.
func (r *Registry) StartSnapshotRefresh(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				r.mu.Lock()
				r.updateSnapshot()
				r.mu.Unlock()
			}
		}
	}()
}

// GetRuntime returns the runtime stats for a backend.
func (r *Registry) GetRuntime(id string) *Runtime {
	snap := r.snapshot.Load().(*registrySnapshot)
	return snap.runtimes[id]
}

// GetAllRuntimes returns all runtime stats.
func (r *Registry) GetAllRuntimes() map[string]*Runtime {
	snap := r.snapshot.Load().(*registrySnapshot)
	return snap.runtimes
}

// HealthCheck performs a health check on a specific backend.
func (r *Registry) HealthCheck(ctx context.Context, b *Backend, maxFails int) {
	req, err := http.NewRequestWithContext(ctx, "GET", b.BaseURL+"/v1/models", nil)
	if err != nil {
		r.markFailure(b, maxFails)
		return
	}
	if b.APIKey != "" {
		switch b.Provider {
		case "anthropic":
			req.Header.Set("x-api-key", b.APIKey)
		default:
			req.Header.Set("Authorization", "Bearer "+b.APIKey)
		}
	}

	resp, err := r.httpClient.Do(req)
	if err != nil {
		r.markFailure(b, maxFails)
		return
	}
	resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		r.markSuccess(b)
	} else {
		r.markFailure(b, maxFails)
	}
}

func (r *Registry) markSuccess(b *Backend) {
	r.mu.Lock()
	defer r.mu.Unlock()
	b.LastHealthCheck = time.Now()
	b.ConsecutiveFails = 0
	if b.Status == StatusDegraded {
		b.Status = StatusActive
	}
	r.updateSnapshot()
}

func (r *Registry) markFailure(b *Backend, maxFails int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	b.LastHealthCheck = time.Now()
	b.ConsecutiveFails++
	if b.ConsecutiveFails >= maxFails {
		if b.Status == StatusActive || b.Status == StatusDegraded {
			b.Status = StatusDegraded
		}
		if b.ConsecutiveFails >= maxFails*3 {
			b.Status = StatusDisabled
		}
	}
	r.updateSnapshot()
}

// StartHealthChecks runs periodic health checks on all backends.
func (r *Registry) StartHealthChecks(ctx context.Context, interval time.Duration, maxFails int) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				for _, b := range r.List() {
					checkCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
					r.HealthCheck(checkCtx, b, maxFails)
					cancel()
				}
			}
		}
	}()
}

// DecryptKeys decrypts all backend API keys into memory.
func (r *Registry) DecryptKeys() {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, b := range r.backends {
		if b.APIKey == "" && b.APIKeyEncrypted != "" {
			key, err := decryptAPIKey(b.APIKeyEncrypted, r.encKey)
			if err == nil {
				b.APIKey = key
			}
		}
	}
	r.updateSnapshot()
}

func (r *Registry) updateSnapshot() {
	// Deep copy
	backendsCopy := make(map[string]*Backend, len(r.backends))
	for k, v := range r.backends {
		clone := *v
		models := make([]string, len(v.Models))
		copy(models, v.Models)
		clone.Models = models
		headers := make(map[string]string, len(v.Headers))
		for hk, hv := range v.Headers {
			headers[hk] = hv
		}
		clone.Headers = headers
		backendsCopy[k] = &clone
	}

	runtimesCopy := make(map[string]*Runtime, len(r.runtimes))
	for k, v := range r.runtimes {
		runtimesCopy[k] = v
	}

	byModelCopy := make(map[string][]string, len(r.byModel))
	for k, v := range r.byModel {
		ids := make([]string, len(v))
		copy(ids, v)
		byModelCopy[k] = ids
	}

	r.snapshot.Store(&registrySnapshot{
		backends: backendsCopy,
		runtimes: runtimesCopy,
		byModel:  byModelCopy,
	})
}

func (r *Registry) removeFromModelIndex(backendID string, models []string) {
	for _, m := range models {
		ids := r.byModel[m]
		for i, id := range ids {
			if id == backendID {
				r.byModel[m] = append(ids[:i], ids[i+1:]...)
				break
			}
		}
	}
}

func appendIfMissing(slice []string, item string) []string {
	for _, s := range slice {
		if s == item {
			return slice
		}
	}
	return append(slice, item)
}

// Snapshot returns the current registry state for lock-free reads.
func (r *Registry) Snapshot() *registrySnapshot {
	return r.snapshot.Load().(*registrySnapshot)
}

// UpdateFromSlice replaces all backends from a slice (e.g., loaded from DB).
func (r *Registry) UpdateFromSlice(backends []*Backend) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.backends = make(map[string]*Backend, len(backends))
	r.byModel = make(map[string][]string)
	for _, b := range backends {
		r.backends[b.ID] = b
		if _, ok := r.runtimes[b.ID]; !ok {
			r.runtimes[b.ID] = &Runtime{}
		}
		for _, m := range b.Models {
			r.byModel[m] = appendIfMissing(r.byModel[m], b.ID)
		}
	}
	r.updateSnapshot()
}

// Ensure Backend implements Stringer.
func (b *Backend) String() string {
	return fmt.Sprintf("Backend{ID:%s Name:%s Provider:%s Status:%s Models:%v}",
		b.ID, b.Name, b.Provider, b.Status, b.Models)
}
