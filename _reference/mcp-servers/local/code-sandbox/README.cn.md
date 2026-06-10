# 代码沙盒 MCP 🐳

[![smithery badge](https://smithery.ai/badge/@Automata-Labs-team/code-sandbox-mcp)](https://smithery.ai/server/@Automata-Labs-team/code-sandbox-mcp)

一个在 Docker 容器内执行代码的安全沙盒环境。此 MCP 服务器为 AI 应用程序提供了一个安全隔离的代码运行环境，通过容器化技术确保安全性。

## 🌟 特性

- **灵活的容器管理**：创建和管理用于代码执行的隔离 Docker 容器
- **自定义环境支持**：使用任何 Docker 镜像作为执行环境
- **文件操作**：在主机和容器之间轻松传输文件和目录
- **命令执行**：在容器化环境中运行任何 shell 命令
- **实时日志**：实时流式传输容器日志和命令输出
- **自动更新**：内置更新检查和自动二进制文件更新
- **多平台支持**：支持 Linux、macOS 和 Windows

## 🚀 安装

### 前置要求

- 已安装并运行 Docker
  - [Linux 安装 Docker](https://docs.docker.com/engine/install/)
  - [macOS 安装 Docker Desktop](https://docs.docker.com/desktop/install/mac/)
  - [Windows 安装 Docker Desktop](https://docs.docker.com/desktop/install/windows-install/)

### 快速安装

#### Linux、MacOS

```bash
curl -fsSL https://raw.githubusercontent.com/Automata-Labs-team/code-sandbox-mcp/main/install.sh | bash
```

#### Windows

```powershell
# 在 PowerShell 中运行
irm https://raw.githubusercontent.com/Automata-Labs-team/code-sandbox-mcp/main/install.ps1 | iex
```

安装程序将：

1. 检查 Docker 安装
2. 下载适合您系统的二进制文件
3. 创建必要的配置文件

### 手动安装

1. 从[发布页面](https://github.com/Automata-Labs-team/code-sandbox-mcp/releases)下载适合您平台的最新版本
2. 将二进制文件放在 PATH 目录中
3. 使其可执行（仅限类 Unix 系统）：

   ```bash
   chmod +x code-sandbox-mcp
   ```

## 🛠️ 可用工具

#### `sandbox_initialize`

初始化用于代码执行的新计算环境。
基于指定的 Docker 镜像创建容器。

**参数：**

- `image`（字符串，可选）：用作基础环境的 Docker 镜像
  - 默认值：'python:3.12-slim-bookworm'

**返回：**

- `container_id`，可与其他工具一起使用来与此环境交互

#### `copy_project`

将目录复制到沙盒文件系统。

**参数：**

- `container_id`（字符串，必需）：初始化调用返回的容器 ID
- `local_src_dir`（字符串，必需）：本地文件系统中目录的路径
- `dest_dir`（字符串，可选）：在沙盒环境中保存源目录的路径

#### `write_file`

将文件写入沙盒文件系统。

**参数：**

- `container_id`（字符串，必需）：初始化调用返回的容器 ID
- `file_name`（字符串，必需）：要创建的文件名
- `file_contents`（字符串，必需）：要写入文件的内容
- `dest_dir`（字符串，可选）：创建文件的目录（默认：${WORKDIR}）

#### `sandbox_exec`

在沙盒环境中执行命令。

**参数：**

- `container_id`（字符串，必需）：初始化调用返回的容器 ID
- `commands`（数组，必需）：在沙盒环境中运行的命令列表
  - 示例：["apt-get update", "pip install numpy", "python script.py"]

#### `copy_file`

将单个文件复制到沙盒文件系统。

**参数：**

- `container_id`（字符串，必需）：初始化调用返回的容器 ID
- `local_src_file`（字符串，必需）：本地文件系统中文件的路径
- `dest_path`（字符串，可选）：在沙盒环境中保存文件的路径

#### `sandbox_stop`

停止并移除正在运行的容器沙盒。

**参数：**

- `container_id`（字符串，必需）：要停止和移除的容器 ID

**描述：**
优雅地停止指定容器（超时时间 10 秒）并移除它及其卷。

#### 容器日志资源

提供访问容器日志的动态资源。

**资源路径：** `containers://{id}/logs`  
**MIME 类型：** `text/plain`  
**描述：** 将指定容器的所有容器日志作为单个文本资源返回。

## 🔐 安全特性

- 使用 Docker 容器的隔离执行环境
- 通过 Docker 容器约束进行资源限制
- 分离的标准输出和标准错误流

## 🔧 配置

### Claude Desktop

安装程序会自动创建配置文件。如果您需要手动配置：

#### Linux

```json
// ~/.config/Claude/claude_desktop_config.json
{
    "mcpServers": {
        "code-sandbox-mcp": {
            "command": "/path/to/code-sandbox-mcp",
            "args": [],
            "env": {}
        }
    }
}
```

#### macOS

```json
// ~/Library/Application Support/Claude/claude_desktop_config.json
{
    "mcpServers": {
        "code-sandbox-mcp": {
            "command": "/path/to/code-sandbox-mcp",
            "args": [],
            "env": {}
        }
    }
}
```

#### Windows

```json
// %APPDATA%\Claude\claude_desktop_config.json
{
    "mcpServers": {
        "code-sandbox-mcp": {
            "command": "C:\\path\\to\\code-sandbox-mcp.exe",
            "args": [],
            "env": {}
        }
    }
}
```

### 其他 AI 应用程序

对于支持 MCP 服务器的其他 AI 应用程序，请将它们配置为使用 `code-sandbox-mcp` 二进制文件作为代码执行后端。
