# Gezhe-MCP-server

## Introduction

Gezhe PPT MCP server, can generate PPTs based on topics.

### Tools

1. `generate_ppt_by_topic`
   - Input:
     - `topic` (string): Topic name
   - Returns: Preview link

## Usage Guide

### Method 1: Streamable HTTP

1. Visit and log in to <https://pro.gezhe.com/settings>
2. Go to the "Settings - MCP Server" page and copy the URL provided on the page.

### Method 2: Local Execution

1. Visit and log in to <https://gezhe.com/>
2. Go to the "Settings - MCP Server" page and copy the URL provided on the page.
3. Copy the following configuration and fill it into Cherry Studio, Cursor, etc.

```json
{
  "mcpServers": {
    "Gezhe PPT": {
      "command": "npx",
      "args": ["-y", "gezhe-mcp-server@latest"],
      "env": {
        "API_KEY": "Replace with the API_KEY you get"
      }
    }
  }
}
```
