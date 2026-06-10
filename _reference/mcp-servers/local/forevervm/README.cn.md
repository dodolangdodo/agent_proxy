# ForeverVM MCP 服务器

ForeverVM 的 MCP 服务器，使 Claude 能够在 Python REPL 中执行代码。

## 工具

1. `create-python-repl`

- 创建一个 Python REPL。
- 返回：新 REPL 的 ID。

2. `run-python-in-repl`
   - 在 Python REPL 中执行代码。
   - 必需输入：
     - `code`（字符串）：Python REPL 将运行的代码。
     - `replId`（字符串）：运行代码的 REPL ID。
   - 返回：执行代码的结果。

## 与 Claude Desktop 一起使用

运行以下命令：

```bash
npx forevervm-mcp install --claude
```

对于其他 MCP 客户端，请参阅[文档](https://forevervm.com/docs/guides/forevervm-mcp-server/)。

## 本地安装（仅用于开发）

在 MCP 客户端中，将命令设置为 `npm`，参数设置为：

```json
["--prefix", "<path/to/this/directory>", "run", "start", "run"]
```
