# Semgrep MCP 服务器

<p align="center">
  <a href="https://semgrep.dev">
    <picture>
      <source media="(prefers-color-scheme: light)" srcset="images/semgrep-logo-light.svg">
      <source media="(prefers-color-scheme: dark)" srcset="images/semgrep-logo-dark.svg">
      <img src="https://raw.githubusercontent.com/semgrep/mcp/main/images/semgrep-logo-light.svg" height="60" alt="Semgrep logo"/>
    </picture>
  </a>
</p>
<p align="center">
  <a href="https://semgrep.dev/docs/">
      <img src="https://img.shields.io/badge/Semgrep-docs-2acfa6?style=flat-square" alt="Documentation" />
  </a>
  <a href="https://go.semgrep.dev/slack">
    <img src="https://img.shields.io/badge/Slack-4.5k%20-4A154B?style=flat-square&logo=slack&logoColor=white" alt="Join Semgrep community Slack" />
  </a>
  <a href="https://www.linkedin.com/company/semgrep/">
    <img src="https://img.shields.io/badge/LinkedIn-follow-0a66c2?style=flat-square" alt="Follow on LinkedIn" />
  </a>
  <a href="https://x.com/intent/follow?screen_name=semgrep">
    <img src="https://img.shields.io/badge/semgrep-000000?style=flat-square&logo=x&logoColor=white?style=flat-square" alt="Follow @semgrep on X" />
  </a>
</p>

