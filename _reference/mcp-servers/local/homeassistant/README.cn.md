# Home Assistant 模型上下文协议服务器

该服务器使用 MCP 协议与 LLM 应用程序共享对本地 Home Assistant 实例的访问。

这是一个强大的桥梁，连接您的 Home Assistant 实例和大型语言模型 (LLM)，通过模型上下文协议 (MCP) 实现智能家居设备的自然语言控制和监控。该服务器提供了管理整个 Home Assistant 生态系统的全面 API，从设备控制到系统管理。

![许可证](https://img.shields.io/badge/license-MIT-blue.svg)
![Node.js](https://img.shields.io/badge/node-%3E%3D20.10.0-green.svg)
![Docker Compose](https://img.shields.io/badge/docker-compose-%3E%3D1.27.0-blue.svg)
![NPM](https://img.shields.io/badge/npm-%3E%3D7.0.0-orange.svg)
![TypeScript](https://img.shields.io/badge/typescript-%5E5.0.0-blue.svg)
![测试覆盖率](https://img.shields.io/badge/coverage-95%25-brightgreen.svg)

## 功能特性

- 🎮 **设备控制**：通过自然语言控制任何 Home Assistant 设备
- 🔄 **实时更新**：通过服务器发送事件 (SSE) 获取即时更新
- 🤖 **自动化管理**：创建、更新和管理自动化
- 📊 **状态监控**：跟踪和查询设备状态
- 🔐 **安全**：基于令牌的身份验证和速率限制
- 📱 **移动就绪**：与任何支持 HTTP 的客户端兼容

## 使用 SSE 的实时更新

服务器包含一个强大的服务器发送事件 (SSE) 系统，提供来自 Home Assistant 实例的实时更新。这允许您：

- 🔄 获取任何设备的即时状态变化
- 📡 监控自动化触发器和执行
- 🎯 订阅特定域或实体
- 📊 跟踪服务调用和脚本执行

### SSE 快速示例

```javascript
const eventSource = new EventSource(
  'http://localhost:3000/subscribe_events?token=YOUR_TOKEN&domain=light'
);

eventSource.onmessage = (event) => {
  const data = JSON.parse(event.data);
  console.log('收到更新:', data);
};
```

有关 SSE 系统的完整文档，请参阅 [SSE_API.md](docs/SSE_API.md)。

## 目录

- [主要功能](#主要功能)
- [先决条件](#先决条件)
- [安装](#安装)
  - [基础设置](#基础设置)
  - [Docker 设置（推荐）](#docker-设置推荐)
- [配置](#配置)
- [开发](#开发)
- [API 参考](#api-参考)
  - [设备控制](#设备控制)
  - [插件管理](#插件管理)
  - [包管理](#包管理)
  - [自动化管理](#自动化管理)
- [自然语言集成](#自然语言集成)
- [故障排除](#故障排除)
- [项目状态](#项目状态)
- [贡献](#贡献)
- [资源](#资源)
- [许可证](#许可证)

## 主要功能

### 核心功能 🎮

- **智能设备控制**
  - 💡 **灯光**：亮度、色温、RGB 颜色
  - 🌡️ **气候**：温度、HVAC 模式、风扇模式、湿度
  - 🚪 **窗帘**：位置和倾斜控制
  - 🔌 **开关**：开/关控制
  - 🚨 **传感器和接触器**：状态监控
  - 🎵 **媒体播放器**：播放控制、音量、源选择
  - 🌪️ **风扇**：速度、摆动、方向
  - 🔒 **锁**：锁定/解锁控制
  - 🧹 **吸尘器**：启动、停止、返回基座
  - 📹 **摄像头**：运动检测、快照

### 系统管理 🛠️

- **插件管理**
  - 浏览可用插件
  - 安装/卸载插件
  - 启动/停止/重启插件
  - 版本管理
  - 配置访问

- **包管理 (HACS)**
  - 与 Home Assistant 社区商店集成
  - 支持多种包类型：
    - 自定义集成
    - 前端主题
    - Python 脚本
    - AppDaemon 应用
    - NetDaemon 应用
  - 版本控制和更新
  - 存储库管理

- **自动化管理**
  - 创建和编辑自动化
  - 高级配置选项：
    - 多种触发器类型
    - 复杂条件
    - 动作序列
    - 执行模式
  - 复制和修改现有自动化
  - 启用/禁用自动化规则
  - 手动触发自动化

### 架构特性 🏗️

- **智能组织**
  - 基于区域和楼层的设备分组
  - 状态监控和查询
  - 智能上下文感知
  - 历史数据访问

- **健壮架构**
  - 全面的错误处理
  - 状态验证
  - 安全 API 集成
  - TypeScript 类型安全
  - 广泛的测试覆盖

## 先决条件

- **Node.js** 20.10.0 或更高版本
- **NPM** 包管理器
- **Docker Compose** 用于容器化
- 运行中的 **Home Assistant** 实例
- Home Assistant 长期访问令牌（[如何获取令牌](https://community.home-assistant.io/t/how-to-get-long-lived-access-token/162159)）
- 已安装 **HACS** 用于包管理功能
- **Supervisor** 访问权限用于插件管理

## 安装

### 基础设置

```bash
# 克隆存储库
git clone https://github.com/jango-blockchained/homeassistant-mcp.git
cd homeassistant-mcp

# 安装依赖项
npm install

# 构建项目
npm run build
```

### Docker 设置（推荐）

项目包含 Docker 支持，便于部署和在不同平台上保持一致的环境。

1. **克隆存储库：**

    ```bash
    git clone https://github.com/jango-blockchained/homeassistant-mcp.git
    cd homeassistant-mcp
    ```

2. **配置环境：**

    ```bash
    cp .env.example .env
    ```

    使用您的 Home Assistant 配置编辑 `.env` 文件：

    ```env
    # Home Assistant 配置
    HASS_HOST=http://homeassistant.local:8123
    HASS_TOKEN=your_home_assistant_token
    HASS_SOCKET_URL=ws://homeassistant.local:8123/api/websocket

    # 服务器配置
    PORT=3000
    NODE_ENV=production
    DEBUG=false
    ```

3. **使用 Docker Compose 构建和运行：**

    ```bash
    # 构建并启动容器
    docker compose up -d

    # 查看日志
    docker compose logs -f

    # 停止服务
    docker compose down
    ```

4. **验证安装：**
    服务器现在应该在 `http://localhost:3000` 运行。您可以在 `http://localhost:3000/health` 检查健康端点。

5. **更新应用程序：**

    ```bash
    # 拉取最新更改
    git pull

    # 重新构建并重启容器
    docker compose up -d --build
    ```

#### Docker 配置

Docker 设置包括：

- 多阶段构建以优化镜像大小
- 容器监控的健康检查
- 环境配置的卷挂载
- 失败时自动重启容器
- 暴露端口 3000 用于 API 访问

#### Docker Compose 环境变量

所有环境变量都可以在 `.env` 文件中配置。支持以下变量：

- `HASS_HOST`：您的 Home Assistant 实例 URL
- `HASS_TOKEN`：Home Assistant 的长期访问令牌
- `HASS_SOCKET_URL`：Home Assistant 的 WebSocket URL
- `PORT`：服务器端口（默认：3000）
- `NODE_ENV`：环境（production/development）
- `DEBUG`：启用调试模式（true/false）

## 配置

### 环境变量

```env
# Home Assistant 配置
HASS_HOST=http://homeassistant.local:8123  # 您的 Home Assistant 实例 URL
HASS_TOKEN=your_home_assistant_token       # 长期访问令牌
HASS_SOCKET_URL=ws://homeassistant.local:8123/api/websocket  # WebSocket URL

# 服务器配置
PORT=3000                # 服务器端口（默认：3000）
NODE_ENV=production     # 环境（production/development）
DEBUG=false            # 启用调试模式

# 测试配置
TEST_HASS_HOST=http://localhost:8123  # 测试实例 URL
TEST_HASS_TOKEN=test_token           # 测试令牌
```

### 配置文件

1. **开发环境**：将 `.env.example` 复制为 `.env.development`
2. **生产环境**：将 `.env.example` 复制为 `.env.production`
3. **测试环境**：将 `.env.example` 复制为 `.env.test`

### 添加到 Claude Desktop（或其他客户端）

要使用您的新 Home Assistant MCP 服务器，您可以添加 Claude Desktop 作为客户端。将以下内容添加到配置中。注意这将在 claude 内运行 MCP，不适用于 Docker 方法。

```
{
  "homeassistant": {
    "command": "node",
    "args": [<path/to/your/dist/folder>]
    "env": {
      NODE_ENV=development
      HASS_HOST=http://homeassistant.local:8123
      HASS_TOKEN=your_home_assistant_token
      PORT=3000
      HASS_SOCKET_URL=ws://homeassistant.local:8123/api/websocket
      LOG_LEVEL=debug
    }
  }
}
```

## API 参考

### 设备控制

#### 通用实体控制

```json
{
  "tool": "control",
  "command": "turn_on",  // 或 "turn_off", "toggle"
  "entity_id": "light.living_room"
}
```

#### 灯光控制

```json
{
  "tool": "control",
  "command": "turn_on",
  "entity_id": "light.living_room",
  "brightness": 128,
  "color_temp": 4000,
  "rgb_color": [255, 0, 0]
}
```

### 插件管理

#### 列出可用插件

```json
{
  "tool": "addon",
  "action": "list"
}
```

#### 安装插件

```json
{
  "tool": "addon",
  "action": "install",
  "slug": "core_configurator",
  "version": "5.6.0"
}
```

#### 管理插件状态

```json
{
  "tool": "addon",
  "action": "start",  // 或 "stop", "restart"
  "slug": "core_configurator"
}
```

### 包管理

#### 列出 HACS 包

```json
{
  "tool": "package",
  "action": "list",
  "category": "integration"  // 或 "plugin", "theme", "python_script", "appdaemon", "netdaemon"
}
```

#### 安装包

```json
{
  "tool": "package",
  "action": "install",
  "category": "integration",
  "repository": "hacs/integration",
  "version": "1.32.0"
}
```

### 自动化管理

#### 创建自动化

```json
{
  "tool": "automation_config",
  "action": "create",
  "config": {
    "alias": "Motion Light",
    "description": "Turn on light when motion detected",
    "mode": "single",
    "trigger": [
      {
        "platform": "state",
        "entity_id": "binary_sensor.motion",
        "to": "on"
      }
    ],
    "action": [
      {
        "service": "light.turn_on",
        "target": {
          "entity_id": "light.living_room"
        }
      }
    ]
  }
}
```

#### 复制自动化

```json
{
  "tool": "automation_config",
  "action": "duplicate",
  "automation_id": "automation.motion_light"
}
```

### 核心功能

#### 状态管理

```http
GET /api/state
POST /api/state
```

管理系统的当前状态。

**示例请求：**

```json
POST /api/state
{
  "context": "living_room",
  "state": {
    "lights": "on",
    "temperature": 22
  }
}
```

#### 上下文更新

```http
POST /api/context
```

使用新信息更新当前上下文。

**示例请求：**

```json
POST /api/context
{
  "user": "john",
  "location": "kitchen",
  "time": "morning",
  "activity": "cooking"
}
```

### 动作端点

#### 执行动作

```http
POST /api/action
```

使用给定参数执行指定动作。

**示例请求：**

```json
POST /api/action
{
  "action": "turn_on_lights",
  "parameters": {
    "room": "living_room",
    "brightness": 80
  }
}
```

#### 批量动作

```http
POST /api/actions/batch
```

按顺序执行多个动作。

**示例请求：**

```json
POST /api/actions/batch
{
  "actions": [
    {
      "action": "turn_on_lights",
      "parameters": {
        "room": "living_room"
      }
    },
    {
      "action": "set_temperature",
      "parameters": {
        "temperature": 22
      }
    }
  ]
}
```

### 查询功能

#### 获取可用动作

```http
GET /api/actions
```

返回所有可用动作的列表。

**示例响应：**

```json
{
  "actions": [
    {
      "name": "turn_on_lights",
      "parameters": ["room", "brightness"],
      "description": "在指定房间打开灯光"
    },
    {
      "name": "set_temperature",
      "parameters": ["temperature"],
      "description": "在当前上下文中设置温度"
    }
  ]
}
```

#### 上下文查询

```http
GET /api/context?type=current
```

检索上下文信息。

**示例响应：**

```json
{
  "current_context": {
    "user": "john",
    "location": "kitchen",
    "time": "morning",
    "activity": "cooking"
  }
}
```

### WebSocket 事件

服务器通过 WebSocket 连接支持实时更新。

```javascript
// 客户端连接示例
const ws = new WebSocket('ws://localhost:3000/ws');

ws.onmessage = (event) => {
  const data = JSON.parse(event.data);
  console.log('收到更新:', data);
};
```

#### 支持的事件

- `state_change`：系统状态变化时发出
- `context_update`：上下文更新时发出
- `action_executed`：动作完成时发出
- `error`：发生错误时发出

**示例事件数据：**

```json
{
  "event": "state_change",
  "data": {
    "previous_state": {
      "lights": "off"
    },
    "current_state": {
      "lights": "on"
    },
    "timestamp": "2024-03-20T10:30:00Z"
  }
}
```

### 错误处理

所有端点返回标准 HTTP 状态码：

- 200：成功
- 400：错误请求
- 401：未授权
- 403：禁止访问
- 404：未找到
- 500：内部服务器错误

**错误响应格式：**

```json
{
  "error": {
    "code": "INVALID_PARAMETERS",
    "message": "缺少必需参数：room",
    "details": {
      "missing_fields": ["room"]
    }
  }
}
```

### 速率限制

API 实现速率限制以防止滥用：

- 常规端点每个 IP 每分钟 100 个请求
- WebSocket 连接每个 IP 每分钟 1000 个请求

当超过速率限制时，服务器返回：

```json
{
  "error": {
    "code": "RATE_LIMIT_EXCEEDED",
    "message": "请求过多",
    "reset_time": "2024-03-20T10:31:00Z"
  }
}
```

### 使用示例

#### 使用 curl

```bash
# 获取当前状态
curl -X GET \
  http://localhost:3000/api/state \
  -H 'Authorization: ApiKey your_api_key_here'

# 执行动作
curl -X POST \
  http://localhost:3000/api/action \
  -H 'Authorization: ApiKey your_api_key_here' \
  -H 'Content-Type: application/json' \
  -d '{
    "action": "turn_on_lights",
    "parameters": {
      "room": "living_room",
      "brightness": 80
    }
  }'
```

#### 使用 JavaScript

```javascript
// 执行动作
async function executeAction() {
  const response = await fetch('http://localhost:3000/api/action', {
    method: 'POST',
    headers: {
      'Authorization': 'ApiKey your_api_key_here',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      action: 'turn_on_lights',
      parameters: {
        room: 'living_room',
        brightness: 80
      }
    })
  });
  
  const data = await response.json();
  console.log('动作结果:', data);
}
```

## 开发

```bash
# 开发模式（热重载）
npm run dev

# 构建项目
npm run build

# 生产模式
npm run start

# 运行测试
npx jest --config=jest.config.cjs

# 运行测试并生成覆盖率报告
npx jest --coverage

# 代码检查
npm run lint

# 代码格式化
npm run format
```

## 故障排除

### 常见问题

1. **Node.js 版本问题（`toSorted is not a function`）**
   - **解决方案：** 更新到 Node.js 20.10.0+

   ```bash
   nvm install 20.10.0
   nvm use 20.10.0
   ```

2. **连接问题**
   - 验证 Home Assistant 正在运行
   - 检查 `HASS_HOST` 可访问性
   - 验证令牌权限
   - 确保 WebSocket 连接用于实时更新

3. **插件管理问题**
   - 验证 Supervisor 访问权限
   - 检查插件兼容性
   - 验证系统资源

4. **HACS 集成问题**
   - 验证 HACS 安装
   - 检查 HACS 集成状态
   - 验证存储库访问

5. **自动化问题**
   - 验证实体可用性
   - 检查触发条件
   - 验证服务调用
   - 监控执行日志

## 项目状态

✅ **已完成**

- 实体、楼层和区域访问
- 设备控制（灯光、气候、窗帘、开关、接触器）
- 插件管理系统
- 通过 HACS 的包管理
- 高级自动化配置
- 基本状态管理
- 错误处理和验证
- Docker 容器化
- Jest 测试设置
- TypeScript 集成
- 环境变量管理
- Home Assistant API 集成
- 项目文档

🚧 **进行中**

- 实时更新的 WebSocket 实现
- 增强的安全功能
- 工具组织优化
- 性能优化
- 资源上下文集成
- API 文档生成
- 多平台桌面集成
- 高级错误恢复
- 自定义提示测试
- 增强的 macOS 集成
- 类型安全改进
- 测试覆盖率扩展

## 贡献

1. Fork 存储库
2. 创建功能分支
3. 实现您的更改
4. 为新功能添加测试
5. 确保所有测试通过
6. 提交拉取请求

## 资源

- [MCP 文档](https://modelcontextprotocol.io/introduction)
- [Home Assistant 文档](https://www.home-assistant.io)
- [HA REST API](https://developers.home-assistant.io/docs/api/rest)
- [HACS 文档](https://hacs.xyz)
- [TypeScript 文档](https://www.typescriptlang.org/docs)
