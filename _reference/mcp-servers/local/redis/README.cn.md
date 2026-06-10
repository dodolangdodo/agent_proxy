# Redis MCP 服务器

> <https://github.com/redis/mcp-redis>

## 概述

Redis MCP 服务器是一个为智能应用程序设计的**自然语言接口**，用于高效管理和搜索 Redis 中的数据。它与 **MCP（模型内容协议）客户端**无缝集成，使 AI 驱动的工作流能够与 Redis 中的结构化和非结构化数据进行交互。使用此 MCP 服务器，您可以询问如下问题：

- "将整个对话存储在流中"
- "缓存此项目"
- "存储带有过期时间的会话"
- "索引并搜索此向量"

## 功能特性

- **自然语言查询**：使 AI 智能体能够使用自然语言查询和更新 Redis。
- **无缝 MCP 集成**：与任何 **MCP 客户端**配合使用，实现流畅通信。
- **完整 Redis 支持**：处理**哈希、列表、集合、有序集合、流**等数据结构。
- **搜索和过滤**：支持在 Redis 中高效的数据检索和搜索。
- **可扩展且轻量级**：专为**高性能**数据操作而设计。

## 工具

此 MCP 服务器提供管理 Redis 中存储数据的工具。

- `string` 工具用于设置、获取带过期时间的字符串。适用于存储简单配置值、会话数据或缓存响应。
- `hash` 工具用于在单个键内存储字段-值对。哈希可以存储向量嵌入。适用于表示具有多个属性的对象、用户配置文件或可以单独访问字段的产品信息。
- `list` 工具提供常见操作来追加和弹出项目。适用于队列、消息代理或维护最近操作列表。
- `set` 工具用于添加、删除和列出集合成员。适用于跟踪唯一值（如用户 ID 或标签），以及执行交集等集合操作。
- `sorted set` 工具用于管理数据，例如排行榜、优先队列或基于分数排序的时间分析。
- `pub/sub` 功能用于向频道发布消息并订阅接收消息。适用于实时通知、聊天应用程序或向多个客户端分发更新。
- `streams` 工具用于添加、读取和删除数据流。适用于事件溯源、活动订阅或支持消费者组的传感器数据记录。
- `JSON` 工具用于在 Redis 中存储、检索和操作 JSON 文档。适用于复杂的嵌套数据结构、文档数据库或具有基于路径访问的配置管理。

其他工具：

- `query engine` 工具用于管理向量索引和执行向量搜索
- `server management` 工具用于检索数据库信息

## 安装

按照以下说明安装服务器。

```sh
# 克隆仓库
git clone https://github.com/redis/mcp-redis.git
cd mcp-redis

# 使用 uv 安装依赖
uv venv
source .venv/bin/activate
uv sync
```

## 配置

要配置此 Redis MCP 服务器，请考虑以下环境变量：

| 名称                 | 描述                                                      | 默认值        |
|----------------------|-----------------------------------------------------------|--------------|
| `REDIS_HOST`         | Redis IP 或主机名                                         | `"127.0.0.1"` |
| `REDIS_PORT`         | Redis 端口                                                | `6379`       |
| `REDIS_DB`           | 数据库                                                    | 0            |
| `REDIS_USERNAME`     | 默认数据库用户名                                           | `"default"`  |
| `REDIS_PWD`          | 默认数据库密码                                             | ""           |
| `REDIS_SSL`          | 启用或禁用 SSL/TLS                                         | `False`      |
| `REDIS_CA_PATH`      | 用于验证服务器的 CA 证书                                   | None         |
| `REDIS_SSL_KEYFILE`  | 客户端用于客户端认证的私钥文件                              | None         |
| `REDIS_SSL_CERTFILE` | 客户端用于客户端认证的证书文件                              | None         |
| `REDIS_CERT_REQS`    | 客户端是否应验证服务器证书                                  | `"required"` |
| `REDIS_CA_CERTS`     | 受信任 CA 证书文件的路径                                   | None         |
| `REDIS_CLUSTER_MODE` | 启用 Redis 集群模式                                       | `False`      |
| `MCP_TRANSPORT`      | 使用 `stdio` 或 `sse` 传输                                | `stdio`      |

