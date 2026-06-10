# 获取 MCP 服务器

> <https://github.com/modelcontextprotocol/servers/tree/main/src/fetch>

一个提供网页内容获取功能的模型上下文协议服务器。此服务器使LLM能够从网页检索和处理内容，将HTML转换为markdown以便于使用。

> [!注意]
> 此服务器可以访问本地/内部IP地址，可能存在安全风险。使用此MCP服务器时请谨慎，确保不会暴露任何敏感数据。

fetch工具会截断响应，但通过使用start_index参数，您可以指定从哪里开始内容提取。这让模型可以分块读取网页，直到找到所需信息。

## 可用工具

- `fetch` - 从互联网获取URL并将其内容提取为markdown。
  - `url` (string, 必需): 要获取的URL
  - `max_length` (integer, 可选): 返回的最大字符数 (默认: 5000)
  - `start_index` (integer, 可选): 从此字符索引开始内容 (默认: 0)
  - `raw` (boolean, 可选): 获取原始内容而不进行markdown转换 (默认: false)
