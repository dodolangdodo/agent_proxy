# Think Split Plugin 配置指南

## 概述

Think Split Plugin 是一个用于处理 AI 模型响应中思考过程的插件。它能够自动识别并分离响应内容中的 `<think>` 标签内容，将思考过程从主要回答中提取到独立的 `reasoning_content` 字段，使得 AI 的思考过程和最终答案能够分开展示和处理。

## 功能特性

- **自动识别**：自动检测响应中的 `<think>...</think>` 标签
- **内容分离**：将思考内容提取到 `reasoning_content` 字段
- **流式支持**：支持流式和非流式响应处理
- **零侵入**：不影响原有响应结构，仅添加新字段
- **高性能**：使用高效的 KMP 算法进行模式匹配

## 工作原理

### 标签识别

插件识别以下格式的思考内容：

```
<think>
这里是 AI 的思考过程...
</think>
```

或

```
\n<think>
这里是 AI 的思考过程...
</think>\n
```

### 内容转换

**转换前：**

```json
{
  "choices": [{
    "message": {
      "content": "\n<think>\n这是一个关于数学的问题...\n</think>\n答案是 42。"
    }
  }]
}
```

**转换后：**

```json
{
  "choices": [{
    "message": {
      "content": "答案是 42。",
      "reasoning_content": "这是一个关于数学的问题..."
    }
  }]
}
```

## 配置示例

### 基础配置

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

### 完整配置示例

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

## 配置字段说明

| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `enable` | bool | 是 | false | 是否启用 Think Split 插件 |

## 与其他插件的配合

Think Split Plugin 可以与其他插件配合使用：

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

## 注意事项

1. **性能影响**：插件使用高效的 KMP 算法，对性能影响极小
2. **内容完整性**：确保 `<think>` 标签正确闭合，否则可能无法正确识别
3. **嵌套处理**：不支持嵌套的 `<think>` 标签
4. **字段冲突**：如果原响应已包含 `reasoning_content` 字段，会被覆盖

## 故障排查

### 插件未生效

1. 检查 `enable` 是否设置为 `true`
2. 确认模型配置正确加载
3. 验证响应中确实包含 `<think>` 标签

### 内容未正确分离

1. 检查 `<think>` 标签格式是否正确
2. 确认标签正确闭合
3. 查看日志中是否有错误信息

### 流式响应问题

1. 确认客户端正确处理流式响应
2. 检查是否正确解析 SSE 格式
3. 验证 `reasoning_content` 字段是否在 delta 中

## API 响应示例

### 非流式响应

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
      "content": "量子纠缠是量子力学中的一种现象...",
      "reasoning_content": "用户询问量子纠缠，这是量子物理学中的核心概念..."
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

### 流式响应

```json
data: {"id":"chatcmpl-123","object":"chat.completion.chunk","created":1677652288,"model":"gpt-4","choices":[{"index":0,"delta":{"reasoning_content":"开始思考问题..."},"finish_reason":null}]}

data: {"id":"chatcmpl-123","object":"chat.completion.chunk","created":1677652288,"model":"gpt-4","choices":[{"index":0,"delta":{"content":"量子纠缠是"},"finish_reason":null}]}

data: [DONE]
```