设置环境变量有几种方法：

1. **使用 `.env` 文件**：  
  在项目目录中放置一个 `.env` 文件，为每个环境变量设置键值对。像 `python-dotenv`、`pipenv` 和 `uv` 这样的工具可以在运行应用程序时自动加载这些变量。这是管理配置的便捷且安全的方法，因为它将敏感数据保留在 shell 历史记录和版本控制之外（如果 `.env` 在 `.gitignore` 中）。

例如，从仓库中提供的 `.env.example` 文件创建一个 `.env` 文件：

  ```bash
cp .env.example .env
  ```

  然后编辑 `.env` 文件以设置您的 Redis 配置：

或者，

2. **在 Shell 中设置变量**：  
  您可以在运行应用程序之前直接在 shell 中导出环境变量。例如：

  ```sh
  export REDIS_HOST=your_redis_host
  export REDIS_PORT=6379
  # 其他变量将类似设置...
  ```

  此方法适用于临时覆盖或快速测试。

## 传输方式

此 MCP 服务器可以配置为本地处理请求，作为进程运行并通过 `stdin` 和 `stdout` 与 MCP 客户端通信。
这是默认配置。`sse` 传输也是可配置的，因此服务器可通过网络访问。
相应地配置 `MCP_TRANSPORT` 变量。

```commandline
export MCP_TRANSPORT="sse"
```

然后启动服务器。

```commandline
uv run src/main.py
```

测试服务器：

```commandline
curl -i http://127.0.0.1:8000/sse
HTTP/1.1 200 OK
```

与您喜欢的工具或客户端集成。GitHub Copilot 的 VS Code 配置是：

```commandline
"mcp": {
    "servers": {
        "redis-mcp": {
            "type": "sse",
            "url": "http://127.0.0.1:8000/sse"
        },
    }
},
```

## 与 OpenAI Agents SDK 集成

