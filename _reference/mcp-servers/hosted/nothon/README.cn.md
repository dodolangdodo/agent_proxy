# Notion MCP 服务器

> <https://github.com/suekou/mcp-notion-server>

一个提供对Notion工作区全面访问的模型上下文协议服务器。此服务器使LLM能够通过官方Notion API与Notion页面、数据库、块和用户进行交互。

## 功能特性

- **页面管理**: 创建、读取、更新和删除Notion页面
- **数据库操作**: 查询数据库、创建数据库项目和管理数据库属性
- **块操作**: 在页面内添加、检索、更新和删除块
- **搜索功能**: 跨页面和数据库搜索
- **用户管理**: 检索用户信息和工作区详情
- **评论系统**: 在页面和块上创建和检索评论
- **富文本支持**: 完全支持Notion的富文本格式
- **Markdown转换**: 可选的将响应转换为可读Markdown格式

## 设置

### 前提条件

1. 在 <https://www.notion.so/profile/integrations> 创建Notion集成
2. 点击 "New Integration"
3. 输入集成名称并选择适当的权限（例如，"Read content", "Update content"）
4. 复制"内部集成令牌"（以`secret_`开头）
5. 与您的集成共享您的Notion页面/数据库

### 配置

服务器需要以下配置：

- `notion-api-token`（必需）：您的Notion API集成令牌
- `enabled-tools`（可选）：要启用的特定工具的逗号分隔列表
- `enable-markdown`（可选）：为响应启用实验性Markdown转换

## 可用工具

### 块操作

- `notion_append_block_children` - 向页面或块添加新块
- `notion_retrieve_block` - 通过ID获取特定块
- `notion_retrieve_block_children` - 获取块的子块
- `notion_update_block` - 更新块内容
- `notion_delete_block` - 删除块

### 页面操作

- `notion_retrieve_page` - 获取页面内容和属性
- `notion_update_page_properties` - 更新页面属性

### 数据库操作

- `notion_query_database` - 使用过滤器和排序查询数据库
- `notion_retrieve_database` - 获取数据库架构和属性
- `notion_create_database_item` - 创建新的数据库条目

### 搜索和用户

- `notion_search` - 搜索页面和数据库
- `notion_list_all_users` - 列出工作区用户（需要企业计划）
- `notion_retrieve_user` - 获取特定用户详情

### 评论

- `notion_create_comment` - 向页面添加评论
- `notion_retrieve_comments` - 从页面获取评论

## 响应格式

服务器支持两种响应格式：

- **JSON**（默认）：用于编程使用的原始Notion API响应
- **Markdown**：用于内容消费的人类可读格式

在工具调用中使用`format`参数来指定您的首选格式。

## 安全性和权限

此服务器需要适当的Notion集成权限：

- 读取内容
- 更新内容
- 插入内容
- 读取评论
- 插入评论
- 读取用户信息（用于用户相关工具）

## 错误处理

服务器为以下情况提供详细的错误消息：

- 缺失或无效的API令牌
- 权限不足
- 无效的块/页面/数据库ID
- 格式错误的请求
- API速率限制

## 速率限制

Notion API有速率限制。当超出限制时，服务器将返回适当的错误消息。考虑在您的应用程序中实现重试逻辑。
