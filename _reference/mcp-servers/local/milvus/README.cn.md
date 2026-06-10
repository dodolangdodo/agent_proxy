# Milvus MCP 服务器

> 模型上下文协议（MCP）是一个开放协议，可以在LLM应用程序与外部数据源和工具之间实现无缝集成。无论您是在构建AI驱动的IDE、增强聊天界面，还是创建自定义AI工作流，MCP都提供了一种标准化的方式来连接LLM与它们所需的上下文。

此仓库包含一个MCP服务器，提供对[Milvus](https://milvus.io/)向量数据库功能的访问。

## 前提条件

在使用此MCP服务器之前，请确保您具备：

- Python 3.10或更高版本
- 正在运行的[Milvus](https://milvus.io/)实例（本地或远程）
- 已安装[uv](https://github.com/astral-sh/uv)（推荐用于运行服务器）

## 使用方法

使用此MCP服务器的推荐方式是直接使用`uv`运行，无需安装。这是下面示例中Claude Desktop和Cursor配置使用的方式。

如果您想克隆仓库：

```bash
git clone https://github.com/zilliztech/mcp-server-milvus.git
cd mcp-server-milvus
```

然后您可以直接运行服务器：

```bash
uv run src/mcp_server_milvus/server.py --milvus-uri http://localhost:19530
```

或者您可以更改`src/mcp_server_milvus/`目录中的.env文件来设置环境变量，并使用以下命令运行服务器：

```bash
uv run src/mcp_server_milvus/server.py
```

### 重要提示：.env文件的优先级高于命令行参数。

### 运行模式

服务器支持两种运行模式：**stdio**（默认）和**SSE**（服务器发送事件）。

### Stdio模式（默认）

- **描述**：通过标准输入/输出与客户端通信。如果未指定模式，这是默认模式。

- 使用方法：

  ```bash
  uv run src/mcp_server_milvus/server.py --milvus-uri http://localhost:19530
  ```

### SSE模式

- **描述**：使用HTTP服务器发送事件进行通信。此模式允许多个客户端通过HTTP连接，适用于基于Web的应用程序。

- **使用方法：**

  ```bash
  uv run src/mcp_server_milvus/server.py --sse --milvus-uri http://localhost:19530 --port 8000
  ```

  - `--sse`：启用SSE模式。
  - `--port`：指定SSE服务器的端口（默认：8000）。

- **SSE模式调试：**

  如果您想在SSE模式下调试，在启动SSE服务后，输入以下命令：

  ```bash
  mcp dev src/mcp_server_milvus/server.py
  ```

  输出将类似于：

  ```plaintext
  % mcp dev src/mcp_server_milvus/merged_server.py
  Starting MCP inspector...
  ⚙️ Proxy server listening on port 6277
  🔍 MCP Inspector is up and running at http://127.0.0.1:6274 🚀
  ```

  然后您可以在`http://127.0.0.1:6274`访问MCP Inspector进行测试。

## 支持的应用程序

此MCP服务器可以与支持模型上下文协议的各种LLM应用程序一起使用：

- **Claude Desktop**：Anthropic的Claude桌面应用程序
- **Cursor**：支持MCP的AI驱动代码编辑器
- **自定义MCP客户端**：任何实现MCP客户端规范的应用程序

## 与Claude Desktop一起使用

### 不同模式的配置

#### SSE模式配置

按照以下步骤为SSE模式配置Claude Desktop：

1. 从https://claude.ai/download安装Claude Desktop。
2. 打开您的Claude Desktop配置文件：
   - **macOS**：`~/Library/Application Support/Claude/claude_desktop_config.json`
3. 为SSE模式添加以下配置：

```json
{
  "mcpServers": {
    "milvus-sse": {
      "url": "http://your_sse_host:port/sse",
      "disabled": false,
      "autoApprove": []
    }
  }
}
```

4. 重启Claude Desktop以应用更改。

#### Stdio模式配置

对于stdio模式，请按照以下步骤：

1. 从https://claude.ai/download安装Claude Desktop。
2. 打开您的Claude Desktop配置文件：
   - **macOS**：`~/Library/Application Support/Claude/claude_desktop_config.json`
3. 为stdio模式添加以下配置：

```json
{
  "mcpServers": {
    "milvus": {
      "command": "/PATH/TO/uv",
      "args": [
        "--directory",
        "/path/to/mcp-server-milvus/src/mcp_server_milvus",
        "run",
        "server.py",
        "--milvus-uri",
        "http://localhost:19530"
      ]
    }
  }
}
```

4. 重启Claude Desktop以应用更改。

## 与Cursor一起使用

[Cursor也支持MCP](https://docs.cursor.com/context/model-context-protocol)工具。您可以按照以下步骤将Milvus MCP服务器与Cursor集成：

### 集成步骤

1. 打开`Cursor Settings` > `MCP`
2. 点击`Add new global MCP server`
3. 点击后，它将自动重定向到`mcp.json`文件，如果不存在将会创建该文件

### 配置`mcp.json`文件

#### Stdio模式：

用以下内容覆盖`mcp.json`文件：

```json
{
  "mcpServers": {
    "milvus": {
      "command": "/PATH/TO/uv",
      "args": [
        "--directory",
        "/path/to/mcp-server-milvus/src/mcp_server_milvus",
        "run",
        "server.py",
        "--milvus-uri",
        "http://127.0.0.1:19530"
      ]
    }
  }
}
```

#### SSE模式：

1. 通过运行以下命令启动服务：

   ```bash
   uv run src/mcp_server_milvus/server.py --sse --milvus-uri http://your_sse_host --port port
   ```

   > **注意**：将`http://your_sse_host`替换为您的实际SSE主机地址，将`port`替换为您使用的特定端口号。

2. 服务启动并运行后，用以下内容覆盖`mcp.json`文件：

   ```json
   {
       "mcpServers": {
         "milvus-sse": {
           "url": "http://your_sse_host:port/sse",
           "disabled": false,
           "autoApprove": []
         }
       }
   }
   ```

### 完成集成

完成上述步骤后，重启Cursor或重新加载窗口以确保配置生效。

## 验证集成

要验证Cursor是否成功与您的Milvus MCP服务器集成：

1. 打开`Cursor Settings` > `MCP`
2. 检查"milvus"或"milvus-sse"是否出现在列表中（取决于您选择的模式）
3. 确认相关工具已列出（例如，milvus_list_collections、milvus_vector_search等）
4. 如果服务器已启用但显示错误，请查看下面的故障排除部分

## 可用工具

服务器提供以下工具：

### 搜索和查询操作

- `milvus_text_search`：使用全文搜索查找文档

  - 参数：
    - `collection_name`：要搜索的集合名称
    - `query_text`：要搜索的文本
    - `limit`：返回的最大结果数（默认：5）
    - `output_fields`：结果中包含的字段
    - `drop_ratio`：要忽略的低频词比例（0.0-1.0）
- `milvus_vector_search`：在集合上执行向量相似性搜索
  - 参数：
    - `collection_name`：要搜索的集合名称
    - `vector`：查询向量
    - `vector_field`：向量搜索的字段名称（默认："vector"）
    - `limit`：返回的最大结果数（默认：5）
    - `output_fields`：结果中包含的字段
    - `filter_expr`：过滤表达式
    - `metric_type`：距离度量（COSINE、L2、IP）（默认："COSINE"）
- `milvus_hybrid_search`：在集合上执行混合搜索
  - 参数：
    - `collection_name`：要搜索的集合名称
    - `query_text`：搜索的文本查询
    - `text_field`：文本搜索的字段名称
    - `vector`：文本查询的向量
    - `vector_field`：向量搜索的字段名称
    - `limit`：返回的最大结果数
    - `output_fields`：结果中包含的字段
    - `filter_expr`：过滤表达式
- `milvus_query`：使用过滤表达式查询集合
  - 参数：
    - `collection_name`：要查询的集合名称
    - `filter_expr`：过滤表达式（例如'age > 20'）
    - `output_fields`：结果中包含的字段
    - `limit`：返回的最大结果数（默认：10）

### 集合管理

- `milvus_list_collections`：列出数据库中的所有集合

- `milvus_create_collection`：使用指定模式创建新集合

  - 参数：
    - `collection_name`：新集合的名称
    - `collection_schema`：集合模式定义
    - `index_params`：可选的索引参数

- `milvus_load_collection`：将集合加载到内存中进行搜索和查询

  - 参数：
    - `collection_name`：要加载的集合名称
    - `replica_number`：副本数量（默认：1）

- `milvus_release_collection`：从内存中释放集合
  - 参数：
    - `collection_name`：要释放的集合名称

- `milvus_get_collection_info`：列出特定集合的详细信息，如模式、属性、集合ID和其他元数据。
  - 参数：
    - `collection_name`：要获取详细信息的集合名称

### 数据操作

- `milvus_insert_data`：向集合中插入数据

  - 参数：
    - `collection_name`：集合名称
    - `data`：将字段名称映射到值列表的字典

- `milvus_delete_entities`：基于过滤表达式从集合中删除实体
  - 参数：
    - `collection_name`：集合名称
    - `filter_expr`：选择要删除实体的过滤表达式

## 环境变量

- `MILVUS_URI`：Milvus服务器URI（可以设置而不是--milvus-uri）
- `MILVUS_TOKEN`：可选的身份验证令牌
- `MILVUS_DB`：数据库名称（默认为"default"）

## 开发

直接运行服务器：

```bash
uv run server.py --milvus-uri http://localhost:19530
```

## 示例

### 使用Claude Desktop

#### 示例1：列出集合

```
我的Milvus数据库中有哪些集合？
```

Claude将使用MCP在您的Milvus数据库中检查此信息。

```
我将检查您的Milvus数据库中可用的集合。

以下是您的Milvus数据库中的集合：

1. rag_demo
2. test
3. chat_messages
4. text_collection
5. image_collection
6. customized_setup
7. streaming_rag_demo
```

#### 示例2：搜索文档

```
在我的text_collection中查找提到"机器学习"的文档
```

Claude将使用Milvus的全文搜索功能查找相关文档：

```
我将在您的text_collection中搜索关于机器学习的文档。

> 查看来自milvus（本地）的milvus-text-search结果

以下是我找到的提到机器学习的文档：
[结果将根据您的实际数据在此处显示]
```

### 使用Cursor

#### 示例：创建集合

在Cursor中，您可以询问：

```
在Milvus中创建一个名为'articles'的新集合，包含标题（字符串）、内容（字符串）和向量字段（128维）
```

Cursor将使用MCP服务器执行此操作：

```
我将创建一个名为'articles'的新集合，包含指定的字段。

集合'articles'已成功创建，具有以下模式：
- title：字符串
- content：字符串
- vector：浮点向量[128]
```

## 故障排除

### 常见问题

#### 连接错误

如果您看到"无法连接到Milvus服务器"等错误：

1. 验证您的Milvus实例正在运行：`docker ps`（如果使用Docker）
2. 检查配置中的URI是否正确
3. 确保没有防火墙规则阻止连接
4. 尝试在URI中使用`127.0.0.1`而不是`localhost`

#### 身份验证问题

如果您看到身份验证错误：

1. 验证您的`MILVUS_TOKEN`是否正确
2. 检查您的Milvus实例是否需要身份验证
3. 确保您对尝试执行的操作具有正确的权限

#### 找不到工具

如果MCP工具未出现在Claude Desktop或Cursor中：

1. 重启应用程序
2. 检查服务器日志是否有任何错误
3. 验证MCP服务器是否正确运行
4. 按下MCP设置中的刷新按钮（对于Cursor）

### 获取帮助

如果您继续遇到问题：

1. 检查[GitHub Issues](https://github.com/zilliztech/mcp-server-milvus/issues)以查找类似问题
2. 加入[Zilliz社区Discord](https://discord.gg/zilliz)获取支持
3. 提交新问题并提供关于您问题的详细信息
