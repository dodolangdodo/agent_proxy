# Cache Plugin 配置指南

## 概述

Cache Plugin 是一个高性能的 AI API 请求缓存解决方案，通过存储和重用相同请求的响应来帮助减少延迟和成本。它支持内存缓存和 Redis，适用于分布式部署。

## 功能特性

- **双重存储**：支持内存缓存和 Redis，提供灵活的部署选项
- **自动降级**：Redis 不可用时自动降级到内存缓存
- **基于内容的缓存**：使用请求体的 SHA256 哈希值生成缓存键
- **可配置 TTL**：为缓存项设置自定义生存时间
- **大小限制**：可配置最大项目大小以防止内存问题
- **缓存头部**：可选的头部信息来指示缓存命中
- **零拷贝设计**：通过缓冲池实现高效的内存使用

## 配置示例

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

## 配置字段说明

### 插件配置

| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `enable` | bool | 是 | false | 是否启用 Cache 插件 |
| `ttl` | int | 否 | 300 | 缓存项的生存时间（秒） |
| `item_max_size` | int | 否 | 1048576 (1MB) | 单个缓存项的最大大小（字节） |
| `add_cache_hit_header` | bool | 否 | false | 是否添加指示缓存命中的头部 |
| `cache_hit_header` | string | 否 | "X-Aiproxy-Cache" | 缓存命中头部的名称 |

## 工作原理

### 缓存键生成

插件基于以下内容生成缓存键：

1. 请求模式（如 chat completions）
2. 请求体的 SHA256 哈希值

这确保了相同的请求会命中缓存，而不同的请求不会相互干扰。

### 缓存存储

插件使用两层缓存策略：

1. **Redis（如果可用）**：分布式缓存的主要存储
2. **内存**：备用存储或未配置 Redis 时的主要存储

### 请求流程

1. **请求阶段**：
   - 插件检查是否启用缓存
   - 从请求体生成缓存键
   - 查找缓存（先查 Redis，再查内存）
   - 如果命中，立即返回缓存的响应
   - 如果未命中，继续请求上游 API

2. **响应阶段**：
   - 捕获响应体和头部
   - 如果响应成功，存储到缓存
   - 遵守大小限制以防止内存问题

## 使用示例

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

## 响应头部示例

当 `add_cache_hit_header` 启用时：

**缓存命中：**

```
X-Aiproxy-Cache: hit
```

**缓存未命中：**

```
X-Aiproxy-Cache: miss
```
