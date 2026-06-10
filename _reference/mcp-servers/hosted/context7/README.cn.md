# Context7 MCP - 为任何提示提供最新的代码文档

[![Website](https://img.shields.io/badge/Website-context7.com-blue)](https://context7.com) [![smithery badge](https://smithery.ai/badge/@upstash/context7-mcp)](https://smithery.ai/server/@upstash/context7-mcp) [<img alt="在 VS Code 中安装 (npx)" src="https://img.shields.io/badge/VS_Code-VS_Code?style=flat-square&label=安装%20Context7%20MCP&color=0098FF">](https://insiders.vscode.dev/redirect?url=vscode%3Amcp%2Finstall%3F%7B%22name%22%3A%22context7%22%2C%22command%22%3A%22npx%22%2C%22args%22%3A%5B%22-y%22%2C%22%40upstash%2Fcontext7-mcp%40latest%22%5D%7D)

## ❌ 没有 Context7 的情况

LLM 依赖于你使用的库的**过时或通用信息**。你会遇到：

- ❌ 代码示例过时，基于一年前的训练数据
- ❌ 幻觉 API（根本不存在的 API）
- ❌ 针对旧版本包的通用答案

## ✅ 使用 Context7 的优势

Context7 MCP 直接从源头拉取**最新的、版本特定的文档和代码示例**——并将它们直接插入到你的提示中。

在 Cursor 中添加 `use context7` 到你的提示：

```txt
创建一个使用 app router 的基础 Next.js 项目。use context7
```

```txt
给定 PostgreSQL 凭证，创建一个删除城市为空字符串的行的脚本。use context7
```

Context7 会将最新的代码示例和文档直接注入到 LLM 的上下文的。

- 1️⃣ 自然地写出你的提示
- 2️⃣ 告诉 LLM 要 `use context7`
- 3️⃣ 获得可运行的代码答案

无需切换标签页，没有不存在的幻觉 API，没有过时的代码生成。

## 📚 添加项目

查看我们的 [项目添加指南](./docs/adding-projects.md)，了解如何将你喜欢的库添加（或更新）到 Context7。

## 🛠️ 安装

### 要求

- Node.js >= v18.0.0
- Cursor、Windsurf、Claude Desktop 或其他 MCP 客户端

<details>
<summary><b>通过 Smithery 安装</b></summary>

要通过 [Smithery](https://smithery.ai/server/@upstash/context7-mcp) 为任何客户端自动安装 Context7 MCP 服务器：

```bash
npx -y @smithery/cli@latest install @upstash/context7-mcp --client <客户端名称> --key <你的 Smithery 密钥>
```

你可以在 [Smithery.ai 网页](https://smithery.ai/server/@upstash/context7-mcp) 中找到你的 Smithery 密钥。

</details>

<details>
<summary><b>在 Cursor 中安装</b></summary>

前往：`设置` -> `Cursor 设置` -> `MCP` -> `添加新的全局 MCP 服务器`

推荐的方式是将以下配置粘贴到你的 Cursor `~/.cursor/mcp.json` 文件中。你也可以通过在项目文件夹中创建 `.cursor/mcp.json` 来在特定项目中安装。查看 [Cursor MCP 文档](https://docs.cursor.com/context/model-context-protocol) 了解更多信息。

> 自 Cursor 1.0 起，你可以点击下方的安装按钮进行一键安装。

#### Cursor 远程服务器连接

[![安装 MCP 服务器](https://cursor.com/deeplink/mcp-install-dark.svg)](https://cursor.com/install-mcp?name=context7&config=eyJ1cmwiOiJodHRwczovL21jcC5jb250ZXh0Ny5jb20vbWNwIn0%3D)

```json
{
  "mcpServers": {
    "context7": {
      "url": "https://mcp.context7.com/mcp"
    }
  }
}
```

#### Cursor 本地服务器连接

[![安装 MCP 服务器](https://cursor.com/deeplink/mcp-install-dark.svg)](https://cursor.com/install-mcp?name=context7&config=eyJjb21tYW5kIjoibnB4IC15IEB1cHN0YXNoL2NvbnRleHQ3LW1jcCJ9)

```json
{
  "mcpServers": {
    "context7": {
      "command": "npx",
      "args": ["-y", "@upstash/context7-mcp"]
    }
  }
}
```

<details>
<summary>替代方案：使用 Bun</summary>

[![安装 MCP 服务器](https://cursor.com/deeplink/mcp-install-dark.svg)](https://cursor.com/install-mcp?name=context7&config=eyJjb21tYW5kIjoiYnVueCAteSBAdXBzdGFzaC9jb250ZXh0Ny1tY3AifQ%3D%3D)

```json
{
  "mcpServers": {
    "context7": {
      "command": "bunx",
      "args": ["-y", "@upstash/context7-mcp"]
    }
  }
}
```

</details>

<details>
<summary>替代方案：使用 Deno</summary>

[![安装 MCP 服务器](https://cursor.com/deeplink/mcp-install-dark.svg)](https://cursor.com/install-mcp?name=context7&config=eyJjb21tYW5kIjoiZGVubyBydW4gLS1hbGxvdy1lbnYgLS1hbGxvdy1uZXQgbnBtOkB1cHN0YXNoL2NvbnRleHQ3LW1jcCJ9)

```json
{
  "mcpServers": {
    "context7": {
      "command": "deno",
      "args": ["run", "--allow-env=NO_DEPRECATION,TRACE_DEPRECATION", "--allow-net", "npm:@upstash/context7-mcp"]
    }
  }
}
```

</details>

</details>

<details>
<summary><b>在 Windsurf 中安装</b></summary>

将此添加到你的 Windsurf MCP 配置文件中。查看 [Windsurf MCP 文档](https://docs.windsurf.com/windsurf/mcp) 了解更多信息。

#### Windsurf 远程服务器连接

```json
{
  "mcpServers": {
    "context7": {
      "serverUrl": "https://mcp.context7.com/sse"
    }
  }
}
```

#### Windsurf 本地服务器连接

```json
{
  "mcpServers": {
    "context7": {
      "command": "npx",
      "args": ["-y", "@upstash/context7-mcp"]
    }
  }
}
```

</details>

<details>
<summary><b>在 VS Code 中安装</b></summary>

[<img alt="在 VS Code 中安装 (npx)" src="https://img.shields.io/badge/VS_Code-VS_Code?style=flat-square&label=安装%20Context7%20MCP&color=0098FF">](https://insiders.vscode.dev/redirect?url=vscode%3Amcp%2Finstall%3F%7B%22name%22%3A%22context7%22%2C%22command%22%3A%22npx%22%2C%22args%22%3A%5B%22-y%22%2C%22%40upstash%2Fcontext7-mcp%40latest%22%5D%7D)
[<img alt="在 VS Code Insiders 中安装 (npx)" src="https://img.shields.io/badge/VS_Code_Insiders-VS_Code_Insiders?style=flat-square&label=安装%20Context7%20MCP&color=24bfa5">](https://insiders.vscode.dev/redirect?url=vscode-insiders%3Amcp%2Finstall%3F%7B%22name%22%3A%22context7%22%2C%22command%22%3A%22npx%22%2C%22args%22%3A%5B%22-y%22%2C%22%40upstash%2Fcontext7-mcp%40latest%22%5D%7D)

将此添加到你的 VS Code MCP 配置文件中。查看 [VS Code MCP 文档](https://code.visualstudio.com/docs/copilot/chat/mcp-servers) 了解更多信息。

#### VS Code 远程服务器连接

```json
"mcp": {
  "servers": {
    "context7": {
      "type": "http",
      "url": "https://mcp.context7.com/mcp"
    }
  }
}
```

#### VS Code 本地服务器连接

```json
"mcp": {
  "servers": {
    "context7": {
      "type": "stdio",
      "command": "npx",
      "args": ["-y", "@upstash/context7-mcp"]
    }
  }
}
```

</details>

<details>
<summary><b>在 Visual Studio 2022 中安装</b></summary>

你可以按照 [Visual Studio MCP 服务器文档](https://learn.microsoft.com/visualstudio/ide/mcp-servers?view=vs-2022) 配置 Context7 MCP。

将此添加到你的 Visual Studio MCP 配置文件中（查看 [Visual Studio 文档](https://learn.microsoft.com/visualstudio/ide/mcp-servers?view=vs-2022) 了解详情）：

```json
{
  "mcp": {
    "servers": {
      "context7": {
        "type": "http",
        "url": "https://mcp.context7.com/mcp"
      }
    }
  }
}
```

或用于本地服务器：

```json
{
  "mcp": {
    "servers": {
      "context7": {
        "type": "stdio",
        "command": "npx",
        "args": ["-y", "@upstash/context7-mcp"]
      }
    }
  }
}
```

如需更多信息和故障排除，请参考 [Visual Studio MCP 服务器文档](https://learn.microsoft.com/visualstudio/ide/mcp-servers?view=vs-2022)。
</details>

<details>
<summary><b>在 Zed 中安装</b></summary>

你可以通过 [Zed 扩展](https://zed.dev/extensions?query=Context7) 安装，或添加以下内容到你的 Zed `settings.json` 中。查看 [Zed 上下文服务器文档](https://zed.dev/docs/assistant/context-servers) 了解更多信息。

```json
{
  "context_servers": {
    "Context7": {
      "command": {
        "path": "npx",
        "args": ["-y", "@upstash/context7-mcp"]
      },
      "settings": {}
    }
  }
}
```

</details>

<details>
<summary><b>在 Claude Code 中安装</b></summary>

运行以下命令。查看 [Claude Code MCP 文档](https://docs.anthropic.com/en/docs/agents-and-tools/claude-code/tutorials#set-up-model-context-protocol-mcp) 了解更多信息。

#### Claude Code 远程服务器连接

```sh
claude mcp add --transport sse context7 https://mcp.context7.com/sse
```

#### Claude Code 本地服务器连接

```sh
claude mcp add context7 -- npx -y @upstash/context7-mcp
```

</details>

<details>
<summary><b>在 Claude Desktop 中安装</b></summary>

将此添加到你的 Claude Desktop `claude_desktop_config.json` 文件中。查看 [Claude Desktop MCP 文档](https://modelcontextprotocol.io/quickstart/user) 了解更多信息。

```json
{
  "mcpServers": {
    "Context7": {
      "command": "npx",
      "args": ["-y", "@upstash/context7-mcp"]
    }
  }
}
```

</details>

<details>
<summary><b>在 BoltAI 中安装</b></summary>

打开应用的“设置”页面，导航到“插件”，并输入以下 JSON：

```json
{
  "mcpServers": {
    "context7": {
      "command": "npx",
      "args": ["-y", "@upstash/context7-mcp"]
    }
  }
}
```

保存后，在聊天中输入 `get-library-docs`  followed by your Context7 documentation ID（例如：`get-library-docs /nuxt/ui`）。更多信息请查看 [BoltAI 文档](https://docs.boltai.com/docs/plugins/mcp-servers)。对于 iOS 上的 BoltAI，请 [查看此指南](https://docs.boltai.com/docs/boltai-mobile/mcp-servers)。

</details>

<details>
<summary><b>使用 Docker 安装</b></summary>

如果你更喜欢在 Docker 容器中运行 MCP 服务器：

1. **构建 Docker 镜像：**

   首先，在项目根目录（或任何你喜欢的位置）创建一个 `Dockerfile`：

   <details>
   <summary>点击查看 Dockerfile 内容</summary>

   ```Dockerfile
   FROM node:18-alpine

   WORKDIR /app

   # 全局安装最新版本
   RUN npm install -g @upstash/context7-mcp

   # 暴露默认端口（可选，取决于 MCP 客户端交互）
   # EXPOSE 3000

   # 运行服务器的默认命令
   CMD ["context7-mcp"]
   ```

   </details>

   然后，使用标签（例如 `context7-mcp`）构建镜像。**确保 Docker Desktop（或 Docker 守护进程）正在运行。** 在保存 `Dockerfile` 的同一目录中运行以下命令：

   ```bash
   docker build -t context7-mcp .
   ```

2. **配置你的 MCP 客户端：**

   更新 MCP 客户端的配置以使用 Docker 命令。

   _例如，cline_mcp_settings.json：_

   ```json
   {
     "mcpServers": {
       "Сontext7": {
         "autoApprove": [],
         "disabled": false,
         "timeout": 60,
         "command": "docker",
         "args": ["run", "-i", "--rm", "context7-mcp"],
         "transportType": "stdio"
       }
     }
   }
   ```

   _注意：这是示例配置。请参考前面 README 中针对你的 MCP 客户端（如 Cursor、VS Code 等）的具体示例来调整结构（例如 `mcpServers` vs `servers`）。同时，确保 `args` 中的镜像名称与 `docker build` 命令中使用的标签一致。_

</details>

<details>
<summary><b>在 Windows 中安装</b></summary>

Windows 上的配置与 Linux 或 macOS 略有不同（_示例中使用 `Cline`_）。其他编辑器的配置原理相同，请参考 `command` 和 `args` 的配置方式。

```json
{
  "mcpServers": {
    "github.com/upstash/context7-mcp": {
      "command": "cmd",
      "args": ["/c", "npx", "-y", "@upstash/context7-mcp@latest"],
      "disabled": false,
      "autoApprove": []
    }
  }
}
```

</details>

<details>
<summary><b>在 Augment Code 中安装</b></summary>

要在 Augment Code 中配置 Context7 MCP，你可以使用图形界面或手动配置。

### **A. 使用 Augment Code UI**

1. 点击汉堡菜单。
2. 选择 **设置**。
3. 导航到 **工具** 部分。
4. 点击 **+ 添加 MCP** 按钮。
5. 输入以下命令：

   ```
   npx -y @upstash/context7-mcp@latest
   ```

6. 命名 MCP：**Context7**。
7. 点击 **添加** 按钮。

添加 MCP 服务器后，你可以直接在 Augment Code 中使用 Context7 的最新代码文档功能。

---

### **B. 手动配置**

1. 按下 Cmd/Ctrl + Shift + P，或前往 Augment 面板的汉堡菜单。
2. 选择 **编辑设置**。
3. 在 **高级** 下，点击 **在 settings.json 中编辑**。
4. 将服务器配置添加到 `augment.advanced` 对象中的 `mcpServers` 数组：

```json
"augment.advanced": {
  "mcpServers": [
    {
      "name": "context7",
      "command": "npx",
      "args": ["-y", "@upstash/context7-mcp"]
    }
  ]
}
```

添加 MCP 服务器后，重启编辑器。如果遇到错误，请检查语法是否缺少闭合括号或逗号。

</details>

<details>
<summary><b>在 Roo Code 中安装</b></summary>

将此添加到你的 Roo Code MCP 配置文件中。查看 [Roo Code MCP 文档](https://docs.roocode.com/features/mcp/using-mcp-in-roo) 了解更多信息。

#### Roo Code 远程服务器连接

```json
{
  "mcpServers": {
    "context7": {
      "type": "streamable-http",
      "url": "https://mcp.context7.com/mcp"
    }
  }
}
```

#### Roo Code 本地服务器连接

```json
{
  "mcpServers": {
    "context7": {
      "command": "npx",
      "args": ["-y", "@upstash/context7-mcp"]
    }
  }
}
```

</details>

<details>
<summary><b>在 Zencoder 中安装</b></summary>

要在 Zencoder 中配置 Context7 MCP，请按照以下步骤操作：

1. 前往 Zencoder 菜单 (...)。
2. 从下拉菜单中选择 **Agent tools**。
3. 点击 **Add custom MCP**。
4. 添加以下名称和服务器配置，并确保点击 **Install** 按钮：

```json
{
    "command": "npx",
    "args": [
        "-y",
        "@upstash/context7-mcp@latest"
    ]
}
```

添加 MCP 服务器后，你可以轻松继续使用它。

</details>

<details>
<summary><b>在 Amazon Q Developer CLI 中安装</b></summary>

将此添加到你的 Amazon Q Developer CLI 配置文件中。查看 [Amazon Q Developer CLI 文档](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/command-line-mcp-configuration.html) 了解更多详情。

```json
{
  "mcpServers": {
    "context7": {
      "command": "npx",
      "args": ["-y", "@upstash/context7-mcp@latest"]
    }
  }
}
```

</details>

<details>
<summary><b>在 Qodo Gen 中安装</b></summary>

查看 [Qodo Gen 文档](https://docs.qodo.ai/qodo-documentation/qodo-gen/qodo-gen-chat/agentic-mode/agentic-tools-mcps) 了解更多详情。

1. 在 VSCode 或 IntelliJ 中打开 Qodo Gen 聊天面板。
2. 点击 **Connect more tools**。
3. 点击 **+ Add new MCP**。
4. 添加以下配置：

```json
{
  "mcpServers": {
    "context7": {
      "url": "https://mcp.context7.com/mcp"
    }
  }
}
```

</details>

## 🔨 可用工具

Context7 MCP 提供以下 LLM 可使用的工具：

- `resolve-library-id`：将通用库名称解析为 Context7 兼容的库 ID。
  - `libraryName`（必填）：要搜索的库名称。

- `get-library-docs`：使用 Context7 兼容的库 ID 获取库文档。
  - `context7CompatibleLibraryID`（必填）：精确的 Context7 兼容库 ID（例如：`/mongodb/docs`、`/vercel/next.js`）。
  - `topic`（可选）：将文档聚焦于特定主题（例如："routing"、"hooks"）。
  - `tokens`（可选，默认 10000）：返回的最大 tokens 数。小于默认值 10000 的值将自动增加到 10000。

## 💻 开发

克隆项目并安装依赖：

```bash
bun i
```

构建：

```bash
bun run build
```

运行服务器：

```bash
bun run dist/index.js
```

### CLI 参数

`context7-mcp` 接受以下 CLI 标志：

- `--transport <stdio|http|sse>` – 使用的传输方式（默认 `stdio`）。
- `--port <number>` – 使用 `http` 或 `sse` 传输时的监听端口（默认 `3000`）。

示例：使用 http 传输和端口 8080：

```bash
bun run dist/index.js --transport http --port 8080
```

<details>
<summary><b>本地配置示例</b></summary>

```json
{
  "mcpServers": {
    "context7": {
      "command": "npx",
      "args": ["tsx", "/path/to/folder/context7-mcp/src/index.ts"]
    }
  }
}
```

</details>

<details>
<summary><b>使用 MCP Inspector 测试</b></summary>

```bash
npx -y @modelcontextprotocol/inspector npx @upstash/context7-mcp
```

</details>

## 🚨 故障排除

<details>
<summary><b>模块未找到错误</b></summary>

如果遇到 `ERR_MODULE_NOT_FOUND`，尝试使用 `bunx` 代替 `npx`：

```json
{
  "mcpServers": {
    "context7": {
      "command": "bunx",
      "args": ["-y", "@upstash/context7-mcp"]
    }
  }
}
```

这通常可以解决 `npx` 无法正确安装或解析包的环境中的模块解析问题。

</details>

<details>
<summary><b>ESM 解析问题</b></summary>

对于类似 `Error: Cannot find module 'uriTemplate.js'` 的错误，尝试添加 `--experimental-vm-modules` 标志：

```json
{
  "mcpServers": {
    "context7": {
      "command": "npx",
      "args": ["-y", "--node-options=--experimental-vm-modules", "@upstash/context7-mcp@1.0.6"]
    }
  }
}
```

</details>

<details>
<summary><b>TLS/证书问题</b></summary>

使用 `--experimental-fetch` 标志绕过 TLS 相关问题：

```json
{
  "mcpServers": {
    "context7": {
      "command": "npx",
      "args": ["-y", "--node-options=--experimental-fetch", "@upstash/context7-mcp"]
    }
  }
}
```

</details>

<details>
<summary><b>通用 MCP 客户端错误</b></summary>

1. 尝试在包名后添加 `@latest`。
2. 使用 `bunx` 作为 `npx` 的替代方案。
3. 考虑使用 `deno` 作为另一个替代方案。
4. 确保使用 Node.js v18 或更高版本以支持原生 fetch。

</details>

## ⚠️ 免责声明

Context7 项目由社区贡献，虽然我们努力保持高质量，但**无法保证所有库文档的准确性、完整性或安全性**。Context7 中列出的项目由其各自的所有者开发和维护，而非 Context7。如果你遇到任何可疑、不当或潜在有害的内容，请使用项目页面上的“报告”按钮立即通知我们。我们会认真对待所有报告，并及时审查标记的内容，以维护平台的完整性和安全性。使用 Context7 即表示你自行承担风险。

## 🤝 联系我们

保持更新并加入我们的社区：

- 📢 关注我们的 [X](https://x.com/contextai) 获取最新新闻和更新。
- 🌐 访问我们的 [网站](https://context7.com)。
- 💬 加入我们的 [Discord 社区](https://upstash.com/discord)。

## 📺 Context7 在媒体中

- [Better Stack：《免费工具让 Cursor 聪明 10 倍》](https://youtu.be/52FC3qObp9E)
- [Cole Medin：《这绝对是 AI 编码助手最好的 MCP 服务器》](https://www.youtube.com/watch?v=G7gK8H6u7Rs)
- [Income Stream Surfers：《Context7 + SequentialThinking MCP：这是 AGI 吗？》](https://www.youtube.com/watch?v=-ggvzyLpK6o)
- [Julian Goldie SEO：《Context7：新 MCP AI 代理更新》](https://www.youtube.com/watch?v=CTZm6fBYisc)
- [JeredBlu：《Context7 MCP：即时获取文档 + VS Code 设置》](https://www.youtube.com/watch?v=-ls0D-rtET4)
- [Income Stream Surfers：《Context7：将改变 AI 编码的新 MCP 服务器》](https://www.youtube.com/watch?v=PS-2Azb-C3M)
- [AICodeKing：《Context7 + Cline & RooCode：这个 MCP 服务器让 CLINE 高效 100 倍！》](https://www.youtube.com/watch?v=qZfENAPMnyo)
- [Sean Kochel：《5 个 MCP 服务器让你轻松编码（只需插入即用）》](https://www.youtube.com/watch?v=LqTQi8qexJM)

## ⭐ Star 历史

[![Star 历史图表](https://api.star-history.com/svg?repos=upstash/context7&type=Date)](https://www.star-history.com/#upstash/context7&Date)

## 📄 许可证

MIT
