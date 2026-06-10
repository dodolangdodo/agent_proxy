# AppleScript MCP 服务器（双重访问：Python 和 Node.js）

[![npm version](https://img.shields.io/npm/v/@peakmojo/applescript-mcp.svg)](https://www.npmjs.com/package/@peakmojo/applescript-mcp) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## 概述

一个模型上下文协议（MCP）服务器，让您可以运行 AppleScript 代码来与 Mac 交互。这个 MCP 有意设计得简单、直观、易懂，并且需要最少的设置。

我简直不敢相信它如此简单而强大。核心代码少于 100 行。

<a href="https://glama.ai/mcp/servers/@peakmojo/applescript-mcp">
  <img width="380" height="200" src="https://glama.ai/mcp/servers/@peakmojo/applescript-mcp/badge" alt="AppleScript Server MCP server" />
</a>

<https://github.com/user-attachments/assets/b85e63ba-fb26-4918-8e6d-2377254ee388>

## 功能特性

* 运行 AppleScript 访问 Mac 应用程序和数据
* 与备忘录、日历、通讯录、信息等应用交互
* 使用 Spotlight 或访达搜索文件
* 读写文件内容并执行 shell 命令
* 通过 SSH 支持远程执行

## 示例提示

```
为我创建一个提醒，明天上午10点给约翰打电话
```

```
在我的日历中为周五下午2-3点添加一个名为"团队评审"的新会议
```

```
创建一个标题为"会议纪要"的新备忘录，包含今天的日期
```

```
显示我下载文件夹中过去一周的所有文件
```

```
我当前的电池电量是多少？
```

```
显示我收件箱中最近的未读邮件
```

```
列出我 Mac 上当前运行的所有应用程序
```

```
在 Apple Music 中播放我的"专注"播放列表
```

```
截取我整个屏幕的截图并保存到桌面
```

```
在我的通讯录中找到约翰·史密斯并显示他的电话号码
```

```
在我的桌面上创建一个名为"项目文件"的文件夹
```

```
打开 Safari 并导航到 apple.com
```

```
告诉我主硬盘还有多少可用空间
```

```
列出我本周所有即将到来的日历事件
```

## 在 Claude Desktop 中使用

### Node.js

```json
{
  "mcpServers": {
    "applescript_execute": {
      "command": "npx",
      "args": [
        "@peakmojo/applescript-mcp"
      ]
    }
  }
}
```

### Python

安装 uv

```
brew install uv
git clone ...
```

运行服务器

```
{
  "mcpServers": {
    "applescript_execute": {
      "command": "uv",
      "args": [
        "--directory",
        "/path/to/your/repo",
        "run",
        "src/applescript_mcp/server.py"
      ]
    }
  }
}
```

## Docker 使用

在 Docker 容器中运行时，您可以使用特殊主机名 `host.docker.internal` 连接到您的 Mac 主机：

### 配置

```json
{
  "mcpServers": {
    "applescript_execute": {
      "command": "npx",
      "args": [
        "@peakmojo/applescript-mcp",
        "--remoteHost", "host.docker.internal",
        "--remoteUser", "yourusername",
        "--remotePassword", "yourpassword"
      ]
    }
  }
}
```

这允许您的 Docker 容器在 Mac 主机系统上执行 AppleScript。请确保：

1. 在您的 Mac 上启用了 SSH（系统设置 → 共享 → 远程登录）
2. 您的用户具有适当的权限
3. 在配置中提供了正确的凭据
