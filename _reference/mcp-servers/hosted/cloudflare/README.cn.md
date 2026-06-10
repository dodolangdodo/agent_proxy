# Cloudflare MCP 服务器

模型上下文协议（MCP）是一个[新的标准化协议](https://modelcontextprotocol.io/introduction)，用于管理大型语言模型（LLM）与外部系统之间的上下文。在这个仓库中，您可以找到多个 MCP 服务器，允许您从 MCP 客户端（例如 Cursor、Claude）连接到 Cloudflare 的服务，并使用自然语言通过您的 Cloudflare 账户完成任务。

这些 MCP 服务器允许您的 [MCP 客户端](https://modelcontextprotocol.io/clients)从您的账户读取配置、处理信息、基于数据提出建议，甚至为您执行这些建议的更改。所有这些操作都可以跨 Cloudflare 的多项服务进行，包括应用程序开发、安全性和性能。

此仓库包含以下服务器：

| 服务器名称                                                      | 描述                                                                                           | 服务器 URL                                      |
| -------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ---------------------------------------------- |
| [**文档服务器**](/apps/docs-vectorize)                         | 获取 Cloudflare 的最新参考信息                                                                  | `https://docs.mcp.cloudflare.com/sse`          |
| [**Workers 绑定服务器**](/apps/workers-bindings)               | 使用存储、AI 和计算原语构建 Workers 应用程序                                                     | `https://bindings.mcp.cloudflare.com/sse`      |
| [**Workers 构建服务器**](/apps/workers-builds)                 | 获取洞察并管理您的 Cloudflare Workers 构建                                                      | `https://builds.mcp.cloudflare.com/sse`        |
| [**可观测性服务器**](/apps/workers-observability)              | 调试并深入了解您应用程序的日志和分析                                                            | `https://observability.mcp.cloudflare.com/sse` |
| [**Radar 服务器**](/apps/radar)                                | 获取全球互联网流量洞察、趋势、URL 扫描和其他实用工具                                            | `https://radar.mcp.cloudflare.com/sse`         |
| [**容器服务器**](/apps/sandbox-container)                      | 启动沙盒开发环境                                                                               | `https://containers.mcp.cloudflare.com/sse`    |
| [**浏览器渲染服务器**](/apps/browser-rendering)                | 获取网页、将其转换为 markdown 并截图                                                           | `https://browser.mcp.cloudflare.com/sse`       |
| [**Logpush 服务器**](/apps/logpush)                            | 获取 Logpush 作业健康状况的快速摘要                                                            | `https://logs.mcp.cloudflare.com/sse`          |
| [**AI Gateway 服务器**](/apps/ai-gateway)                      | 搜索您的日志，获取提示和响应的详细信息                                                          | `https://ai-gateway.mcp.cloudflare.com/sse`    |
| [**AutoRAG 服务器**](/apps/autorag)                            | 列出并搜索您的 AutoRAG 上的文档                                                               | `https://autorag.mcp.cloudflare.com/sse`       |
| [**审计日志服务器**](/apps/auditlogs)                          | 查询审计日志并生成报告供审查                                                                   | `https://auditlogs.mcp.cloudflare.com/sse`     |
| [**DNS 分析服务器**](/apps/dns-analytics)                      | 基于当前设置优化 DNS 性能并调试问题                                                            | `https://dns-analytics.mcp.cloudflare.com/sse` |
| [**数字体验监控服务器**](/apps/dex-analysis)                   | 快速洞察您组织的关键应用程序                                                                   | `https://dex.mcp.cloudflare.com/sse`           |
| [**Cloudflare One CASB 服务器**](/apps/cloudflare-one-casb)    | 快速识别 SaaS 应用程序的任何安全配置错误，以保护用户和数据                                       | `https://casb.mcp.cloudflare.com/sse`          |
| [**GraphQL 服务器**](/apps/graphql/)                           | 使用 Cloudflare 的 GraphQL API 获取分析数据                                                   | `https://graphql.mcp.cloudflare.com/sse`       |

## 从任何 MCP 客户端访问远程 MCP 服务器

如果您的 MCP 客户端对远程 MCP 服务器有一流的支持，客户端将提供一种在其界面内直接接受服务器 URL 的方式（例如 [Cloudflare AI Playground](https://playground.ai.cloudflare.com/)）

如果您的客户端尚不支持远程 MCP 服务器，您需要使用 mcp-remote (<https://www.npmjs.com/package/mcp-remote>) 设置其相应的配置文件，以指定您的客户端可以访问哪些服务器。

```json
{
 "mcpServers": {
  "cloudflare-observability": {
   "command": "npx",
   "args": ["mcp-remote", "https://observability.mcp.cloudflare.com/sse"]
  },
  "cloudflare-bindings": {
   "command": "npx",
   "args": ["mcp-remote", "https://bindings.mcp.cloudflare.com/sse"]
  }
 }
}
```

## 从 OpenAI Responses API 使用 Cloudflare 的 MCP 服务器

要在 [OpenAI 的 responses API](https://openai.com/index/new-tools-and-features-in-the-responses-api/) 中使用 Cloudflare 的 MCP 服务器之一，您需要为 Responses API 提供一个具有该特定 MCP 服务器所需范围（权限）的 API 令牌。

例如，要在 OpenAI 中使用[浏览器渲染 MCP 服务器](https://github.com/cloudflare/mcp-server-cloudflare/tree/main/apps/browser-rendering)，请在 Cloudflare 仪表板[这里](https://dash.cloudflare.com/profile/api-tokens)创建一个 API 令牌，具有以下权限：

<img width="937" alt="Screenshot 2025-05-21 at 10 38 02 AM" src="https://github.com/user-attachments/assets/872e253f-23ce-43b3-983c-45f9d0f66100" />

## 需要访问更多 Cloudflare 工具？

我们正在继续为这个远程 MCP 服务器仓库添加更多功能。如果您想留下反馈、报告错误或提供功能请求，[请在此仓库上开启一个 issue](https://github.com/cloudflare/mcp-server-cloudflare/issues/new/choose)

## 故障排除

"Claude 的响应被中断..."

如果您看到此消息，Claude 可能达到了其上下文长度限制并在回复中途停止。这在触发许多链式工具调用的服务器（如可观测性服务器）上最常发生。

要减少遇到此问题的机会：

- 尽量具体，保持查询简洁。
- 如果单个请求调用多个工具，尝试将其分解为几个较小的工具调用，以保持响应简短。

## 付费功能

某些功能可能需要付费的 Cloudflare Workers 计划。确保您的 Cloudflare 账户具有您打算使用的功能所需的订阅级别。

## 贡献

有兴趣贡献并在本地运行此服务器？请参阅 [CONTRIBUTING.md](CONTRIBUTING.md) 开始使用。
