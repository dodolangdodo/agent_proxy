# 文件系统 MCP 服务器

> <https://github.com/modelcontextprotocol/servers/tree/main/src/filesystem>

实现模型上下文协议 (MCP) 的 Node.js 服务器，用于文件系统操作。

## 功能特性

- 读取/写入文件
- 创建/列出/删除目录
- 移动文件/目录
- 搜索文件
- 获取文件元数据

**注意**：服务器只允许在通过 `args` 指定的目录内进行操作。

## API

### 资源

- `file://system`：文件系统操作接口

### 工具

- **read_file**
  - 读取文件的完整内容
  - 输入：`path`（字符串）
  - 使用 UTF-8 编码读取完整文件内容

- **read_multiple_files**
  - 同时读取多个文件
  - 输入：`paths`（字符串数组）
  - 读取失败不会停止整个操作

- **write_file**
  - 创建新文件或覆盖现有文件（使用时请谨慎）
  - 输入：
    - `path`（字符串）：文件位置
    - `content`（字符串）：文件内容

- **edit_file**
  - 使用高级模式匹配和格式化进行选择性编辑
  - 功能特性：
    - 基于行和多行内容匹配
    - 空白符规范化并保持缩进
    - 多个同时编辑并正确定位
    - 缩进样式检测和保持
    - Git 风格的差异输出和上下文
    - 预览更改的试运行模式
  - 输入：
    - `path`（字符串）：要编辑的文件
    - `edits`（数组）：编辑操作列表
      - `oldText`（字符串）：要搜索的文本（可以是子字符串）
      - `newText`（字符串）：要替换的文本
    - `dryRun`（布尔值）：预览更改而不应用（默认：false）
  - 试运行时返回详细的差异和匹配信息，否则应用更改
  - 最佳实践：始终先使用 dryRun 预览更改，然后再应用

- **create_directory**
  - 创建新目录或确保其存在
  - 输入：`path`（字符串）
  - 如需要会创建父目录
  - 如果目录已存在则静默成功

- **list_directory**
  - 列出目录内容，带有 [FILE] 或 [DIR] 前缀
  - 输入：`path`（字符串）

- **move_file**
  - 移动或重命名文件和目录
  - 输入：
    - `source`（字符串）
    - `destination`（字符串）
  - 如果目标已存在则失败

- **search_files**
  - 递归搜索文件/目录
  - 输入：
    - `path`（字符串）：起始目录
    - `pattern`（字符串）：搜索模式
    - `excludePatterns`（字符串数组）：排除模式。支持 Glob 格式。
  - 不区分大小写匹配
  - 返回匹配项的完整路径

- **get_file_info**
  - 获取详细的文件/目录元数据
  - 输入：`path`（字符串）
  - 返回：
    - 大小
    - 创建时间
    - 修改时间
    - 访问时间
    - 类型（文件/目录）
    - 权限

- **list_allowed_directories**
  - 列出服务器允许访问的所有目录
  - 无需输入
  - 返回：
    - 此服务器可以读取/写入的目录

## 与 Claude Desktop 一起使用

将此内容添加到您的 `claude_desktop_config.json`：

注意：您可以通过将沙盒目录挂载到 `/projects` 来为服务器提供沙盒目录。添加 `ro` 标志将使目录对服务器只读。

### Docker

注意：默认情况下，所有目录都必须挂载到 `/projects`。

```json
{
  "mcpServers": {
    "filesystem": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "--mount", "type=bind,src=/Users/username/Desktop,dst=/projects/Desktop",
        "--mount", "type=bind,src=/path/to/other/allowed/dir,dst=/projects/other/allowed/dir,ro",
        "--mount", "type=bind,src=/path/to/file.txt,dst=/projects/path/to/file.txt",
        "mcp/filesystem",
        "/projects"
      ]
    }
  }
}
```

### NPX

```json
{
  "mcpServers": {
    "filesystem": {
      "command": "npx",
      "args": [
        "-y",
        "@modelcontextprotocol/server-filesystem",
        "/Users/username/Desktop",
        "/path/to/other/allowed/dir"
      ]
    }
  }
}
```

## 与 VS Code 一起使用

快速安装，请点击下面的安装按钮...

[![在 VS Code 中使用 NPX 安装](https://img.shields.io/badge/VS_Code-NPM-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=filesystem&config=%7B%22command%22%3A%22npx%22%2C%22args%22%3A%5B%22-y%22%2C%22%40modelcontextprotocol%2Fserver-filesystem%22%2C%22%24%7BworkspaceFolder%7D%22%5D%7D) [![在 VS Code Insiders 中使用 NPX 安装](https://img.shields.io/badge/VS_Code_Insiders-NPM-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=filesystem&config=%7B%22command%22%3A%22npx%22%2C%22args%22%3A%5B%22-y%22%2C%22%40modelcontextprotocol%2Fserver-filesystem%22%2C%22%24%7BworkspaceFolder%7D%22%5D%7D&quality=insiders)

[![在 VS Code 中使用 Docker 安装](https://img.shields.io/badge/VS_Code-Docker-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=filesystem&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%22-i%22%2C%22--rm%22%2C%22--mount%22%2C%22type%3Dbind%2Csrc%3D%24%7BworkspaceFolder%7D%2Cdst%3D%2Fprojects%2Fworkspace%22%2C%22mcp%2Ffilesystem%22%2C%22%2Fprojects%22%5D%7D) [![在 VS Code Insiders 中使用 Docker 安装](https://img.shields.io/badge/VS_Code_Insiders-Docker-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=filesystem&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%22-i%22%2C%22--rm%22%2C%22--mount%22%2C%22type%3Dbind%2Csrc%3D%24%7BworkspaceFolder%7D%2Cdst%3D%2Fprojects%2Fworkspace%22%2C%22mcp%2Ffilesystem%22%2C%22%2Fprojects%22%5D%7D&quality=insiders)

手动安装时，将以下 JSON 块添加到 VS Code 的用户设置 (JSON) 文件中。您可以按 `Ctrl + Shift + P` 并输入 `Preferences: Open Settings (JSON)` 来执行此操作。

可选地，您可以将其添加到工作区中名为 `.vscode/mcp.json` 的文件中。这将允许您与他人共享配置。

> 注意在 `.vscode/mcp.json` 文件中不需要 `mcp` 键。

您可以通过将沙盒目录挂载到 `/projects` 来为服务器提供沙盒目录。添加 `ro` 标志将使目录对服务器只读。

### Docker

注意：默认情况下，所有目录都必须挂载到 `/projects`。

```json
{
  "mcp": {
    "servers": {
      "filesystem": {
        "command": "docker",
        "args": [
          "run",
          "-i",
          "--rm",
          "--mount", "type=bind,src=${workspaceFolder},dst=/projects/workspace",
          "mcp/filesystem",
          "/projects"
        ]
      }
    }
  }
}
```

### NPX

```json
{
  "mcp": {
    "servers": {
      "filesystem": {
        "command": "npx",
        "args": [
          "-y",
          "@modelcontextprotocol/server-filesystem",
          "${workspaceFolder}"
        ]
      }
    }
  }
}
```

## 构建

Docker 构建：

```bash
docker build -t mcp/filesystem -f src/filesystem/Dockerfile .
```

## 许可证

此 MCP 服务器根据 MIT 许可证授权。这意味着您可以自由使用、修改和分发软件，但需遵守 MIT 许可证的条款和条件。有关更多详细信息，请参阅项目存储库中的 LICENSE 文件。
