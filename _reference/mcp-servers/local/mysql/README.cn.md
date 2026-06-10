![Tests](https://github.com/designcomputer/mysql_mcp_server/actions/workflows/test.yml/badge.svg)
![PyPI - Downloads](https://img.shields.io/pypi/dm/mysql-mcp-server)
[![smithery badge](https://smithery.ai/badge/mysql-mcp-server)](https://smithery.ai/server/mysql-mcp-server)
[![MseeP.ai Security Assessment Badge](https://mseep.net/mseep-audited.png)](https://mseep.ai/app/designcomputer-mysql-mcp-server)

# MySQL MCP 服务器

一个模型上下文协议（MCP）实现，可以安全地与 MySQL 数据库进行交互。该服务器组件通过受控接口促进 AI 应用程序（主机/客户端）与 MySQL 数据库之间的通信，使数据库探索和分析更加安全和结构化。

> **注意**：MySQL MCP 服务器不是设计为独立服务器使用，而是作为 AI 应用程序与 MySQL 数据库之间的通信协议实现。

## 功能特性

- 将可用的 MySQL 表列为资源
- 读取表内容
- 执行 SQL 查询并进行适当的错误处理
- 通过环境变量实现安全的数据库访问
- 全面的日志记录

## 安装

### 手动安装

```bash
pip install mysql-mcp-server
```

### 通过 Smithery 安装

通过 [Smithery](https://smithery.ai/server/mysql-mcp-server) 自动为 Claude Desktop 安装 MySQL MCP 服务器：

```bash
npx -y @smithery/cli install mysql-mcp-server --client claude
```

## 配置

设置以下环境变量：

```bash
MYSQL_HOST=localhost     # 数据库主机
MYSQL_PORT=3306         # 可选：数据库端口（如果未指定，默认为 3306）
MYSQL_USER=your_username
MYSQL_PASSWORD=your_password
MYSQL_DATABASE=your_database
```

## 使用方法

### 与 Claude Desktop 一起使用

将以下内容添加到您的 `claude_desktop_config.json`：

```json
{
  "mcpServers": {
    "mysql": {
      "command": "uv",
      "args": [
        "--directory",
        "path/to/mysql_mcp_server",
        "run",
        "mysql_mcp_server"
      ],
      "env": {
        "MYSQL_HOST": "localhost",
        "MYSQL_PORT": "3306",
        "MYSQL_USER": "your_username",
        "MYSQL_PASSWORD": "your_password",
        "MYSQL_DATABASE": "your_database"
      }
    }
  }
}
```

### 与 Visual Studio Code 一起使用

将以下内容添加到您的 `mcp.json`：

```json
{
  "servers": {
      "mysql": {
            "type": "stdio",
            "command": "uvx",
            "args": [
                "--from",
                "mysql-mcp-server",
                "mysql_mcp_server"
            ],
      "env": {
        "MYSQL_HOST": "localhost",
        "MYSQL_PORT": "3306",
        "MYSQL_USER": "your_username",
        "MYSQL_PASSWORD": "your_password",
        "MYSQL_DATABASE": "your_database"
      }
    }
  }
}
```

注意：需要安装 uv 才能正常工作

### 使用 MCP Inspector 进行调试

虽然 MySQL MCP 服务器不是设计为独立运行或直接从命令行使用 Python 运行，但您可以使用 MCP Inspector 来调试它。

MCP Inspector 提供了一种方便的方式来测试和调试您的 MCP 实现：

```bash
# 安装依赖项
pip install -r requirements.txt
# 使用 MCP Inspector 进行调试（不要直接用 Python 运行）
```

MySQL MCP 服务器设计为与 Claude Desktop 等 AI 应用程序集成，不应直接作为独立的 Python 程序运行。

## 开发

```bash
# 克隆仓库
git clone https://github.com/designcomputer/mysql_mcp_server.git
cd mysql_mcp_server
# 创建虚拟环境
python -m venv venv
source venv/bin/activate  # 或在 Windows 上使用 `venv\Scripts\activate`
# 安装开发依赖项
pip install -r requirements-dev.txt
# 运行测试
pytest
```

## 安全注意事项

- 永远不要提交环境变量或凭据
- 使用具有最小必需权限的数据库用户
- 考虑在生产环境中实施查询白名单
- 监控和记录所有数据库操作

## 安全最佳实践

此 MCP 实现需要数据库访问权限才能运行。为了安全：

1. **创建专用的 MySQL 用户**，具有最小权限
2. **永远不要使用 root 凭据**或管理员账户
3. **限制数据库访问**仅限于必要的操作
4. **启用日志记录**用于审计目的
5. **定期安全审查**数据库访问

详细说明请参见 [MySQL 安全配置指南](https://github.com/designcomputer/mysql_mcp_server/blob/main/SECURITY.md)：

- 创建受限的 MySQL 用户
- 设置适当的权限
- 监控数据库访问
- 安全最佳实践

⚠️ 重要提示：在配置数据库访问时，始终遵循最小权限原则。

## 许可证

MIT 许可证 - 详情请参见 LICENSE 文件。

## 贡献

1. Fork 仓库
2. 创建您的功能分支（`git checkout -b feature/amazing-feature`）
3. 提交您的更改（`git commit -m 'Add some amazing feature'`）
4. 推送到分支（`git push origin feature/amazing-feature`）
5. 打开一个 Pull Request
