# MCP 股票分析服务器 by Akshay Bavkar

这是一个MCP服务器，使用Yahoo Finance API提供实时和历史印度股票数据访问。它允许通过Claude Desktop、Cursor和其他兼容MCP的代理将股票数据检索用作本地LLM的上下文。

## 可用功能

- **getStockQuote**: 获取印度股票的当前报价。
- **getHistoricalData**: 获取印度股票的历史数据，支持自定义时间间隔和周期。

---

## 安装

```bash
npm install mcp-stock-analysis
```

## 在主机中使用

配置您的MCP客户端（例如Claude Desktop）连接到服务器：

```JSON
{
  "mcpServers": {
    "mcp-stock-analysis": {
      "command": "npx",
      "args": ["-y", "mcp-stock-analysis"],
    }
  }
}
```

## 工具

### `getStockQuote`

获取股票的当前报价。

输入：

`symbol`: 股票代码（例如：RELIANCE.NS）

输出：

```JSON
{
  "symbol": "RELIANCE.NS",
  "price": 2748.15,
  "name": "Reliance Industries Ltd"
}
```

### `getHistoricalData`

获取股票的历史数据。

输入：

- `symbol`: 股票代码（例如：RELIANCE.NS）
- `interval`: 数据的时间间隔（`daily`、`weekly`或`monthly`）（可选，默认：`daily`）

输出：

```JSON
{
    "date": "2025-03-21T00:00:00+05:30",
    "open": 2735,
    "high": 2750,
    "low": 2725,
    "close": 2748.15,
    "volume": 21780769
}
```

包含历史数据的JSON对象。输出结构取决于interval参数。

## 贡献

欢迎贡献！请提交issue或pull request。

### 许可证

MIT
