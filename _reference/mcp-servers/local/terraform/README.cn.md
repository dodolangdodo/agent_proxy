# <img src="public/images/Terraform-LogoMark_onDark.svg" width="30" align="left" style="margin-right: 12px;"/> Terraform MCP 服务器

Terraform MCP 服务器是一个[模型上下文协议 (MCP)](https://modelcontextprotocol.io/introduction) 服务器，提供与 Terraform Registry APIs 的无缝集成，为基础设施即代码 (IaC) 开发启用高级自动化和交互功能。

## 功能特性

- **双重传输支持**：同时支持 Stdio 和 StreamableHTTP 传输
- **Terraform 提供者发现**：查询和探索 Terraform 提供者及其文档
- **模块搜索与分析**：搜索并检索 Terraform 模块的详细信息
- **Registry 集成**：直接集成 Terraform Registry APIs
- **容器就绪**：支持 Docker 便于部署

> **注意：** MCP 服务器提供的输出和建议是动态生成的，可能会根据查询、模型和连接的 MCP 服务器而有所不同。用户应该**彻底审查所有输出/建议**，确保它们符合组织的**安全最佳实践**、**成本效率目标**和**合规要求**，然后再实施。

## 前置条件

1. 要在容器中运行服务器，您需要安装 [Docker](https://www.docker.com/)。
2. 安装 Docker 后，您需要确保 Docker 正在运行。

## 传输支持

Terraform MCP 服务器支持多种传输协议：

### 1. Stdio 传输（默认）

使用 JSON-RPC 消息的标准输入/输出通信。适用于本地开发和与 MCP 客户端的直接集成。

### 2. StreamableHTTP 传输

现代的基于 HTTP 的传输，支持直接 HTTP 请求和服务器发送事件 (SSE) 流。这是远程/分布式设置的推荐传输方式。

**功能特性：**

- **端点**：`http://{hostname}:8080/mcp`
- **健康检查**：`http://{hostname}:8080/health`
- **环境配置**：设置 `MODE=http` 或 `PORT=8080` 来启用

**环境变量：**

| 变量 | 描述 | 默认值 |
|----------|-------------|---------|
| `MODE` | 设置为 `http` 启用 HTTP 传输 | `stdio` |
| `PORT` | HTTP 服务器端口 | `8080` |

## 命令行选项

```bash
# Stdio 模式
terraform-mcp-server stdio [--log-file /path/to/log]

# HTTP 模式
terraform-mcp-server http [--port 8080] [--host 0.0.0.0] [--log-file /path/to/log]
```

## 安装

### 在 VS Code 中使用

将以下 JSON 块添加到 VS Code 的用户设置 (JSON) 文件中。您可以按 `Ctrl + Shift + P` 并输入 `Preferences: Open User Settings (JSON)` 来执行此操作。

更多关于在 VS Code 的[代理模式文档](https://code.visualstudio.com/docs/copilot/chat/mcp-servers)中使用 MCP 服务器工具的信息。

```json
{
  "mcp": {
    "servers": {
      "terraform": {
        "command": "docker",
        "args": [
          "run",
          "-i",
          "--rm",
          "hashicorp/terraform-mcp-server"
        ]
      }
    }
  }
}
```

可选地，您可以将类似的示例（即不包含 mcp 键）添加到工作区中名为 `.vscode/mcp.json` 的文件中。这将允许您与他人共享配置。

```json
{
  "servers": {
    "terraform": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "hashicorp/terraform-mcp-server"
      ]
    }
  }
}
```

### 在 Claude Desktop / Amazon Q Developer / Amazon Q CLI 中使用

更多关于在 Claude Desktop [用户文档](https://modelcontextprotocol.io/quickstart/user)中使用 MCP 服务器工具的信息。
从[文档](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/qdev-mcp.html)了解更多关于在 Amazon Q 中使用 MCP 服务器的信息。

```json
{
  "mcpServers": {
    "terraform": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "hashicorp/terraform-mcp-server"
      ]
    }
  }
}
```

## 工具配置

### 可用工具集

以下工具集可用：

| 工具集     | 工具                   | 描述                                                                                                                                                                                                                                                    |
|-------------|------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `providers` | `resolveProviderDocID` | 查询 Terraform Registry 以查找并列出特定提供者的可用文档，使用指定的 `serviceSlug`。返回带有标题和类别的提供者文档 ID 列表，用于资源、数据源、函数或指南。 |
| `providers` | `getProviderDocs`      | 使用从 `resolveProviderDocID` 工具获得的文档 ID 获取特定提供者资源、数据源或函数的完整文档内容。返回 markdown 格式的原始文档。                                     |
| `modules`   | `searchModules`        | 基于指定的 `moduleQuery` 在 Terraform Registry 中搜索模块，支持分页。返回模块 ID 列表，包括其名称、描述、下载次数、验证状态和发布日期                                             |
| `modules`   | `moduleDetails`        | 使用从 `searchModules` 工具获得的模块 ID 检索模块的详细文档，包括输入、输出、配置、子模块和示例。                                                                                     |
| `policies`  | `searchPolicies`       | 查询 Terraform Registry 以查找并列出基于提供的查询 `policyQuery` 的适当 Sentinel 策略。返回匹配策略列表，包含 terraformPolicyID 及其名称、标题和下载次数。                             |
| `policies`  | `policyDetails`        | 使用从 `searchPolicies` 工具获得的 terraformPolicyID 检索策略集的详细文档，包括策略 readme 和实现详情。                                                                                        |

### 从源码安装

使用最新发布版本：

```console
go install github.com/hashicorp/terraform-mcp-server/cmd/terraform-mcp-server@latest
```

使用主分支：

```console
go install github.com/hashicorp/terraform-mcp-server/cmd/terraform-mcp-server@main
```

```json
{
  "mcp": {
    "servers": {
      "terraform": {
        "command": "/path/to/terraform-mcp-server",
        "args": ["stdio"]
      }
    }
  }
}
```

## 本地构建 Docker 镜像

在使用服务器之前，您需要本地构建 Docker 镜像：

1. 克隆仓库：

```bash
git clone https://github.com/hashicorp/terraform-mcp-server.git
cd terraform-mcp-server
```

2. 构建 Docker 镜像：

```bash
make docker-build
```

3. 这将创建一个本地 Docker 镜像，您可以在以下配置中使用。

```bash
# 以 stdio 模式运行
docker run -i --rm terraform-mcp-server:dev

# 以 http 模式运行
docker run -p 8080:8080 --rm -e MODE=http terraform-mcp-server:dev
```

4. （可选）在 http 模式下测试连接
  
```bash
# 测试连接
curl http://localhost:8080/health
```

5. 您可以在 AI 助手中按如下方式使用它：

```json
{
  "mcpServers": {
    "terraform": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "terraform-mcp-server:dev"
      ]
    }
  }
}
```

## 开发

### 可用的 Make 命令

| 命令 | 描述 |
|---------|-------------|
| `make build` | 构建二进制文件 |
| `make test` | 运行所有测试 |
| `make test-e2e` | 运行端到端测试 |
| `make docker-build` | 构建 Docker 镜像 |
| `make run-http` | 本地运行 HTTP 服务器 |
| `make docker-run-http` | 在 Docker 中运行 HTTP 服务器 |
| `make test-http` | 测试 HTTP 健康端点 |
| `make clean` | 删除构建产物 |
| `make help` | 显示所有可用命令 |

## 贡献

1. Fork 仓库
2. 创建您的功能分支
3. 进行更改
4. 运行测试
5. 提交拉取请求

## 安全

对于安全问题，请联系 <security@hashicorp.com> 或遵循我们的[安全政策](https://www.hashicorp.com/en/trust/security/vulnerability-management)。

## 支持

对于错误报告和功能请求，请在 GitHub 上开启 issue。

对于一般问题和讨论，请开启 GitHub Discussion。
