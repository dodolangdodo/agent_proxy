# OpenMemory

> <https://github.com/mem0ai/mem0/tree/main/openmemory>

OpenMemory 是您的个人化 LLM 记忆层 - 私有、便携且开源。您的记忆数据存储在本地，让您完全控制自己的数据。在保持数据安全的同时，构建具有个性化记忆的 AI 应用程序。

![OpenMemory](https://github.com/user-attachments/assets/3c701757-ad82-4afa-bfbe-e049c2b4320b)

## 简易安装

### 前置要求

- Docker
- OpenAI API 密钥

您可以通过运行以下命令快速运行 OpenMemory：

```bash
curl -sL https://raw.githubusercontent.com/mem0ai/mem0/main/openmemory/run.sh | bash
```

您应该将 `OPENAI_API_KEY` 设置为全局环境变量：

```bash
export OPENAI_API_KEY=your_api_key
```

您也可以将 `OPENAI_API_KEY` 作为脚本参数设置：

```bash
curl -sL https://raw.githubusercontent.com/mem0ai/mem0/main/openmemory/run.sh | OPENAI_API_KEY=your_api_key bash
```

## 前置要求

- Docker 和 Docker Compose
- Python 3.9+（用于后端开发）
- Node.js（用于前端开发）
- OpenAI API 密钥（LLM 交互必需，运行 `cp api/.env.example api/.env` 然后将 **OPENAI_API_KEY** 改为您的密钥）

## 快速开始

### 1. 设置环境变量

在运行项目之前，您需要为 API 和 UI 配置环境变量。

您可以通过以下方式之一完成：

- **手动方式**：  
  在以下目录中分别创建 `.env` 文件：
  - `/api/.env`
  - `/ui/.env`

- **使用 `.env.example` 文件**：  
  复制并重命名示例文件：

  ```bash
  cp api/.env.example api/.env
  cp ui/.env.example ui/.env
  ```

- **使用 Makefile**（如果支持）：  
    运行：
  
   ```bash
   make env
   ```

- #### `/api/.env` 示例

```env
OPENAI_API_KEY=sk-xxx
USER=<user-id> # 您想要关联记忆的用户 ID
```

- #### `/ui/.env` 示例

```env
NEXT_PUBLIC_API_URL=http://localhost:8765
NEXT_PUBLIC_USER_ID=<user-id> # 与 api 环境变量中的用户 ID 相同
```

### 2. 构建并运行项目

您可以使用以下两个命令运行项目：

```bash
make build # 构建 mcp 服务器和 ui
make up    # 运行 openmemory mcp 服务器和 ui
```

运行这些命令后，您将拥有：

- OpenMemory MCP 服务器运行在：<http://localhost:8765（API> 文档可在 <http://localhost:8765/docs> 查看）
- OpenMemory UI 运行在：<http://localhost:3000>

#### UI 在 `localhost:3000` 无法正常工作？

如果 UI 在 [http://localhost:3000](http://localhost:3000) 无法正常启动，请尝试手动运行：

```bash
cd ui
pnpm install
pnpm dev
```

## 项目结构

- `api/` - 后端 API + MCP 服务器
- `ui/` - 前端 React 应用程序

## 贡献

我们是一个对 AI 和开源软件未来充满热情的开发者团队。凭借在这两个领域多年的经验，我们相信社区驱动开发的力量，并致力于构建让 AI 更加易用和个性化的工具。

我们欢迎各种形式的贡献：

- 错误报告和功能请求
- 文档改进
- 代码贡献
- 测试和反馈
- 社区支持

如何贡献：

1. Fork 仓库
2. 创建您的功能分支（`git checkout -b openmemory/feature/amazing-feature`）
3. 提交您的更改（`git commit -m 'Add some amazing feature'`）
4. 推送到分支（`git push origin openmemory/feature/amazing-feature`）
5. 开启 Pull Request

加入我们，共同构建 AI 记忆管理的未来！您的贡献将帮助 OpenMemory 为每个人变得更好。
