# Playwright Browserbase MCP 服务器

> <https://github.com/browserbase/mcp-server-browserbase/tree/main/browserbase>

模型上下文协议（MCP）是一个开放协议，能够在LLM应用程序与外部数据源和工具之间实现无缝集成。无论您是在构建AI驱动的IDE、增强聊天界面，还是创建自定义AI工作流，MCP都提供了一种标准化的方式来连接LLM与其所需的上下文。

## 如何在MCP JSON中设置

您可以使用我们托管在NPM上的服务器，或者通过克隆此仓库完全在本地运行。

### 在NPM上运行（推荐）

进入您的MCP配置JSON并添加Browserbase服务器：

```json
{
   "mcpServers": {
      "browserbase": {
         "command": "npx",
         "args" : ["@browserbasehq/mcp"],
         "env": {
            "BROWSERBASE_API_KEY": "",
            "BROWSERBASE_PROJECT_ID": ""
         }
      }
   }
}
```

就是这样！重新加载您的MCP客户端，Claude就能够使用Browserbase了。

### 100%本地运行

```bash
# 克隆仓库
git clone https://github.com/browserbase/mcp-server-browserbase.git

# 在正确的目录中安装依赖项并构建项目
cd browserbase
npm install && npm run build
```

然后在您的MCP配置JSON中运行服务器。要在本地运行，我们可以使用STDIO或通过SSE自托管。

### STDIO

在您的MCP配置JSON文件中添加以下内容：

```json
{
"mcpServers": {
   "browserbase": {
      "command" : "node",
      "args" : ["/path/to/mcp-server-browserbase/browserbase/cli.js"],
      "env": {
         "BROWSERBASE_API_KEY": "",
         "BROWSERBASE_PROJECT_ID": ""
         }
      }
   }
}
```

### SSE

在终端中运行以下命令。您可以添加任何标志（见下面的选项）来自定义您的配置。

```bash
   node cli.js --port 8931
```

然后在您的MCP配置JSON文件中添加以下内容：

```json
   {
      "mcpServers": {
         "browserbase": {
            "url": "http://localhost:8931/sse",
            "env": {
               "BROWSERBASE_API_KEY": "",
               "BROWSERBASE_PROJECT_ID": ""
            }
         }
      }
   }
```

然后重新加载您的MCP客户端，您就可以开始使用了！

## 标志说明

Browserbase MCP服务器接受以下命令行标志：

| 标志 | 描述 |
|------|-------------|
| `--browserbaseApiKey <key>` | 用于身份验证的Browserbase API密钥 |
| `--browserbaseProjectId <id>` | 您的Browserbase项目ID |
| `--proxies` | 为会话启用Browserbase代理 |
| `--advancedStealth` | 启用Browserbase高级隐身模式（仅限Scale计划用户） |
| `--contextId <contextId>` | 指定要使用的Browserbase上下文ID |
| `--persist [boolean]` | 是否持久化Browserbase上下文（默认：true） |
| `--port <port>` | HTTP/SSE传输监听端口 |
| `--host <host>` | 服务器绑定主机（默认：localhost，使用0.0.0.0表示所有接口） |
| `--cookies [json]` | 要注入到浏览器中的cookies的JSON数组 |
| `--browserWidth <width>` | 浏览器视口宽度（默认：1024） |
| `--browserHeight <height>` | 浏览器视口高度（默认：768） |

这些标志可以直接传递给CLI或在您的MCP配置文件中配置。

### 注意

目前，这些标志只能与本地服务器（npx @browserbasehq/mcp）一起使用。

____

## 标志和示例配置

### 代理

