# `workers-mcp`

> <https://github.com/cloudflare/workers-mcp>

> **让 Claude Desktop 与 Cloudflare Worker 对话！**

> [!WARNING]  
> 你应该从[这里](https://developers.cloudflare.com/agents/guides/remote-mcp-server/)开始 — 构建一个**远程** MCP 服务器
>
> 你可以[使用 mcp-remote](https://developers.cloudflare.com/agents/guides/test-remote-mcp-server/) 从 Claude Desktop、Cursor 和其他客户端连接到远程 MCP 服务器。

### 什么是 `workers-mcp`？

这个包提供了 CLI 工具和 Worker 内部逻辑，用于将 Claude Desktop（或任何 [MCP 客户端](https://modelcontextprotocol.io/)）连接到你账户中的 Cloudflare Worker，这样你就可以根据需要进行自定义。它通过构建步骤工作，可以将你的 Worker 的 TypeScript 方法转换，例如：

```ts
export class ExampleWorkerMCP extends WorkerEntrypoint<Env> {
  /**
   * 生成一个随机数。这个随机数特别随机，因为它必须一路传输到
   * 你最近的 Cloudflare PoP 来计算，这... 某种程度上与熔岩灯有关？
   *
   * @return {string} 包含超级随机数的消息
   * */
  async getRandomNumber() {
    return `你的随机数是 ${Math.random()}`
  }
  
  // ...等等
}
```

...转换为本地 Node.js 服务器可以向 MCP 客户端公开的 MCP 工具。Node.js 服务器充当代理，在本地处理 stdio 传输，并调用运行在 Cloudflare 上的 Worker 的相关方法。这允许你将应用程序中的任何函数或 API，或 [Cloudflare 开发者平台](https://developers.cloudflare.com/products/)中的任何服务，暴露给编码代理、Claude Desktop 或其他 MCP 客户端中的 LLM。

![image](https://github.com/user-attachments/assets/c16b2631-4eba-4914-8e26-d6ccea0fc578)

> <sub>是的，我知道 `Math.random()` 在 Worker 上的工作方式与在本地机器上相同，但别告诉 Claude</sub> 🤫

## 使用方法

### 步骤 1：生成新的 Worker

使用 `create-cloudflare` 生成新的 Worker。

```shell
npx create-cloudflare@latest my-new-worker
```

我建议选择 `Hello World` worker。

### 步骤 2：安装 `workers-mcp`

```shell
cd my-new-worker # 我总是忘记这一步
npm install workers-mcp
```

### 步骤 3：运行 `setup` 命令

```shell
npx workers-mcp setup
```

注意：如果出现问题，运行 `npx workers-mcp help`

### 步骤 4..♾️：迭代

更改 Worker 代码后，你只需要运行 `npm run deploy` 来同时更新 Claude 关于你函数的元数据和你的实时 Worker 实例。

但是，如果你更改了方法的名称、参数，或者添加/删除了方法，Claude 不会看到更新，直到你重启它。

你应该永远不需要重新运行 `npx workers-mcp install:claude`，但如果你想排除 Claude 配置作为错误源，这样做是安全的。

## 与其他 MCP 客户端一起使用

### Cursor

要让你的 Cloudflare MCP 服务器在 Cursor 中工作，你需要将配置文件中的 'command' 和 'args' 合并成单个字符串，并使用类型 'command'。

例如，如果你的配置文件如下所示：

```json
{
  "mcpServers": {
    "your-mcp-server-name": {
      "command": "/path/to/workers-mcp",
      "args": [
        "run",
        "your-mcp-server-name",
        "https://your-server-url.workers.dev",
        "/path/to/your/project"
      ],
      "env": {}
    }
  }
}
```

在 Cursor 中，创建一个 MCP 服务器条目：

* type: `command`
* command: `/path/to/workers-mcp run your-mcp-server-name https://your-server-url.workers.dev /path/to/your/project`

### 其他 MCP 客户端

对于 Windsurf 和其他 MCP 客户端，更新你的配置文件以包含你的 worker，这样你就可以直接从客户端使用这些工具：

```json
{
  "mcpServers": {
    "your-mcp-server-name": {
      "command": "/path/to/workers-mcp",
      "args": [
        "run",
        "your-mcp-server-name",
        "https://your-server-url.workers.dev",
        "/path/to/your/project"
      ],
      "env": {}
    }
  }
}
```

确保用你的实际服务器名称、URL 和项目路径替换占位符。

## 示例

查看 `examples` 目录以获取一些使用想法：

* `examples/01-hello-world` 是按照上述安装说明后的快照
* `examples/02-image-generation` 使用 Workers AI 运行 Flux 图像生成模型。Claude 非常擅长建议提示，实际上可以解释结果并决定尝试什么新提示来实现你想要的结果。
* TODO 浏览器渲染
* TODO Durable Objects
