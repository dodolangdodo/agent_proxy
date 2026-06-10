# Everything 搜索 MCP 服务器

[![smithery badge](https://smithery.ai/badge/mcp-server-everything-search)](https://smithery.ai/server/mcp-server-everything-search)

一个提供跨 Windows、macOS 和 Linux 快速文件搜索功能的 MCP 服务器。在 Windows 上，它使用 [Everything](https://www.voidtools.com/) SDK。在 macOS 上，它使用内置的 `mdfind` 命令。在 Linux 上，它使用 `locate`/`plocate` 命令。

## 工具

### search

在您的系统中搜索文件和文件夹。搜索功能和语法支持因平台而异：

- Windows：完整的 Everything SDK 功能（请参阅下面的语法指南）
- macOS：使用 Spotlight 数据库进行基本文件名和内容搜索
- Linux：使用 locate 数据库进行基本文件名搜索

参数：

- `query`（必需）：搜索查询字符串。请参阅下面的平台特定说明。
- `max_results`（可选）：返回的最大结果数（默认：100，最大：1000）
- `match_path`（可选）：匹配完整路径而不是仅文件名（默认：false）
- `match_case`（可选）：启用区分大小写搜索（默认：false）
- `match_whole_word`（可选）：仅匹配完整单词（默认：false）
- `match_regex`（可选）：启用正则表达式搜索（默认：false）
- `sort_by`（可选）：结果排序顺序（默认：1）。可用选项：

```
  - 1: 按文件名排序（A 到 Z）
  - 2: 按文件名排序（Z 到 A）
  - 3: 按路径排序（A 到 Z）
  - 4: 按路径排序（Z 到 A）
  - 5: 按大小排序（最小优先）
  - 6: 按大小排序（最大优先）
  - 7: 按扩展名排序（A 到 Z）
  - 8: 按扩展名排序（Z 到 A）
  - 11: 按创建日期排序（最旧优先）
  - 12: 按创建日期排序（最新优先）
  - 13: 按修改日期排序（最旧优先）
  - 14: 按修改日期排序（最新优先）
```

示例：

```json
{
  "query": "*.py",
  "max_results": 50,
  "sort_by": 6
}
```

```json
{
  "query": "ext:py datemodified:today",
  "max_results": 10
}
```

响应包括：

- 文件/文件夹路径
- 文件大小（字节）
- 最后修改日期

### 搜索语法指南

有关每个平台（Windows、macOS 和 Linux）支持的搜索语法的详细信息，请参阅 [SEARCH_SYNTAX.md](SEARCH_SYNTAX.md)。

## 先决条件

### Windows

1. [Everything](https://www.voidtools.com/) 搜索工具：
   - 从 <https://www.voidtools.com/> 下载并安装
   - **确保 Everything 服务正在运行**
2. Everything SDK：
   - 从 <https://www.voidtools.com/support/everything/sdk/> 下载
   - 将 SDK 文件解压到系统上的某个位置

### Linux

1. 安装并初始化 `locate` 或 `plocate` 命令：
   - Ubuntu/Debian：`sudo apt-get install plocate` 或 `sudo apt-get install mlocate`
   - Fedora：`sudo dnf install mlocate`
2. 安装后，更新数据库：
   - 对于 plocate：`sudo updatedb`
   - 对于 mlocate：`sudo /etc/cron.daily/mlocate`

### macOS

无需额外设置。服务器使用内置的 `mdfind` 命令。

## 安装

### 通过 Smithery 安装

通过 [Smithery](https://smithery.ai/server/mcp-server-everything-search) 自动为 Claude Desktop 安装 Everything Search：

```bash
npx -y @smithery/cli install mcp-server-everything-search --client claude
```

### 使用 uv（推荐）

使用 [`uv`](https://docs.astral.sh/uv/) 时无需特定安装。我们将使用 [`uvx`](https://docs.astral.sh/uv/guides/tools/) 直接运行 _mcp-server-everything-search_。

### 使用 PIP

或者，您可以通过 pip 安装 `mcp-server-everything-search`：

```
pip install mcp-server-everything-search
```

安装后，您可以使用以下命令作为脚本运行：

```
python -m mcp_server_everything_search
```

## 配置

### Windows

服务器需要 Everything SDK DLL 可用：

环境变量：

```
EVERYTHING_SDK_PATH=path\to\Everything-SDK\dll\Everything64.dll
```

### Linux 和 macOS

无需额外配置。

### 与 Claude Desktop 一起使用

根据您的平台，将以下配置之一添加到您的 `claude_desktop_config.json`：

<details>
<summary>Windows（使用 uvx）</summary>

```json
"mcpServers": {
  "everything-search": {
    "command": "uvx",
    "args": ["mcp-server-everything-search"],
    "env": {
      "EVERYTHING_SDK_PATH": "path/to/Everything-SDK/dll/Everything64.dll"
    }
  }
}
```

</details>

<details>
<summary>Windows（使用 pip 安装）</summary>

```json
"mcpServers": {
  "everything-search": {
    "command": "python",
    "args": ["-m", "mcp_server_everything_search"],
    "env": {
      "EVERYTHING_SDK_PATH": "path/to/Everything-SDK/dll/Everything64.dll"
    }
  }
}
```

</details>

<details>
<summary>Linux 和 macOS</summary>

```json
"mcpServers": {
  "everything-search": {
    "command": "uvx",
    "args": ["mcp-server-everything-search"]
  }
}
```

或者如果使用 pip 安装：

```json
"mcpServers": {
  "everything-search": {
    "command": "python",
    "args": ["-m", "mcp_server_everything_search"]
  }
}
```

</details>

## 调试

您可以使用 MCP 检查器来调试服务器。对于 uvx 安装：

```
npx @modelcontextprotocol/inspector uvx mcp-server-everything-search
```

或者如果您已将包安装在特定目录中或正在开发它：

```
git clone https://github.com/mamertofabian/mcp-everything-search.git
cd mcp-everything-search/src/mcp_server_everything_search
npx @modelcontextprotocol/inspector uv run mcp-server-everything-search
```

查看服务器日志：

Linux/macOS：

```bash
tail -f ~/.config/Claude/logs/mcp*.log
```

Windows（PowerShell）：

```powershell
Get-Content -Path "$env:APPDATA\Claude\logs\mcp*.log" -Tail 20 -Wait
```

## 开发

如果您正在进行本地开发，有两种方法测试您的更改：

1. 运行 MCP 检查器来测试您的更改。有关运行说明，请参阅[调试](#调试)。

2. 使用 Claude 桌面应用程序进行测试。将以下内容添加到您的 `claude_desktop_config.json`：

```json
"everything-search": {
  "command": "uv",
  "args": [
    "--directory",
    "/path/to/mcp-everything-search/src/mcp_server_everything_search",
    "run",
    "mcp-server-everything-search"
  ],
  "env": {
    "EVERYTHING_SDK_PATH": "path/to/Everything-SDK/dll/Everything64.dll"
  }
}
```

## 许可证

此 MCP 服务器采用 MIT 许可证。这意味着您可以自由使用、修改和分发软件，但需遵守 MIT 许可证的条款和条件。有关更多详细信息，请参阅项目存储库中的 LICENSE 文件。

## 免责声明

此项目与 voidtools（Everything 搜索工具的创建者）无关，未得到其认可或赞助。这是一个独立项目，使用公开可用的 Everything SDK。
