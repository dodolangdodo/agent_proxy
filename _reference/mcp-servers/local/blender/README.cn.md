# BlenderMCP - Blender模型上下文协议集成

> <https://github.com/ahujasid/blender-mcp>

BlenderMCP通过模型上下文协议(MCP)将Blender连接到Claude AI，允许Claude直接与Blender交互并控制Blender。这种集成实现了提示辅助的3D建模、场景创建和操作。

[完整教程](https://www.youtube.com/watch?v=lCyQ717DuzQ)

### 加入社区

提供反馈、获得灵感，并在MCP基础上构建：[Discord](https://discord.gg/z5apgR8TFU)

### 支持者

**顶级支持者：**

[CodeRabbit](https://www.coderabbit.ai/)

**所有支持者：**

[支持此项目](https://github.com/sponsors/ahujasid)

## 发布说明 (1.2.0)

- 查看Blender视口截图以更好地理解场景
- 搜索和下载Sketchfab模型

### 之前添加的功能

- 通过API支持Poly Haven资源
- 支持使用Hyper3D Rodin生成3D模型
- 新用户可以直接跳到安装部分。现有用户请参考以下要点
- 下载最新的addon.py文件并替换旧文件，然后将其添加到Blender
- 从Claude中删除MCP服务器并重新添加，就可以正常使用了！

## 功能特性

- **双向通信**：通过基于socket的服务器将Claude AI连接到Blender
- **对象操作**：在Blender中创建、修改和删除3D对象
- **材质控制**：应用和修改材质和颜色
- **场景检查**：获取当前Blender场景的详细信息
- **代码执行**：从Claude在Blender中运行任意Python代码

## 组件

系统由两个主要组件组成：

1. **Blender插件 (`addon.py`)**：在Blender中创建socket服务器以接收和执行命令的Blender插件
2. **MCP服务器 (`src/blender_mcp/server.py`)**：实现模型上下文协议并连接到Blender插件的Python服务器

## 安装

### 前置要求

- Blender 3.0或更新版本
- Python 3.10或更新版本
- uv包管理器：

**如果你使用Mac，请按如下方式安装uv**

```bash
brew install uv
```

**在Windows上**

```bash
powershell -c "irm https://astral.sh/uv/install.ps1 | iex" 
```

然后

```bash
set Path=C:\Users\nntra\.local\bin;%Path%
```

其他安装说明请参考官网：[安装uv](https://docs.astral.sh/uv/getting-started/installation/)

**⚠️ 安装UV之前请勿继续**

### Claude桌面版集成

[观看设置说明视频](https://www.youtube.com/watch?v=neoK_WMq92g)（假设你已经安装了uv）

转到Claude > 设置 > 开发者 > 编辑配置 > claude_desktop_config.json，包含以下内容：

```json
{
    "mcpServers": {
        "blender": {
            "command": "uvx",
            "args": [
                "blender-mcp"
            ]
        }
    }
}
```

### Cursor集成

对于Mac用户，转到设置 > MCP并粘贴以下内容

- 要用作全局服务器，使用"添加新的全局MCP服务器"按钮并粘贴
- 要用作项目特定服务器，在项目根目录创建`.cursor/mcp.json`并粘贴

```json
{
    "mcpServers": {
        "blender": {
            "command": "uvx",
            "args": [
                "blender-mcp"
            ]
        }
    }
}
```

对于Windows用户，转到设置 > MCP > 添加服务器，使用以下设置添加新服务器：

```json
{
    "mcpServers": {
        "blender": {
            "command": "cmd",
            "args": [
                "/c",
                "uvx",
                "blender-mcp"
            ]
        }
    }
}
```

[Cursor设置视频](https://www.youtube.com/watch?v=wgWsJshecac)

**⚠️ 只运行一个MCP服务器实例（Cursor或Claude桌面版），不要同时运行两个**

### 安装Blender插件

1. 从此仓库下载`addon.py`文件
2. 打开Blender
3. 转到编辑 > 首选项 > 插件
4. 点击"安装..."并选择`addon.py`文件
5. 通过勾选"Interface: Blender MCP"旁边的复选框来启用插件

## 使用方法

### 开始连接

1. 在Blender中，转到3D视图侧边栏（如果不可见请按N键）
2. 找到"BlenderMCP"选项卡
3. 如果你想要来自Poly Haven API的资源，请打开Poly Haven复选框（可选）
4. 点击"连接到Claude"
5. 确保MCP服务器在你的终端中运行

### 与Claude一起使用

在Claude上设置配置文件并在Blender上运行插件后，你将看到一个锤子图标，其中包含Blender MCP的工具。

#### 功能

- 获取场景和对象信息
- 创建、删除和修改形状
- 为对象应用或创建材质
- 在Blender中执行任何Python代码
- 通过[Poly Haven](https://polyhaven.com/)下载正确的模型、资源和HDRI
- 通过[Hyper3D Rodin](https://hyper3d.ai/)生成AI 3D模型

### 示例命令

以下是一些你可以要求Claude执行的示例：

- "在地牢中创建一个低多边形场景，有一条龙守护着一罐金子" [演示](https://www.youtube.com/watch?v=DqgKuLYUv00)
- "使用HDRI、纹理和来自Poly Haven的岩石和植被等模型创建海滩氛围" [演示](https://www.youtube.com/watch?v=I29rn92gkC4)
- 给出参考图像，并从中创建Blender场景 [演示](https://www.youtube.com/watch?v=FDRb03XPiRo)
- "通过Hyper3D生成花园小矮人的3D模型"
- "获取当前场景信息，并从中制作threejs草图" [演示](https://www.youtube.com/watch?v=jxbNI5L7AH8)
- "让这辆车变成红色和金属质感"
- "创建一个球体并将其放在立方体上方"
- "让照明像工作室一样"
- "将相机指向场景，并使其等距"

## Hyper3D集成

Hyper3D的免费试用密钥允许你每天生成有限数量的模型。如果达到每日限制，你可以等待第二天重置或从hyper3d.ai和fal.ai获取自己的密钥。

## 故障排除

- **连接问题**：确保Blender插件服务器正在运行，并且MCP服务器在Claude上配置，不要在终端中运行uvx命令。有时，第一个命令不会通过，但之后就开始工作了。
- **超时错误**：尝试简化你的请求或将其分解为较小的步骤
- **Poly Haven集成**：Claude有时行为不稳定
- **你试过关机再开机吗？**：如果你仍然有连接错误，尝试重启Claude和Blender服务器

## 技术细节

### 通信协议

系统使用基于TCP socket的简单JSON协议：

- **命令**作为具有`type`和可选`params`的JSON对象发送
- **响应**是具有`status`和`result`或`message`的JSON对象

## 限制和安全考虑

- `execute_blender_code`工具允许在Blender中运行任意Python代码，这可能很强大但也可能很危险。在生产环境中请谨慎使用。使用前请始终保存你的工作。
- Poly Haven需要下载模型、纹理和HDRI图像。如果你不想使用它，请在Blender中的复选框中关闭它。
- 复杂操作可能需要分解为较小的步骤

## 贡献

欢迎贡献！请随时提交Pull Request。

## 免责声明

这是第三方集成，不是由Blender制作的。由[Siddharth](https://x.com/sidahuj)制作
