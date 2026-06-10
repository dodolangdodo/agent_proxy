# ScreenshotOne MCP 服务器

[ScreenshotOne](https://screenshotone.com) 的官方 [MCP (模型上下文协议)](https://modelcontextprotocol.io/) 服务器实现。

[关于为什么构建此服务器以及对 MCP 未来的一些思考](https://screenshotone.com/blog/mcp-server/)。

<a href="https://glama.ai/mcp/servers/nq85q0596a">
  <img width="380" height="200" src="https://glama.ai/mcp/servers/nq85q0596a/badge" alt="ScreenshotOne Server MCP server" />
</a>

## 工具

- `render-website-screenshot`: 渲染网站截图并将其作为图像返回。

## 使用方法

### 构建项目

首先安装依赖并构建项目：

```bash
npm install && npm run build
```

### 获取您的 ScreenshotOne API 密钥

在 [ScreenshotOne](https://screenshotone.com) 注册并获取您的 API 密钥。

### 与 Claude for Desktop 一起使用

将以下内容添加到您的 `~/Library/Application\ Support/Claude/claude_desktop_config.json` 文件中：

```json
{
    "mcpServers": {
        "screenshotone": {
            "command": "node",
            "args": ["path/to/screenshotone/mcp/build/index.js"],
            "env": {
                "SCREENSHOTONE_API_KEY": "<您的 API 密钥>"
            }
        }
    }
}
```

### 独立运行或用于其他项目

```bash
SCREENSHOTONE_API_KEY=your_api_key && node build/index.js
```
