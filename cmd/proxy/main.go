package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"agent-proxy-v2/internal/api"
	"agent-proxy-v2/internal/backend"
	"agent-proxy-v2/internal/config"
	"agent-proxy-v2/internal/proxy"
	"agent-proxy-v2/internal/store"
	"agent-proxy-v2/internal/strategy"
)

func main() {
	cfg := config.Get()

	// Initialize store (JSON file)
	s, err := store.New("proxy_state.json")
	if err != nil {
		log.Fatalf("Failed to open store: %v", err)
	}
	defer s.Close()

	// Initialize backend registry
	encKey := (cfg.AdminAPIKey + "32byteencryptionkey!!____")[:32]
	reg := backend.NewRegistry(encKey)

	// Load existing backends from store
	backends, _ := s.ListBackends()
	reg.UpdateFromSlice(backends)
	reg.DecryptKeys()
	log.Printf("Loaded %d backends", len(backends))

	// Start health checks
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reg.StartHealthChecks(ctx, time.Duration(cfg.HealthCheckFreq)*time.Second, 3)
	reg.StartSnapshotRefresh(ctx, 5*time.Second)

	// Load strategy configs into an in-memory cache to avoid store lock on every request
	configs, _ := s.ListStrategyConfigs()
	log.Printf("Loaded %d strategy configs", len(configs))

	var strategyCacheMu sync.RWMutex
	strategyCache := make(map[string]strategy.LoadBalanceStrategy, len(configs))
	for model, name := range configs {
		strategyCache[model] = strategy.Get(name)
	}
	// Background refresh every 5 seconds
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				newConfigs, _ := s.ListStrategyConfigs()
				newCache := make(map[string]strategy.LoadBalanceStrategy, len(newConfigs))
				for model, name := range newConfigs {
					newCache[model] = strategy.Get(name)
				}
				strategyCacheMu.Lock()
				strategyCache = newCache
				strategyCacheMu.Unlock()
			}
		}
	}()

	// Shared HTTP client with connection pooling for backend requests
	sharedClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
			ForceAttemptHTTP2:   true,
		},
	}

	// Create proxy handler with hot-reloadable strategy lookup
	handler := &proxy.Handler{
		Registry:   reg,
		HTTPClient: sharedClient,
		GetStrategy: func(model string) strategy.LoadBalanceStrategy {
			strategyCacheMu.RLock()
			defer strategyCacheMu.RUnlock()
			return strategyCache[model]
		},
		DefaultStrategy: strategy.Get("priority"),
	}

	// Set up admin API
	apiServer := api.New(s, reg, cfg.AdminAPIKey, encKey)
	apiMux := apiServer.Mux()

	// Main router: combine proxy + admin API + static UI
	mainMux := http.NewServeMux()

	// Logging middleware
	logged := func(name string, h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			h(w, r)
			log.Printf("[http] %s %s %s %v", r.Method, r.URL.Path, name, time.Since(start))
		}
	}

	// Proxy endpoints - both OpenAI and Anthropic formats
	mainMux.HandleFunc("POST /v1/chat/completions", logged("openai-proxy", handler.HandleRequest))
	mainMux.HandleFunc("POST /v1/messages", logged("anthropic-proxy", handler.HandleRequest))
	mainMux.HandleFunc("GET /v1/models", logged("models", handler.HandleModels))

	// Admin API endpoints
	mainMux.HandleFunc("GET /api/", logged("api", apiMux.ServeHTTP))
	mainMux.HandleFunc("POST /api/", logged("api", apiMux.ServeHTTP))
	mainMux.HandleFunc("PUT /api/", logged("api", apiMux.ServeHTTP))
	mainMux.HandleFunc("DELETE /api/", logged("api", apiMux.ServeHTTP))
	mainMux.HandleFunc("OPTIONS /api/", logged("api", apiMux.ServeHTTP))

	// Serve static web UI if available
	if _, err := os.Stat("web/dist"); err == nil {
		// SPA fallback: serve index.html for any non-file path under /ui/
		spaHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := strings.TrimPrefix(r.URL.Path, "/ui")
			if path == "" || path == "/" {
				path = "/index.html"
			}
			fullPath := "web/dist" + path
			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				http.ServeFile(w, r, "web/dist/index.html")
				return
			}
			http.ServeFile(w, r, fullPath)
		})
		mainMux.Handle("GET /ui/", spaHandler)
		// Vite builds use absolute paths — serve assets/locales/logo at root
		mainMux.Handle("GET /assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("web/dist/assets"))))
		mainMux.Handle("GET /locales/", http.StripPrefix("/locales", http.FileServer(http.Dir("web/dist/locales"))))
		mainMux.HandleFunc("GET /logo.svg", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "web/dist/logo.svg")
		})
		log.Printf("Web UI: http://localhost%s/ui", cfg.ListenAddr)
	}

	// Start server
	srv := &http.Server{
		Addr:        cfg.ListenAddr,
		Handler:     mainMux,
		ReadTimeout: 30 * time.Second,
	}

	go func() {
		log.Printf("Proxy listening on %s", cfg.ListenAddr)
		log.Printf("Admin API: http://localhost%s/api/health", cfg.ListenAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down...")
	cancel()
	srv.Shutdown(context.Background())
}
