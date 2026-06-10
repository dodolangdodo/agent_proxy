# AllVoiceLab MCP 服务器

> <https://github.com/allvoicelab/AllVoiceLab-MCP>

官方 AllVoiceLab 模型上下文协议 (MCP) 服务器，支持与强大的文本转语音和视频翻译 API 交互。使 MCP 客户端如 Claude Desktop、Cursor、Windsurf、OpenAI Agents 能够生成语音、翻译视频并执行智能语音转换。适用于短剧全球市场本地化、AI 生成有声书、AI 驱动的影视旁白制作等场景。

## 为什么选择 AllVoiceLab MCP 服务器？

- 多引擎技术解锁语音无限可能：通过简单的文本输入，您可以访问视频生成、语音合成、语音克隆等功能。
- AI 语音生成器 (TTS)：支持 30+ 种语言的自然语音生成，具有超高真实感
- 变声器：实时语音转换，非常适合游戏、直播和隐私保护
- 人声分离：超快 5ms 人声和背景音乐分离，具有行业领先的精度
- 多语言配音：一键翻译和配音短视频/电影，保持情感语调和节奏
- 语音转文本 (STT)：AI 驱动的多语言字幕生成，准确率超过 98%
- 字幕移除：无缝硬字幕擦除，即使在复杂背景上也能完美处理
- 语音克隆：3 秒超快克隆，具有类人语音合成

## 快速开始

1. 从 [AllVoiceLab](https://www.allvoicelab.com/) 获取您的 API 密钥。
2. 安装 `uv`（Python 包管理器），使用 `curl -LsSf https://astral.sh/uv/install.sh | sh` 安装
3. **重要**：不同地区 API 的服务器地址需要与相应地区的密钥匹配，否则会出现工具不可用的错误。

|地区| 全球  | 中国大陆  |
|:--|:-----|:-----|
|ALLVOICELAB_API_KEY| 从 [AllVoiceLab](https://www.allvoicelab.com/workbench/api-keys) 获取 | 从 [AllVoiceLab](https://www.allvoicelab.cn/workbench/api-keys) 获取 |
|ALLVOICELAB_API_DOMAIN| <https://api.allvoicelab.com> | <https://api.allvoicelab.cn> |

### Claude Desktop

转到 Claude > 设置 > 开发者 > 编辑配置 > claude_desktop_config.json 包含以下内容：

```json
{
  "mcpServers": {
    "AllVoiceLab": {
      "command": "uvx",
      "args": ["allvoicelab-mcp"],
      "env": {
        "ALLVOICELAB_API_KEY": "<在此插入您的API密钥>",
        "ALLVOICELAB_API_DOMAIN": "<在此插入API域名>",
        "ALLVOICELAB_BASE_PATH":"可选，默认为用户主目录。用于存储输出文件。"
      }
    }
  }
}
```

如果您使用的是 Windows，您需要在 Claude Desktop 中启用"开发者模式"才能使用 MCP 服务器。点击左上角汉堡菜单中的"帮助"并选择"启用开发者模式"。

### Cursor

转到 Cursor -> 首选项 -> Cursor 设置 -> MCP -> 添加新的全局 MCP 服务器，添加上述配置。

就是这样。您的 MCP 客户端现在可以与 AllVoiceLab 交互了。

## 可用方法

| 方法 | 简要描述 |
| --- | --- |
| text_to_speech | 将文本转换为语音 |
| speech_to_speech | 将音频转换为另一种声音，同时保留语音内容 |
| isolate_human_voice | 通过去除背景噪音和非语音声音提取清晰的人声 |
| clone_voice | 通过从音频样本克隆创建自定义语音配置文件 |
| remove_subtitle | 使用 OCR 从视频中移除硬编码字幕 |
| video_translation_dubbing | 将视频语音翻译并配音为不同语言 |
| text_translation | 将文本文件翻译为另一种语言 |
| subtitle_extraction | 使用 OCR 从视频中提取字幕 |
