# SQLite MCP 服务器

> <https://github.com/modelcontextprotocol/servers-archived/tree/main/src/sqlite>

## 概述

一个模型上下文协议 (MCP) 服务器实现，通过 SQLite 提供数据库交互和商业智能功能。该服务器支持运行 SQL 查询、分析业务数据，并自动生成商业洞察备忘录。

## 组件

### 资源

服务器公开一个动态资源：

- `memo://insights`：持续更新的商业洞察备忘录，汇总分析过程中发现的洞察
  - 当通过 append-insight 工具发现新洞察时自动更新

### 提示

服务器提供一个演示提示：

- `mcp-demo`：引导用户进行数据库操作的交互式提示
  - 必需参数：`topic` - 要分析的业务领域
  - 生成适当的数据库模式和示例数据
  - 引导用户进行分析和洞察生成
  - 与商业洞察备忘录集成

### 工具

服务器提供六个核心工具：

#### 查询工具

- `read_query`
  - 执行 SELECT 查询从数据库读取数据
  - 输入：
    - `query`（字符串）：要执行的 SELECT SQL 查询
  - 返回：查询结果作为对象数组

- `write_query`
  - 执行 INSERT、UPDATE 或 DELETE 查询
  - 输入：
    - `query`（字符串）：SQL 修改查询
  - 返回：`{ affected_rows: number }`

- `create_table`
  - 在数据库中创建新表
  - 输入：
    - `query`（字符串）：CREATE TABLE SQL 语句
  - 返回：表创建确认

#### 模式工具

- `list_tables`
  - 获取数据库中所有表的列表
  - 无需输入
  - 返回：表名数组

- `describe-table`
  - 查看特定表的模式信息
  - 输入：
    - `table_name`（字符串）：要描述的表名
  - 返回：包含名称和类型的列定义数组

#### 分析工具

- `append_insight`
  - 向备忘录资源添加新的商业洞察
  - 输入：
    - `insight`（字符串）：从数据分析中发现的商业洞察
  - 返回：洞察添加确认
  - 触发 memo://insights 资源更新

## 在 Claude Desktop 中使用

### uv

```bash
# 将服务器添加到您的 claude_desktop_config.json
"mcpServers": {
  "sqlite": {
    "command": "uv",
    "args": [
      "--directory",
      "parent_of_servers_repo/servers/src/sqlite",
      "run",
      "mcp-server-sqlite",
      "--db-path",
      "~/test.db"
    ]
  }
}
```

### Docker

```json
# 将服务器添加到您的 claude_desktop_config.json
"mcpServers": {
  "sqlite": {
    "command": "docker",
    "args": [
      "run",
      "--rm",
      "-i",
      "-v",
      "mcp-test:/mcp",
      "mcp/sqlite",
      "--db-path",
      "/mcp/test.db"
    ]
  }
}
```

## 在 VS Code 中使用

快速安装，请点击下方安装按钮：

[![在 VS Code 中使用 UV 安装](https://img.shields.io/badge/VS_Code-UV-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=sqlite&inputs=%5B%7B%22type%22%3A%22promptString%22%2C%22id%22%3A%22db_path%22%2C%22description%22%3A%22SQLite%20Database%20Path%22%2C%22default%22%3A%22%24%7BworkspaceFolder%7D%2Fdb.sqlite%22%7D%5D&config=%7B%22command%22%3A%22uvx%22%2C%22args%22%3A%5B%22mcp-server-sqlite%22%2C%22--db-path%22%2C%22%24%7Binput%3Adb_path%7D%22%5D%7D) [![在 VS Code Insiders 中使用 UV 安装](https://img.shields.io/badge/VS_Code_Insiders-UV-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=sqlite&inputs=%5B%7B%22type%22%3A%22promptString%22%2C%22id%22%3A%22db_path%22%2C%22description%22%3A%22SQLite%20Database%20Path%22%2C%22default%22%3A%22%24%7BworkspaceFolder%7D%2Fdb.sqlite%22%7D%5D&config=%7B%22command%22%3A%22uvx%22%2C%22args%22%3A%5B%22mcp-server-sqlite%22%2C%22--db-path%22%2C%22%24%7Binput%3Adb_path%7D%22%5D%7D&quality=insiders)

[![在 VS Code 中使用 Docker 安装](https://img.shields.io/badge/VS_Code-Docker-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=sqlite&inputs=%5B%7B%22type%22%3A%22promptString%22%2C%22id%22%3A%22db_path%22%2C%22description%22%3A%22SQLite%20Database%20Path%20(within%20container)%22%2C%22default%22%3A%22%2Fmcp%2Fdb.sqlite%22%7D%5D&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%22-i%22%2C%22--rm%22%2C%22-v%22%2C%22mcp-sqlite%3A%2Fmcp%22%2C%22mcp%2Fsqlite%22%2C%22--db-path%22%2C%22%24%7Binput%3Adb_path%7D%22%5D%7D) [![在 VS Code Insiders 中使用 Docker 安装](https://img.shields.io/badge/VS_Code_Insiders-Docker-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=sqlite&inputs=%5B%7B%22type%22%3A%22promptString%22%2C%22id%22%3A%22db_path%22%2C%22description%22%3A%22SQLite%20Database%20Path%20(within%20container)%22%2C%22default%22%3A%22%2Fmcp%2Fdb.sqlite%22%7D%5D&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%22-i%22%2C%22--rm%22%2C%22-v%22%2C%22mcp-sqlite%3A%2Fmcp%22%2C%22mcp%2Fsqlite%22%2C%22--db-path%22%2C%22%24%7Binput%3Adb_path%7D%22%5D%7D&quality=insiders)

手动安装时，请将以下 JSON 块添加到 VS Code 的用户设置 (JSON) 文件中。您可以按 `Ctrl + Shift + P` 并输入 `Preferences: Open Settings (JSON)` 来完成此操作。

或者，您可以将其添加到工作区中名为 `.vscode/mcp.json` 的文件中。这将允许您与他人共享配置。

> 注意：使用 `mcp.json` 文件时需要 `mcp` 键。

### uv

```json
{
  "mcp": {
    "inputs": [
      {
        "type": "promptString",
        "id": "db_path",
        "description": "SQLite 数据库路径",
        "default": "${workspaceFolder}/db.sqlite"
      }
    ],
    "servers": {
      "sqlite": {
        "command": "uvx",
        "args": [
          "mcp-server-sqlite",
          "--db-path",
          "${input:db_path}"
        ]
      }
    }
  }
}
```

### Docker

```json
{
  "mcp": {
    "inputs": [
      {
        "type": "promptString",
        "id": "db_path",
        "description": "SQLite 数据库路径（容器内）",
        "default": "/mcp/db.sqlite"
      }
    ],
    "servers": {
      "sqlite": {
        "command": "docker",
        "args": [
          "run",
          "-i",
          "--rm",
          "-v",
          "mcp-sqlite:/mcp",
          "mcp/sqlite",
          "--db-path",
          "${input:db_path}"
        ]
      }
    }
  }
}
```

## 构建

Docker：

```bash
docker build -t mcp/sqlite .
```

## 使用 MCP 检查器测试

```bash
uv add "mcp[cli]"
mcp dev src/mcp_server_sqlite/server.py:wrapper  
```

## 许可证

此 MCP 服务器采用 MIT 许可证。这意味着您可以自由使用、修改和分发软件，但需遵守 MIT 许可证的条款和条件。更多详情，请参阅项目仓库中的 LICENSE 文件。
