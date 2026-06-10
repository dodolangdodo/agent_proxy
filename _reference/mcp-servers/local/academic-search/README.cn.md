# 学术论文搜索 MCP 服务器

一个[模型上下文协议 (MCP)](https://www.anthropic.com/news/model-context-protocol) 服务器，可以从多个来源搜索和检索学术论文信息。

该服务器为大语言模型提供：

- 实时学术论文搜索功能
- 访问论文元数据和摘要
- 在可用时检索全文内容的能力
- 遵循 MCP 规范的结构化数据响应

虽然主要设计用于与 Anthropic 的 Claude Desktop 客户端集成，但 MCP 规范允许与其他支持工具/函数调用功能的 AI 模型和客户端潜在兼容（例如 OpenAI 的 API）。

**注意**：此软件正在积极开发中。功能和特性可能会发生变化。

<a href="https://glama.ai/mcp/servers/kzsu1zzz9j"><img width="380" height="200" src="https://glama.ai/mcp/servers/kzsu1zzz9j/badge" alt="学术论文搜索服务器 MCP 服务器" /></a>

## 功能特性

该服务器提供以下工具：

- `search_papers`：跨多个来源搜索学术论文
  - 参数：
    - `query`（字符串）：搜索查询文本
    - `limit`（整数，可选）：返回结果的最大数量（默认：10）
  - 返回：包含论文详情的格式化字符串
  
- `fetch_paper_details`：检索特定论文的详细信息
  - 参数：
    - `paper_id`（字符串）：论文标识符（DOI 或 Semantic Scholar ID）
    - `source`（字符串，可选）：数据源（"crossref" 或 "semantic_scholar"，默认："crossref"）
  - 返回：包含全面论文元数据的格式化字符串，包括：
    - 标题、作者、年份、DOI
    - 期刊、开放获取状态、PDF URL（仅限 Semantic Scholar）
    - 摘要和 TL;DR 总结（如有）

- `search_by_topic`：按主题搜索论文，可选择日期范围过滤
  - 参数：
    - `topic`（字符串）：搜索查询文本（限制为 300 个字符）
    - `year_start`（整数，可选）：日期范围的开始年份
    - `year_end`（整数，可选）：日期范围的结束年份
    - `limit`（整数，可选）：返回结果的最大数量（默认：10）
  - 返回：包含搜索结果的格式化字符串，包括：
    - 论文标题、作者和年份
    - 摘要和 TL;DR 总结（如有）
    - 期刊和开放获取信息

## 设置

### 通过 Smithery 安装

通过 [Smithery](https://smithery.ai/server/@afrise/academic-search-mcp-server) 自动为 Claude Desktop 安装学术论文搜索服务器：

```bash
npx -y @smithery/cli install @afrise/academic-search-mcp-server --client claude
```

***注意*** 此方法基本未经测试，因为他们的服务器似乎存在问题。您可以按照独立安装说明进行操作，直到 smithery 修复为止。

### 通过 uv 安装（手动安装）

1. 安装依赖项：

```sh
uv add "mcp[cli]" httpx
```

2. 在环境或 `.env` 文件中设置所需的 API 密钥：

```sh
# 这些实际上尚未实现
SEMANTIC_SCHOLAR_API_KEY=your_key_here 
CROSSREF_API_KEY=your_key_here  # 可选但推荐
```

3. 运行服务器：

```sh
uv run server.py
```

## 与 Claude Desktop 一起使用

1. 将服务器添加到您的 Claude Desktop 配置（`claude_desktop_config.json`）：

```json
{
  "mcpServers": {
    "academic-search": {
      "command": "uv",
      "args": ["run ", "/path/to/server/server.py"],
      "env": {
        "SEMANTIC_SCHOLAR_API_KEY": "your_key_here",
        "CROSSREF_API_KEY": "your_key_here"
      }
    }
  }
}
```

2. 重启 Claude Desktop

## 开发

该服务器使用以下技术构建：

- Python MCP SDK
- FastMCP 用于简化服务器实现
- httpx 用于 API 请求

## API 来源

- Semantic Scholar API
- Crossref API

## 许可证

本项目采用 GNU Affero 通用公共许可证 v3.0 (AGPL-3.0) 许可。此许可证确保：

- 您可以自由使用、修改和分发此软件
- 任何修改都必须在相同许可证下开源
- 任何使用此软件提供网络服务的人都必须提供源代码
- 允许商业使用，但软件和任何衍生作品必须保持免费和开源

请参阅 [LICENSE](LICENSE) 文件获取完整许可证文本。

## 贡献

欢迎贡献！以下是您可以帮助的方式：

1. Fork 仓库
2. 创建功能分支（`git checkout -b feature/amazing-feature`）
3. 提交您的更改（`git commit -m 'Add amazing feature'`）
4. 推送到分支（`git push origin feature/amazing-feature`）
5. 打开 Pull Request

请注意：

- 遵循现有的代码风格和约定
- 为任何新功能添加测试
- 根据需要更新文档
- 确保您的更改遵守 AGPL-3.0 许可证条款

通过为此项目做出贡献，您同意您的贡献将在 AGPL-3.0 许可证下获得许可。
