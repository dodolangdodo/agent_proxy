<div align="center">
  <h1>Agent Proxy V2</h1>
  <p>Lightweight AI gateway with OpenAI & Anthropic protocol translation</p>

  [![Go Version](https://img.shields.io/github/go-mod/go-version/dodolangdodo/agent_proxy)](https://github.com/dodolangdodo/agent_proxy/blob/main/go.mod)
  [![License](https://img.shields.io/github/license/dodolangdodo/agent_proxy)](https://github.com/dodolangdodo/agent_proxy/blob/main/LICENSE)
</div>

---

## Overview

Agent Proxy V2 is a lightweight, zero-dependency Go proxy that sits between your AI clients and multiple LLM providers. It translates between OpenAI and Anthropic API formats, load-balances across backends, and provides a web-based admin dashboard — all in a single binary with no external services required.

## Features

### Protocol Translation
- **Canonical IR pattern**: Parse once into an intermediate representation, then serialize to any target format
- **Auto-detection**: Identifies client protocol from URL path (`/v1/messages` = Anthropic) or headers
- **Bidirectional**: OpenAI ↔ Anthropic, or pass-through when formats match

### Load Balancing
Seven strategies assignable per-model via the admin API:

| Strategy | Mechanism |
|---|---|
| `round_robin` | Atomic counter mod N |
| `weighted_random` | Random proportional to weight |
| `weighted_round_robin` | Smooth WRR (avoids burst) |
| `least_latency` | Lowest EWMA latency |
| `priority` | Highest weight always wins (default) |
| `adaptive` | Score = weight × (minLatency/latency) × (1-errorRate) |
| `least_usage` | Lowest token usage ratio |

### Backend Management
- **Health checks**: Automatic degradation (`active` → `degraded` → `disabled`)
- **Runtime stats**: EWMA latency, error rate, token consumption
- **Auto-quota discovery**: Parses rate-limit headers from OpenAI, Anthropic, DeepSeek
- **Token balance tracking**: Manual budgets + auto-discovered quotas
- **Model mapping**: Proxy model name → real backend model name per-backend

### Filter Pipeline
1. Model resolution (exact → fuzzy `gpt-*`/`claude-*` → `auto` fallback)
2. Context-length filter (skips backends with insufficient `MaxContextTokens`)
3. Cost-tier preference (`prepaid` → `pay_per_token`)
4. Active/quota filter (skips unhealthy or exhausted backends)

### Admin Interface
- REST API for backend CRUD and strategy configuration
- Single-page web UI (`/ui/`) — dark theme, vanilla JS, no build step
- Bearer token authentication
- JSON file persistence (`proxy_state.json`)

## Quick Start

### Build

Zero external dependencies. Pure Go standard library.

```bash
go build -o proxy ./cmd/proxy/
```

### Run

```bash
# With defaults (port 9999, SQLite file storage)
PROXY_ADMIN_KEY=your-admin-key ./proxy

# Custom config via env vars
PROXY_LISTEN_ADDR=:9090 PROXY_ADMIN_KEY=mykey ./proxy
```

### Environment Variables

| Variable | Default | Description |
|---|---|---|
| `PROXY_LISTEN_ADDR` | `:9999` | HTTP listen address |
| `PROXY_ADMIN_KEY` | `admin-change-me` | Admin API bearer token |
| `PROXY_LOG_LEVEL` | `info` | Log verbosity |

## API Endpoints

### Proxy (no auth)
- `POST /v1/chat/completions` — OpenAI format
- `POST /v1/messages` — Anthropic format
- `GET /v1/models` — Model listing

### Admin (Bearer token auth)
- `GET /api/health` — Health check (no auth)
- `GET/POST /api/backends` — Backend CRUD
- `PUT/DELETE /api/backends/{id}` — Update/delete backend
- `GET/POST /api/strategies` — Strategy management
- `GET /api/strategies/config` — List strategy configs

### Web UI
- `GET /ui/` — Admin dashboard

## Architecture

```
Client (OpenAI or Anthropic format)
  → net/http ServeMux routing
  → Proxy Handler: detect format → parse to canonical IR
  → Strategy Engine: select backend
  → Provider Adapter: translate IR → backend-native format
  → HTTP call to backend
  → Response Adapter: translate backend response → IR → client format
```

### Request Flow

1. **Format detection**: URL path or `anthropic-version` header
2. **Parse to IR**: `RequestIR` with model, messages, stream flag
3. **Resolve backends**: Model → candidate backends via registry snapshot
4. **Filter pipeline**: Context window, cost tier, active status
5. **Strategy selection**: Per-model strategy from config or default
6. **Backend selection**: Strategy picks one backend from filtered pool
7. **Model resolution**: Apply per-backend `model_mapping` if configured
8. **Build request**: IR → provider-native HTTP request
9. **Execute**: HTTP call with shared connection-pooled client
10. **Response handling**: Parse → translate → serialize, or fast-path pass-through when formats match

### Key Design Decisions

- **Atomic snapshot swaps**: Registry uses `atomic.Value` for lock-free reads; all proxy handlers read without locks
- **Shared HTTP client**: Connection pooling with HTTP/2 support, not per-request clients
- **Async snapshot refresh**: Token recording and rate-limit capture no longer trigger deep-copy on the hot path
- **In-memory strategy cache**: Avoids store lock on every request; background refresh every 5s
- **Fast-path optimization**: When client format matches backend provider, skips full IR translation and forwards raw JSON with only model replacement

## Backend Configuration

Example backend JSON for the admin API:

```json
{
  "name": "deepseek-pro",
  "provider": "anthropic",
  "base_url": "https://api.deepseek.com/anthropic",
  "api_key": "sk-...",
  "models": ["deepseek-v4-pro"],
  "weight": 10,
  "max_context_tokens": 64000,
  "timeout_seconds": 120,
  "model_mapping": {
    "deepseek-v4-pro": "deepseek-chat"
  }
}
```

### Backend Fields

| Field | Type | Description |
|---|---|---|
| `name` | string | Display name |
| `provider` | string | `"openai"`, `"anthropic"`, or `"custom"` |
| `base_url` | string | Provider API base URL |
| `api_key` | string | Provider API key (encrypted at rest) |
| `models` | []string | Models this backend serves |
| `model_mapping` | map | Proxy model → real backend model |
| `weight` | int | Priority/weight for load balancing |
| `cost_tier` | string | `"prepaid"` or `"pay_per_token"` |
| `max_context_tokens` | int | Maximum context window (0 = unlimited) |
| `skip_context_filter` | bool | Bypass context-length filter |
| `max_rpm` | int | Rate limit (not yet enforced) |
| `max_concurrent` | int | Concurrency limit (not yet enforced) |
| `timeout_seconds` | int | Request timeout |
| `token_balance` | int64 | Manual token budget (0 = unlimited) |

## Project Structure

```
.
├── cmd/proxy/           # Entry point, server setup
├── internal/
│   ├── config/          # Env var loading, atomic config snapshot
│   ├── protocol/        # Canonical IR, parsers, serializers, adapters
│   ├── backend/         # Backend model, registry, runtime stats, health checker
│   ├── strategy/        # Load balancing strategies
│   ├── proxy/           # Main request handler
│   ├── api/             # REST admin API
│   └── store/           # JSON file persistence
├── web/dist/            # Single-file SPA admin UI
└── _reference/          # Original labring/aiproxy codebase (reference)
```

## Performance Notes

The proxy is optimized for minimal overhead on the hot path:

- Lock-free backend reads via atomic snapshot
- Connection reuse via shared `http.Client` with pooling
- Request/response fast-path when client and backend formats match (skips full IR translation)
- Async registry snapshot refresh (avoids deep-copy per request)
- In-memory strategy cache (avoids store lock per request)

For Anthropic → Anthropic routing (the most common case), the proxy adds only:
1. Model name extraction from JSON
2. Backend selection via registry snapshot
3. Model replacement in raw JSON body
4. Forward response (streaming or buffered)

## Development

### Requirements
- Go 1.22+

### Build
```bash
go build -o proxy ./cmd/proxy/
```

### Run tests
```bash
go test ./...
```

## License

MIT License — see [LICENSE](LICENSE) file.
