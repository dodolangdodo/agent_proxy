# Project Description

## Installation

### Install ChromeDriver

1. Find your Chrome browser version (e.g., "134.0.6998.166")
2. Run the command to download the corresponding version:

   ```bash
   npx @puppeteer/browsers install chromedriver@134.0.6998.166
   ```

3. Copy ChromeDriver to system path or add the path to environment variables

## Login

Run the following command in terminal (please use absolute path for `PATH_TO_STORE_YOUR_COOKIES`, e.g., `/Users/Bruce/`. This MCP server will store your cookies in this path):

```bash
env phone=YOUR_PHONE_NUMBER json_path=PATH_TO_STORE_YOUR_COOKIES uvx --from xhs_mcp_server@latest login
```

The terminal will display:

```
Invalid cookies, cleared
Please enter verification code:
```

At this point, you need to enter the received verification code in the terminal and press Enter.

## Verify Login

```bash
env phone=YOUR_PHONE_NUMBER json_path=PATH_TO_STORE_YOUR_COOKIES uvx --from xhs_mcp_server@latest login
```

If successful, it will display:

```
Login successful using cookies
```

## Inspector

Start the inspector tool in terminal:

```bash
npx @modelcontextprotocol/inspector -e phone=YOUR_PHONE_NUMBER -e json_path=PATH_TO_STORE_YOUR_COOKIES uvx xhs_mcp_server@latest
```

In the inspector tool, you can use local images:

- Enter image path (e.g., `["C:\path\to\your\image.jpg"]`), image paths need to be wrapped in double quotes.

> **Note:** You may see "Error Request timed out" warning when sending, but the post will still be published successfully.

## Start Server

### Method 1: Direct Command

```bash
env phone=YOUR_PHONE_NUMBER json_path=PATH_TO_STORE_YOUR_COOKIES uvx xhs_mcp_server@latest
```

### Method 2: Configuration File Setup

Add the following content to your configuration file:

```json
{
  "mcpServers": {
    "xhs-mcp-server": {
      "command": "uvx",
      "args": [
        "xhs_mcp_server@latest"
      ],
      "env": {
        "phone": "YOUR_PHONE_NUMBER",
        "json_path": "PATH_TO_STORE_YOUR_COOKIES"
      }
    }
  }
}
```

## Important Notes

This MCP server is for research purposes only and is prohibited for commercial use.
