package config

import (
	"os"
	"sync/atomic"
)

// Config holds all runtime configuration as an atomic snapshot.
// Handlers read from it without locks; writers atomically swap.
type Config struct {
	ListenAddr      string `json:"listen_addr"`
	DBDSN           string `json:"db_dsn"`
	LogLevel        string `json:"log_level"`
	DefaultTimeout  int    `json:"default_timeout_seconds"`
	MaxRetries      int    `json:"max_retries"`
	AdminAPIKey     string `json:"admin_api_key"`
	HealthCheckFreq int    `json:"health_check_frequency_seconds"`
}

var global atomic.Value

func init() {
	global.Store(defaultConfig())
}

func defaultConfig() *Config {
	return &Config{
		ListenAddr:      envOrDefault("PROXY_LISTEN_ADDR", ":9999"),
		DBDSN:           envOrDefault("PROXY_DB_DSN", "file:proxy.db?cache=shared&_journal_mode=WAL"),
		LogLevel:        envOrDefault("PROXY_LOG_LEVEL", "info"),
		DefaultTimeout:  120,
		MaxRetries:      3,
		AdminAPIKey:     envOrDefault("PROXY_ADMIN_KEY", "admin-change-me"),
		HealthCheckFreq: 30,
	}
}

func Get() *Config {
	return global.Load().(*Config)
}

func Update(c *Config) {
	global.Store(c)
}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
