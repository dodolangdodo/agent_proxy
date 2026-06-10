# Puppeteer

> <https://github.com/modelcontextprotocol/servers-archived/tree/main/src/puppeteer>

一个模型上下文协议服务器，使用 Puppeteer 提供浏览器自动化功能。该服务器使 LLM 能够与网页交互、截图并在真实浏览器环境中执行 JavaScript。

> [!CAUTION]
> 该服务器可以访问本地文件和本地/内部 IP 地址，因为它在您的机器上运行浏览器。使用此 MCP 服务器时请谨慎，确保不会暴露任何敏感数据。

## 组件

### 工具

- **puppeteer_navigate**

  - 在浏览器中导航到任何 URL
  - 输入参数：
    - `url`（字符串，必需）：要导航到的 URL
    - `launchOptions`（对象，可选）：PuppeteerJS 启动选项。默认为 null。如果更改且不为 null，浏览器将重启。示例：`{ headless: true, args: ['--user-data-dir="C:/Data"'] }`
    - `allowDangerous`（布尔值，可选）：允许降低安全性的危险启动选项。当为 false 时，危险参数如 `--no-sandbox`、`--disable-web-security` 将抛出错误。默认为 false。

- **puppeteer_screenshot**

  - 捕获整个页面或特定元素的截图
  - 输入参数：
    - `name`（字符串，必需）：截图名称
    - `selector`（字符串，可选）：要截图的元素的 CSS 选择器
    - `width`（数字，可选，默认：800）：截图宽度
    - `height`（数字，可选，默认：600）：截图高度
    - `encoded`（布尔值，可选）：如果为 true，将截图捕获为 base64 编码的数据 URI（作为文本）而不是二进制图像内容。默认为 false。

- **puppeteer_click**

  - 点击页面上的元素
  - 输入参数：`selector`（字符串）：要点击的元素的 CSS 选择器

- **puppeteer_hover**

  - 悬停页面上的元素
  - 输入参数：`selector`（字符串）：要悬停的元素的 CSS 选择器

- **puppeteer_fill**

  - 填写输入字段
  - 输入参数：
    - `selector`（字符串）：输入字段的 CSS 选择器
    - `value`（字符串）：要填入的值

- **puppeteer_select**

  - 选择带有 SELECT 标签的元素
  - 输入参数：
    - `selector`（字符串）：要选择的元素的 CSS 选择器
    - `value`（字符串）：要选择的值

- **puppeteer_evaluate**
  - 在浏览器控制台中执行 JavaScript
  - 输入参数：`script`（字符串）：要执行的 JavaScript 代码

### 资源

服务器提供两种类型的资源访问：

1. **控制台日志**（`console://logs`）

   - 以文本格式显示浏览器控制台输出
   - 包括来自浏览器的所有控制台消息

2. **截图**（`screenshot://<name>`）
   - 捕获的截图的 PNG 图像
   - 通过捕获时指定的截图名称访问

## 主要功能

- 浏览器自动化
- 控制台日志监控
- 截图功能
- JavaScript 执行
- 基本网页交互（导航、点击、表单填写）
- 可自定义的 Puppeteer 启动选项

## Puppeteer 服务器配置使用

### 与 Claude Desktop 一起使用

以下是使用 Puppeteer 服务器的 Claude Desktop 配置：

### Docker

**注意** Docker 实现将使用无头 Chromium，而 NPX 版本将打开浏览器窗口。

```json
{
  "mcpServers": {
    "puppeteer": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "--init",
        "-e",
        "DOCKER_CONTAINER=true",
        "mcp/puppeteer"
      ]
    }
  }
}
```

### NPX

```json
{
  "mcpServers": {
    "puppeteer": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-puppeteer"]
    }
  }
}
```

### 与 VS Code 一起使用

快速安装，请使用下面的一键安装按钮...

[![在 VS Code 中使用 NPX 安装](https://img.shields.io/badge/VS_Code-NPM-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=puppeteer&config=%7B%22command%22%3A%22npx%22%2C%22args%22%3A%5B%22-y%22%2C%22%40modelcontextprotocol%2Fserver-puppeteer%22%5D%7D) [![在 VS Code Insiders 中使用 NPX 安装](https://img.shields.io/badge/VS_Code_Insiders-NPM-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=puppeteer&config=%7B%22command%22%3A%22npx%22%2C%22args%22%3A%5B%22-y%22%2C%22%40modelcontextprotocol%2Fserver-puppeteer%22%5D%7D&quality=insiders)

[![在 VS Code 中使用 Docker 安装](https://img.shields.io/badge/VS_Code-Docker-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=puppeteer&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%22-i%22%2C%22--rm%22%2C%22--init%22%2C%22-e%22%2C%22DOCKER_CONTAINER%3Dtrue%22%2C%22mcp%2Fpuppeteer%22%5D%7D) [![在 VS Code Insiders 中使用 Docker 安装](https://img.shields.io/badge/VS_Code_Insiders-Docker-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=puppeteer&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%22-i%22%2C%22--rm%22%2C%22--init%22%2C%22-e%22%2C%22DOCKER_CONTAINER%3Dtrue%22%2C%22mcp%2Fpuppeteer%22%5D%7D&quality=insiders)

手动安装，请将以下 JSON 块添加到 VS Code 的用户设置（JSON）文件中。您可以按 `Ctrl + Shift + P` 并输入 `Preferences: Open User Settings (JSON)` 来完成此操作。

可选地，您可以将其添加到工作区中名为 `.vscode/mcp.json` 的文件中。这将允许您与他人共享配置。

> 注意，在 `.vscode/mcp.json` 文件中不需要 `mcp` 键。

NPX 安装（打开浏览器窗口）：

```json
{
  "mcp": {
    "servers": {
      "puppeteer": {
        "command": "npx",
        "args": ["-y", "@modelcontextprotocol/server-puppeteer"]
      }
    }
  }
}
```

Docker 安装（使用无头 Chromium）：

```json
{
  "mcp": {
    "servers": {
      "puppeteer": {
        "command": "docker",
        "args": [
          "run",
          "-i",
          "--rm",
          "--init",
          "-e",
          "DOCKER_CONTAINER=true",
          "mcp/puppeteer"
        ]
      }
    }
  }
}
```

### 启动选项

您可以通过两种方式自定义 Puppeteer 的浏览器行为：

1. **环境变量**：在 MCP 配置的 `env` 参数中设置带有 JSON 编码字符串的 `PUPPETEER_LAUNCH_OPTIONS`：

   ```json
   {
     "mcpServers": {
       "mcp-puppeteer": {
         "command": "npx",
         "args": ["-y", "@modelcontextprotocol/server-puppeteer"],
         "env": {
           "PUPPETEER_LAUNCH_OPTIONS": "{ \"headless\": false, \"executablePath\": \"C:/Program Files/Google/Chrome/Application/chrome.exe\", \"args\": [] }",
           "ALLOW_DANGEROUS": "true"
         }
       }
     }
   }
   ```

2. **工具调用参数**：将 `launchOptions` 和 `allowDangerous` 参数传递给 `puppeteer_navigate` 工具：

   ```json
   {
     "url": "https://example.com",
     "launchOptions": {
       "headless": false,
       "defaultViewport": { "width": 1280, "height": 720 }
     }
   }
   ```

## 构建

Docker 构建：

```bash
docker build -t mcp/puppeteer -f src/puppeteer/Dockerfile .
```

## 许可证

此 MCP 服务器采用 MIT 许可证。这意味着您可以自由使用、修改和分发软件，但需遵守 MIT 许可证的条款和条件。更多详情，请参阅项目仓库中的 LICENSE 文件。
