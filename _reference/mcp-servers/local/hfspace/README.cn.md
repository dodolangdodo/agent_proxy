# mcp-hfspace MCP 服务器 🤗

> [!TIP]
>
> 您可以直接在 <https://hf.co/mcp> 访问和配置 Hugging Face MCP 服务，包括 Gradio spaces。
>
> 此项目已被官方的 [Hugging Face MCP Server](https://github.com/evalstate/hf-mcp-server) 和 [Gradio MCP Endpoints](https://huggingface.co/blog/gradio-mcp) 所取代。
>
> 或者您可以在本地运行 hf-mcp-server 作为 STDIO 服务器，或者使用对 SSE、流式 HTTP 和流式 HTTP JSON 模式的强大支持。这还运行一个本地 UI 来选择工具和端点，并且也支持 `ToolListChangedNotifications`。

## hf.co/mcp

![image](https://github.com/user-attachments/assets/9cbf407b-2330-4330-8274-e47305a555b9)

## mcp-hfspace

在这里阅读介绍 [llmindset.co.uk/resources/mcp-hfspace/](https://llmindset.co.uk/resources/mcp-hfspace/)

连接到 [Hugging Face Spaces](https://huggingface.co/spaces)，只需最少的设置 - 只需添加您的 spaces 即可开始！

默认情况下，它连接到 `black-forest-labs/FLUX.1-schnell`，为 Claude Desktop 提供图像生成功能。

![默认设置](./images/2024-12-09-flower.png)

## Gradio MCP 支持

> [!TIP]
> Gradio 5.28 现在通过 SSE 集成了 MCP 支持：<https://huggingface.co/blog/gradio-mcp>。检查您的目标 Space 是否启用了 MCP！

## 安装

NPM 包名为 `@llmindset/mcp-hfspace`。

为您的平台安装最新版本的 [NodeJS](https://nodejs.org/en/download)，然后将以下内容添加到您的 `claude_desktop_config.json` 文件的 `mcpServers` 部分：

```json
    "mcp-hfspace": {
      "command": "npx",
      "args": [
        "-y",
        "@llmindset/mcp-hfspace"
      ]
    }
```

请确保您使用的是 Claude Desktop 0.78 或更高版本。

这将为您启动一个图像生成器。

### 基本设置

在参数中提供 HuggingFace spaces 列表。mcp-hfspace 将找到最合适的端点并自动配置它以供使用。下面提供了一个示例 `claude_desktop_config.json`（[见安装部分](#installation)）。

默认情况下，当前工作目录用于文件上传/下载。在 Windows 上，这是位于 `\users\<username>\AppData\Roaming\Claude\<version.number\` 的读/写文件夹，在 MacOS 上是只读根目录：`/`。

建议覆盖此设置并设置一个工作目录来处理图像和其他基于文件的内容的上传和下载。指定 `--work-dir=/your_directory` 参数或 `MCP_HF_WORK_DIR` 环境变量。

以下是使用现代图像生成器、视觉模型和文本转语音，并设置工作目录的配置示例：

```json
    "mcp-hfspace": {
      "command": "npx",
      "args": [
        "-y",
        "@llmindset/mcp-hfspace",
        "--work-dir=/Users/evalstate/mcp-store",
        "shuttleai/shuttle-jaguar",
        "styletts2/styletts2",
        "Qwen/QVQ-72B-preview"
      ]
    }
```

要使用私有 spaces，请使用 `--hf-token=hf_...` 参数或 `HF_TOKEN` 环境变量提供您的 Hugging Face Token。

如果需要，可以运行多个服务器实例来使用不同的工作目录和令牌。

## 文件处理和 Claude Desktop 模式

默认情况下，服务器在 _Claude Desktop 模式_ 下运行。在此模式下，图像在工具响应中返回，而其他文件保存在工作文件夹中，它们的文件路径作为消息返回。如果使用 Claude Desktop 作为客户端，这通常会提供最佳体验。

URL 也可以作为输入提供：内容会传递给 Space。

有一个"可用资源"提示，它向 Claude 提供工作目录中的可用文件和 MIME 类型。这目前是管理文件的最佳方式。

### 示例 1 - 图像生成（下载图像 / Claude 视觉）

我们将使用 Claude 比较由 `shuttleai/shuttle-3.1-aesthetic` 和 `FLUX.1-schnell` 创建的图像。图像保存到工作目录，同时包含在 Claude 的上下文窗口中 - 因此 Claude 可以使用其视觉功能。

![图像生成比较](./images/2024-12-05-flux-shuttle.png)

### 示例 2 - 视觉模型（上传图像）

我们将使用 `merve/paligemma2-vqav2` [space 链接](https://huggingface.co/spaces/merve/paligemma2-vqav2) 来查询图像。在这种情况下，我们指定工作目录中可用的文件名：我们不想将图像直接上传到 Claude 的上下文窗口。因此，我们可以提示 Claude：

`use paligemma to find out who is in "test_gemma.jpg"` -> `文本输出：david bowie`
![视觉 - 文件上传](./images/2024-12-09-bowie.png)

_如果您要上传某些内容到 Claude 的上下文，请使用回形针附件按钮，否则指定文件名让服务器直接发送。_

我们也可以提供 URL。例如：`use paligemma to detect humans in https://e3.365dm.com/24/12/1600x900/skynews-taylor-swift-eras-tour_6771083.jpg?20241209000914` -> `图像中检测到一个人 - 舞台上的 Taylor Swift。`

### 示例 3 - 文本转语音（下载音频）

在 _Claude Desktop 模式_ 下，音频文件保存在 WORK_DIR 中，并通知 Claude 创建。如果不在桌面模式下，文件作为 base64 编码资源返回给客户端（如果支持嵌入式音频附件则很有用）。

![语音制作](./images/2024-12-08-mcp-parler.png)

### 示例 4 - 语音转文本（上传音频）

在这里，我们使用 `hf-audio/whisper-large-v3-turbo` 来转录一些音频，并使其对 Claude 可用。

![音频转录](./images/2024-12-09-transcribe.png)

### 示例 5 - 图像到图像

在此示例中，我们为 `microsoft/OmniParser` 指定要使用的文件名，并获得返回的注释图像和 2 个单独的文本片段：描述和坐标。使用的提示是 `use omniparser to analyse ./screenshot.png` 和 `use the analysis to produce an artifact that reproduces that screen`。`DawnC/Pawmatch` 在这方面也很出色。

![Omniparser 和 Artifact](./images/2024-12-08-mcp-omni-artifact.png)

### 示例 6 - 聊天

在此示例中，Claude 为 Qwen 设置了一些推理谜题，并提出后续问题以进行澄清。

![Qwen 推理测试](./images/2024-12-09-qwen-reason.png)

### 指定 API 端点

如果需要，您可以通过将其添加到 spacename 来指定特定的 API 端点。因此，不是传入 `Qwen/Qwen2.5-72B-Instruct`，您应该使用 `Qwen/Qwen2.5-72B-Instruct/model_chat`。

### Claude Desktop 模式

这可以通过选项 --desktop-mode=false 或环境变量 CLAUDE_DESKTOP_MODE=false 来禁用。在这种情况下，内容作为嵌入式 Base64 编码资源返回。

## 推荐的 Spaces

一些推荐尝试的 spaces：

### 图像生成

- shuttleai/shuttle-3.1-aesthetic
- black-forest-labs/FLUX.1-schnell
- yanze/PuLID-FLUX
- gokaygokay/Inspyrenet-Rembg（背景移除）
- diyism/Datou1111-shou_xin - [美丽的铅笔画](https://x.com/ClementDelangue/status/1867318931502895358)

### 聊天

- Qwen/Qwen2.5-72B-Instruct
- prithivMLmods/Mistral-7B-Instruct-v0.3

### 文本转语音 / 音频生成

- fantaxy/Sound-AI-SFX
- parler-tts/parler_tts

### 语音转文本

- hf-audio/whisper-large-v3-turbo
- （openai 模型使用未命名参数，因此无法工作）

### 文本转音乐

- haoheliu/audioldm2-text2audio-text2music

### 视觉任务

- microsoft/OmniParser
- merve/paligemma2-vqav2
- merve/paligemma-doc
- DawnC/PawMatchAI
- DawnC/PawMatchAI/on_find_match_click - 用于交互式狗狗推荐

## 其他功能

### 提示

为每个 Space 生成提示，并提供输入机会。请记住，Spaces 通常没有配置特别有用的标签等。Claude 实际上非常擅长解决这个问题，工具描述非常丰富（但在 Claude Desktop 中不可见）。

### 资源

返回 WORK_DIR 中的文件列表，并方便地将名称返回为"使用文件..."文本。如果您想向 Claude 的上下文添加某些内容，请使用回形针 - 否则为 MCP 服务器指定文件名。Claude 不支持从上下文内传输资源。

### 私有 Spaces

使用 HuggingFace 令牌支持私有 Spaces。令牌用于下载和保存生成的内容。

### 使用 Claude Desktop

要与 Claude Desktop 一起使用，请添加服务器配置：

在 MacOS 上：`~/Library/Application Support/Claude/claude_desktop_config.json`
在 Windows 上：`%APPDATA%/Claude/claude_desktop_config.json`

```json
{
  "mcpServers": {
    "mcp-hfspace": {
      "command": "npx",
      "args": [
        "-y",
        "@llmindset/mcp-hfspace",
        "--work-dir=~/mcp-files/ 或 x:/temp/mcp-files/",
        "--HF_TOKEN=HF_{可选令牌}",
        "Qwen/Qwen2-72B-Instruct",
        "black-forest-labs/FLUX.1-schnell",
        "space/example/specific-endpint"
        "（... 等等）"
        ]
    }
  }
}
```

## 已知问题和限制

### mcp-hfspace

- 目前不支持具有未命名参数的端点。
- 从一些复杂的 Python 类型到合适的 MCP 格式的完整转换。

### Claude Desktop

- Claude Desktop 0.75 似乎不响应来自 MCP 服务器的错误，而是超时。对于持续问题，请使用 MCP Inspector 来更好地诊断问题。如果某些功能突然停止工作，可能是由于耗尽了您的 HuggingFace ZeroGPU 配额 - 请在短暂时间后重试，或设置您自己的 Space 进行托管。
- Claude Desktop 似乎使用 60 秒的硬超时值，并且似乎不使用进度通知来管理 UX 或保持连接。如果您使用 ZeroGPU spaces，大型/重型作业可能会超时。不过请检查 WORK_DIR 的结果；如果产生了结果，MCP 服务器仍会捕获并保存结果。
- Claude Desktop 的服务器状态、日志记录等报告不是很好 - 使用 [@modelcontextprotocol/inspector](https://github.com/modelcontextprotocol/inspector) 来帮助诊断问题。

### HuggingFace Spaces

- 如果 ZeroGPU 配额或队列太长，请尝试复制 space。如果您的作业耗时少于六十秒，您通常可以在 `app.py` 中更改函数装饰器 `@spaces.GPU(duration=20)` 以在运行作业时请求更少的配额。
- 传递 HF_TOKEN 将使 ZeroGPU 配额适用于您的（Pro）HF 账户
- 如果您有私有 space 和专用硬件，您的 HF_TOKEN 将为您提供直接访问权限 - 不适用配额。如果您将其用于任何类型的生产任务，我建议这样做。

## 第三方 MCP 服务

<a href="https://glama.ai/mcp/servers/s57c80wvgq"><img width="380" height="200" src="https://glama.ai/mcp/servers/s57c80wvgq/badge" alt="mcp-hfspace MCP server" /></a>
