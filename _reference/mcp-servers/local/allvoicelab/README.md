# AllVoiceLab MCP Server

> <https://github.com/allvoicelab/AllVoiceLab-MCP>

Official AllVoiceLab Model Context Protocol (MCP) server, supporting interaction with powerful text-to-speech and video translation APIs. Enables MCP clients like Claude Desktop, Cursor, Windsurf, OpenAI Agents to generate speech, translate videos, and perform intelligent voice conversion. Serves scenarios such as short drama localization for global markets, AI-Generated audiobooks, AI-Powered production of film/TV narration.

## Why Choose AllVoiceLab MCP Server?

- Multi-engine technology unlocks infinite possibilities for voice: With simple text input, you can access video generation, speech synthesis, voice cloning, and more.
- AI Voice Generator (TTS): Natural voice generation in 30+ languages with ultra-high realism
- Voice Changer: Real-time voice conversion, ideal for gaming, live streaming, and privacy protection
- Vocal Separation: Ultra-fast 5ms separation of vocals and background music, with industry-leading precision
- Multilingual Dubbing: One-click translation and dubbing for short videos/films, preserving emotional tone and rhythm
- Speech-to-Text (STT): AI-powered multilingual subtitle generation with over 98% accuracy
- Subtitle Removal: Seamless hard subtitle erasure, even on complex backgrounds
- Voice Cloning: 3-Second Ultra-Fast Cloning with Human-like Voice Synthesis

## Quickstart

1. Get your API key from [AllVoiceLab](https://www.allvoicelab.com/).
2. Install `uv` (Python package manager), install with `curl -LsSf https://astral.sh/uv/install.sh | sh`
3. **Important**: The server addresses of APIs in different regions need to match the keys of the corresponding regions, otherwise there will be an error that the tool is unavailable.

|Region| Global  | Mainland  |
|:--|:-----|:-----|
|ALLVOICELAB_API_KEY| go get from [AllVoiceLab](https://www.allvoicelab.com/workbench/api-keys) | go get from [AllVoiceLab](https://www.allvoicelab.cn/workbench/api-keys) |
|ALLVOICELAB_API_DOMAIN| https://api.allvoicelab.com | https://api.allvoicelab.cn |

### Claude Desktop

Go to Claude > Settings > Developer > Edit Config > claude_desktop_config.json to include the following:
```json
{
  "mcpServers": {
    "AllVoiceLab": {
      "command": "uvx",
      "args": ["allvoicelab-mcp"],
      "env": {
        "ALLVOICELAB_API_KEY": "<insert-your-api-key-here>",
        "ALLVOICELAB_API_DOMAIN": "<insert-api-domain-here>",
        "ALLVOICELAB_BASE_PATH":"optional, default is user home directory.This is uesd to store the output files."
      }
    }
  }
}
```

If you're using Windows, you will have to enable "Developer Mode" in Claude Desktop to use the MCP server. Click "Help" in the hamburger menu in the top left and select "Enable Developer Mode".

### Cursor
Go to Cursor -> Preferences -> Cursor Settings -> MCP -> Add new global MCP Server to add above config.

That's it. Your MCP client can now interact with AllVoiceLab.


## Available methods

| Methods | Brief description |
| --- | --- |
| text_to_speech | Convert text to speech |
| speech_to_speech | Convert audio to another voice while preserving the speech content |
| isolate_human_voice | Extract clean human voice by removing background noise and non-speech sounds |
| clone_voice | Create a custom voice profile by cloning from an audio sample |
| remove_subtitle | Remove hardcoded subtitles from a video using OCR |
| video_translation_dubbing | Translate and dub video speech into different languages ​​|
| text_translation | Translate a text file into another language |
| subtitle_extraction | Extract subtitles from a video using OCR |
