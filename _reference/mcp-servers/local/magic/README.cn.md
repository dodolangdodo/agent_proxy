# 21st.dev Magic AI 智能代理

> <https://github.com/21st-dev/magic-mcp>

![MCP Banner](https://21st.dev/magic-agent-og-image.png)

Magic Component Platform (MCP) 是一个强大的 AI 驱动工具，帮助开发者通过自然语言描述即时创建美观、现代的 UI 组件。它与流行的 IDE 无缝集成，为 UI 开发提供流畅的工作流程。

## 🌟 功能特性

- **AI 驱动的 UI 生成**：通过自然语言描述创建 UI 组件
- **多 IDE 支持**：
  - [Cursor](https://cursor.com) IDE 集成
  - [Windsurf](https://windsurf.ai) 支持
  - [VSCode](https://code.visualstudio.com/) 支持
  - [VSCode + Cline](https://cline.bot) 集成 (测试版)
- **现代组件库**：访问受 [21st.dev](https://21st.dev) 启发的大量预构建可定制组件
- **实时预览**：创建组件时即时查看效果
- **TypeScript 支持**：完整的 TypeScript 支持，确保类型安全开发
- **SVGL 集成**：访问大量专业品牌资产和标志
- **组件增强**：使用高级功能和动画改进现有组件（即将推出）

## 🎯 工作原理

1. **告诉代理您的需求**

   - 在您的 AI 代理聊天中，只需输入 `/ui` 并描述您需要的组件
   - 示例：`/ui 创建一个具有响应式设计的现代导航栏`

2. **让 Magic 创建它**

   - 您的 IDE 会提示您使用 Magic
   - Magic 立即构建一个精美的 UI 组件
   - 组件受 21st.dev 库的启发

3. **无缝集成**
   - 组件自动添加到您的项目中
   - 立即开始使用您的新 UI 组件
   - 所有组件都完全可定制

## 🚀 快速开始

### 前置要求

- Node.js（推荐最新 LTS 版本）
- 支持的 IDE 之一：
  - Cursor
  - Windsurf
  - VSCode（带 Cline 扩展）

### 安装

1. **生成 API 密钥**

   - 访问 [21st.dev Magic 控制台](https://21st.dev/magic/console)
   - 生成新的 API 密钥

2. **选择安装方法**

#### 方法 1：CLI 安装（推荐）

一条命令即可为您的 IDE 安装和配置 MCP：

```bash
npx @21st-dev/cli@latest install <client> --api-key <key>
```

支持的客户端：cursor、windsurf、cline、claude

#### 方法 2：手动配置

如果您更喜欢手动设置，请将此内容添加到您的 IDE 的 MCP 配置文件中：

```json
{
  "mcpServers": {
    "@21st-dev/magic": {
      "command": "npx",
      "args": ["-y", "@21st-dev/magic@latest", "API_KEY=\"your-api-key\""]
    }
  }
}
```

配置文件位置：

- Cursor：`~/.cursor/mcp.json`
- Windsurf：`~/.codeium/windsurf/mcp_config.json`
- Cline：`~/.cline/mcp_config.json`
- Claude：`~/.claude/mcp_config.json`

#### 方法 3：VS Code 安装

一键安装，点击下面的安装按钮：

[![在 VS Code 中使用 NPX 安装](https://img.shields.io/badge/VS_Code-NPM-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=%4021st-dev%2Fmagic&config=%7B%22command%22%3A%22npx%22%2C%22args%22%3A%5B%22-y%22%2C%22%4021st-dev%2Fmagic%40latest%22%5D%2C%22env%22%3A%7B%22API_KEY%22%3A%22%24%7Binput%3AapiKey%7D%22%7D%7D&inputs=%5B%7B%22type%22%3A%22promptString%22%2C%22id%22%3A%22apiKey%22%2C%22description%22%3A%2221st.dev+Magic+API+Key%22%2C%22password%22%3Atrue%7D%5D) [![在 VS Code Insiders 中使用 NPX 安装](https://img.shields.io/badge/VS_Code_Insiders-NPM-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=%4021st-dev%2Fmagic&config=%7B%22command%22%3A%22npx%22%2C%22args%22%3A%5B%22-y%22%2C%22%4021st-dev%2Fmagic%40latest%22%5D%2C%22env%22%3A%7B%22API_KEY%22%3A%22%24%7Binput%3AapiKey%7D%22%7D%7D&inputs=%5B%7B%22type%22%3A%22promptString%22%2C%22id%22%3A%22apiKey%22%2C%22description%22%3A%2221st.dev+Magic+API+Key%22%2C%22password%22%3Atrue%7D%5D&quality=insiders)

##### VS Code 手动设置

首先，请检查上面的安装按钮进行一键安装。手动设置：

将以下 JSON 块添加到 VS Code 的用户设置（JSON）文件中。您可以按 `Ctrl + Shift + P` 并输入 `首选项：打开用户设置（JSON）`：

```json
{
  "mcp": {
    "inputs": [
      {
        "type": "promptString",
        "id": "apiKey",
        "description": "21st.dev Magic API 密钥",
        "password": true
      }
    ],
    "servers": {
      "@21st-dev/magic": {
        "command": "npx",
        "args": ["-y", "@21st-dev/magic@latest"],
        "env": {
          "API_KEY": "${input:apiKey}"
        }
      }
    }
  }
}
```

或者，您可以将其添加到工作区中名为 `.vscode/mcp.json` 的文件中：

```json
{
  "inputs": [
    {
      "type": "promptString",
      "id": "apiKey",
      "description": "21st.dev Magic API 密钥",
      "password": true
    }
  ],
  "servers": {
    "@21st-dev/magic": {
      "command": "npx",
      "args": ["-y", "@21st-dev/magic@latest"],
      "env": {
        "API_KEY": "${input:apiKey}"
      }
    }
  }
}
```

## ❓ 常见问题

### Magic AI 代理如何处理我的代码库？

Magic AI 代理只会编写或修改与其生成的组件相关的文件。它遵循您项目的代码风格和结构，与现有代码库无缝集成，不会影响应用程序的其他部分。

### 我可以自定义生成的组件吗？

可以！所有生成的组件都完全可编辑，并具有良好的代码结构。您可以像修改代码库中的任何其他 React 组件一样修改样式、功能和行为。

### 如果我用完了生成次数会怎样？

如果您超过了每月生成限制，系统会提示您升级计划。您可以随时升级以继续生成组件。您现有的组件将保持完全功能。

### 新组件多久会添加到 21st.dev 的库中？

作者可以随时将组件发布到 21st.dev，Magic 代理将立即访问它们。这意味着您将始终能够访问社区中最新的组件和设计模式。

### 组件复杂度有限制吗？

Magic AI 代理可以处理各种复杂度的组件，从简单的按钮到复杂的交互式表单。但是，为了获得最佳效果，我们建议将非常复杂的 UI 分解为更小、更易管理的组件。

## 🛠️ 开发

### 项目结构

```
mcp/
├── app/
│   └── components/     # 核心 UI 组件
├── types/             # TypeScript 类型定义
├── lib/              # 实用函数
└── public/           # 静态资源
```

### 关键组件

- `IdeInstructions`：不同 IDE 的设置说明
- `ApiKeySection`：API 密钥管理界面
- `WelcomeOnboarding`：新用户引导流程

## 🤝 贡献

我们欢迎贡献！请加入我们的 [Discord 社区](https://discord.gg/Qx4rFunHfm) 并提供反馈以帮助改进 Magic 代理。源代码可在 [GitHub](https://github.com/serafimcloud/21st) 上获得。

## 👥 社区与支持

- [Discord 社区](https://discord.gg/Qx4rFunHfm) - 加入我们活跃的社区
- [Twitter](https://x.com/serafimcloud) - 关注我们获取更新

## ⚠️ 测试版声明

Magic 代理目前处于测试版。在此期间所有功能都是免费的。我们感谢您的反馈和耐心，我们将继续改进平台。

## 📝 许可证

MIT 许可证

## 🙏 致谢

- 感谢我们的测试用户和社区成员
- 特别感谢 Cursor、Windsurf 和 Cline 团队的合作
- 与 [21st.dev](https://21st.dev) 集成获得组件灵感
- [SVGL](https://svgl.app) 提供标志和品牌资产集成

---

更多信息，请加入我们的 [Discord 社区](https://discord.gg/Qx4rFunHfm) 或访问 [21st.dev/magic](https://21st.dev/magic)。
