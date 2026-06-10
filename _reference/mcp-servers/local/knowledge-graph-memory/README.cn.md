# 知识图谱记忆服务器

> <https://github.com/modelcontextprotocol/servers/tree/main/src/memory>

使用本地知识图谱的持久化内存基本实现。这让 Claude 能够跨聊天记住用户的信息。

## 核心概念

### 实体

实体是知识图谱中的主要节点。每个实体包含：

- 唯一名称（标识符）
- 实体类型（例如："person"、"organization"、"event"）
- 观察列表

示例：

```json
{
  "name": "John_Smith",
  "entityType": "person",
  "observations": ["说流利的西班牙语"]
}
```

### 关系

关系定义实体之间的有向连接。它们总是以主动语态存储，描述实体之间的交互或关联方式。

示例：

```json
{
  "from": "John_Smith",
  "to": "Anthropic",
  "relationType": "works_at"
}
```

### 观察

观察是关于实体的离散信息片段。它们具有以下特点：

- 以字符串形式存储
- 附属于特定实体
- 可以独立添加或删除
- 应该是原子性的（每个观察一个事实）

示例：

```json
{
  "entityName": "John_Smith",
  "observations": [
    "说流利的西班牙语",
    "2019年毕业",
    "喜欢上午开会"
  ]
}
```

## API

### 工具

- **create_entities**
  - 在知识图谱中创建多个新实体
  - 输入：`entities`（对象数组）
    - 每个对象包含：
      - `name`（字符串）：实体标识符
      - `entityType`（字符串）：类型分类
      - `observations`（字符串数组）：关联观察
  - 忽略已存在名称的实体

- **create_relations**
  - 在实体之间创建多个新关系
  - 输入：`relations`（对象数组）
    - 每个对象包含：
      - `from`（字符串）：源实体名称
      - `to`（字符串）：目标实体名称
      - `relationType`（字符串）：主动语态的关系类型
  - 跳过重复关系

- **add_observations**
  - 向现有实体添加新观察
  - 输入：`observations`（对象数组）
    - 每个对象包含：
      - `entityName`（字符串）：目标实体
      - `contents`（字符串数组）：要添加的新观察
  - 返回每个实体添加的观察
  - 如果实体不存在则失败

- **delete_entities**
  - 删除实体及其关系
  - 输入：`entityNames`（字符串数组）
  - 级联删除关联关系
  - 如果实体不存在则静默操作

- **delete_observations**
  - 从实体中删除特定观察
  - 输入：`deletions`（对象数组）
    - 每个对象包含：
      - `entityName`（字符串）：目标实体
      - `observations`（字符串数组）：要删除的观察
  - 如果观察不存在则静默操作

- **delete_relations**
  - 从图谱中删除特定关系
  - 输入：`relations`（对象数组）
    - 每个对象包含：
      - `from`（字符串）：源实体名称
      - `to`（字符串）：目标实体名称
      - `relationType`（字符串）：关系类型
  - 如果关系不存在则静默操作

- **read_graph**
  - 读取整个知识图谱
  - 无需输入
  - 返回包含所有实体和关系的完整图谱结构

- **search_nodes**
  - 基于查询搜索节点
  - 输入：`query`（字符串）
  - 搜索范围：
    - 实体名称
    - 实体类型
    - 观察内容
  - 返回匹配的实体及其关系

- **open_nodes**
  - 按名称检索特定节点
  - 输入：`names`（字符串数组）
  - 返回：
    - 请求的实体
    - 请求实体之间的关系
  - 静默跳过不存在的节点

# 在 Claude Desktop 中使用

### 设置

将以下内容添加到您的 claude_desktop_config.json：

#### Docker

```json
{
  "mcpServers": {
    "memory": {
      "command": "docker",
      "args": ["run", "-i", "-v", "claude-memory:/app/dist", "--rm", "mcp/memory"]
    }
  }
}
```

#### NPX

```json
{
  "mcpServers": {
    "memory": {
      "command": "npx",
      "args": [
        "-y",
        "@modelcontextprotocol/server-memory"
      ]
    }
  }
}
```

#### NPX 自定义设置

服务器可以使用以下环境变量进行配置：

