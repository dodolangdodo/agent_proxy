package api

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"agent-proxy-v2/internal/backend"
	"agent-proxy-v2/internal/store"
	"agent-proxy-v2/internal/strategy"
)

// Server wraps the admin REST API using net/http ServeMux (Go 1.22+).
type Server struct {
	store    *store.Store
	registry *backend.Registry
	adminKey string
	encKey   string
	mux      *http.ServeMux
}

// BackendInput is the JSON payload for creating/updating a backend.
type BackendInput struct {
	Name          string            `json:"name"`
	Provider      string            `json:"provider"`
	BaseURL       string            `json:"base_url"`
	APIKey        string            `json:"api_key"`
	Models        []string          `json:"models"`
	Weight        int               `json:"weight"`
	MaxRPM        int               `json:"max_rpm"`
	MaxConcurrent int               `json:"max_concurrent"`
	Timeout       int               `json:"timeout_seconds"`
	Headers       map[string]string `json:"headers"`
	ModelMapping  map[string]string `json:"model_mapping"` // proxy model → real backend model
	TokenBalance     int64             `json:"token_balance"`      // 0 = unlimited
	CostTier         string            `json:"cost_tier"`           // "prepaid" or "pay_per_token"
	MaxContextTokens int               `json:"max_context_tokens"`  // 0 = unlimited
}

// StrategyInput sets the strategy for a model.
type StrategyInput struct {
	Model    string `json:"model"`
	Strategy string `json:"strategy"`
}

// New creates the admin API server.
func New(s *store.Store, reg *backend.Registry, adminKey, encKey string) *Server {
	api := &Server{
		store:    s,
		registry: reg,
		adminKey: adminKey,
		encKey:   encKey,
		mux:      http.NewServeMux(),
	}

	// Health check (no auth)
	api.mux.HandleFunc("GET /api/health", api.handleHealth)

	// Backend CRUD (auth required)
	api.mux.HandleFunc("GET /api/backends", api.withAuth(api.listBackends))
	api.mux.HandleFunc("POST /api/backends", api.withAuth(api.createBackend))
	api.mux.HandleFunc("PUT /api/backends/{id}", api.withAuth(api.updateBackend))
	api.mux.HandleFunc("DELETE /api/backends/{id}", api.withAuth(api.deleteBackend))
	api.mux.HandleFunc("GET /api/backends/{id}/key", api.withAuth(api.getBackendKey))

	// Strategies
	api.mux.HandleFunc("GET /api/strategies", api.withAuth(api.listStrategies))
	api.mux.HandleFunc("POST /api/strategies", api.withAuth(api.setStrategy))
	api.mux.HandleFunc("GET /api/strategies/config", api.withAuth(api.listStrategyConfigs))

	// Models
	api.mux.HandleFunc("GET /api/models", api.listModels) // no auth needed for proxy clients

	// CORS preflight
	api.mux.HandleFunc("OPTIONS /api/", api.handleCORS)

	return api
}

func (s *Server) Mux() *http.ServeMux { return s.mux }

// --- Auth middleware ---

func (s *Server) withAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("Authorization")
		key = strings.TrimPrefix(key, "Bearer ")
		if key != s.adminKey {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
			return
		}
		next(w, r)
	}
}

// --- Handlers ---

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{"status": "ok", "time": time.Now().UTC()})
}

func (s *Server) listBackends(w http.ResponseWriter, r *http.Request) {
	backends, err := s.store.ListBackends()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	// APIKey has json:"-" so it's not serialized; no need to mutate here.
	writeJSON(w, http.StatusOK, backends)
}

