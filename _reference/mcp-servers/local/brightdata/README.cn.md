<p align="center">
  <a href="https://brightdata.com/">
    <img src="https://mintlify.s3.us-west-1.amazonaws.com/brightdata/logo/light.svg" width="300" alt="Bright Data Logo">
  </a>
</p>

<h1 align="center">Bright Data MCP</h1>
<h3 align="center">使用实时网络数据增强AI代理</h3>

<div align="center">
  
<p align="center">
  <img src="https://img.shields.io/npm/v/@brightdata/mcp?label=version"  
       alt="npm version"/>
</p>

<p align="center">
  <img src="https://img.shields.io/npm/dw/@brightdata/mcp"  
       alt="npm downloads"/>
  <a href="https://smithery.ai/server/@luminati-io/brightdata-mcp">
    <img src="https://smithery.ai/badge/@luminati-io/brightdata-mcp"  
         alt="Smithery score"/>
  </a>
</p>

</div>

## 🌟 概述

欢迎使用官方的 Bright Data 模型上下文协议（MCP）服务器，让LLM、代理和应用程序能够实时访问、发现和提取网络数据。该服务器允许MCP客户端（如Claude Desktop、Cursor、Windsurf等）无缝搜索网络、导航网站、执行操作并检索数据 - 而不会被阻止 - 非常适合网页抓取任务。