```json
{
  "mcpServers": {
    "memory": {
      "command": "npx",
      "args": [
        "-y",
        "@modelcontextprotocol/server-memory"
      ],
      "env": {
        "MEMORY_FILE_PATH": "/path/to/custom/memory.json"
      }
    }
  }
}
```

- `MEMORY_FILE_PATH`：内存存储 JSON 文件的路径（默认：服务器目录中的 `memory.json`）

# VS Code 安装说明

快速安装，请使用下面的一键安装按钮：

[![在 VS Code 中使用 NPX 安装](https://img.shields.io/badge/VS_Code-NPM-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=memory&config=%7B%22command%22%3A%22npx%22%2C%22args%22%3A%5B%22-y%22%2C%22%40modelcontextprotocol%2Fserver-memory%22%5D%7D) [![在 VS Code Insiders 中使用 NPX 安装](https://img.shields.io/badge/VS_Code_Insiders-NPM-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=memory&config=%7B%22command%22%3A%22npx%22%2C%22args%22%3A%5B%22-y%22%2C%22%40modelcontextprotocol%2Fserver-memory%22%5D%7D&quality=insiders)

[![在 VS Code 中使用 Docker 安装](https://img.shields.io/badge/VS_Code-Docker-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=memory&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%22-i%22%2C%22-v%22%2C%22claude-memory%3A%2Fapp%2Fdist%22%2C%22--rm%22%2C%22mcp%2Fmemory%22%5D%7D) [![在 VS Code Insiders 中使用 Docker 安装](https://img.shields.io/badge/VS_Code_Insiders-Docker-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=memory&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%22-i%22%2C%22-v%22%2C%22claude-memory%3A%2Fapp%2Fdist%22%2C%22--rm%22%2C%22mcp%2Fmemory%22%5D%7D&quality=insiders)

手动安装时，将以下 JSON 块添加到 VS Code 的用户设置（JSON）文件中。您可以按 `Ctrl + Shift + P` 并输入 `Preferences: Open Settings (JSON)` 来实现。

可选地，您可以将其添加到工作区中名为 `.vscode/mcp.json` 的文件中。这将允许您与他人共享配置。

> 注意：在 `.vscode/mcp.json` 文件中不需要 `mcp` 键。

#### NPX

```json
{
  "mcp": {
    "servers": {
      "memory": {
        "command": "npx",
        "args": [
          "-y",
          "@modelcontextprotocol/server-memory"
        ]
      }
    }
  }
}
```

#### Docker

```json
{
  "mcp": {
    "servers": {
      "memory": {
        "command": "docker",
        "args": [
          "run",
          "-i",
          "-v",
          "claude-memory:/app/dist",
          "--rm",
          "mcp/memory"
        ]
      }
    }
  }
}
```

### 系统提示

使用内存的提示取决于用例。更改提示将帮助模型确定创建内存的频率和类型。

以下是聊天个性化的示例提示。您可以在 [Claude.ai 项目](https://www.anthropic.com/news/projects) 的"自定义指令"字段中使用此提示。

```
对于每次交互，请遵循以下步骤：

1. 用户识别：
   - 您应该假设您正在与 default_user 交互
   - 如果您尚未识别 default_user，请主动尝试识别

2. 内存检索：
   - 始终在聊天开始时只说"正在回忆..."并从知识图谱中检索所有相关信息
   - 始终将您的知识图谱称为您的"记忆"

3. 记忆
   - 在与用户对话时，请注意属于以下类别的任何新信息：
     a) 基本身份（年龄、性别、位置、职位、教育水平等）
     b) 行为（兴趣、习惯等）
     c) 偏好（沟通风格、首选语言等）
     d) 目标（目标、目的、愿望等）
     e) 关系（最多三度分离的个人和职业关系）

4. 记忆更新：
   - 如果在交互过程中收集到任何新信息，请按以下方式更新您的记忆：
     a) 为经常出现的组织、人员和重要事件创建实体
     b) 使用关系将它们连接到当前实体
     c) 将关于它们的事实存储为观察
```

## 构建

Docker：

```sh
docker build -t mcp/memory -f src/memory/Dockerfile . 
```

## 许可证

此 MCP 服务器采用 MIT 许可证。这意味着您可以自由使用、修改和分发软件，但需遵守 MIT 许可证的条款和条件。更多详情，请参阅项目仓库中的 LICENSE 文件。
