# Figma MCP 服务器

> <https://github.com/GLips/Figma-Context-MCP>

提供对Figma设计文件访问的模型上下文协议服务器。此服务器使LLM能够检索和处理Figma设计数据，将其转换为简化格式以便于使用。

## 功能特性

- **设计文件访问**: 检索完整的Figma设计文件或特定节点
- **简化数据格式**: 将复杂的Figma数据转换为简化的、AI友好的格式
- **图像下载**: 从Figma设计中下载SVG和PNG图像
- **灵活认证**: 支持个人访问令牌和OAuth认证
- **多种输出格式**: 支持YAML和JSON输出格式

## 设置

### 前提条件

1. 在 <https://www.figma.com/developers/api#access-tokens> 创建Figma个人访问令牌
2. 或为您的应用程序设置OAuth认证

### 配置

服务器需要以下配置：

- `figma-api-key` (如果不使用OAuth则必需): 您的Figma个人访问令牌
- `figma-oauth-token` (可选): 您的Figma OAuth Bearer令牌（优先于API密钥）
- `output-format` (可选): 设计数据的输出格式（yaml或json，默认：yaml）

## 可用工具

### get_figma_data

检索Figma文件或特定节点的布局信息。

**参数:**

- `fileKey` (必需): 要获取的Figma文件的键
- `nodeId` (可选): 要获取的特定节点的ID
- `depth` (可选): 遍历节点树的深度级别

### download_figma_images

从Figma设计中下载SVG和PNG图像。

**参数:**

- `fileKey` (必需): 包含节点的Figma文件的键
- `nodes` (必需): 要作为图像获取的节点数组
- `localPath` (必需): 图像保存的目录路径
- `pngScale` (可选): PNG图像的导出比例（默认：2）
- `svgOptions` (可选): SVG导出选项

## 认证

服务器支持两种认证方法：

1. **个人访问令牌**: 设置 `figma-api-key` 配置
2. **OAuth Bearer令牌**: 设置 `figma-oauth-token` 配置（优先级更高）

## 输出格式

服务器可以输出两种格式的数据：

- **YAML** (默认): 人类可读格式
- **JSON**: 机器可读格式

设置 `output-format` 配置来选择您喜欢的格式。
