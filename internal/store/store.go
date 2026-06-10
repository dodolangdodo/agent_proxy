package store

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"agent-proxy-v2/internal/backend"
)

// Store persists backends and configuration to a JSON file.
type Store struct {
	mu       sync.RWMutex
	filePath string
	data     storeData
}

type storeData struct {
	Backends  []*backend.Backend       `json:"backends"`
	Strategies map[string]string       `json:"strategies"` // model -> strategy name
}

func New(filePath string) (*Store, error) {
	s := &Store{
		filePath: filePath,
		data: storeData{
			Backends:   make([]*backend.Backend, 0),
			Strategies: make(map[string]string),
		},
	}
	if err := s.load(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	return s, nil
}

func (s *Store) load() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &s.data)
}

func (s *Store) save() error {
	data, err := json.MarshalIndent(s.data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filePath, data, 0600)
}

// SaveBackend inserts or updates a backend.
func (s *Store) SaveBackend(b *backend.Backend) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.data.Backends == nil {
		s.data.Backends = make([]*backend.Backend, 0)
	}

	for i, existing := range s.data.Backends {
		if existing.ID == b.ID {
			b.CreatedAt = existing.CreatedAt
			b.UpdatedAt = time.Now().UTC()
			s.data.Backends[i] = b
			return s.save()
		}
	}
	// New backend
	b.CreatedAt = time.Now().UTC()
	b.UpdatedAt = time.Now().UTC()
	s.data.Backends = append(s.data.Backends, b)
	return s.save()
}

// DeleteBackend removes a backend by ID.
func (s *Store) DeleteBackend(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, b := range s.data.Backends {
		if b.ID == id {
			s.data.Backends = append(s.data.Backends[:i], s.data.Backends[i+1:]...)
			return s.save()
		}
	}
	return nil
}

// ListBackends returns all persisted backends.
func (s *Store) ListBackends() ([]*backend.Backend, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]*backend.Backend, len(s.data.Backends))
	copy(result, s.data.Backends)
	return result, nil
}

// GetBackend returns a single backend by ID.
func (s *Store) GetBackend(id string) (*backend.Backend, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, b := range s.data.Backends {
		if b.ID == id {
			return b, nil
		}
	}
	return nil, os.ErrNotExist
}

// SaveStrategyConfig persists a model-to-strategy mapping.
func (s *Store) SaveStrategyConfig(model, strategy string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.data.Strategies == nil {
		s.data.Strategies = make(map[string]string)
	}
	s.data.Strategies[model] = strategy
	return s.save()
}

// GetStrategyConfig returns the strategy name for a model.
func (s *Store) GetStrategyConfig(model string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data.Strategies[model]
}

// ListStrategyConfigs returns all model→strategy mappings.
func (s *Store) ListStrategyConfigs() (map[string]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make(map[string]string, len(s.data.Strategies))
	for k, v := range s.data.Strategies {
		result[k] = v
	}
	return result, nil
}

// Close is a no-op for file-based storage.
func (s *Store) Close() error {
	return nil
}
