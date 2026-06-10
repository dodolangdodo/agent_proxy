# Cache Plugin Configuration Guide

## Overview

The Cache Plugin is a high-performance caching solution for AI API requests that helps reduce latency and costs by storing and reusing responses for identical requests. It supports both in-memory caching and Redis, making it suitable for distributed deployments.

## Features

- **Dual Storage**: Supports both in-memory cache and Redis for flexible deployment options
- **Automatic Fallback**: Automatically falls back to in-memory cache when Redis is unavailable
- **Content-Based Caching**: Uses SHA256 hash of request body to generate cache keys
- **Configurable TTL**: Set custom time-to-live for cached items
- **Size Limits**: Configurable maximum item size to prevent memory issues
- **Cache Headers**: Optional headers to indicate cache hits
- **Zero-Copy Design**: Efficient memory usage through buffer pooling

## Configuration Example

```json
{
    "model": "gpt-4",
    "type": 1,
    "plugin": {
        "cache": {
            "enable": true,
            "ttl": 300,
            "item_max_size": 1048576,
            "add_cache_hit_header": true,
            "cache_hit_header": "X-Cache-Status"
        }
    }
}
```

## Configuration Fields

### Plugin Configuration

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `enable` | bool | Yes | false | Whether to enable the Cache plugin |
| `ttl` | int | No | 300 | Time-to-live for cached items (in seconds) |
| `item_max_size` | int | No | 1048576 (1MB) | Maximum size of a single cached item (in bytes) |
| `add_cache_hit_header` | bool | No | false | Whether to add a header indicating cache hit |
| `cache_hit_header` | string | No | "X-Aiproxy-Cache" | Name of the cache hit header |

## How It Works

### Cache Key Generation

The plugin generates cache keys based on:

1. Request pattern (e.g., chat completions)
2. SHA256 hash of the request body

This ensures identical requests hit the cache while different requests don't interfere with each other.

### Cache Storage

The plugin uses a two-tier caching strategy:

1. **Redis (if available)**: Primary storage for distributed caching
2. **Memory**: Fallback storage or primary when Redis is not configured

### Request Flow

1. **Request Phase**:
   - Plugin checks if caching is enabled
   - Generates cache key from request body
   - Looks up cache (Redis first, then memory)
   - If hit, immediately returns cached response
   - If miss, continues to upstream API

2. **Response Phase**:
   - Captures response body and headers
   - If response is successful, stores in cache
   - Respects size limits to prevent memory issues

## Usage Example

```json
{
    "plugin": {
        "cache": {
            "enable": true,
            "ttl": 60,
            "item_max_size": 524288,
            "add_cache_hit_header": true
        }
    }
}
```

## Response Header Example

When `add_cache_hit_header` is enabled:

**Cache Hit:**

```
X-Aiproxy-Cache: hit
```

**Cache Miss:**

```
X-Aiproxy-Cache: miss
```
