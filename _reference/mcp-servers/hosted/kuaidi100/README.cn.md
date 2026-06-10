# 快递100 MCP Server

快递100 MCP Server 是一个 [Model Context Protocol (MCP)](https://modelcontextprotocol.io/introduction) 服务器，通过快递100 API 提供全面的快递跟踪服务。

## 使用方法

### 注册快递100API开放平台

为了使用快递100服务，您需要先[注册](https://api.kuaidi100.com/register)一个快递100账户，如果已经注册过，可以直接跳转到第二步。

### 获取授权key

注册快递100账户后，进入[企业管理后台](https://api.kuaidi100.com/manager/v2/query/overview)。在授权参数位置可以看到快递100的授权key。

### 配置MCP服务

您需要通过MCP客户端才可以使用MCP服务，快递100支持所有可以使用MCP-SSE模式的MCP客户端。您需要按以下格式进行配置。

```json
{
  "mcpServers": {
    "kuaidi100-server": {
      "url": "http://api.kuaidi100.com/mcp/sse?key=***********"
    }
  }
}
```

需要根据您在快递100企业管理后台中获取的key，把url中的key进行替换。例如您在快递100企业管理后台中获取的key为A1234321，您需要配置的url地址为<http://api.kuaidi100.com/mcp/sse?key=A1234321>
