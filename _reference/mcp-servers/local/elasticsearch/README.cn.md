# Elasticsearch MCP 服务器

> <https://github.com/elastic/mcp-server-elasticsearch>

此存储库包含用于研究和评估的实验性功能，不适用于生产环境。

使用模型上下文协议（MCP）从任何 MCP 客户端（如 Claude Desktop）直接连接到您的 Elasticsearch 数据。

此服务器使用模型上下文协议将智能体连接到您的 Elasticsearch 数据。它允许您通过自然语言对话与 Elasticsearch 索引进行交互。

## 可用工具

* `list_indices`：列出所有可用的 Elasticsearch 索引
* `get_mappings`：获取特定 Elasticsearch 索引的字段映射
* `search`：使用提供的查询 DSL 执行 Elasticsearch 搜索
* `get_shards`：获取所有或特定索引的分片信息

## 先决条件

* 一个 Elasticsearch 实例
* Elasticsearch 身份验证凭据（API 密钥或用户名/密码）
* MCP 客户端（例如 Claude Desktop）

## 演示

<https://github.com/user-attachments/assets/5dd292e1-a728-4ca7-8f01-1380d1bebe0c>

## 安装和设置

### 通过 Smithery 安装

通过 [Smithery](https://smithery.ai/server/@elastic/mcp-server-elasticsearch) 自动为 Claude Desktop 安装 Elasticsearch MCP 服务器：

```bash
npx -y @smithery/cli install @elastic/mcp-server-elasticsearch --client claude
```

### 使用已发布的 NPM 包

> [!TIP]
> 使用 Elasticsearch MCP 服务器的最简单方法是通过已发布的 npm 包。

1. **配置 MCP 客户端**
   * 打开您的 MCP 客户端。查看 [MCP 客户端列表](https://modelcontextprotocol.io/clients)，这里我们配置 Claude Desktop。
   * 转到 **设置 > 开发者 > MCP 服务器**
   * 点击 `编辑配置` 并添加一个新的 MCP 服务器，配置如下：

   ```json
   {
     "mcpServers": {
       "elasticsearch-mcp-server": {
         "command": "npx",
         "args": [
           "-y",
           "@elastic/mcp-server-elasticsearch"
         ],
         "env": {
           "ES_URL": "your-elasticsearch-url",
           "ES_API_KEY": "your-api-key"
         }
       }
     }
   }
   ```

2. **开始对话**
   * 在您的 MCP 客户端中打开新对话
   * MCP 服务器应该会自动连接
   * 现在您可以询问关于 Elasticsearch 数据的问题

### 配置选项

Elasticsearch MCP 服务器支持配置选项来连接到您的 Elasticsearch：

> [!NOTE]
> 您必须提供 API 密钥或用户名和密码来进行身份验证。

| 环境变量 | 描述 | 必需 |
|---------|------|------|
| `ES_URL` | 您的 Elasticsearch 实例 URL | 是 |
| `ES_API_KEY` | 用于身份验证的 Elasticsearch API 密钥 | 否 |
| `ES_USERNAME` | 用于基本身份验证的 Elasticsearch 用户名 | 否 |
| `ES_PASSWORD` | 用于基本身份验证的 Elasticsearch 密码 | 否 |
| `ES_CA_CERT` | Elasticsearch SSL/TLS 自定义 CA 证书路径 | 否 |
| `ES_SSL_SKIP_VERIFY` | 设置为 '1' 或 'true' 以跳过 SSL 证书验证 | 否 |
| `ES_PATH_PREFIX` | 在非根路径暴露的 Elasticsearch 实例的路径前缀 | 否 |
| `ES_VERSION` | 服务器假设 Elasticsearch 9.x。设置为 `8` 以目标 Elasticsearch 8.x | 否 |

### 本地开发

> [!NOTE]
> 如果您想修改或扩展 MCP 服务器，请按照这些本地开发步骤。

1. **使用正确的 Node.js 版本**

   ```bash
   nvm use
   ```

2. **安装依赖**

   ```bash
   npm install
   ```

3. **构建项目**

   ```bash
   npm run build
   ```

4. **在 Claude Desktop 应用中本地运行**
   * 打开 **Claude Desktop 应用**
   * 转到 **设置 > 开发者 > MCP 服务器**
   * 点击 `编辑配置` 并添加一个新的 MCP 服务器，配置如下：

   ```json
   {
     "mcpServers": {
       "elasticsearch-mcp-server-local": {
         "command": "node",
         "args": [
           "/path/to/your/project/dist/index.js"
         ],
         "env": {
           "ES_URL": "your-elasticsearch-url",
           "ES_API_KEY": "your-api-key"
         }
       }
     }
   }
   ```

5. **使用 MCP Inspector 调试**

   ```bash
   ES_URL=your-elasticsearch-url ES_API_KEY=your-api-key npm run inspector
   ```

   这将启动 MCP Inspector，允许您调试和分析请求。您应该会看到：

   ```bash
   Starting MCP inspector...
   Proxy server listening on port 3000

   🔍 MCP Inspector is up and running at http://localhost:5173 🚀
   ```

#### Docker 镜像

如果您想在容器中构建和运行服务器，可以使用 `Dockerfile`。要构建，运行：

```sh
docker build -t mcp-server-elasticsearch .
```

要运行，不使用上面的 `npx` 命令或自定义的 `node` 或 `npm` 命令，而是运行：

```sh
docker run -i \
  -e ES_URL=<url> \
  -e ES_API_KEY=<key> \
  mcp-server-elasticsearch
```

## 贡献

我们欢迎社区的贡献！有关如何贡献的详细信息，请参阅[贡献指南](/docs/CONTRIBUTING.md)。

## 示例问题

> [!TIP]
> 这里是一些您可以在 MCP 客户端中尝试的自然语言查询。

* "我的 Elasticsearch 集群中有哪些索引？"
* "显示 'products' 索引的字段映射。"
* "查找上个月所有超过 $500 的订单。"
* "哪些产品收到了最多的 5 星评价？"

## 工作原理

1. MCP 客户端分析您的请求并确定需要哪些 Elasticsearch 操作。
2. MCP 服务器执行这些操作（列出索引、获取映射、执行搜索）。
3. MCP 客户端处理结果并以用户友好的格式呈现。

## 安全最佳实践

> [!WARNING]
> 避免使用集群管理员权限。创建具有有限范围的专用 API 密钥，并在索引级别应用细粒度访问控制以防止未经授权的数据访问。

您可以创建一个具有最小权限的专用 Elasticsearch API 密钥来控制对数据的访问：

```
POST /_security/api_key
{
  "name": "es-mcp-server-access",
  "role_descriptors": {
    "mcp_server_role": {
      "cluster": [
        "monitor"
      ],
      "indices": [
        {
          "names": [
            "index-1",
            "index-2",
            "index-pattern-*"
          ],
          "privileges": [
            "read",
            "view_index_metadata"
          ]
        }
      ]
    }
  }
}
```

## 许可证

此项目使用 Apache License 2.0 许可证。

## 故障排除

* 确保您的 MCP 配置正确。
* 验证您的 Elasticsearch URL 可以从您的机器访问。
* 检查您的身份验证凭据（API 密钥或用户名/密码）具有必要的权限。
* 如果使用带有自定义 CA 的 SSL/TLS，验证证书路径正确且文件可读。
* 查看终端输出中的错误消息。

如果遇到问题，请随时在 GitHub 存储库上开启 issue。
