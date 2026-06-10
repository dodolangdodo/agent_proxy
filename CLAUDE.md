# CLAUDE.md

This file provides guidance to Claude Code when working with code in this repository.

## Build & Run

```bash
# Build (zero external dependencies, pure Go stdlib)
export PATH=$HOME/go-local/go/bin:$PATH
export GOTOOLCHAIN=local
go build -o proxy ./cmd/proxy/

# Run with defaults (port 8080, SQLite file storage)
PROXY_ADMIN_KEY=your-admin-key ./proxy

# Custom config via env vars
PROXY_LISTEN_ADDR=:9090 PROXY_ADMIN_KEY=mykey ./proxy
```

The admin API key is set via `PROXY_ADMIN_KEY` env var (default: `admin-change-me`).

## Architecture

### Request Flow

```
Client (OpenAI or Anthropic format)
  → net/http ServeMux routing
  → Proxy Handler: detect format → parse to canonical IR
  → Strategy Engine: select backend (round_robin, weighted, adaptive, etc.)
  → Provider Adapter: translate IR → backend-native format
  → HTTP call to backend
  → Response Adapter: translate backend response → IR → client format
```

### Package Map

| Package | Purpose |
|---|---|
| `cmd/proxy/` | Entry point, server setup, wiring |
| `internal/config/` | Atomic config snapshot, env var loading |
| `internal/protocol/` | Canonical IR, OpenAI/Anthropic parsers/serializers, ProviderAdapter interface |
| `internal/backend/` | Backend model, registry (CRUD + snapshot), runtime stats (EWMA latency/error), health checker |
| `internal/strategy/` | Load balancing: round_robin, weighted_random, weighted_round_robin, least_latency, priority, adaptive |
| `internal/proxy/` | Main request handler: format detection, routing, orchestration |
| `internal/api/` | REST admin API (backends CRUD, strategy config, model listing) |
| `internal/store/` | JSON file persistence for backends and strategy configs |
| `web/dist/` | Single-file SPA admin UI (dark theme, vanilla JS, no build step) |

### Protocol Translation

Uses a canonical Intermediate Representation (IR) pattern — not N×N direct translation.

```
OpenAI Request   → ParseOpenAIRequest()   → RequestIR
Anthropic Request → ParseAnthropicRequest() → RequestIR
RequestIR → OpenAIAdapter.BuildRequest()    → OpenAI HTTP call
RequestIR → AnthropicAdapter.BuildRequest() → Anthropic HTTP call
Response  → ParseResponse() → ResponseIR → SerializeOpenAIResponse/SerializeAnthropicResponse()
```

Protocol detection: URL path (`/v1/messages` = Anthropic) or `anthropic-version` header.

### Load Balancing Strategies

Six strategies implementing `LoadBalanceStrategy` interface (`Select` + `RecordResult`). Assigned per-model via `/api/strategies`. Default: `priority`.

| Strategy | Mechanism |
|---|---|
| `round_robin` | Atomic counter mod N |
| `weighted_random` | Random proportional to backend weight |
| `weighted_round_robin` | Smooth WRR (avoids burst) |
| `least_latency` | Lowest EWMA latency backend |
| `priority` | Highest weight always wins (hot-standby) |
| `adaptive` | Score = weight × (minLatency/latency) × (1-errorRate) |

### Configuration

Three tiers: env vars > config struct > defaults. Dynamic config (backends, strategies) stored in `proxy_state.json`, loaded at startup, persisted via REST API. Backend registry uses atomic snapshot swaps — all handlers read lock-free.

### API Endpoints

**Proxy** (no auth):
- `POST /v1/chat/completions` — OpenAI format
- `POST /v1/messages` — Anthropic format
- `GET /v1/models` — Model listing

**Admin** (Bearer token auth):
- `GET /api/health` — Health check (no auth)
- `GET/POST /api/backends`, `PUT/DELETE /api/backends/{id}` — Backend CRUD
- `GET/POST /api/strategies`, `GET /api/strategies/config` — Strategy management

**UI**: `/ui/` — Single-page admin dashboard

## Reference Code

`_reference/` contains the original labring/aiproxy codebase (Go 1.26, gin, GORM, 40+ provider adaptors). Used for architecture reference. Our implementation is a from-scratch reimplementation focused on the five core requirements.