这里是我们关于[代理](https://docs.browserbase.com/features/proxies)的文档。

要在STDIO中使用代理，请在您的MCP配置中设置--proxies标志：

```json
{
   "mcpServers": {
      "browserbase": {
         "command" : "npx",
         "args" : ["@browserbasehq/mcp", "--proxies"],
         "env": {
            "BROWSERBASE_API_KEY": "",
            "BROWSERBASE_PROJECT_ID": ""
         }
      }
   }
}
```

### 高级隐身

这里是我们关于[高级隐身](https://docs.browserbase.com/features/stealth-mode#advanced-stealth-mode)的文档。

要在STDIO中使用高级隐身，请在您的MCP配置中设置--advancedStealth标志：

```json
{
   "mcpServers": {
      "browserbase": {
         "command" : "npx",
         "args" : ["@browserbasehq/mcp", "--advancedStealth"],
         "env": {
            "BROWSERBASE_API_KEY": "",
            "BROWSERBASE_PROJECT_ID": ""
         }
      }
   }
}
```

### 上下文

这里是我们关于[上下文](https://docs.browserbase.com/features/contexts)的文档。

要在STDIO中使用上下文，请在您的MCP配置中设置--contextId标志：

```json
{
   "mcpServers": {
      "browserbase": {
         "command" : "npx",
         "args" : ["@browserbasehq/mcp", "--contextId", "<YOUR_CONTEXT_ID>"],
         "env": {
            "BROWSERBASE_API_KEY": "",
            "BROWSERBASE_PROJECT_ID": ""
         }
      }
   }
}
```

### Cookie注入

为什么需要注入cookies？我们的上下文API目前适用于持久性cookies，但不适用于会话cookies。所以有时我们的持久身份验证可能不起作用（我们正在努力添加此功能）。

您可以通过将cookies.json添加到您的MCP配置中来将cookies标记到MCP中。

要在STDIO中使用代理，请在您的MCP配置中设置--proxies标志。您的cookies JSON必须是[Playwright Cookies](https://playwright.dev/docs/api/class-browsercontext#browser-context-cookies)类型：

```json
{
   "mcpServers": {
      "browserbase" {
         "command" : "npx",
         "args" : [
            "@browserbasehq/mcp", "--cookies", 
            '{
               "cookies": json,
            }'
         ],
         "env": {
            "BROWSERBASE_API_KEY": "",
            "BROWSERBASE_PROJECT_ID": ""
         }
      }
   }
}
```

### 浏览器视口大小

浏览器会话的默认视口大小为1024 x 768。您可以使用browserWidth和browserHeight标志调整浏览器视口大小。

以下是如何使用自定义浏览器大小。我们建议坚持使用16:9的宽高比（即：1920 x 1080、1280 x 720、1024 x 768）：

```json
{
   "mcpServers": {
      "browserbase": {
         "command" : "npx",
         "args" : [
            "@browserbasehq/mcp",
            "--browserHeight 1080",
            "--browserWidth 1920",
         ],
         "env": {
            "BROWSERBASE_API_KEY": "",
            "BROWSERBASE_PROJECT_ID": ""
         }
      }
   }
}
```

## 结构

* `src/`: TypeScript源代码
  * `index.ts`: 主入口点，环境检查，关闭
  * `server.ts`: MCP服务器设置和请求路由
  * `sessionManager.ts`: 处理Browserbase会话创建/管理
  * `tools/`: 工具定义和实现
  * `resources/`: 资源（截图）处理
  * `types.ts`: 共享TypeScript类型
* `dist/`: 编译的JavaScript输出
* `tests/`: 测试占位符
* `utils/`: 实用程序脚本占位符
* `Dockerfile`: 用于构建Docker镜像
* 配置文件（`.json`、`.ts`、`.mjs`、`.npmignore`）

## 持久化上下文

此服务器支持Browserbase的上下文功能，允许在浏览器会话之间持久化cookies、身份验证和缓存数据：

1. **创建上下文**：

   ```
   browserbase_context_create: 创建新上下文，可选择使用友好名称
   ```

2. **在会话中使用上下文**：

   ```
   browserbase_session_create: 现在接受'context'参数：
     - id: 要使用的上下文ID
     - name: ID的替代方案，上下文的友好名称
     - persist: 是否将更改（cookies、缓存）保存回上下文（默认：true）
   ```

3. **删除上下文**：

   ```
   browserbase_context_delete: 当您不再需要时删除上下文
   ```

上下文使以下操作变得更加容易：

* 在会话之间维护登录状态
* 通过保留缓存减少页面加载时间
* 通过重用浏览器指纹避免验证码和检测

## Cookie管理

此服务器还提供直接的cookie管理功能：

1. **添加Cookies**：

   ```
   browserbase_cookies_add: 向当前浏览器会话添加cookies，完全控制属性
   ```

2. **获取Cookies**：

   ```
   browserbase_cookies_get: 查看当前会话中的所有cookies（可选择按URL过滤）
   ```

3. **删除Cookies**：

   ```
   browserbase_cookies_delete: 删除特定cookies或清除会话中的所有cookies
   ```

这些工具对以下用途很有用：

* 无需导航到登录页面即可设置身份验证cookies
* 备份和恢复cookie状态
* 调试与cookie相关的问题
* 操作cookie属性（过期时间、安全标志等）

## 待办事项/路线图

* 为click、type、drag、hover、select_option实现真正的基于`ref`的交互逻辑。
* 使用`ref`实现元素特定的截图。
* 添加更多标准MCP工具（标签页、导航等）。
* 添加测试。
