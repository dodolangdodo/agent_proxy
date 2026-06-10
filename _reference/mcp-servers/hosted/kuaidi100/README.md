# Kuaidi100 MCP Server

Kuaidi100 MCP Server is a [Model Context Protocol (MCP)](https://modelcontextprotocol.io/introduction) server that provides comprehensive express delivery tracking services through the Kuaidi100 API.

## Usage

### Register for Kuaidi100 API Open Platform

To use Kuaidi100 services, you need to first [register](https://api.kuaidi100.com/register) for a Kuaidi100 account. If you have already registered, you can skip directly to step two.

### Obtain Authorization Key

After registering your Kuaidi100 account, go to the [Enterprise Management Console](https://api.kuaidi100.com/manager/v2/query/overview). You can find the Kuaidi100 authorization key in the authorization parameters section.

### Configure MCP Service

You need to use an MCP client to access the MCP service. Kuaidi100 supports all MCP clients that can use MCP-SSE mode. You need to configure it in the following format:

```json
{
  "mcpServers": {
    "kuaidi100-server": {
      "url": "http://api.kuaidi100.com/mcp/sse?key=***********"
    }
  }
}
```

You need to replace the key in the URL with the key you obtained from the Kuaidi100 Enterprise Management Console. For example, if the key you obtained from the Kuaidi100 Enterprise Management Console is A1234321, the URL address you need to configure would be <http://api.kuaidi100.com/mcp/sse?key=A1234321>
