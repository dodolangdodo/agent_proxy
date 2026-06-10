package strategy

import (
	"context"
	"math/rand/v2"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"agent-proxy-v2/internal/backend"
)

// LoadBalanceStrategy selects a backend from a pool for a given request.
type LoadBalanceStrategy interface {
	// Name returns the strategy identifier (for config and UI).
	Name() string
	// Select picks a backend from the slice. Returns nil if none are available.
	Select(ctx context.Context, backends []*backend.Backend, registry *backend.Registry) *backend.Backend
	// RecordResult feeds back the outcome for adaptive strategies.
	RecordResult(backendID string, registry *backend.Registry)
}

// registry holds named strategy constructors.
var constructors = map[string]func() LoadBalanceStrategy{}
var instances sync.Map // string → LoadBalanceStrategy (lazy singletons)

func Register(name string, fn func() LoadBalanceStrategy) {
	constructors[name] = fn
}

func Get(name string) LoadBalanceStrategy {
	if inst, ok := instances.Load(name); ok {
		return inst.(LoadBalanceStrategy)
	}
	if fn, ok := constructors[name]; ok {
		inst := fn()
		instances.Store(name, inst)
		return inst
	}
	return Get("priority") // default fallback
}

func Names() []string {
	names := make([]string, 0, len(constructors))
	for n := range constructors {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}

func init() {
	Register("round_robin", func() LoadBalanceStrategy { return &RoundRobin{} })
	Register("weighted_random", func() LoadBalanceStrategy { return &WeightedRandom{} })
	Register("weighted_round_robin", func() LoadBalanceStrategy { return &WeightedRoundRobin{} })
	Register("least_latency", func() LoadBalanceStrategy { return &LeastLatency{} })
	Register("priority", func() LoadBalanceStrategy { return &Priority{} })
	Register("adaptive", func() LoadBalanceStrategy { return &Adaptive{} })
	Register("least_usage", func() LoadBalanceStrategy { return &LeastUsage{} })
}

// --- Round Robin ---

type RoundRobin struct {
	counter atomic.Uint64
}

func (s *RoundRobin) Name() string { return "round_robin" }

func (s *RoundRobin) Select(ctx context.Context, backends []*backend.Backend, reg *backend.Registry) *backend.Backend {
	active := filterActive(backends)
	if len(active) == 0 {
		return nil
	}
	idx := s.counter.Add(1) % uint64(len(active))
	return active[idx]
}

func (s *RoundRobin) RecordResult(string, *backend.Registry) {}

// --- Weighted Random ---

type WeightedRandom struct {
	rng *rand.Rand
	mu  sync.Mutex
}

func (s *WeightedRandom) Name() string { return "weighted_random" }

func (s *WeightedRandom) Select(ctx context.Context, backends []*backend.Backend, reg *backend.Registry) *backend.Backend {
	active := filterActive(backends)
	if len(active) == 0 {
		return nil
	}
	totalWeight := 0
	for _, b := range active {
		w := b.Weight
		if w <= 0 {
			w = 1
		}
		totalWeight += w
	}
	if totalWeight == 0 {
		return active[0]
	}
	r := rand.IntN(totalWeight)
	for _, b := range active {
		w := b.Weight
		if w <= 0 {
			w = 1
		}
		r -= w
		if r < 0 {
			return b
		}
	}
	return active[len(active)-1]
}

func (s *WeightedRandom) RecordResult(string, *backend.Registry) {}

// --- Weighted Round Robin ---

type WeightedRoundRobin struct {
	mu       sync.Mutex
	states   map[string]*wrrState
}

type wrrState struct {
	currentWeight int
}

func (s *WeightedRoundRobin) Name() string { return "weighted_round_robin" }

func (s *WeightedRoundRobin) Select(ctx context.Context, backends []*backend.Backend, reg *backend.Registry) *backend.Backend {
	active := filterActive(backends)
	if len(active) == 0 {
		return nil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.states == nil {
		s.states = make(map[string]*wrrState)
	}

	// Smooth weighted round-robin
	totalWeight := 0
	for _, b := range active {
		w := b.Weight
		if w <= 0 {
			w = 1
		}
		totalWeight += w
	}

	var selected *backend.Backend
	maxWeight := -1
	for _, b := range active {
		st, ok := s.states[b.ID]
		if !ok {
			st = &wrrState{}
			s.states[b.ID] = st
		}
		w := b.Weight
		if w <= 0 {
			w = 1
		}
		st.currentWeight += w
		if st.currentWeight > maxWeight {
			maxWeight = st.currentWeight
			selected = b
		}
	}

	if selected != nil {
		s.states[selected.ID].currentWeight -= totalWeight
	}

	return selected
}

func (s *WeightedRoundRobin) RecordResult(string, *backend.Registry) {}

// --- Least Latency ---

type LeastLatency struct{}

func (s *LeastLatency) Name() string { return "least_latency" }

func (s *LeastLatency) Select(ctx context.Context, backends []*backend.Backend, reg *backend.Registry) *backend.Backend {
	active := filterActive(backends)
	if len(active) == 0 {
		return nil
	}

	var best *backend.Backend
	var bestLatency int64
	// Fall back to random among untested backends
	untested := make([]*backend.Backend, 0)

	for _, b := range active {
		rt := reg.GetRuntime(b.ID)
		if rt == nil {
			untested = append(untested, b)
			continue
		}
		lat := rt.GetLatency()
		if lat == 0 {
			untested = append(untested, b)
			continue
		}
		if best == nil || int64(lat) < bestLatency {
			best = b
			bestLatency = int64(lat)
		}
	}

	if best != nil {
		return best
	}
	if len(untested) > 0 {
		return untested[rand.IntN(len(untested))]
	}
	return active[0]
}

func (s *LeastLatency) RecordResult(string, *backend.Registry) {}

// --- Priority ---

type Priority struct{}

func (s *Priority) Name() string { return "priority" }

func (s *Priority) Select(ctx context.Context, backends []*backend.Backend, reg *backend.Registry) *backend.Backend {
	active := filterActive(backends)
	if len(active) == 0 {
		return nil
	}
	sort.Slice(active, func(i, j int) bool {
		return active[i].Weight > active[j].Weight
	})
	return active[0]
}

func (s *Priority) RecordResult(string, *backend.Registry) {}

// --- Adaptive ---

type Adaptive struct{}

func (s *Adaptive) Name() string { return "adaptive" }

func (s *Adaptive) Select(ctx context.Context, backends []*backend.Backend, reg *backend.Registry) *backend.Backend {
	active := filterActive(backends)
	if len(active) == 0 {
		return nil
	}

	// Find minimum latency for normalization
	var minLatency time.Duration
	for _, b := range active {
		rt := reg.GetRuntime(b.ID)
		if rt != nil {
			lat := rt.GetLatency()
			if minLatency == 0 || (lat > 0 && lat < minLatency) {
				minLatency = lat
			}
		}
	}
	if minLatency == 0 {
		minLatency = time.Millisecond // fallback
	}

	var best *backend.Backend
	var bestScore float64
	for _, b := range active {
		rt := reg.GetRuntime(b.ID)
		if rt == nil {
			continue
		}
		w := b.Weight
		if w <= 0 {
			w = 1
		}
		score := rt.Score(w, minLatency)
		if best == nil || score > bestScore {
			best = b
			bestScore = score
		}
	}
	if best == nil {
		return active[0]
	}
	return best
}

func (s *Adaptive) RecordResult(string, *backend.Registry) {}

// --- Least Usage ---

// LeastUsage selects the backend with the lowest token usage ratio.
// Backends with unlimited balance (TokenBalance==0) are treated as having ratio 0.
type LeastUsage struct{}

func (s *LeastUsage) Name() string { return "least_usage" }

func (s *LeastUsage) Select(ctx context.Context, backends []*backend.Backend, reg *backend.Registry) *backend.Backend {
	active := filterActive(backends)
	if len(active) == 0 {
		return nil
	}

	var best *backend.Backend
	var minRatio float64 = 2.0 // start above max (1.0)
	for _, b := range active {
		ratio := b.UsageRatio()
		// Unlimited: no manual balance and no auto-quota
		if b.TokenBalance <= 0 && b.QuotaTokensTotal <= 0 {
			ratio = 0
		}
		if best == nil || ratio < minRatio {
			best = b
			minRatio = ratio
		}
	}
	return best
}

func (s *LeastUsage) RecordResult(string, *backend.Registry) {}

// --- Filter Pipeline ---

// FilterByContextWindow removes backends whose MaxContextTokens is smaller
// than estimatedTokens. Backends with MaxContextTokens == 0 pass through.
// If all backends would be filtered, returns the single backend with the
// largest context window so the request can still be attempted.
func FilterByContextWindow(backends []*backend.Backend, estTokens int) []*backend.Backend {
	if estTokens <= 0 || len(backends) == 0 {
		return backends
	}
	result := make([]*backend.Backend, 0, len(backends))
	for _, b := range backends {
		if b.SkipContextFilter || b.MaxContextTokens <= 0 || b.MaxContextTokens >= estTokens {
			result = append(result, b)
		}
	}
	if len(result) == 0 {
		// All too small — pick the one with the largest context window
		var best *backend.Backend
		for _, b := range backends {
			if best == nil || b.MaxContextTokens > best.MaxContextTokens {
				best = b
			}
		}
		if best != nil {
			return []*backend.Backend{best}
		}
		return backends
	}
	return result
}

// FilterByCostTier selects the highest-priority cost tier with at least one
// usable backend. Prepaid backends are preferred; pay-per-token (or empty
// cost_tier) is the fallback. Returns nil if nothing is usable.
func FilterByCostTier(backends []*backend.Backend) []*backend.Backend {
	if len(backends) == 0 {
		return nil
	}
	prepaid := make([]*backend.Backend, 0)
	payPerToken := make([]*backend.Backend, 0)
	for _, b := range backends {
		if b.CostTier == "prepaid" {
			prepaid = append(prepaid, b)
		} else {
			payPerToken = append(payPerToken, b)
		}
	}
	if active := filterActive(prepaid); len(active) > 0 {
		return active
	}
	if active := filterActive(payPerToken); len(active) > 0 {
		return active
	}
	return nil
}

// --- Helpers ---

// filterActive returns backends that are active AND have remaining token balance/quota.
func filterActive(backends []*backend.Backend) []*backend.Backend {
	result := make([]*backend.Backend, 0, len(backends))
	for _, b := range backends {
		if b.Status != backend.StatusActive {
			continue
		}
		// Skip if auto-discovered quota is exhausted
		if b.QuotaTokensTotal > 0 && b.QuotaTokensRemaining <= 0 {
			continue
		}
		// Skip if manual token balance is exhausted
		if b.TokenBalance > 0 && b.TokenUsed >= b.TokenBalance {
			continue
		}
		result = append(result, b)
	}
	return result
}
