# mcp-server-git: Git MCP 服务器

> <https://github.com/modelcontextprotocol/servers/tree/main/src/git>

## 概述

一个用于 Git 仓库交互和自动化的模型上下文协议服务器。该服务器提供工具，通过大型语言模型来读取、搜索和操作 Git 仓库。

请注意，mcp-server-git 目前处于早期开发阶段。随着我们继续开发和改进服务器，功能和可用工具可能会发生变化和扩展。

### 工具

1. `git_status`
   - 显示工作树状态
   - 输入：
     - `repo_path` (字符串)：Git 仓库路径
   - 返回：工作目录的当前状态文本输出

2. `git_diff_unstaged`
   - 显示工作目录中尚未暂存的更改
   - 输入：
     - `repo_path` (字符串)：Git 仓库路径
   - 返回：未暂存更改的差异输出

3. `git_diff_staged`
   - 显示已暂存等待提交的更改
   - 输入：
     - `repo_path` (字符串)：Git 仓库路径
   - 返回：已暂存更改的差异输出

4. `git_diff`
   - 显示分支或提交之间的差异
   - 输入：
     - `repo_path` (字符串)：Git 仓库路径
     - `target` (字符串)：要比较的目标分支或提交
   - 返回：当前状态与目标的差异输出

5. `git_commit`
   - 将更改记录到仓库
   - 输入：
     - `repo_path` (字符串)：Git 仓库路径
     - `message` (字符串)：提交消息
   - 返回：带有新提交哈希的确认信息

6. `git_add`
   - 将文件内容添加到暂存区
   - 输入：
     - `repo_path` (字符串)：Git 仓库路径
     - `files` (字符串数组)：要暂存的文件路径数组
   - 返回：已暂存文件的确认信息

7. `git_reset`
   - 取消暂存所有已暂存的更改
   - 输入：
     - `repo_path` (字符串)：Git 仓库路径
   - 返回：重置操作的确认信息

8. `git_log`
   - 显示提交日志
   - 输入：
     - `repo_path` (字符串)：Git 仓库路径
     - `max_count` (数字，可选)：要显示的最大提交数量（默认：10）
   - 返回：包含哈希、作者、日期和消息的提交条目数组

9. `git_create_branch`
   - 创建新分支
   - 输入：
     - `repo_path` (字符串)：Git 仓库路径
     - `branch_name` (字符串)：新分支名称
     - `start_point` (字符串，可选)：新分支的起始点
   - 返回：分支创建的确认信息

10. `git_checkout`
    - 切换分支
    - 输入：
      - `repo_path` (字符串)：Git 仓库路径
      - `branch_name` (字符串)：要切换到的分支名称
    - 返回：分支切换的确认信息

11. `git_show`
    - 显示提交的内容
    - 输入：
      - `repo_path` (字符串)：Git 仓库路径
      - `revision` (字符串)：要显示的修订版本（提交哈希、分支名称、标签）
    - 返回：指定提交的内容

12. `git_init`
    - 初始化 Git 仓库
    - 输入：
      - `repo_path` (字符串)：要初始化 git 仓库的目录路径
    - 返回：仓库初始化的确认信息

## 安装

### 使用 uv（推荐）