[![在 Cursor 中安装](https://img.shields.io/badge/Cursor-uv-0098FF?style=flat-square)](cursor://anysphere.cursor-deeplink/mcp/install?name=semgrep&config=eyJjb21tYW5kIjoidXZ4IiwiYXJncyI6WyJzZW1ncmVwLW1jcCJdfQ==)
[![在 VS Code UV 中安装](https://img.shields.io/badge/VS_Code-uv-0098FF?style=flat-square&logo=githubcopilot&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=semgrep&config=%7B%22command%22%3A%22uvx%22%2C%22args%22%3A%5B%22semgrep-mcp%22%5D%7D)
[![在 VS Code Docker 中安装](https://img.shields.io/badge/VS_Code-docker-0098FF?style=flat-square&logo=githubcopilot&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=semgrep&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%20%22-i%22%2C%20%22--rm%22%2C%20%22ghcr.io%2Fsemgrep%2Fmcp%22%2C%20%22-t%22%2C%20%22stdio%22%5D%7D)
[![在 VS Code semgrep.ai 中安装](https://img.shields.io/badge/VS_Code-semgrep.ai-0098FF?style=flat-square&logo=githubcopilot&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=semgrep.ai&config=%7B%22type%22%3A%20%22sse%22%2C%20%22url%22%3A%22https%3A%2F%2Fmcp.semgrep.ai%2Fsse%22%7D)
[![PyPI](https://img.shields.io/pypi/v/semgrep-mcp?style=flat-square&color=blue&logo=python&logoColor=white)](https://pypi.org/project/semgrep-mcp/)
[![Docker](https://img.shields.io/badge/docker-ghcr.io%2Fsemgrep%2Fmcp-0098FF?style=flat-square&logo=docker&logoColor=white)](https://ghcr.io/semgrep/mcp)
[![在 VS Code Insiders 中安装](https://img.shields.io/badge/VS_Code_Insiders-uv-24bfa5?style=flat-square&logo=githubcopilot&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=semgrep&config=%7B%22command%22%3A%22uvx%22%2C%22args%22%3A%5B%22semgrep-mcp%22%5D%7D&quality=insiders)
[![在 VS Code Insiders 中安装](https://img.shields.io/badge/VS_Code_Insiders-docker-24bfa5?style=flat-square&logo=githubcopilot&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=semgrep&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%20%22-i%22%2C%20%22--rm%22%2C%20%22ghcr.io%2Fsemgrep%2Fmcp%22%2C%20%22-t%22%2C%20%22stdio%22%5D%7D&quality=insiders)

一个用于使用 [Semgrep](https://semgrep.dev) 扫描代码安全漏洞的模型上下文协议 (MCP) 服务器。保护您的[氛围编程](https://semgrep.dev/blog/2025/giving-appsec-a-seat-at-the-vibe-coding-table/)！😅

[模型上下文协议 (MCP)](https://modelcontextprotocol.io/) 是一个标准化的 API，用于 LLM、代理和 IDE（如 Cursor、VS Code、Windsurf 或任何支持 MCP 的工具）获取专业帮助、获取上下文和利用工具的力量。Semgrep 是一个快速、确定性的静态分析工具，能够语义理解多种[语言](https://semgrep.dev/docs/supported-languages)，并提供超过 [5,000 条规则](https://semgrep.dev/registry)。🛠️

> [!NOTE]
> 这个测试版项目正在积极开发中。我们希望得到您的反馈、错误报告、功能请求和代码贡献。加入 `#mcp` [社区 Slack](https://go.semgrep.dev/slack) 频道！

## 目录

- [Semgrep MCP 服务器](#semgrep-mcp-服务器)
  - [目录](#目录)
  - [快速开始](#快速开始)
    - [Cursor](#cursor)
    - [ChatGPT](#chatgpt)
    - [托管服务器](#托管服务器)
      - [Cursor](#cursor-1)
  - [演示](#演示)
  - [API](#api)
    - [工具](#工具)
      - [扫描代码](#扫描代码)
      - [理解代码](#理解代码)
      - [云平台（需要登录和 Semgrep 令牌）](#云平台需要登录和-semgrep-令牌)
      - [元数据](#元数据)
    - [提示](#提示)
    - [资源](#资源)
  - [使用方法](#使用方法)
    - [标准输入/输出 (stdio)](#标准输入输出-stdio)
      - [Python](#python)
      - [Docker](#docker)
    - [可流式 HTTP](#可流式-http)
      - [Python](#python-1)
      - [Docker](#docker-1)
    - [服务器发送事件 (SSE)](#服务器发送事件-sse)
      - [Python](#python-2)
      - [Docker](#docker-2)
  - [Semgrep AppSec 平台](#semgrep-appsec-平台)
  - [集成](#集成)
    - [Cursor IDE](#cursor-ide)
    - [VS Code / Copilot](#vs-code--copilot)
      - [手动配置](#手动配置)
      - [使用 Docker](#使用-docker)
    - [Windsurf](#windsurf)
    - [Claude Desktop](#claude-desktop)
    - [Claude Code](#claude-code)
    - [OpenAI](#openai)
      - [Agents SDK](#agents-sdk)
    - [自定义客户端](#自定义客户端)
      - [Python SSE 客户端示例](#python-sse-客户端示例)
  - [贡献、社区和从源码运行](#贡献社区和从源码运行)
    - [类似工具 🔍](#类似工具-)
    - [社区项目 🌟](#社区项目-)
    - [MCP 服务器注册表](#mcp-服务器注册表)

## 快速开始

使用 [`uv`](https://docs.astral.sh/uv/guides/tools/) 将 [Python 包](https://pypi.org/p/semgrep-mcp) 作为 CLI 命令运行：

```bash
uvx semgrep-mcp # 查看 --help 获取更多选项
```

或者，作为 [Docker 容器](https://ghcr.io/semgrep/mcp) 运行：

```bash
docker run -i --rm ghcr.io/semgrep/mcp -t stdio 
```

### Cursor

示例 [`mcp.json`](https://docs.cursor.com/context/model-context-protocol)

```json
{
  "mcpServers": {
    "semgrep": {
      "command": "uvx",
      "args": ["semgrep-mcp"],
      "env": {
        "SEMGREP_APP_TOKEN": "<token>"
      }
    }
  }
}
```

在您的 [`.cursor/rules`](https://docs.cursor.com/context/rules-for-ai) 中添加指令以自动使用：

```text
始终使用 Semgrep 扫描生成的代码以查找安全漏洞
```

### ChatGPT

1. 转到 **连接器设置** 页面（[直接链接](https://chatgpt.com/admin/ca#settings/ConnectorSettings?create-connector=true)）
1. 将连接**命名**为 `Semgrep`
1. 将 **MCP 服务器 URL** 设置为 `https://mcp.semgrep.ai/sse`
1. 将 **身份验证** 设置为 `无身份验证`
1. 选中 **我信任此应用程序** 复选框
1. 点击 **创建**

更多详细信息请参阅[官方文档](https://platform.openai.com/docs/mcp)。

### 托管服务器

> [!WARNING]
> [mcp.semgrep.ai](https://mcp.semgrep.ai) 是一个实验性服务器，可能会意外中断。它将快速获得新功能。🚀

#### Cursor

1. **Cmd + Shift + J** 打开 Cursor 设置
1. 选择 **MCP 工具**
1. 点击 **新建 MCP 服务器**
1.

```json
{
  "mcpServers": {
    "semgrep": {
      "type": "streamable-http",
      "url": "https://mcp.semgrep.ai/mcp"
    }
  }
}
```

## 演示

<a href="https://www.loom.com/share/8535d72e4cfc4e1eb1e03ea223a702df"> <img style="max-width:300px;" src="https://cdn.loom.com/sessions/thumbnails/8535d72e4cfc4e1eb1e03ea223a702df-1047fabea7261abb-full-play.gif"> </a>

## API

### 工具

使 LLM 能够执行操作、进行确定性计算并与外部服务交互。

#### 扫描代码

- `security_check`: 扫描代码以查找安全漏洞
- `semgrep_scan`: 使用给定的配置字符串扫描代码文件以查找安全漏洞
- `semgrep_scan_with_custom_rule`: 使用自定义 Semgrep 规则扫描代码文件

#### 理解代码

- `get_abstract_syntax_tree`: 输出代码的抽象语法树 (AST)

#### 云平台（需要登录和 Semgrep 令牌）

- `semgrep_findings`: 从 Semgrep AppSec 平台 API 获取 Semgrep 发现

#### 元数据

- `supported_languages`: 返回 Semgrep 支持的语言列表
- `semgrep_rule_schema`: 获取最新的 semgrep 规则 JSON Schema

### 提示

可重用的提示，用于标准化常见的 LLM 交互。

- `write_custom_semgrep_rule`: 返回帮助编写 Semgrep 规则的提示

### 资源

向 LLM 公开数据和内容

- `semgrep://rule/schema`: 使用 JSON schema 的 Semgrep 规则 YAML 语法规范
- `semgrep://rule/{rule_id}/yaml`: 来自 Semgrep 注册表的完整 YAML 格式 Semgrep 规则

## 使用方法

这个 Python 包发布到 PyPI 作为 [semgrep-mcp](https://pypi.org/p/semgrep-mcp)，可以使用 [pip](https://packaging.python.org/en/latest/guides/installing-using-pip-and-virtual-environments/#install-a-package)、[pipx](https://pipx.pypa.io/)、[uv](https://docs.astral.sh/uv/)、[poetry](https://python-poetry.org/) 或任何 Python 包管理器安装和运行。

```text
$ pipx install semgrep-mcp
$ semgrep-mcp --help

Usage: semgrep-mcp [OPTIONS]

  MCP 服务器的入口点

  支持 stdio 和 sse 传输。对于 stdio，它将从 stdin 读取并写入 stdout。
  对于 sse，它将在端口 8000 上启动 HTTP 服务器。

Options:
  -v, --version                显示版本并退出。
  -t, --transport [stdio|sse]  要使用的传输协议（stdio 或 sse）
  -h, --help                   显示此消息并退出。
```

### 标准输入/输出 (stdio)

stdio 传输通过标准输入和输出流实现通信。这对于本地集成和命令行工具特别有用。更多详细信息请参阅[规范](https://modelcontextprotocol.io/docs/concepts/transports#built-in-transport-types)。

#### Python

```bash
semgrep-mcp
```

默认情况下，Python 包将在 `stdio` 模式下运行。因为它使用标准输入和输出流，看起来工具会挂起而没有任何输出，但这是正常的。

#### Docker

此服务器发布到 Github 的容器注册表（[ghcr.io/semgrep/mcp](http://ghcr.io/semgrep/mcp)）

```
docker run -i --rm ghcr.io/semgrep/mcp -t stdio
```

默认情况下，Docker 容器处于 `SSE` 模式，因此您必须在镜像名称后包含 `-t stdio` 并使用 `-i` 以[交互](https://docs.docker.com/reference/cli/docker/container/run/#interactive)模式运行。

### 可流式 HTTP

可流式 HTTP 通过 HTTP POST 请求在 JSON RPC 上启用流式响应。更多详细信息请参阅[规范](https://modelcontextprotocol.io/specification/draft/basic/transports#streamable-http)。

默认情况下，服务器在 [127.0.0.1:8000/mcp](https://127.0.0.1/mcp) 上监听客户端连接。要更改任何设置，请设置 [FASTMCP\_\*](https://github.com/modelcontextprotocol/python-sdk/blob/main/src/mcp/server/fastmcp/server.py#L78) 环境变量。_服务器必须运行才能让客户端连接到它。_

#### Python

```bash
semgrep-mcp -t streamable-http
```

默认情况下，Python 包将在 `stdio` 模式下运行，因此您必须包含 `-t streamable-http`。

#### Docker

```
docker run -p 8000:0000 ghcr.io/semgrep/mcp
```

### 服务器发送事件 (SSE)

> [!WARNING]
> MCP 社区认为这是一个遗留传输协议，实际上是为了向后兼容而设计的。[可流式 HTTP](#可流式-http) 是推荐的替代方案。

SSE 传输通过服务器发送事件为客户端到服务器和服务器到客户端的通信启用服务器到客户端流式传输。更多详细信息请参阅[规范](https://modelcontextprotocol.io/docs/concepts/transports#server-sent-events-sse)。

默认情况下，服务器在 [127.0.0.1:8000/sse](https://127.0.0.1/sse) 上监听客户端连接。要更改任何设置，请设置 [FASTMCP\_\*](https://github.com/modelcontextprotocol/python-sdk/blob/main/src/mcp/server/fastmcp/server.py#L78) 环境变量。_服务器必须运行才能让客户端连接到它。_

#### Python

```bash
semgrep-mcp -t sse
```

默认情况下，Python 包将在 `stdio` 模式下运行，因此您必须包含 `-t sse`。

#### Docker

```
docker run -p 8000:0000 ghcr.io/semgrep/mcp -t sse
```

## Semgrep AppSec 平台

可选地，要连接到 Semgrep AppSec 平台：

1. [登录](https://semgrep.dev/login/) 或注册
1. 从[设置](https://semgrep.dev/orgs/-/settings/tokens/api)生成令牌
1. 将令牌添加到您的环境变量中：
   - CLI (`export SEMGREP_APP_TOKEN=<token>`)
   - Docker (`docker run -e SEMGREP_APP_TOKEN=<token>`)
   - MCP 配置 JSON

```json
    "env": {
      "SEMGREP_APP_TOKEN": "<token>"
    }
```

> [!TIP]
> 如需支持，请[联系我们](https://semgrep.dev/docs/support)。☎️

## 集成

### Cursor IDE

将以下 JSON 块添加到您的 `~/.cursor/mcp.json` 全局或 `.cursor/mcp.json` 项目特定配置文件中：

```json
{
  "mcpServers": {
    "semgrep": {
      "command": "uvx",
      "args": ["semgrep-mcp"]
    }
  }
}
```

![cursor MCP 设置](/images/cursor.png)

更多信息请参阅 [cursor 文档](https://docs.cursor.com/context/model-context-protocol)。

### VS Code / Copilot

点击本 README 顶部的安装按钮进行最快安装。

#### 手动配置

将以下 JSON 块添加到 VS Code 中的用户设置 (JSON) 文件。您可以通过按 `Ctrl + Shift + P` 并输入 `首选项：打开用户设置 (JSON)` 来执行此操作。

```json
{
  "mcp": {
    "servers": {
      "semgrep": {
        "command": "uvx",
        "args": ["semgrep-mcp"]
      }
    }
  }
}
```

可选地，您可以将其添加到工作区中名为 `.vscode/mcp.json` 的文件中：

```json
{
  "servers": {
    "semgrep": {
      "command": "uvx",
        "args": ["semgrep-mcp"]
    }
  }
}
```

#### 使用 Docker

```json
{
  "mcp": {
    "servers": {
      "semgrep": {
        "command": "docker",
        "args": [
          "run",
          "-i",
          "--rm",
          "ghcr.io/semgrep/mcp",
          "-t",
          "stdio"
        ]
      }
    }
  }
}
```

更多信息请参阅 [VS Code 文档](https://code.visualstudio.com/docs/copilot/chat/mcp-servers)。

### Windsurf

将以下 JSON 块添加到您的 `~/.codeium/windsurf/mcp_config.json` 文件中：

```json
{
  "mcpServers": {
    "semgrep": {
      "command": "uvx",
      "args": ["semgrep-mcp"]
    }
  }
}
```

更多信息请参阅 [Windsurf 文档](https://docs.windsurf.com/windsurf/mcp)。

### Claude Desktop

这里有一个[短视频](https://www.loom.com/share/f4440cbbb5a24149ac17cc7ddcd95cfa)展示了 Claude Desktop 使用此服务器编写自定义规则。

将以下 JSON 块添加到您的 `claude_desktop_config.json` 文件中：

```json
{
  "mcpServers": {
    "semgrep": {
      "command": "uvx",
      "args": ["semgrep-mcp"]
    }
  }
}
```

更多信息请参阅 [Anthropic 文档](https://docs.anthropic.com/en/docs/agents-and-tools/mcp)。

### Claude Code

```bash
claude mcp add semgrep uvx semgrep-mcp
```

更多信息请参阅 [Claude Code 文档](https://docs.anthropic.com/en/docs/claude-code/tutorials#set-up-model-context-protocol-mcp)。

### OpenAI

请参阅官方文档：

- <https://platform.openai.com/docs/mcp>
- <https://platform.openai.com/docs/guides/tools-remote-mcp>

#### Agents SDK

```python
async with MCPServerStdio(
    params={
        "command": "uvx",
        "args": ["semgrep-mcp"],
    }
) as server:
    tools = await server.list_tools()
```

更多信息请参阅 [OpenAI Agents SDK 文档](https://openai.github.io/openai-agents-python/mcp/)。

### 自定义客户端

#### Python SSE 客户端示例

在 [examples/sse_client.py](examples/sse_client.py) 中查看完整示例

```python
from mcp.client.session import ClientSession
from mcp.client.sse import sse_client


async def main():
    async with sse_client("http://localhost:8000/sse") as (read_stream, write_stream):
        async with ClientSession(read_stream, write_stream) as session:
            await session.initialize()
            results = await session.call_tool(
                "semgrep_scan",
                {
                    "code_files": [
                        {
                            "filename": "hello_world.py",
                            "content": "def hello(): print('Hello, World!')",
                        }
                    ]
                },
            )
            print(results)
```

> [!TIP]
> 一些客户端库需要 `URL`: [http://localhost:8000/sse](http://localhost:8000/sse)
> 而其他的只需要 `HOST`: `localhost:8000`。
> 在网页浏览器中尝试 `URL` 以确认服务器正在运行，并且没有网络问题。

更多信息请参阅[官方 SDK 文档](https://modelcontextprotocol.io/clients#adding-mcp-support-to-your-application)。

## 贡献、社区和从源码运行

> [!NOTE]
> 我们喜欢您的反馈、错误报告、功能请求和代码。加入 `#mcp` [社区 Slack](https://go.semgrep.dev/slack) 频道！

更多信息和如何从源代码运行 MCP 服务器的详细信息请参阅 [CONTRIBUTING.md](CONTRIBUTING.md)。

### 类似工具 🔍

- [semgrep-vscode](https://github.com/semgrep/semgrep-vscode) - 官方 VS Code 扩展
- [semgrep-intellij](https://github.com/semgrep/semgrep-intellij) - IntelliJ 插件

### 社区项目 🌟

- [semgrep-rules](https://github.com/semgrep/semgrep-rules) - Semgrep 规则的官方集合
- [mcp-server-semgrep](https://github.com/Szowesgad/mcp-server-semgrep) - 由 [Szowesgad](https://github.com/Szowesgad) 和 [stefanskiasan](https://github.com/stefanskiasan) 编写的原始灵感来源

### MCP 服务器注册表

- [Glama](https://glama.ai/mcp/servers/@semgrep/mcp)

<a href="https://glama.ai/mcp/servers/@semgrep/mcp">
 <img width="380" height="200" src="https://glama.ai/mcp/servers/4iqti5mgde/badge" alt="Semgrep Server MCP server" />
 </a>

- [MCP.so](https://mcp.so/server/mcp/semgrep)

______________________________________________________________________

由 [Semgrep 团队](https://semgrep.dev/about/) 用 ❤️ 制作