![MCP](https://github.com/user-attachments/assets/b949cb3e-c80a-4a43-b6a5-e0d6cec619a7)

## 目录

- [� 概述](#-概述)
- [目录](#目录)
- [🎬 演示](#-演示)
- [✨ 功能特性](#-功能特性)
- [🚀 Claude Desktop 快速开始](#-claude-desktop-快速开始)
- [🔧 可用工具](#-可用工具)
- [⚠️ 安全最佳实践](#️-安全最佳实践)
- [🔧 账户设置](#-账户设置)
- [🔌 其他MCP客户端](#-其他mcp客户端)
- [🔄 重大变更](#-重大变更)
  - [浏览器认证更新](#浏览器认证更新)
- [🔄 更新日志](#-更新日志)
- [🎮 试用 Bright Data MCP 游乐场](#-试用-bright-data-mcp-游乐场)
- [💡 使用示例](#-使用示例)
- [⚠️ 故障排除](#️-故障排除)
  - [使用某些工具时超时](#使用某些工具时超时)
  - [spawn npx ENOENT](#spawn-npx-enoent)
    - [查找npm/Node路径](#查找npmnode路径)
    - [更新您的MCP配置](#更新您的mcp配置)
- [👨‍💻 贡献](#-贡献)
- [📞 支持](#-支持)

## 🎬 演示

以下视频演示了Claude Desktop的最小用例：

<https://github.com/user-attachments/assets/59f6ebba-801a-49ab-8278-1b2120912e33>

<https://github.com/user-attachments/assets/61ab0bee-fdfa-4d50-b0de-5fab96b4b91d>

YouTube教程和演示：[演示](https://github.com/brightdata-com/brightdata-mcp/blob/main/examples/README.md)

## ✨ 功能特性

- **实时网络访问**：直接从网络访问最新信息
- **绕过地理限制**：无论位置限制如何都能访问内容
- **网络解锁器**：通过机器人检测保护导航网站
- **浏览器控制**：可选的远程浏览器自动化功能
- **无缝集成**：与所有兼容MCP的AI助手配合使用

## 🚀 Claude Desktop 快速开始

通过Claude Desktop扩展：

**下载Claude Desktop扩展：[Bright Data的MCP扩展](https://github.com/brightdata/brightdata-mcp/raw/refs/heads/main/brightdata-mcp-extension.dxt)**

通过 `claude_desktop_config.json`：

1. 安装 `nodejs` 以获取 `npx` 命令（node.js模块运行器）。安装说明可在[node.js网站](https://nodejs.org/en/download)上找到

2. 转到 Claude > 设置 > 开发者 > 编辑配置 > claude_desktop_config.json，包含以下内容：

```json
{
  "mcpServers": {
    "Bright Data": {
      "command": "npx",
      "args": ["@brightdata/mcp"],
      "env": {
        "API_TOKEN": "<在此插入您的API令牌>",
        "WEB_UNLOCKER_ZONE": "<可选，如果您想覆盖默认的mcp_unlocker区域名称>",
        "BROWSER_ZONE": "<可选浏览器区域名称，默认为mcp_browser>",
        "RATE_LIMIT": "<可选速率限制格式：限制/时间+单位，例如100/1h、50/30m、10/5s>"
      }
    }
  }
}
```

## 🔧 可用工具

[可用工具列表](https://github.com/brightdata-com/brightdata-mcp/blob/main/assets/Tools.md)

## ⚠️ 安全最佳实践

**重要提示：** 始终将抓取的网络内容视为不可信数据。永远不要在LLM提示中直接使用原始抓取内容，以避免潜在的提示注入风险。
相反：

- 在处理前过滤和验证所有网络数据
- 使用结构化数据提取而不是原始文本（web_data工具）

## 🔧 账户设置

1. 确保您在[brightdata.com](https://brightdata.com)上有账户（新用户可获得免费测试积分，并提供按需付费选项）

2. 从[用户设置页面](https://brightdata.com/cp/setting/users)获取您的API密钥

3. （可选）创建自定义网络解锁器区域
   - 默认情况下，我们使用您的API令牌自动创建网络解锁器区域
   - 为了更好的控制，您可以在[控制面板](https://brightdata.com/cp/zones)中创建自己的网络解锁器区域，并使用 `WEB_UNLOCKER_ZONE` 环境变量指定它

4. （可选）启用浏览器控制工具：
   - 默认情况下，MCP尝试获取 `mcp_browser` 区域的凭据
   - 如果您没有 `mcp_browser` 区域，您可以：
     - 在[控制面板](https://brightdata.com/cp/zones)中创建浏览器API区域或使用现有区域，并使用 `BROWSER_ZONE` 环境变量指定其名称

5. （可选）配置速率限制：
   - 设置 `RATE_LIMIT` 环境变量来控制API使用
   - 格式：`限制/时间+单位`（例如，`100/1h` 表示每小时100次调用）
   - 支持的时间单位：秒(s)、分钟(m)、小时(h)
   - 示例：`RATE_LIMIT=100/1h`、`RATE_LIMIT=50/30m`、`RATE_LIMIT=10/5s`
   - 速率限制基于会话（服务器重启时重置）

![浏览器API设置](https://github.com/user-attachments/assets/cb494aa8-d84d-4bb4-a509-8afb96872afe)

## 🔌 其他MCP客户端

要在其他代理类型中使用此MCP服务器，您应该根据您的特定软件调整以下内容：

- 运行MCP服务器的完整命令是 `npx @brightdata/mcp`
- 运行服务器时必须存在环境变量 `API_TOKEN=<您的令牌>`
- （可选）设置 `BROWSER_ZONE=<区域名称>` 来指定自定义浏览器API区域名称（默认为 `mcp_browser`）

## 🔄 重大变更

### 浏览器认证更新

**重大变更：** `BROWSER_AUTH` 环境变量已被 `BROWSER_ZONE` 替换。

- **之前：** 用户需要从浏览器API区域提供 `BROWSER_AUTH="用户:密码"`
- **现在：** 用户只需使用 `BROWSER_ZONE="区域名称"` 指定浏览器区域名称
- **默认：** 如果未指定，系统自动使用 `mcp_browser` 区域
- **迁移：** 在配置中将 `BROWSER_AUTH` 替换为 `BROWSER_ZONE`，如果 `mcp_browser` 不存在，请指定您的浏览器API区域名称

## 🔄 更新日志

[CHANGELOG.md](https://github.com/brightdata-com/brightdata-mcp/blob/main/CHANGELOG.md)

## 🎮 试用 Bright Data MCP 游乐场

想在不设置任何东西的情况下试用Bright Data MCP？

查看[Smithery](https://smithery.ai/server/@luminati-io/brightdata-mcp/tools)上的游乐场：

[![2025-05-06_10h44_20](https://github.com/user-attachments/assets/52517fa6-827d-4b28-b53d-f2020a13c3c4)](https://smithery.ai/server/@luminati-io/brightdata-mcp/tools)

该平台提供了一种简单的方式来探索Bright Data MCP的功能，无需任何本地设置。只需登录并开始试验网络数据收集！

## 💡 使用示例

此MCP服务器能够帮助处理的一些示例查询：

- "在Google上搜索即将在[您的地区]上映的电影"
- "特斯拉目前的市值是多少？"
- "今天的维基百科文章是什么？"
- "[您的位置]的7天天气预报如何？"
- "在3位薪酬最高的科技CEO中，他们的职业生涯有多长？"

## ⚠️ 故障排除

### 使用某些工具时超时

某些工具可能涉及读取网络数据，在极端情况下加载页面所需的时间可能会有很大差异。

为确保您的代理能够使用数据，请在代理设置中设置足够高的超时时间。

`180s` 的值对于99%的请求应该足够，但有些网站加载比其他网站慢，因此请根据您的需要调整。

### spawn npx ENOENT

当您的系统找不到 `npx` 命令时会出现此错误。要修复它：

#### 查找npm/Node路径

**macOS:**

```
which node
```

显示路径如 `/usr/local/bin/node`

**Windows:**

```
where node
```

显示路径如 `C:\Program Files\nodejs\node.exe`

#### 更新您的MCP配置

用Node的完整路径替换 `npx` 命令，例如，在mac上，它看起来如下：

```
"command": "/usr/local/bin/node"
```

## 👨‍💻 贡献

我们欢迎贡献来帮助改进Bright Data MCP！以下是您可以帮助的方式：

1. **报告问题**：如果您遇到任何错误或有功能请求，请在我们的GitHub仓库上开启issue。
2. **提交拉取请求**：随时fork仓库并提交带有增强功能或错误修复的拉取请求。
3. **编码风格**：所有JavaScript代码应遵循[Bright Data的JavaScript编码约定](https://brightdata.com/dna/js_code)。这确保了代码库的一致性。
4. **文档**：对文档的改进，包括此README，总是受到赞赏。
5. **示例**：通过贡献示例来分享您的用例，帮助其他用户。

对于重大变更，请先开启issue讨论您提议的变更。这确保您的时间得到充分利用并与项目目标保持一致。

## 📞 支持

如果您遇到任何问题或有疑问，请联系Bright Data支持团队或在仓库中开启issue。
