# 时间 MCP 服务器

> <https://github.com/modelcontextprotocol/servers/tree/main/src/time>

一个提供时间和时区转换功能的模型上下文协议服务器。此服务器使LLM能够使用IANA时区名称获取当前时间信息并执行时区转换，具有自动系统时区检测功能。

### 可用工具

- `get_current_time` - 获取特定时区或系统时区的当前时间。
  - 必需参数:
    - `timezone` (string): IANA时区名称 (例如, 'America/New_York', 'Europe/London')

- `convert_time` - 在时区之间转换时间。
  - 必需参数:
    - `source_timezone` (string): 源IANA时区名称
    - `time` (string): 24小时格式的时间 (HH:MM)
    - `target_timezone` (string): 目标IANA时区名称
