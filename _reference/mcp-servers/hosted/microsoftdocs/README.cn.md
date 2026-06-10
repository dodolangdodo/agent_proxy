# 🌟 Microsoft Learn 文档 MCP 服务器
[![在 VS Code 中安装](https://img.shields.io/badge/VS_Code-Install_Microsoft_Docs_MCP-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://vscode.dev/redirect/mcp/install?name=microsoft.docs.mcp&config=%7B%22type%22%3A%22http%22%2C%22url%22%3A%22https%3A%2F%2Flearn.microsoft.com%2Fapi%2Fmcp%22%7D) [![在 VS Code Insiders 中安装](https://img.shields.io/badge/VS_Code_Insiders-Install_Microsoft_Docs_MCP-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=microsoft.docs.mcp&config=%7B%22type%22%3A%22http%22%2C%22url%22%3A%22https%3A%2F%2Flearn.microsoft.com%2Fapi%2Fmcp%22%7D&quality=insiders)

Microsoft Learn 文档 MCP 服务器实现了[模型上下文协议 (MCP)](https://modelcontextprotocol.io) 服务器，为 AI 助手提供对官方 [Microsoft 文档](https://learn.microsoft.com)的实时访问。

> 请注意，此项目处于公开预览阶段，在正式发布之前实现可能会发生重大变化。

## 📑 目录
- [🌟 Microsoft Learn 文档 MCP 服务器](#-microsoft-learn-文档-mcp-服务器)
  - [📑 目录](#-目录)
  - [🎯 概述](#-概述)
    - [✨ 什么是 Microsoft Learn 文档 MCP 服务器？](#-什么是-microsoft-learn-文档-mcp-服务器)
    - [📊 主要功能](#-主要功能)
  - [🌐 Microsoft Learn 文档 MCP 服务器端点](#-microsoft-learn-文档-mcp-服务器端点)
  - [🛠️ 当前支持的工具](#️-当前支持的工具)
  - [🔌 安装和入门](#-安装和入门)
    - [替代安装（适用于旧版客户端或本地配置）](#替代安装适用于旧版客户端或本地配置)
    - [▶️ 入门指南](#️-入门指南)
  - [❓ 故障排除](#-故障排除)
    - [💻 系统提示](#-系统提示)
  - [查询 Microsoft 文档](#查询-microsoft-文档)
    - [⚠️ 常见问题](#️-常见问题)
    - [🆘 获取支持](#-获取支持)
  - [🔮 未来增强功能](#-未来增强功能)
  - [📚 其他资源](#-其他资源)

## 🎯 概述

### ✨ 什么是 Microsoft Learn 文档 MCP 服务器？

Microsoft 文档 MCP 服务器是一个云托管服务，使 GitHub Copilot 和 Cursor 等 MCP 主机能够直接从 Microsoft 官方文档中搜索和检索准确信息。通过实现标准化的模型上下文协议 (MCP)，此服务允许任何兼容的 AI 系统基于权威的 Microsoft 内容来支撑其响应。

### 📊 主要功能

- **高质量内容检索**：从 Microsoft Learn、Azure 文档、Microsoft 365 文档和其他官方 Microsoft 来源搜索和检索相关内容。
- **语义理解**：使用先进的向量搜索为任何查询找到最具上下文相关性的文档。
- **优化分块**：返回最多 10 个高质量内容块（每个最多 500 个令牌），包含文章标题、URL 和自包含的内容摘录。
- **实时更新**：访问最新发布的 Microsoft 文档。

## 🌐 Microsoft Learn 文档 MCP 服务器端点

Microsoft Learn 文档 MCP 服务器可供任何支持模型上下文协议 (MCP) 的 IDE、代理或工具访问。任何兼容的客户端都可以连接到以下**远程 MCP 端点**：

```
<https://learn.microsoft.com/api/mcp>

```
> **注意：** 此端点专为 MCP 客户端通过可流式 HTTP 进行程序化访问而设计。它不支持从 Web 浏览器直接访问，如果手动访问可能会返回 `405 Method Not Allowed` 错误。

**示例 JSON 配置：**
```json
{
  "microsoft.docs.mcp": {
    "type": "http",
    "url": "https://learn.microsoft.com/api/mcp"
  }
}
```

## 🛠️ 当前支持的工具

| 工具名称 | 描述 | 输入参数 |
|-----------|-------------|------------------|
| `microsoft_docs_search` | 对 Microsoft 官方技术文档执行语义搜索 | `query`（字符串）：用于检索的搜索查询 |

## 🔌 安装和入门

Microsoft Learn 文档 MCP 服务器支持在多个开发环境中快速安装。选择您首选的客户端以进行简化设置：

| 客户端 | 一键安装 | MCP 指南 |
|--------|----------------------|-------------------|
| **VS Code** | [![在 VS Code 中安装](https://img.shields.io/badge/VS_Code-Install_Microsoft_Docs_MCP-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://vscode.dev/redirect/mcp/install?name=microsoft.docs.mcp&config=%7B%22type%22%3A%22http%22%2C%22url%22%3A%22https%3A%2F%2Flearn.microsoft.com%2Fapi%2Fmcp%22%7D) [![在 VS Code Insiders 中安装](https://img.shields.io/badge/VS_Code_Insiders-Install_Microsoft_Docs_MCP-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=microsoft.docs.mcp&config=%7B%22type%22%3A%22http%22%2C%22url%22%3A%22https%3A%2F%2Flearn.microsoft.com%2Fapi%2Fmcp%22%7D&quality=insiders) | [VS Code MCP 官方指南](https://code.visualstudio.com/docs/copilot/chat/mcp-servers) |
| **Claude Desktop** | <details><summary>查看说明</summary>1. 打开 Claude Desktop<br/>2. 转到 **设置 → 集成**<br/>3. 点击 **添加集成**<br/>4. 输入 URL：`https://learn.microsoft.com/api/mcp`<br/>5. 点击 **连接**</details> | [Claude Desktop 远程 MCP 指南](https://support.anthropic.com/en/articles/11503834-building-custom-integrations-via-remote-mcp-servers) |
| **Visual Studio** | 需要手动配置<br/>使用 `"type": "http"` | [Visual Studio MCP 官方指南](https://learn.microsoft.com/en-us/visualstudio/ide/mcp-servers?view=vs-2022) |
| **Cursor IDE** | [![在 Cursor 中安装](https://cursor.com/deeplink/mcp-install-dark.svg)](https://cursor.com/install-mcp?name=microsoft.docs.mcp&config=eyJ0eXBlIjoiaHR0cCIsInVybCI6Imh0dHBzOi8vbGVhcm4ubWljcm9zb2Z0LmNvbS9hcGkvbWNwIn0%3D) | [Cursor MCP 官方指南](https://docs.cursor.com/context/model-context-protocol) |
| **Roo Code** | 需要手动配置<br/>使用 `"type": "streamable-http"` | [Roo Code MCP 官方指南](https://docs.roocode.com/features/mcp/using-mcp-in-roo) |
| **Cline** | 需要手动配置<br/>使用 `"type": "streamableHttp"` | [Cline MCP 官方指南](https://docs.cline.bot/mcp/connecting-to-a-remote-server) |

### 替代安装（适用于旧版客户端或本地配置）

对于不支持原生远程 MCP 服务器的客户端或如果您更喜欢本地配置，可以使用 `mcp-remote` 作为代理：

| 客户端 | 手动配置 | MCP 指南 |
|--------|----------------------|-----------|
| **Claude Desktop（旧版配置）** | <details><summary>查看配置</summary>**注意**：仅在设置 → 集成不起作用时使用<br/><pre>{<br/>  "microsoft.docs.mcp": {<br/>    "command": "npx",<br/>    "args": [<br/>      "-y",<br/>      "mcp-remote",<br/>      "https://learn.microsoft.com/api/mcp"<br/>    ]<br/>  }<br/>}</pre>添加到 `claude_desktop_config.json`</details>| [Claude Desktop MCP 指南](https://modelcontextprotocol.io/quickstart/user) |
| **Windsurf** | <details><summary>查看配置</summary><pre>{<br/>  "microsoft.docs.mcp": {<br/>    "command": "npx",<br/>    "args": [<br/>      "-y",<br/>      "mcp-remote",<br/>      "https://learn.microsoft.com/api/mcp"<br/>    ]<br/>  }<br/>}</pre> </details>| [Windsurf MCP 指南](https://docs.windsurf.com/windsurf/cascade/mcp) |

### ▶️ 入门指南

1. **对于 VS Code**：在 VS Code 中打开 GitHub Copilot 并[切换到代理模式](https://code.visualstudio.com/docs/copilot/chat/chat-agent-mode)
2. **对于 Claude Desktop**：添加集成后，您将在聊天界面中看到 MCP 工具图标
3. 您应该在可用工具列表中看到 Learn 文档 MCP 服务器
4. 尝试一个提示，告诉代理使用文档 MCP 服务器，例如"根据官方 Microsoft Learn 文档，创建 Azure 容器应用的 az cli 命令是什么？"
5. 代理应该能够使用文档 MCP 服务器工具来完成您的查询

## ❓ 故障排除

### 💻 系统提示

即使是像 Claude Sonnet 4.0 这样的工具友好模型通常也不会默认调用 MCP 工具 - 它们需要通过"系统提示"的形式给予一些鼓励。

以下是 Cursor 规则（系统提示）的示例，它将使 LLM 更频繁地使用 `microsoft.docs.mcp`：

## 查询 Microsoft 文档

您可以访问名为 `microsoft.docs.mcp` 的 MCP 服务器 - 此工具允许您搜索 Microsoft 最新的官方文档，该信息可能比您的训练数据集中的信息更详细或更新。

在处理有关如何使用原生 Microsoft 技术的问题时，例如 C#、F#、ASP.NET Core、Microsoft.Extensions、NuGet、Entity Framework、`dotnet` 运行时 - 在处理可能出现的特定/狭义定义的问题时，请使用此工具进行研究。

### ⚠️ 常见问题

| 问题 | 可能的解决方案 |
|-------|-------------------|
| 连接错误 | 验证您的网络连接并确保服务器 URL 输入正确 |
| 未返回结果 | 尝试使用更具体的技术术语重新表述您的查询 |
| 工具未在 VS Code 中显示 | 重启 VS Code 或检查 MCP 扩展是否正确安装 |
| HTTP 状态 405 | 当浏览器尝试连接到端点时会发生方法不允许错误。请尝试通过 VS Code GitHub Copilot 或 [MCP Inspector](https://modelcontextprotocol.io/docs/tools/inspector) 使用文档 MCP 服务器。 |

### 🆘 获取支持

- [提问、分享想法](https://github.com/MicrosoftDocs/mcp/discussions)
- [创建问题](https://github.com/MicrosoftDocs/mcp/issues)

## 🔮 未来增强功能

Microsoft Learn 文档 MCP 服务器团队正在开发几项增强功能：

- 扩展覆盖范围到其他 Microsoft 文档来源
- 改进查询理解以获得更精确的结果

## 📚 其他资源

- [Microsoft MCP 服务器](https://github.com/microsoft/mcp)
- [Microsoft Learn](https://learn.microsoft.com)
- [模型上下文协议规范](https://modelcontextprotocol.io)
