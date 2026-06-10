# mcp-server-flomo MCP 服务器

> <https://github.com/chatmcp/mcp-server-flomo>

向Flomo写入笔记。

这是一个基于TypeScript的MCP服务器，帮助您向Flomo写入笔记。

## 功能特性

### 工具

- `write_note` - 向Flomo写入文本笔记
  - 需要content作为必需参数

## 配置

- `flomo_api_url` - Flomo API URL
  - 必需
  - 示例: `https://flomoapp.com/iwh/xxx`
  - 描述: 用于写入笔记的Flomo API webhook URL
  - 在[flomo](https://v.flomoapp.com/mine?source=incoming_webhook)中找到您的api url