func (s *Server) createBackend(w http.ResponseWriter, r *http.Request) {
	var input BackendInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if input.Models == nil {
		input.Models = []string{}
	}
	if input.Headers == nil {
		input.Headers = map[string]string{}
	}
	if input.Weight <= 0 {
		input.Weight = 10
	}
	if input.Timeout <= 0 {
		input.Timeout = 120
	}

	b := &backend.Backend{
		ID:          newUUID(),
		Name:        input.Name,
		Provider:    input.Provider,
		BaseURL:     input.BaseURL,
		APIKey:      input.APIKey,
		Models:      input.Models,
		Weight:      input.Weight,
		MaxRPM:      input.MaxRPM,
		MaxConcurrent: input.MaxConcurrent,
		Timeout:     input.Timeout,
		Headers:      input.Headers,
		ModelMapping: input.ModelMapping,
		TokenBalance:     input.TokenBalance,
		CostTier:         input.CostTier,
		MaxContextTokens: input.MaxContextTokens,
		Status:           backend.StatusActive,
	}

	// Encrypt API key for storage
	enc, _ := backend.EncryptAPIKey(input.APIKey, s.encKey)
	b.APIKeyEncrypted = enc

	if err := s.store.SaveBackend(b); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	s.registry.Add(b)
	writeJSON(w, http.StatusCreated, b)
}

func (s *Server) updateBackend(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	// Read from registry to get live TokenUsed, not stale store value
	regBackend := s.registry.GetLive(id)
	if regBackend == nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "backend not found"})
		return
	}

	var input BackendInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	regBackend.Name = input.Name
	regBackend.Provider = input.Provider
	regBackend.BaseURL = input.BaseURL
	if input.APIKey != "" {
		enc, _ := backend.EncryptAPIKey(input.APIKey, s.encKey)
		regBackend.APIKeyEncrypted = enc
		regBackend.APIKey = input.APIKey
	}
	regBackend.Models = input.Models
	regBackend.Weight = input.Weight
	regBackend.MaxRPM = input.MaxRPM
	regBackend.MaxConcurrent = input.MaxConcurrent
	regBackend.Timeout = input.Timeout
	regBackend.Headers = input.Headers
	regBackend.ModelMapping = input.ModelMapping
	regBackend.TokenBalance = input.TokenBalance
	regBackend.CostTier = input.CostTier
	regBackend.MaxContextTokens = input.MaxContextTokens
	if regBackend.Headers == nil {
		regBackend.Headers = map[string]string{}
	}

	if err := s.store.SaveBackend(regBackend); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	s.registry.Add(regBackend)
	writeJSON(w, http.StatusOK, regBackend)
}

func (s *Server) deleteBackend(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.store.DeleteBackend(id); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	s.registry.Remove(id)
	writeJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func (s *Server) getBackendKey(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	b := s.registry.Get(id)
	if b == nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "backend not found"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{
		"id":     b.ID,
		"name":   b.Name,
		"api_key": b.APIKey,
	})
}

func (s *Server) listStrategies(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"strategies": strategy.Names(),
	})
}

func (s *Server) setStrategy(w http.ResponseWriter, r *http.Request) {
	var input StrategyInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	valid := false
	for _, n := range strategy.Names() {
		if n == input.Strategy {
			valid = true
			break
		}
	}
	if !valid {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "unknown strategy: " + input.Strategy})
		return
	}
	if err := s.store.SaveStrategyConfig(input.Model, input.Strategy); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "saved"})
}

func (s *Server) listStrategyConfigs(w http.ResponseWriter, r *http.Request) {
	configs, err := s.store.ListStrategyConfigs()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, configs)
}

func (s *Server) listModels(w http.ResponseWriter, r *http.Request) {
	backends := s.registry.List()
	modelSet := make(map[string]bool)
	for _, b := range backends {
		if b.Status == backend.StatusActive {
			for _, m := range b.Models {
				modelSet[m] = true
			}
		}
	}
	type modelEntry struct {
		ID     string `json:"id"`
		Object string `json:"object"`
	}
	models := make([]modelEntry, 0, len(modelSet))
	for m := range modelSet {
		models = append(models, modelEntry{ID: m, Object: "model"})
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"object": "list",
		"data":   models,
	})
}

// --- Helpers ---

func (s *Server) handleCORS(w http.ResponseWriter, r *http.Request) {
	setCORS(w)
	w.WriteHeader(http.StatusNoContent)
}

func setCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Max-Age", "86400")
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	setCORS(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func newUUID() string {
	b := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		// fallback: use timestamp-based
		return fmt.Sprintf("%016x-%04x-%04x-%04x-%012x",
			time.Now().UnixNano(), 0, 0, 0, 0)
	}
	b[6] = (b[6] & 0x0f) | 0x40 // version 4
	b[8] = (b[8] & 0x3f) | 0x80 // variant 10
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}