将此 MCP 服务器与 OpenAI Agents SDK 集成。阅读[文档](https://openai.github.io/openai-agents-python/mcp/)以了解更多关于 SDK 与 MCP 集成的信息。

安装 Python SDK。

```commandline
pip install openai-agents
```

配置 OpenAI 令牌：

```commandline
export OPENAI_API_KEY="<openai_token>"
```

并运行[应用程序](./examples/redis_assistant.py)。

```commandline
python3.13 redis_assistant.py
```

您可以使用 [OpenAI 仪表板](https://platform.openai.com/traces/)对智能体工作流进行故障排除。

## 与 Claude Desktop 集成

### 通过 Smithery

如果您想测试[由 Smithery 部署](https://smithery.ai/docs/deployments)的 [Redis MCP 服务器](https://smithery.ai/server/@redis/mcp-redis)，您可以自动配置 Claude Desktop：

```bash
npx -y @smithery/cli install @redis/mcp-redis --client claude
```

按照提示提供详细信息来配置服务器并连接到 Redis（例如使用 Redis Cloud 数据库）。
该过程将在 `claude_desktop_config.json` 配置文件中创建适当的配置。

### 手动配置

您可以配置 Claude Desktop 使用此 MCP 服务器。

1. 指定您的 Redis 凭据和 TLS 配置
2. 获取您的 `uv` 命令完整路径（例如 `which uv`）
3. 编辑 `claude_desktop_config.json` 配置文件
   - 在 MacOS 上，位于 `~/Library/Application\ Support/Claude/`

```commandline
{
    "mcpServers": {
        "redis": {
            "command": "<full_path_uv_command>",
            "args": [
                "--directory",
                "<your_mcp_server_directory>",
                "run",
                "src/main.py"
            ],
            "env": {
                "REDIS_HOST": "<your_redis_database_hostname>",
                "REDIS_PORT": "<your_redis_database_port>",
                "REDIS_PWD": "<your_redis_database_password>",
                "REDIS_SSL": True|False,
                "REDIS_CA_PATH": "<your_redis_ca_path>",
                "REDIS_CLUSTER_MODE": True|False
            }
        }
    }
}
```

### 使用 Docker

您可以使用此服务器的 Docker 化部署。您可以构建自己的镜像或使用官方 [Redis MCP Docker](https://hub.docker.com/r/mcp/redis) 镜像。

如果您想构建自己的镜像，Redis MCP 服务器提供了 Dockerfile。使用以下命令构建此服务器的镜像：

```commandline
docker build -t mcp-redis .
```

最后，配置 Claude Desktop 在启动时创建容器。编辑 `claude_desktop_config.json` 并添加：

```commandline
{
  "mcpServers": {
    "redis": {
      "command": "docker",
      "args": ["run",
                "--rm",
                "--name",
                "redis-mcp-server",
                "-i",
                "-e", "REDIS_HOST=<redis_hostname>",
                "-e", "REDIS_PORT=<redis_port>",
                "-e", "REDIS_USERNAME=<redis_username>",
                "-e", "REDIS_PWD=<redis_password>",
                "mcp-redis"]
    }
  }
}
```

要使用官方 [Redis MCP Docker](https://hub.docker.com/r/mcp/redis) 镜像，只需将您的镜像名称（上例中的 `mcp-redis`）替换为 `mcp/redis`。

### 故障排除

您可以通过查看日志文件来排除问题。

```commandline
tail -f ~/Library/Logs/Claude/mcp-server-redis.log
```

## 与 VS Code 集成

要在 VS Code 中使用 Redis MCP 服务器，您需要：

1. 启用[智能体模式](https://code.visualstudio.com/docs/copilot/chat/chat-agent-mode)工具。将以下内容添加到您的 `settings.json`：

```commandline
{
  "chat.agent.enabled": true
}
```

2. 将 Redis MCP 服务器配置添加到您的 `mcp.json` 或 `settings.json`：

```commandline
// 示例 .vscode/mcp.json
{
  "servers": {
    "redis": {
      "type": "stdio",
      "command": "<full_path_uv_command>",
      "args": [
        "--directory",
        "<your_mcp_server_directory>",
        "run",
        "src/main.py"
      ],
      "env": {
        "REDIS_HOST": "<your_redis_database_hostname>",
        "REDIS_PORT": "<your_redis_database_port>",
        "REDIS_USERNAME": "<your_redis_database_username>",
        "REDIS_PWD": "<your_redis_database_password>",
      }
    }
  }
}
```

```commandline
// 示例 settings.json
{
  "mcp": {
    "servers": {
      "redis": {
        "type": "stdio",
        "command": "<full_path_uv_command>",
        "args": [
          "--directory",
          "<your_mcp_server_directory>",
          "run",
          "src/main.py"
        ],
        "env": {
          "REDIS_HOST": "<your_redis_database_hostname>",
          "REDIS_PORT": "<your_redis_database_port>",
          "REDIS_USERNAME": "<your_redis_database_username>",
          "REDIS_PWD": "<your_redis_database_password>",
        }
      }
    }
  }
}
```

更多信息请参阅 [VS Code 文档](https://code.visualstudio.com/docs/copilot/chat/mcp-servers)。

## 测试

您可以使用 [MCP Inspector](https://modelcontextprotocol.io/docs/tools/inspector) 对此 MCP 服务器进行可视化调试。

```sh
npx @modelcontextprotocol/inspector uv run src/main.py
```

## 示例用例

- **AI 助手**：使 LLM 能够在 Redis 中获取、存储和处理数据。
- **聊天机器人和虚拟智能体**：检索会话数据、管理队列并个性化响应。
- **数据搜索和分析**：查询 Redis 以获得**实时洞察和快速查找**。
- **事件处理**：使用 **Redis Streams** 管理事件流。

## 贡献

1. Fork 仓库
2. 创建新分支（`feature-branch`）
3. 提交您的更改
4. 推送到您的分支并提交 PR！

## 许可证

此项目采用 **MIT 许可证**。

## 徽章

<a href="https://glama.ai/mcp/servers/@redis/mcp-redis">
  <img width="380" height="200" src="https://glama.ai/mcp/servers/@redis/mcp-redis/badge" alt="Redis Server MCP server" />
</a>

## 联系方式

如有问题或需要支持，请通过 [GitHub Issues](https://github.com/redis/mcp-redis/issues) 联系我们。
