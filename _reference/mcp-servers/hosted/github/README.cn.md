# GitHub MCP 服务器

GitHub MCP 服务器是一个[模型上下文协议 (MCP)](https://modelcontextprotocol.io/introduction) 服务器，提供与 GitHub API 的无缝集成，为开发者和工具提供高级自动化和交互功能。

### 使用场景

- 自动化 GitHub 工作流程和流程
- 从 GitHub 仓库中提取和分析数据
- 构建与 GitHub 生态系统交互的 AI 驱动工具和应用程序

---

## 远程 GitHub MCP 服务器

[![在 VS Code 中安装](https://img.shields.io/badge/VS_Code-Install_Server-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=github&config=%7B%22type%22%3A%20%22http%22%2C%22url%22%3A%20%22https%3A%2F%2Fapi.githubcopilot.com%2Fmcp%2F%22%7D) [![在 VS Code Insiders 中安装](https://img.shields.io/badge/VS_Code_Insiders-Install_Server-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=github&config=%7B%22type%22%3A%20%22http%22%2C%22url%22%3A%20%22https%3A%2F%2Fapi.githubcopilot.com%2Fmcp%2F%22%7D&quality=insiders)

远程 GitHub MCP 服务器由 GitHub 托管，提供最简单的启动和运行方法。如果您的 MCP 主机不支持远程 MCP 服务器，别担心！您可以使用[本地版本的 GitHub MCP 服务器](https://github.com/github/github-mcp-server?tab=readme-ov-file#local-github-mcp-server)。

## 前提条件

1. 支持最新 MCP 规范和远程服务器的 MCP 主机，例如 [VS Code](https://code.visualstudio.com/)。

## 安装

### 在 VS Code 中使用

快速安装请使用上面的一键安装按钮。完成该流程后，切换代理模式（位于 Copilot Chat 文本输入框旁边），服务器将启动。确保您使用的是 [VS Code 1.101](https://code.visualstudio.com/updates/v1_101) 或[更高版本](https://code.visualstudio.com/updates)以支持远程 MCP 和 OAuth。

或者，要手动配置 VS Code，请从下面的示例中选择适当的 JSON 块并将其添加到您的主机配置中：

<table>
<tr><th>使用 OAuth</th><th>使用 GitHub PAT</th></tr>
<tr><th align=left colspan=2>VS Code（1.101 或更高版本）</th></tr>
<tr valign=top>
<td>
  
```json
{
  "servers": {
    "github-remote": {
      "type": "http",
      "url": "https://api.githubcopilot.com/mcp/"
    }
  }
}
```

</td>
<td>

```json
{
  "servers": {
    "github-remote": {
      "type": "http",
      "url": "https://api.githubcopilot.com/mcp/",
      "headers": {
        "Authorization": "Bearer ${input:github_mcp_pat}"
      }
    }
  },
  "inputs": [
    {
      "type": "promptString",
      "id": "github_mcp_pat",
      "description": "GitHub Personal Access Token",
      "password": true
    }
  ]
}
```

</td>
</tr>
</table>

### 在其他 MCP 主机中使用

对于[兼容远程 MCP](docs/host-integration.md) 的 MCP 主机，请从下面的示例中选择适当的 JSON 块并将其添加到您的主机配置中：

<table>
<tr><th>使用 OAuth</th><th>使用 GitHub PAT</th></tr>
<tr valign=top>
<td>
  
```json
{
  "mcpServers": {
    "github-remote": {
      "url": "https://api.githubcopilot.com/mcp/"
    }
  }
}
```

</td>
<td>

```json
{
  "mcpServers": {
    "github-remote": {
      "url": "https://api.githubcopilot.com/mcp/",
      "authorization_token": "Bearer <your GitHub PAT>"
    }
  }
}
```

</td>
</tr>
</table>

> **注意：** 确切的配置格式可能因主机而异。请参考您的主机文档了解远程 MCP 服务器设置的正确语法和位置。

### 配置

请参阅[远程服务器文档](docs/remote-server.md)了解如何向远程 GitHub MCP 服务器传递其他配置设置。

---

## 本地 GitHub MCP 服务器

[![使用 Docker 在 VS Code 中安装](https://img.shields.io/badge/VS_Code-Install_Server-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=github&inputs=%5B%7B%22id%22%3A%22github_token%22%2C%22type%22%3A%22promptString%22%2C%22description%22%3A%22GitHub%20Personal%20Access%20Token%22%2C%22password%22%3Atrue%7D%5D&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%22-i%22%2C%22--rm%22%2C%22-e%22%2C%22GITHUB_PERSONAL_ACCESS_TOKEN%22%2C%22ghcr.io%2Fgithub%2Fgithub-mcp-server%22%5D%2C%22env%22%3A%7B%22GITHUB_PERSONAL_ACCESS_TOKEN%22%3A%22%24%7Binput%3Agithub_token%7D%22%7D%7D) [![使用 Docker 在 VS Code Insiders 中安装](https://img.shields.io/badge/VS_Code_Insiders-Install_Server-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=github&inputs=%5B%7B%22id%22%3A%22github_token%22%2C%22type%22%3A%22promptString%22%2C%22description%22%3A%22GitHub%20Personal%20Access%20Token%22%2C%22password%22%3Atrue%7D%5D&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%22-i%22%2C%22--rm%22%2C%22-e%22%2C%22GITHUB_PERSONAL_ACCESS_TOKEN%22%2C%22ghcr.io%2Fgithub%2Fgithub-mcp-server%22%5D%2C%22env%22%3A%7B%22GITHUB_PERSONAL_ACCESS_TOKEN%22%3A%22%24%7Binput%3Agithub_token%7D%22%7D%7D&quality=insiders)

## 前提条件

1. 要在容器中运行服务器，您需要安装 [Docker](https://www.docker.com/)。
2. 安装 Docker 后，您还需要确保 Docker 正在运行。镜像是公开的；如果拉取时出现错误，您可能有过期的令牌，需要执行 `docker logout ghcr.io`。
3. 最后，您需要[创建 GitHub 个人访问令牌](https://github.com/settings/personal-access-tokens/new)。
MCP 服务器可以使用许多 GitHub API，因此请启用您愿意授予 AI 工具的权限（要了解更多关于访问令牌的信息，请查看[文档](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens)）。

## 安装

### 在 VS Code 中使用

快速安装请使用一键安装按钮。完成该流程后，切换代理模式（位于 Copilot Chat 文本输入框旁边），服务器将启动。

### 在其他 MCP 主机中使用

将以下 JSON 块添加到您的 IDE MCP 设置中。

```json
{
  "mcp": {
    "inputs": [
      {
        "type": "promptString",
        "id": "github_token",
        "description": "GitHub Personal Access Token",
        "password": true
      }
    ],
    "servers": {
      "github": {
        "command": "docker",
        "args": [
          "run",
          "-i",
          "--rm",
          "-e",
          "GITHUB_PERSONAL_ACCESS_TOKEN",
          "ghcr.io/github/github-mcp-server"
        ],
        "env": {
          "GITHUB_PERSONAL_ACCESS_TOKEN": "${input:github_token}"
        }
      }
    }
  }
}
```

可选地，您可以在工作区中创建一个名为 `.vscode/mcp.json` 的文件，并添加类似的示例（即不包含 mcp 键）。这将允许您与他人共享配置。

```json
{
  "inputs": [
    {
      "type": "promptString",
      "id": "github_token",
      "description": "GitHub Personal Access Token",
      "password": true
    }
  ],
  "servers": {
    "github": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "-e",
        "GITHUB_PERSONAL_ACCESS_TOKEN",
        "ghcr.io/github/github-mcp-server"
      ],
      "env": {
        "GITHUB_PERSONAL_ACCESS_TOKEN": "${input:github_token}"
      }
    }
  }
}
```

更多关于在 VS Code 的[代理模式文档](https://code.visualstudio.com/docs/copilot/chat/mcp-servers)中使用 MCP 服务器工具的信息。

### 在 Claude Desktop 中使用

```json
{
  "mcpServers": {
    "github": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "-e",
        "GITHUB_PERSONAL_ACCESS_TOKEN",
        "ghcr.io/github/github-mcp-server"
      ],
      "env": {
        "GITHUB_PERSONAL_ACCESS_TOKEN": "<YOUR_TOKEN>"
      }
    }
  }
}
```

### 从源码构建

如果您没有 Docker，可以使用 `go build` 在 `cmd/github-mcp-server` 目录中构建二进制文件，并使用设置了 `GITHUB_PERSONAL_ACCESS_TOKEN` 环境变量的 `github-mcp-server stdio` 命令。要指定构建的输出位置，请使用 `-o` 标志。您应该配置服务器使用构建的可执行文件作为其 `command`。例如：

```JSON
{
  "mcp": {
    "servers": {
      "github": {
        "command": "/path/to/github-mcp-server",
        "args": ["stdio"],
        "env": {
          "GITHUB_PERSONAL_ACCESS_TOKEN": "<YOUR_TOKEN>"
        }
      }
    }
  }
}
```

## 工具配置

GitHub MCP 服务器支持通过 `--toolsets` 标志启用或禁用特定功能组。这允许您控制哪些 GitHub API 功能可用于您的 AI 工具。仅启用您需要的工具集可以帮助 LLM 进行工具选择并减少上下文大小。

_工具集不仅限于工具。相关的 MCP 资源和提示也会在适用的地方包含。_

### 可用工具集

以下工具集可用（默认全部开启）：

| 工具集                 | 描述                                                   |
| ----------------------- | ----------------------------------------------------- |
| `context`               | **强烈推荐**：提供当前用户和您正在操作的 GitHub 上下文信息的工具 |
| `code_security`         | 代码扫描警报和安全功能                                    |
| `issues`                | 问题相关工具（创建、读取、更新、评论）                       |
| `notifications`         | GitHub 通知相关工具                                      |
| `pull_requests`         | 拉取请求操作（创建、合并、审查）                             |
| `repos`                 | 仓库相关工具（文件操作、分支、提交）                         |
| `secret_protection`     | 密钥保护相关工具，如 GitHub 密钥扫描                       |
| `users`                 | 与 GitHub 用户相关的任何内容                              |
| `experiments`           | 实验性功能（不被认为是稳定的）                              |

#### 指定工具集

要指定您希望 LLM 可用的工具集，您可以通过两种方式传递允许列表：

1. **使用命令行参数**：

   ```bash
   github-mcp-server --toolsets repos,issues,pull_requests,code_security
   ```

2. **使用环境变量**：

   ```bash
   GITHUB_TOOLSETS="repos,issues,pull_requests,code_security" ./github-mcp-server
   ```

如果同时提供了两者，环境变量 `GITHUB_TOOLSETS` 优先于命令行参数。

### 在 Docker 中使用工具集

使用 Docker 时，您可以将工具集作为环境变量传递：

```bash
docker run -i --rm \
  -e GITHUB_PERSONAL_ACCESS_TOKEN=<your-token> \
  -e GITHUB_TOOLSETS="repos,issues,pull_requests,code_security,experiments" \
  ghcr.io/github/github-mcp-server
```

### "all" 工具集

特殊工具集 `all` 可以提供来启用所有可用工具集，无论任何其他配置：

```bash
./github-mcp-server --toolsets all
```

或使用环境变量：

```bash
GITHUB_TOOLSETS="all" ./github-mcp-server
```

## 动态工具发现

**注意**：此功能目前处于测试阶段，可能不在所有环境中可用。请测试并告诉我们您是否遇到任何问题。

您可以开启动态工具集发现，而不是从启用所有工具开始。动态工具集允许 MCP 主机响应用户提示来列出和启用工具集。这应该有助于避免模型因可用工具数量过多而感到困惑的情况。

### 使用动态工具发现

使用二进制文件时，您可以传递 `--dynamic-toolsets` 标志。

```bash
./github-mcp-server --dynamic-toolsets
```

使用 Docker 时，您可以将工具集作为环境变量传递：

```bash
docker run -i --rm \
  -e GITHUB_PERSONAL_ACCESS_TOKEN=<your-token> \
  -e GITHUB_DYNAMIC_TOOLSETS=1 \
  ghcr.io/github/github-mcp-server
```

## 只读模式

要在只读模式下运行服务器，您可以使用 `--read-only` 标志。这将只提供只读工具，防止对仓库、问题、拉取请求等进行任何修改。

```bash
./github-mcp-server --read-only
```

使用 Docker 时，您可以将只读模式作为环境变量传递：

```bash
docker run -i --rm \
  -e GITHUB_PERSONAL_ACCESS_TOKEN=<your-token> \
  -e GITHUB_READ_ONLY=1 \
  ghcr.io/github/github-mcp-server
```

## GitHub Enterprise Server 和带数据驻留的 Enterprise Cloud (ghe.com)

标志 `--gh-host` 和环境变量 `GITHUB_HOST` 可用于设置 GitHub Enterprise Server 或带数据驻留的 GitHub Enterprise Cloud 的主机名。

- 对于 GitHub Enterprise Server，请在主机名前加上 `https://` URI 方案，否则默认为 `http://`，而 GitHub Enterprise Server 不支持。
- 对于带数据驻留的 GitHub Enterprise Cloud，请使用 `https://YOURSUBDOMAIN.ghe.com` 作为主机名。

``` json
"github": {
    "command": "docker",
    "args": [
    "run",
    "-i",
    "--rm",
    "-e",
    "GITHUB_PERSONAL_ACCESS_TOKEN",
    "-e",
    "GITHUB_HOST",
    "ghcr.io/github/github-mcp-server"
    ],
    "env": {
        "GITHUB_PERSONAL_ACCESS_TOKEN": "${input:github_token}",
        "GITHUB_HOST": "https://<your GHES or ghe.com domain name>"
    }
}
```

## 国际化 / 覆盖描述

可以通过在二进制文件的同一目录中创建 `github-mcp-server-config.json` 文件来覆盖工具的描述。

该文件应包含一个 JSON 对象，以工具名称作为键，新描述作为值。例如：

```json
{
  "TOOL_ADD_ISSUE_COMMENT_DESCRIPTION": "替代描述",
  "TOOL_CREATE_BRANCH_DESCRIPTION": "在 GitHub 仓库中创建新分支"
}
```

您可以通过使用 `--export-translations` 标志运行二进制文件来创建当前翻译的导出。

此标志将保留您已做的任何翻译/覆盖，同时添加自上次导出以来添加到二进制文件中的任何新翻译。

```sh
./github-mcp-server --export-translations
cat github-mcp-server-config.json
```

您也可以使用环境变量来覆盖描述。环境变量名称与 JSON 文件中的键相同，前缀为 `GITHUB_MCP_` 并全部大写。

例如，要覆盖 `TOOL_ADD_ISSUE_COMMENT_DESCRIPTION` 工具，您可以设置以下环境变量：

```sh
export GITHUB_MCP_TOOL_ADD_ISSUE_COMMENT_DESCRIPTION="替代描述"
```

## 工具

### 用户

- **get_me** - 获取已认证用户的详细信息
  - 无需参数

### 问题

- **get_issue** - 获取仓库中问题的内容

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `issue_number`: 问题编号（数字，必需）

- **get_issue_comments** - 获取 GitHub 问题的评论

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `issue_number`: 问题编号（数字，必需）

- **create_issue** - 在 GitHub 仓库中创建新问题

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `title`: 问题标题（字符串，必需）
  - `body`: 问题正文内容（字符串，可选）
  - `assignees`: 分配给此问题的用户名（字符串数组，可选）
  - `labels`: 应用于此问题的标签（字符串数组，可选）

- **add_issue_comment** - 为问题添加评论

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `issue_number`: 问题编号（数字，必需）
  - `body`: 评论文本（字符串，必需）

- **list_issues** - 列出和过滤仓库问题

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `state`: 按状态过滤（'open'、'closed'、'all'）（字符串，可选）
  - `labels`: 按标签过滤（字符串数组，可选）
  - `sort`: 排序方式（'created'、'updated'、'comments'）（字符串，可选）
  - `direction`: 排序方向（'asc'、'desc'）（字符串，可选）
  - `since`: 按日期过滤（ISO 8601 时间戳）（字符串，可选）
  - `page`: 页码（数字，可选）
  - `perPage`: 每页结果数（数字，可选）

- **update_issue** - 更新 GitHub 仓库中的现有问题

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `issue_number`: 要更新的问题编号（数字，必需）
  - `title`: 新标题（字符串，可选）
  - `body`: 新描述（字符串，可选）
  - `state`: 新状态（'open' 或 'closed'）（字符串，可选）
  - `labels`: 新标签（字符串数组，可选）
  - `assignees`: 新分配者（字符串数组，可选）
  - `milestone`: 新里程碑编号（数字，可选）

- **search_issues** - 搜索问题和拉取请求
  - `query`: 搜索查询（字符串，必需）
  - `sort`: 排序字段（字符串，可选）
  - `order`: 排序顺序（字符串，可选）
  - `page`: 页码（数字，可选）
  - `perPage`: 每页结果数（数字，可选）

- **assign_copilot_to_issue** - 将 Copilot 分配给 GitHub 仓库中的特定问题

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `issueNumber`: 问题编号（数字，必需）
  - _注意_：此工具可以帮助创建包含源代码更改的拉取请求来解决问题。更多信息可以在 [GitHub Copilot 文档](https://docs.github.com/en/copilot/using-github-copilot/using-copilot-coding-agent-to-work-on-tasks/about-assigning-tasks-to-copilot)中找到

### 拉取请求

- **get_pull_request** - 获取特定拉取请求的详细信息

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `pullNumber`: 拉取请求编号（数字，必需）

- **list_pull_requests** - 列出和过滤仓库拉取请求

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `state`: PR 状态（字符串，可选）
  - `sort`: 排序字段（字符串，可选）
  - `direction`: 排序方向（字符串，可选）
  - `perPage`: 每页结果数（数字，可选）
  - `page`: 页码（数字，可选）

- **merge_pull_request** - 合并拉取请求

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `pullNumber`: 拉取请求编号（数字，必需）
  - `commit_title`: 合并提交的标题（字符串，可选）
  - `commit_message`: 合并提交的消息（字符串，可选）
  - `merge_method`: 合并方法（字符串，可选）

- **get_pull_request_files** - 获取拉取请求中更改的文件列表

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `pullNumber`: 拉取请求编号（数字，必需）

- **get_pull_request_status** - 获取拉取请求所有状态检查的综合状态

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `pullNumber`: 拉取请求编号（数字，必需）

- **update_pull_request_branch** - 使用基础分支的最新更改更新拉取请求分支

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `pullNumber`: 拉取请求编号（数字，必需）
  - `expectedHeadSha`: 拉取请求 HEAD 引用的预期 SHA（字符串，可选）

- **get_pull_request_comments** - 获取拉取请求的审查评论

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `pullNumber`: 拉取请求编号（数字，必需）

- **get_pull_request_reviews** - 获取拉取请求的审查

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `pullNumber`: 拉取请求编号（数字，必需）

- **get_pull_request_diff** - 获取拉取请求的差异

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `pullNumber`: 拉取请求编号（数字，必需）

- **create_pull_request_review** - 为拉取请求创建审查

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `pullNumber`: 拉取请求编号（数字，必需）
  - `body`: 审查评论文本（字符串，可选）
  - `event`: 审查操作（'APPROVE'、'REQUEST_CHANGES'、'COMMENT'）（字符串，必需）
  - `commitId`: 要审查的提交 SHA（字符串，可选）
  - `comments`: 在拉取请求更改上放置评论的行特定评论对象数组（数组，可选）
    - 对于内联评论：提供 `path`、`position`（或 `line`）和 `body`
    - 对于多行评论：提供 `path`、`start_line`、`line`、可选的 `side`/`start_side` 和 `body`

- **create_pending_pull_request_review** - 为拉取请求创建可稍后提交的待审查

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `pullNumber`: 拉取请求编号（数字，必需）
  - `commitID`: 要审查的提交 SHA（字符串，可选）

- **add_pull_request_review_comment_to_pending_review** - 向请求者的最新待审查拉取请求添加评论

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `pullNumber`: 拉取请求编号（数字，必需）
  - `path`: 需要评论的文件的相对路径（字符串，必需）
  - `body`: 审查评论的文本（字符串，必需）
  - `subjectType`: 评论针对的级别（字符串，必需）
    - 枚举："FILE"、"LINE"
  - `line`: 评论适用的拉取请求差异中 blob 的行（数字，可选）
  - `side`: 要评论的差异侧（字符串，可选）
    - 枚举："LEFT"、"RIGHT"
  - `startLine`: 对于多行评论，范围的第一行（数字，可选）
  - `startSide`: 对于多行评论，差异的起始侧（字符串，可选）
    - 枚举："LEFT"、"RIGHT"

- **submit_pending_pull_request_review** - 提交请求者的最新待审查拉取请求

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `pullNumber`: 拉取请求编号（数字，必需）
  - `event`: 要执行的事件（字符串，必需）
    - 枚举："APPROVE"、"REQUEST_CHANGES"、"COMMENT"
  - `body`: 审查评论的文本（字符串，可选）

- **delete_pending_pull_request_review** - 删除请求者的最新待审查拉取请求

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `pullNumber`: 拉取请求编号（数字，必需）

- **create_and_submit_pull_request_review** - 为拉取请求创建并提交不包含审查评论的审查

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `pullNumber`: 拉取请求编号（数字，必需）
  - `body`: 审查评论文本（字符串，必需）
  - `event`: 审查操作（'APPROVE'、'REQUEST_CHANGES'、'COMMENT'）（字符串，必需）
  - `commitID`: 要审查的提交 SHA（字符串，可选）

- **create_pull_request** - 创建新的拉取请求

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `title`: PR 标题（字符串，必需）
  - `body`: PR 描述（字符串，可选）
  - `head`: 包含更改的分支（字符串，必需）
  - `base`: 要合并到的分支（字符串，必需）
  - `draft`: 创建为草稿 PR（布尔值，可选）
  - `maintainer_can_modify`: 允许维护者编辑（布尔值，可选）

- **add_pull_request_review_comment** - 向拉取请求添加审查评论或回复现有评论

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `pull_number`: 拉取请求编号（数字，必需）
  - `body`: 审查评论的文本（字符串，必需）
  - `commit_id`: 要评论的提交的 SHA（字符串，除非使用 in_reply_to 否则必需）
  - `path`: 需要评论的文件的相对路径（字符串，除非使用 in_reply_to 否则必需）
  - `line`: 评论适用的拉取请求差异中 blob 的行（数字，可选）
  - `side`: 要评论的差异侧（LEFT 或 RIGHT）（字符串，可选）
  - `start_line`: 对于多行评论，范围的第一行（数字，可选）
  - `start_side`: 对于多行评论，差异的起始侧（LEFT 或 RIGHT）（字符串，可选）
  - `subject_type`: 评论针对的级别（line 或 file）（字符串，可选）
  - `in_reply_to`: 要回复的审查评论的 ID（数字，可选）。指定时，只需要 body，其他参数将被忽略。

- **update_pull_request** - 更新 GitHub 仓库中的现有拉取请求

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `pullNumber`: 要更新的拉取请求编号（数字，必需）
  - `title`: 新标题（字符串，可选）
  - `body`: 新描述（字符串，可选）
  - `state`: 新状态（'open' 或 'closed'）（字符串，可选）
  - `base`: 新基础分支名称（字符串，可选）
  - `maintainer_can_modify`: 允许维护者编辑（布尔值，可选）

- **request_copilot_review** - 为拉取请求请求 GitHub Copilot 审查（实验性；取决于 GitHub API 支持）

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `pullNumber`: 拉取请求编号（数字，必需）
  - _注意_：目前，此工具仅适用于 github.com

### 仓库

- **create_or_update_file** - 在仓库中创建或更新单个文件
  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `path`: 文件路径（字符串，必需）
  - `message`: 提交消息（字符串，必需）
  - `content`: 文件内容（字符串，必需）
  - `branch`: 分支名称（字符串，可选）
  - `sha`: 更新时的文件 SHA（字符串，可选）

- **delete_file** - 从 GitHub 仓库删除文件
  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `path`: 要删除的文件路径（字符串，必需）
  - `message`: 提交消息（字符串，必需）
  - `branch`: 要从中删除文件的分支（字符串，必需）

- **list_branches** - 列出 GitHub 仓库中的分支
  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `page`: 页码（数字，可选）
  - `perPage`: 每页结果数（数字，可选）

- **push_files** - 在单个提交中推送多个文件
  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `branch`: 要推送到的分支（字符串，必需）
  - `files`: 要推送的文件，每个都有路径和内容（数组，必需）
  - `message`: 提交消息（字符串，必需）

- **search_repositories** - 搜索 GitHub 仓库
  - `query`: 搜索查询（字符串，必需）
  - `sort`: 排序字段（字符串，可选）
  - `order`: 排序顺序（字符串，可选）
  - `page`: 页码（数字，可选）
  - `perPage`: 每页结果数（数字，可选）

- **create_repository** - 创建新的 GitHub 仓库
  - `name`: 仓库名称（字符串，必需）
  - `description`: 仓库描述（字符串，可选）
  - `private`: 仓库是否为私有（布尔值，可选）
  - `autoInit`: 自动使用 README 初始化（布尔值，可选）

- **get_file_contents** - 获取文件或目录的内容
  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `path`: 文件路径（字符串，必需）
  - `ref`: Git 引用（字符串，可选）

- **fork_repository** - 分叉仓库
  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `organization`: 目标组织名称（字符串，可选）

- **create_branch** - 创建新分支
  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `branch`: 新分支名称（字符串，必需）
  - `sha`: 创建分支的 SHA（字符串，必需）

- **list_commits** - 获取仓库分支的提交列表
  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `sha`: 分支名称、标签或提交 SHA（字符串，可选）
  - `path`: 仅包含此文件路径的提交（字符串，可选）
  - `page`: 页码（数字，可选）
  - `perPage`: 每页结果数（数字，可选）

- **get_commit** - 获取仓库提交的详细信息
  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `sha`: 提交 SHA、分支名称或标签名称（字符串，必需）
  - `page`: 页码，用于提交中的文件（数字，可选）
  - `perPage`: 每页结果数，用于提交中的文件（数字，可选）

- **get_tag** - 获取 GitHub 仓库中特定 git 标签的详细信息
  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `tag`: 标签名称（字符串，必需）

- **list_tags** - 列出 GitHub 仓库中的 git 标签
  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `page`: 页码（数字，可选）
  - `perPage`: 每页结果数（数字，可选）

- **search_code** - 在 GitHub 仓库中搜索代码
  - `query`: 搜索查询（字符串，必需）
  - `sort`: 排序字段（字符串，可选）
  - `order`: 排序顺序（字符串，可选）
  - `page`: 页码（数字，可选）
  - `perPage`: 每页结果数（数字，可选）

### 用户

- **search_users** - 搜索 GitHub 用户
  - `q`: 搜索查询（字符串，必需）
  - `sort`: 排序字段（字符串，可选）
  - `order`: 排序顺序（字符串，可选）
  - `page`: 页码（数字，可选）
  - `perPage`: 每页结果数（数字，可选）

### 代码扫描

- **get_code_scanning_alert** - 获取代码扫描警报

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `alertNumber`: 警报编号（数字，必需）

- **list_code_scanning_alerts** - 列出仓库的代码扫描警报
  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `ref`: Git 引用（字符串，可选）
  - `state`: 警报状态（字符串，可选）
  - `severity`: 警报严重性（字符串，可选）
  - `tool_name`: 用于代码扫描的工具名称（字符串，可选）

### 密钥扫描

- **get_secret_scanning_alert** - 获取密钥扫描警报

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `alertNumber`: 警报编号（数字，必需）

- **list_secret_scanning_alerts** - 列出仓库的密钥扫描警报
  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `state`: 警报状态（字符串，可选）
  - `secret_type`: 要在逗号分隔列表中过滤的密钥类型（字符串，可选）
  - `resolution`: 解决状态（字符串，可选）

### 通知

- **list_notifications** – 列出 GitHub 用户的通知
  - `filter`: 应用于响应的过滤器（`default`、`include_read_notifications`、`only_participating`）
  - `since`: 仅显示在给定时间之后更新的通知（ISO 8601 格式）
  - `before`: 仅显示在给定时间之前更新的通知（ISO 8601 格式）
  - `owner`: 可选的仓库所有者（字符串）
  - `repo`: 可选的仓库名称（字符串）
  - `page`: 页码（数字，可选）
  - `perPage`: 每页结果数（数字，可选）

- **get_notification_details** – 获取特定 GitHub 通知的详细信息
  - `notificationID`: 通知的 ID（字符串，必需）

- **dismiss_notification** – 通过将通知标记为已读或完成来关闭通知
  - `threadID`: 通知线程的 ID（字符串，必需）
  - `state`: 通知的新状态（`read` 或 `done`）

- **mark_all_notifications_read** – 将所有通知标记为已读
  - `lastReadAt`: 描述上次检查通知的时间点（可选，RFC3339/ISO8601 字符串，默认：现在）
  - `owner`: 可选的仓库所有者（字符串）
  - `repo`: 可选的仓库名称（字符串）

- **manage_notification_subscription** – 管理通知线程的通知订阅（忽略、关注或删除）
  - `notificationID`: 通知线程的 ID（字符串，必需）
  - `action`: 要执行的操作：`ignore`、`watch` 或 `delete`（字符串，必需）

- **manage_repository_notification_subscription** – 管理仓库通知订阅（忽略、关注或删除）
  - `owner`: 仓库的账户所有者（字符串，必需）
  - `repo`: 仓库的名称（字符串，必需）
  - `action`: 要执行的操作：`ignore`、`watch` 或 `delete`（字符串，必需）

## 资源

### 仓库内容

- **获取仓库内容**
  检索仓库在特定路径的内容。

  - **模板**: `repo://{owner}/{repo}/contents{/path*}`
  - **参数**:
    - `owner`: 仓库所有者（字符串，必需）
    - `repo`: 仓库名称（字符串，必需）
    - `path`: 文件或目录路径（字符串，可选）

- **获取特定分支的仓库内容**
  检索给定分支在特定路径的仓库内容。

  - **模板**: `repo://{owner}/{repo}/refs/heads/{branch}/contents{/path*}`
  - **参数**:
    - `owner`: 仓库所有者（字符串，必需）
    - `repo`: 仓库名称（字符串，必需）
    - `branch`: 分支名称（字符串，必需）
    - `path`: 文件或目录路径（字符串，可选）

- **获取特定提交的仓库内容**
  检索给定提交在特定路径的仓库内容。

  - **模板**: `repo://{owner}/{repo}/sha/{sha}/contents{/path*}`
  - **参数**:
    - `owner`: 仓库所有者（字符串，必需）
    - `repo`: 仓库名称（字符串，必需）
    - `sha`: 提交 SHA（字符串，必需）
    - `path`: 文件或目录路径（字符串，可选）

- **获取特定标签的仓库内容**
  检索给定标签在特定路径的仓库内容。

  - **模板**: `repo://{owner}/{repo}/refs/tags/{tag}/contents{/path*}`
  - **参数**:
    - `owner`: 仓库所有者（字符串，必需）
    - `repo`: 仓库名称（字符串，必需）
    - `tag`: 标签名称（字符串，必需）
    - `path`: 文件或目录路径（字符串，可选）

- **获取特定拉取请求的仓库内容**
  检索给定拉取请求在特定路径的仓库内容。

  - **模板**: `repo://{owner}/{repo}/refs/pull/{prNumber}/head/contents{/path*}`
  - **参数**:
    - `owner`: 仓库所有者（字符串，必需）
    - `repo`: 仓库名称（字符串，必需）
    - `prNumber`: 拉取请求编号（字符串，必需）
    - `path`: 文件或目录路径（字符串，可选）

## 库使用

此模块导出的 Go API 目前应被视为不稳定，可能会有破坏性更改。将来，我们可能会提供稳定性；如果有用例需要这种稳定性，请提交问题。