使用 [`uv`](https://docs.astral.sh/uv/) 时不需要特定的安装。我们将使用 [`uvx`](https://docs.astral.sh/uv/guides/tools/) 直接运行 *mcp-server-git*。

### 使用 PIP

或者您可以通过 pip 安装 `mcp-server-git`：

```
pip install mcp-server-git
```

安装后，您可以使用以下命令作为脚本运行：

```
python -m mcp_server_git
```

## 配置

### 与 Claude Desktop 一起使用

将以下内容添加到您的 `claude_desktop_config.json`：

<details>
<summary>使用 uvx</summary>

```json
"mcpServers": {
  "git": {
    "command": "uvx",
    "args": ["mcp-server-git", "--repository", "path/to/git/repo"]
  }
}
```

</details>

<details>
<summary>使用 docker</summary>

- 注意：将 '/Users/username' 替换为您希望此工具可访问的路径

```json
"mcpServers": {
  "git": {
    "command": "docker",
    "args": ["run", "--rm", "-i", "--mount", "type=bind,src=/Users/username,dst=/Users/username", "mcp/git"]
  }
}
```

</details>

<details>
<summary>使用 pip 安装</summary>

```json
"mcpServers": {
  "git": {
    "command": "python",
    "args": ["-m", "mcp_server_git", "--repository", "path/to/git/repo"]
  }
}
```

</details>

### 与 VS Code 一起使用

要快速安装，请使用下面的一键安装按钮...

[![在 VS Code 中使用 UV 安装](https://img.shields.io/badge/VS_Code-UV-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=git&config=%7B%22command%22%3A%22uvx%22%2C%22args%22%3A%5B%22mcp-server-git%22%5D%7D) [![在 VS Code Insiders 中使用 UV 安装](https://img.shields.io/badge/VS_Code_Insiders-UV-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=git&config=%7B%22command%22%3A%22uvx%22%2C%22args%22%3A%5B%22mcp-server-git%22%5D%7D&quality=insiders)

[![在 VS Code 中使用 Docker 安装](https://img.shields.io/badge/VS_Code-Docker-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=git&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%22--rm%22%2C%22-i%22%2C%22--mount%22%2C%22type%3Dbind%2Csrc%3D%24%7BworkspaceFolder%7D%2Cdst%3D%2Fworkspace%22%2C%22mcp%2Fgit%22%5D%7D) [![在 VS Code Insiders 中使用 Docker 安装](https://img.shields.io/badge/VS_Code_Insiders-Docker-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=git&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%22--rm%22%2C%22-i%22%2C%22--mount%22%2C%22type%3Dbind%2Csrc%3D%24%7BworkspaceFolder%7D%2Cdst%3D%2Fworkspace%22%2C%22mcp%2Fgit%22%5D%7D&quality=insiders)

要手动安装，请将以下 JSON 块添加到 VS Code 中的用户设置（JSON）文件。您可以按 `Ctrl + Shift + P` 并输入 `Preferences: Open Settings (JSON)` 来执行此操作。

或者，您可以将其添加到工作区中名为 `.vscode/mcp.json` 的文件中。这将允许您与他人共享配置。

> 注意在 `.vscode/mcp.json` 文件中不需要 `mcp` 键。

```json
{
  "mcp": {
    "servers": {
      "git": {
        "command": "uvx",
        "args": ["mcp-server-git"]
      }
    }
  }
}
```

Docker 安装：

```json
{
  "mcp": {
    "servers": {
      "git": {
        "command": "docker",
        "args": [
          "run",
          "--rm",
          "-i",
          "--mount", "type=bind,src=${workspaceFolder},dst=/workspace",
          "mcp/git"
        ]
      }
    }
  }
}
```

### 与 [Zed](https://github.com/zed-industries/zed) 一起使用

添加到您的 Zed settings.json：

<details>
<summary>使用 uvx</summary>

```json
"context_servers": [
  "mcp-server-git": {
    "command": {
      "path": "uvx",
      "args": ["mcp-server-git"]
    }
  }
],
```

</details>

<details>
<summary>使用 pip 安装</summary>

```json
"context_servers": {
  "mcp-server-git": {
    "command": {
      "path": "python",
      "args": ["-m", "mcp_server_git"]
    }
  }
},
```

</details>

## 调试

您可以使用 MCP 检查器来调试服务器。对于 uvx 安装：

```
npx @modelcontextprotocol/inspector uvx mcp-server-git
```

或者如果您在特定目录中安装了包或正在开发它：

```
cd path/to/servers/src/git
npx @modelcontextprotocol/inspector uv run mcp-server-git
```

运行 `tail -n 20 -f ~/Library/Logs/Claude/mcp*.log` 将显示服务器的日志，可能有助于您调试任何问题。

## 开发

如果您正在进行本地开发，有两种方法可以测试您的更改：

1. 运行 MCP 检查器来测试您的更改。运行说明请参见[调试](#调试)。

2. 使用 Claude 桌面应用程序进行测试。将以下内容添加到您的 `claude_desktop_config.json`：

### Docker

```json
{
  "mcpServers": {
    "git": {
      "command": "docker",
      "args": [
        "run",
        "--rm",
        "-i",
        "--mount", "type=bind,src=/Users/username/Desktop,dst=/projects/Desktop",
        "--mount", "type=bind,src=/path/to/other/allowed/dir,dst=/projects/other/allowed/dir,ro",
        "--mount", "type=bind,src=/path/to/file.txt,dst=/projects/path/to/file.txt",
        "mcp/git"
      ]
    }
  }
}
```

### UVX

```json
{
"mcpServers": {
  "git": {
    "command": "uv",
    "args": [ 
      "--directory",
      "/<path to mcp-servers>/mcp-servers/src/git",
      "run",
      "mcp-server-git"
    ]
  }
}
```

## 构建

Docker 构建：

```bash
cd src/git
docker build -t mcp/git .
```

## 许可证

此 MCP 服务器根据 MIT 许可证授权。这意味着您可以自由使用、修改和分发软件，但需遵守 MIT 许可证的条款和条件。有关更多详细信息，请参阅项目仓库中的 LICENSE 文件。
