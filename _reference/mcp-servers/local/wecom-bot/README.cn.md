# WeCom Bot MCP 服务器

<div align="center">
    <img src="wecom.png" alt="WeCom Bot Logo" width="200"/>
</div>

一个符合模型上下文协议 (MCP) 的企业微信机器人服务器实现。

<a href="https://glama.ai/mcp/servers/amr2j23lbk"><img width="380" height="200" src="https://glama.ai/mcp/servers/amr2j23lbk/badge" alt="WeCom Bot Server MCP server" /></a>

## 功能特性

- 支持多种消息类型：
  - 文本消息
  - Markdown 消息
  - 图片消息（base64）
  - 文件消息
- @提醒支持（通过用户 ID 或手机号）
- 消息历史记录追踪
- 可配置的日志系统
- 完整的类型注解
- 基于 Pydantic 的数据验证

## 系统要求

- Python 3.10+
- 企业微信机器人 Webhook URL（从企业微信群组设置中获取）

## 安装

有多种方式安装 WeCom Bot MCP 服务器：

### 1. 自动安装（推荐）

#### 使用 Smithery（适用于 Claude Desktop）

```bash
npx -y @smithery/cli install wecom-bot-mcp-server --client claude
```

#### 使用 VSCode 配合 Cline 扩展

1. 从 VSCode 市场安装 [Cline 扩展](https://marketplace.visualstudio.com/items?itemName=saoudrizwan.claude-dev)
2. 打开命令面板（Ctrl+Shift+P / Cmd+Shift+P）
3. 搜索 "Cline: Install Package"
4. 输入 "wecom-bot-mcp-server" 并按回车

### 2. 手动安装

#### 从 PyPI 安装

```bash
pip install wecom-bot-mcp-server
```

#### 手动配置 MCP

创建或更新您的 MCP 配置文件：

```json
// Windsurf 配置文件: ~/.windsurf/config.json
{
  "mcpServers": {
    "wecom": {
      "command": "uvx",
      "args": [
        "wecom-bot-mcp-server"
      ],
      "env": {
        "WECOM_WEBHOOK_URL": "your-webhook-url"
      }
    }
  }
}
```

## 配置

### 设置环境变量

```bash
# Windows PowerShell
$env:WECOM_WEBHOOK_URL = "your-webhook-url"

# 可选配置
$env:MCP_LOG_LEVEL = "DEBUG"  # 日志级别: DEBUG, INFO, WARNING, ERROR, CRITICAL
$env:MCP_LOG_FILE = "path/to/custom/log/file.log"  # 自定义日志文件路径
```

### 日志管理

日志系统使用 `platformdirs.user_log_dir()` 进行跨平台日志文件管理：

- Windows: `C:\Users\<username>\AppData\Local\hal\wecom-bot-mcp-server`
- Linux: `~/.local/share/hal/wecom-bot-mcp-server`
- macOS: `~/Library/Application Support/hal/wecom-bot-mcp-server`

日志文件名为 `mcp_wecom.log`，存储在上述目录中。

## 使用方法

### 启动服务器

```bash
wecom-bot-mcp-server
```

### 使用示例（配合 MCP）

```python
# 场景 1: 发送天气信息到企业微信
用户: "深圳今天天气怎么样？发送到企业微信"
助手: "我来查看深圳的天气并发送到企业微信"

await mcp.send_message(
    content="深圳天气:\n- 温度: 25°C\n- 天气: 晴天\n- 空气质量: 良好",
    msg_type="markdown"
)

# 场景 2: 发送会议提醒并@相关人员
用户: "发送下午3点项目评审会议提醒，提醒张三和李四参加"
助手: "我来发送会议提醒"

await mcp.send_message(
    content="## 项目评审会议提醒\n\n时间: 今天下午3:00\n地点: 会议室A\n\n请准时参加！",
    msg_type="markdown",
    mentioned_list=["zhangsan", "lisi"]
)

# 场景 3: 发送文件
用户: "把这个周报发送到企业微信群"
助手: "我来发送周报"

await mcp.send_message(
    content=Path("weekly_report.docx"),
    msg_type="file"
)
```

### 直接 API 使用

#### 发送消息

```python
from wecom_bot_mcp_server import mcp

# 发送 markdown 消息
await mcp.send_message(
    content="**你好世界！**", 
    msg_type="markdown"
)

# 发送文本消息并@用户
await mcp.send_message(
    content="你好 @user1 @user2",
    msg_type="text",
    mentioned_list=["user1", "user2"]
)
```

#### 发送文件

```python
from wecom_bot_mcp_server import send_wecom_file

# 发送文件
await send_wecom_file("/path/to/file.txt")
```

#### 发送图片

```python
from wecom_bot_mcp_server import send_wecom_image

# 发送本地图片
await send_wecom_image("/path/to/image.png")

# 发送网络图片
await send_wecom_image("https://example.com/image.png")
```

## 开发

### 设置开发环境

1. 克隆仓库：

```bash
git clone https://github.com/loonghao/wecom-bot-mcp-server.git
cd wecom-bot-mcp-server
```

2. 创建虚拟环境并安装依赖：

```bash
# 使用 uv（推荐）
pip install uv
uv venv
uv pip install -e ".[dev]"

# 或使用传统方法
python -m venv venv
source venv/bin/activate  # Windows 系统: venv\Scripts\activate
pip install -e ".[dev]"
```

### 测试

```bash
# 使用 uv（推荐）
uvx nox -s pytest

# 或使用传统方法
nox -s pytest
```

### 代码风格

```bash
# 检查代码
uvx nox -s lint

# 自动修复代码风格问题
uvx nox -s lint_fix
```

### 构建和发布

```bash
# 构建包
uv build

# 构建并发布到 PyPI
uv build && twine upload dist/*
```

## 项目结构

```
wecom-bot-mcp-server/
├── src/
│   └── wecom_bot_mcp_server/
│       ├── __init__.py
│       ├── server.py
│       ├── message.py
│       ├── file.py
│       ├── image.py
│       ├── utils.py
│       └── errors.py
├── tests/
│   ├── test_server.py
│   ├── test_message.py
│   ├── test_file.py
│   └── test_image.py
├── docs/
├── pyproject.toml
├── noxfile.py
└── README.md
```

## 许可证

本项目采用 MIT 许可证 - 详情请见 [LICENSE](LICENSE) 文件。

## 联系方式

- 作者：longhao
- 邮箱：<hal.long@outlook.com>
