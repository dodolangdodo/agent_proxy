# Think Split Plugin Configuration Guide

## Overview

Think Split Plugin is a plugin designed to handle thinking processes in AI model responses. It automatically identifies and separates `<think>` tag content in responses, extracting the thinking process from the main answer into a separate `reasoning_content` field, allowing AI's thinking process and final answer to be displayed and processed separately.

## Features

- **Automatic Recognition**: Automatically detects `<think>...</think>` tags in responses
- **Content Separation**: Extracts thinking content to `reasoning_content` field
- **Streaming Support**: Supports both streaming and non-streaming response processing
- **Zero Intrusion**: Doesn't affect original response structure, only adds new fields
- **High Performance**: Uses efficient KMP algorithm for pattern matching

## How It Works

### Tag Recognition

The plugin recognizes thinking content in the following formats:

```
<think>
AI's thinking process here...
</think>
```

or

```
\n<think>
AI's thinking process here...
</think>\n
```

### Content Transformation

**Before:**

```json
{
  "choices": [{
    "message": {
      "content": "\n<think>\nThis is a math question...\n</think>\nThe answer is 42."
    }
  }]
}
```

**After:**

```json
{
  "choices": [{
    "message": {
      "content": "The answer is 42.",
      "reasoning_content": "This is a math question..."
    }
  }]
}
```

## Configuration Examples

### Basic Configuration

```json
{
  "model": "gpt-4",
  "type": 1,
  "plugin": {
    "think-split": {
      "enable": true
    }
  }
}
```

### Complete Configuration Example

```json
{
  "model": "claude-3-opus",
  "type": 1,
  "retry_times": 3,
  "plugin": {
    "think-split": {
      "enable": true
    }
  }
}
```

## Configuration Field Description

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `enable` | bool | Yes | false | Whether to enable Think Split plugin |

## Working with Other Plugins

Think Split Plugin can work together with other plugins:

```json
{
  "model": "gpt-4",
  "plugin": {
    "think-split": {
      "enable": true
    },
    "cache": {
      "enable": true,
      "ttl": 300
    },
    "web-search": {
      "enable": true
    }
  }
}
```

## Important Notes

1. **Performance Impact**: The plugin uses efficient KMP algorithm with minimal performance impact
2. **Content Integrity**: Ensure `<think>` tags are properly closed, otherwise recognition may fail
3. **Nested Handling**: Nested `<think>` tags are not supported
4. **Field Conflicts**: If original response already contains `reasoning_content` field, it will be overwritten

## Troubleshooting

### Plugin Not Working

1. Check if `enable` is set to `true`
2. Confirm model configuration is loaded correctly
3. Verify response actually contains `<think>` tags

### Content Not Properly Separated

1. Check if `<think>` tag format is correct
2. Confirm tags are properly closed
3. Check logs for error messages

### Streaming Response Issues

1. Confirm client properly handles streaming responses
2. Check if SSE format is parsed correctly
3. Verify `reasoning_content` field is in delta

## API Response Examples

### Non-Streaming Response

```json
{
  "id": "chatcmpl-123",
  "object": "chat.completion",
  "created": 1677652288,
  "model": "gpt-4",
  "choices": [{
    "index": 0,
    "message": {
      "role": "assistant",
      "content": "Quantum entanglement is a phenomenon in quantum mechanics...",
      "reasoning_content": "The user is asking about quantum entanglement, which is a core concept in quantum physics..."
    },
    "finish_reason": "stop"
  }],
  "usage": {
    "prompt_tokens": 9,
    "completion_tokens": 12,
    "total_tokens": 21
  }
}
```

### Streaming Response

```json
data: {"id":"chatcmpl-123","object":"chat.completion.chunk","created":1677652288,"model":"gpt-4","choices":[{"index":0,"delta":{"reasoning_content":"Starting to think about the problem..."},"finish_reason":null}]}

data: {"id":"chatcmpl-123","object":"chat.completion.chunk","created":1677652288,"model":"gpt-4","choices":[{"index":0,"delta":{"content":"Quantum entanglement is"},"finish_reason":null}]}

data: [DONE]
```
