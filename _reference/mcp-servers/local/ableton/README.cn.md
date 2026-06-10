# AbletonMCP - Ableton Live 模型上下文协议集成

> <https://github.com/ahujasid/ableton-mcp>

AbletonMCP 通过模型上下文协议（MCP）将 Ableton Live 连接到 Claude AI，允许 Claude 直接与 Ableton Live 交互并控制它。这种集成实现了提示辅助的音乐制作、音轨创建和 Live 会话操作。

### 加入社区

提供反馈、获得灵感，并在 MCP 基础上构建：[Discord](https://discord.gg/3ZrMyGKnaU)。由 [Siddharth](https://x.com/sidahuj) 制作

## 功能特性

- **双向通信**：通过基于套接字的服务器将 Claude AI 连接到 Ableton Live
- **音轨操作**：创建、修改和操作 MIDI 和音频音轨
- **乐器和效果选择**：Claude 可以访问并从 Ableton 的音色库中加载正确的乐器、效果和声音
- **片段创建**：创建和编辑带有音符的 MIDI 片段
- **会话控制**：开始和停止播放、触发片段以及控制传输

## 组件

系统由两个主要组件组成：

1. **Ableton 远程脚本** (`Ableton_Remote_Script/__init__.py`)：Ableton Live 的 MIDI 远程脚本，创建套接字服务器来接收和执行命令
2. **MCP 服务器** (`server.py`)：实现模型上下文协议并连接到 Ableton 远程脚本的 Python 服务器

## 安装

### 通过 Smithery 安装

通过 [Smithery](https://smithery.ai/server/@ahujasid/ableton-mcp) 自动为 Claude Desktop 安装 Ableton Live 集成：

```bash
npx -y @smithery/cli install @ahujasid/ableton-mcp --client claude
```

### 前提条件

- Ableton Live 10 或更新版本
- Python 3.8 或更新版本
- [uv 包管理器](https://astral.sh/uv)

如果您使用 Mac，请按以下方式安装 uv：

```
brew install uv
```

否则，请从 [uv 官方网站](https://docs.astral.sh/uv/getting-started/installation/) 安装

⚠️ 安装 UV 之前请勿继续

### Claude Desktop 集成

[观看设置说明视频](https://youtu.be/iJWJqyVuPS8)

1. 转到 Claude > 设置 > 开发者 > 编辑配置 > claude_desktop_config.json，包含以下内容：

```json
{
    "mcpServers": {
        "AbletonMCP": {
            "command": "uvx",
            "args": [
                "ableton-mcp"
            ]
        }
    }
}
```

### Cursor 集成

通过 uvx 运行 ableton-mcp 而无需永久安装。转到 Cursor 设置 > MCP 并将此作为命令粘贴：

```
uvx ableton-mcp
```

⚠️ 只运行一个 MCP 服务器实例（Claude Desktop 或 Cursor），不要同时运行两个

### 安装 Ableton 远程脚本

[观看设置说明视频](https://youtu.be/iJWJqyVuPS8)

1. 从此仓库下载 `AbletonMCP_Remote_Script/__init__.py` 文件

2. 将文件夹复制到 Ableton 的 MIDI 远程脚本目录。不同的操作系统和版本有不同的位置。**其中一个应该有效，您可能需要查找**：

   **对于 macOS：**
   - 方法 1：转到应用程序 > 右键单击 Ableton Live 应用 → 显示包内容 → 导航到：
     `Contents/App-Resources/MIDI Remote Scripts/`
   - 方法 2：如果第一种方法中没有，请使用直接路径（将 XX 替换为您的版本号）：
     `/Users/[用户名]/Library/Preferences/Ableton/Live XX/User Remote Scripts`

   **对于 Windows：**
   - 方法 1：
     C:\Users\[用户名]\AppData\Roaming\Ableton\Live x.x.x\Preferences\User Remote Scripts
   - 方法 2：
     `C:\ProgramData\Ableton\Live XX\Resources\MIDI Remote Scripts\`
   - 方法 3：
     `C:\Program Files\Ableton\Live XX\Resources\MIDI Remote Scripts\`
   *注意：将 XX 替换为您的 Ableton 版本号（例如，10、11、12）*

3. 在远程脚本目录中创建一个名为 'AbletonMCP' 的文件夹，并粘贴下载的 '\_\_init\_\_.py' 文件

4. 启动 Ableton Live

5. 转到设置/首选项 → 链接、节拍和 MIDI

6. 在控制面板下拉菜单中，选择 "AbletonMCP"

7. 将输入和输出设置为 "无"

## 使用方法

### 启动连接

1. 确保 Ableton 远程脚本已在 Ableton Live 中加载
2. 确保 MCP 服务器已在 Claude Desktop 或 Cursor 中配置
3. 当您与 Claude 交互时，连接应该会自动建立

### 与 Claude 一起使用

一旦在 Claude 上设置了配置文件，并且远程脚本在 Ableton 中运行，您将看到一个带有 Ableton MCP 工具的锤子图标。

## 功能

- 获取会话和音轨信息
- 创建和修改 MIDI 和音频音轨
- 创建、编辑和触发片段
- 控制播放
- 从 Ableton 的浏览器加载乐器和效果
- 向 MIDI 片段添加音符
- 更改节拍和其他会话参数

## 示例命令

以下是您可以要求 Claude 执行的一些示例：

- "创建一个 80 年代合成波音轨" [演示](https://youtu.be/VH9g66e42XA)
- "创建一个 Metro Boomin 风格的嘻哈节拍"
- "创建一个带有合成贝斯乐器的新 MIDI 音轨"
- "为我的鼓添加混响"
- "创建一个 4 小节的 MIDI 片段，包含简单的旋律"
- "获取当前 Ableton 会话的信息"
- "将 808 鼓架加载到选定的音轨中"
- "向音轨 1 的片段添加爵士和弦进行"
- "将节拍设置为 120 BPM"
- "播放音轨 2 中的片段"

## 故障排除

- **连接问题**：确保 Ableton 远程脚本已加载，并且 MCP 服务器已在 Claude 上配置
- **超时错误**：尝试简化您的请求或将其分解为更小的步骤
- **试过重启吗？**：如果您仍然遇到连接错误，请尝试重启 Claude 和 Ableton Live

## 技术细节

### 通信协议

系统使用基于 TCP 套接字的简单 JSON 协议：

- 命令作为带有 `type` 和可选 `params` 的 JSON 对象发送
- 响应是带有 `status` 和 `result` 或 `message` 的 JSON 对象

### 限制和安全考虑

- 创建复杂的音乐编排可能需要分解为更小的步骤
- 该工具设计用于 Ableton 的默认设备和浏览器项目
- 在进行大量实验之前，请始终保存您的工作

## 贡献

欢迎贡献！请随时提交拉取请求。

## 免责声明

这是第三方集成，不是由 Ableton 制作的。
