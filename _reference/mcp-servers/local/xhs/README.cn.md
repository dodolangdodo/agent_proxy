# 项目说明

## 安装

### 安装 ChromeDriver

1. 查找你的 Chrome 浏览器版本（例如 "134.0.6998.166"）
2. 运行命令下载对应版本：

   ```bash
   npx @puppeteer/browsers install chromedriver@134.0.6998.166
   ```

3. 将 ChromeDriver 复制到系统路径或添加路径到环境变量中

## 登录

在终端运行以下命令（请使用绝对路径指定 `PATH_TO_STORE_YOUR_COOKIES`，例如 `/Users/Bruce/`。该 MCP 服务器会将你的 cookie 存储在此路径）：

```bash
env phone=YOUR_PHONE_NUMBER json_path=PATH_TO_STORE_YOUR_COOKIES uvx --from xhs_mcp_server@latest login
```

终端会显示：

```
无效的 cookies，已清理
请输入验证码:
```

此时需在终端输入接收到的验证码并回车。

## 验证登录

```bash
env phone=YOUR_PHONE_NUMBER json_path=PATH_TO_STORE_YOUR_COOKIES uvx --from xhs_mcp_server@latest login
```

若成功会显示：

```
使用 cookies 登录成功
```

## 检查工具（Inspector）

在终端启动检查工具：

```bash
npx @modelcontextprotocol/inspector -e phone=YOUR_PHONE_NUMBER -e json_path=PATH_TO_STORE_YOUR_COOKIES uvx xhs_mcp_server@latest
```

在检查工具中，你可以使用本地图片：

- 输入图片路径（例如 `["C:\路径\到\你的\图片.jpg"]`），图片路径需要用双引号包裹。

> **注：** 发送时可能会显示 "Error Request timed out" 警告，但实际帖子仍会发布成功。

## 启动服务器

### 方式1：直接运行命令

```bash
env phone=YOUR_PHONE_NUMBER json_path=PATH_TO_STORE_YOUR_COOKIES uvx xhs_mcp_server@latest
```

### 方式2：配置文件设置

在配置文件中添加以下内容：

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

## 注意事项

此 MCP 服务器仅限研究用途，禁止用于商业目的。
