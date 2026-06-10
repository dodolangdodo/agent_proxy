# Markdownify MCP 服务器

> 求助！我需要有 Windows 电脑的人帮助我为 Markdownify-MCP 添加 Windows 支持。已有 PR 但我无法测试。如有兴趣请在[这里](https://github.com/zcaceres/markdownify-mcp/issues/18)留言。

Markdownify 是一个模型上下文协议 (MCP) 服务器，可以将各种文件类型和网页内容转换为 Markdown 格式。它提供了一套工具，可以将 PDF、图片、音频文件、网页等转换为易于阅读和分享的 Markdown 文本。

<a href="https://glama.ai/mcp/servers/bn5q4b0ett"><img width="380" height="200" src="https://glama.ai/mcp/servers/bn5q4b0ett/badge" alt="Markdownify Server MCP server" /></a>

## 功能特性

- 将多种文件类型转换为 Markdown：
  - PDF
  - 图片
  - 音频（带转录功能）
  - DOCX
  - XLSX
  - PPTX
- 将网页内容转换为 Markdown：
  - YouTube 视频转录
  - 必应搜索结果
  - 一般网页
- 检索现有的 Markdown 文件

## 快速开始

1. 克隆此仓库
2. 安装依赖：

   ```
   pnpm install
   ```

注意：这也会安装 `uv` 和相关的 Python 依赖。

3. 构建项目：

   ```
   pnpm run build
   ```

4. 启动服务器：

   ```
   pnpm start
   ```

## 开发

- 使用 `pnpm run dev` 以监视模式启动 TypeScript 编译器
- 修改 `src/server.ts` 来自定义服务器行为
- 在 `src/tools.ts` 中添加或修改工具

## 与桌面应用集成

要将此服务器与桌面应用集成，请在应用的服务器配置中添加以下内容：

```js
{
  "mcpServers": {
    "markdownify": {
      "command": "node",
      "args": [
        "{此处填写绝对路径}/dist/index.js"
      ],
      "env": {
        // 默认情况下，服务器将使用 `uv` 的默认安装位置
        "UV_PATH": "/path/to/uv"
      }
    }
  }
}
```

## 可用工具

- `youtube-to-markdown`：将 YouTube 视频转换为 Markdown
- `pdf-to-markdown`：将 PDF 文件转换为 Markdown
- `bing-search-to-markdown`：将必应搜索结果转换为 Markdown
- `webpage-to-markdown`：将网页转换为 Markdown
- `image-to-markdown`：将图片转换为带元数据的 Markdown
- `audio-to-markdown`：将音频文件转换为带转录的 Markdown
- `docx-to-markdown`：将 DOCX 文件转换为 Markdown
- `xlsx-to-markdown`：将 XLSX 文件转换为 Markdown
- `pptx-to-markdown`：将 PPTX 文件转换为 Markdown
- `get-markdown-file`：检索现有的 Markdown 文件。文件扩展名必须以：*.md、*.markdown 结尾。
  
  可选：设置 `MD_SHARE_DIR` 环境变量来限制可检索文件的目录，例如 `MD_SHARE_DIR=[某个路径] pnpm run start`

## 贡献

欢迎贡献！请随时提交 Pull Request。

## 许可证

本项目采用 MIT 许可证 - 详情请参阅 [LICENSE](LICENSE) 文件。
