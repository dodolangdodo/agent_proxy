# 天气 MCP 服务器

[![smithery badge](https://smithery.ai/badge/@CodeByWaqas/weather-mcp-server)](https://smithery.ai/server/@CodeByWaqas/weather-mcp-server)

一个使用 OpenWeatherMap API 提供天气信息的现代代码协议（MCP）服务器。

## 功能特性

- 实时天气数据获取
- 温度采用公制单位
- 详细的天气信息，包括：
  - 温度
  - 湿度
  - 风速
  - 日出/日落时间
  - 天气描述

## 系统要求

- Python 3.12 或更高版本
- OpenWeatherMap API 密钥

## 安装方法

### 通过 Smithery 安装

通过 [Smithery](https://smithery.ai/server/@CodeByWaqas/weather-mcp-server) 自动为 Claude Desktop 安装天气 MCP 服务器：

```bash
npx -y @smithery/cli install @CodeByWaqas/weather-mcp-server --client claude
```

### 手动安装

1. 克隆仓库
2. 创建虚拟环境：

```bash
python -m venv .venv
source .venv/bin/activate  # Windows 系统使用：.venv\Scripts\activate
```

3. 安装依赖：

```bash
pip install -e .
```

## 配置说明

### 在 Claude Desktop 中配置

```json
# claude_desktop_config.json
# 可通过以下路径找到配置文件位置：
# Claude -> Settings -> Developer -> Edit Config
{
  "mcpServers": {
      "mcp-weather-project": {
          "command": "uv",
          "args": [
              "--directory",
              "/<绝对路径>/weather-mcp-server/src/resources",
              "run",
              "server.py"
          ],
          "env": {
            "WEATHER_API_KEY": "您的API密钥"
          }
      }
  }
}
```

## 本地/开发环境配置说明

### 克隆仓库

`git clone https://github.com/CodeByWaqas/weather-mcp-server`

### 安装依赖

安装 MCP 服务器依赖：

```bash
cd weather-mcp-server

# 创建虚拟环境并激活
uv venv

source .venv/bin/activate # MacOS/Linux
# 或者
.venv/Scripts/activate # Windows

# 安装依赖
uv add "mcp[cli]" python-dotenv requests httpx
```

## 配置

1. 将 `src/resources/env.example` 复制为 `src/resources/.env`
2. 在 `.env` 文件中添加您的 OpenWeatherMap API 密钥：

```
WEATHER_API_KEY=您的API密钥
```

## 使用方法

运行 Claude Desktop 并使用 LLM 获取天气信息

## 许可证

本项目采用 MIT 许可证 - 详情请参见 LICENSE 文件。
