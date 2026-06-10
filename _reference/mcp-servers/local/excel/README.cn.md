# Excel MCP 服务器

> <https://github.com/negokaz/excel-mcp-server>

<img src="https://github.com/negokaz/excel-mcp-server/blob/main/docs/img/icon-800.png?raw=true" width="128">

[![NPM Version](https://img.shields.io/npm/v/@negokaz/excel-mcp-server)](https://www.npmjs.com/package/@negokaz/excel-mcp-server)
[![smithery badge](https://smithery.ai/badge/@negokaz/excel-mcp-server)](https://smithery.ai/server/@negokaz/excel-mcp-server)

一个用于读取和写入 MS Excel 数据的模型上下文协议 (MCP) 服务器。

## 功能特性

- 读取/写入文本值
- 读取/写入公式
- 创建新工作表

**🪟仅限 Windows：**

- 实时编辑
- 从工作表捕获屏幕图像

更多详细信息，请参见[工具](#tools)部分。

## 系统要求

- Node.js 20.x 或更高版本

## 支持的文件格式

- xlsx (Excel 工作簿)
- xlsm (Excel 启用宏的工作簿)
- xltx (Excel 模板)
- xltm (Excel 启用宏的模板)

## 安装

### 通过 NPM 安装

通过在 MCP 服务器配置中添加以下配置，excel-mcp-server 会自动安装。

Windows 系统：

```json
{
    "mcpServers": {
        "excel": {
            "command": "cmd",
            "args": ["/c", "npx", "--yes", "@negokaz/excel-mcp-server"],
            "env": {
                "EXCEL_MCP_PAGING_CELLS_LIMIT": "4000"
            }
        }
    }
}
```

其他平台：

```json
{
    "mcpServers": {
        "excel": {
            "command": "npx",
            "args": ["--yes", "@negokaz/excel-mcp-server"],
            "env": {
                "EXCEL_MCP_PAGING_CELLS_LIMIT": "4000"
            }
        }
    }
}
```

### 通过 Smithery 安装

通过 [Smithery](https://smithery.ai/server/@negokaz/excel-mcp-server) 为 Claude Desktop 自动安装 Excel MCP 服务器：

```bash
npx -y @smithery/cli install @negokaz/excel-mcp-server --client claude
```

<h2 id="tools">工具</h2>

### `excel_describe_sheets`

列出指定 Excel 文件的所有工作表信息。

**参数：**

- `fileAbsolutePath`
  - Excel 文件的绝对路径

### `excel_read_sheet`

分页读取 Excel 工作表中的值。

**参数：**

- `fileAbsolutePath`
  - Excel 文件的绝对路径
- `sheetName`
  - Excel 文件中的工作表名称
- `range`
  - 要读取的 Excel 工作表中的单元格范围（例如："A1:C10"）。[默认：第一个分页范围]
- `knownPagingRanges`
  - 已读取的分页范围列表
- `showFormula`
  - 显示公式而不是值

### `excel_screen_capture`

**[仅限 Windows]** 分页截取 Excel 工作表的屏幕截图。

**参数：**

- `fileAbsolutePath`
  - Excel 文件的绝对路径
- `sheetName`
  - Excel 文件中的工作表名称
- `range`
  - 要读取的 Excel 工作表中的单元格范围（例如："A1:C10"）。[默认：第一个分页范围]
- `knownPagingRanges`
  - 已读取的分页范围列表

### `excel_write_to_sheet`

向 Excel 工作表写入值。

**参数：**

- `fileAbsolutePath`
  - Excel 文件的绝对路径
- `sheetName`
  - Excel 文件中的工作表名称
- `newSheet`
  - 如果为 true 则创建新工作表，否则写入现有工作表
- `range`
  - 要读取的 Excel 工作表中的单元格范围（例如："A1:C10"）
- `values`
  - 要写入 Excel 工作表的值。如果值是公式，应以"="开头

### `excel_create_table`

在 Excel 工作表中创建表格

**参数：**

- `fileAbsolutePath`
  - Excel 文件的绝对路径
- `sheetName`
  - 创建表格的工作表名称
- `range`
  - 要成为表格的范围（例如："A1:C10"）
- `tableName`
  - 要创建的表格名称

### `excel_copy_sheet`

将现有工作表复制到新工作表

**参数：**

- `fileAbsolutePath`
  - Excel 文件的绝对路径
- `srcSheetName`
  - Excel 文件中的源工作表名称
- `dstSheetName`
  - 要复制到的工作表名称

<h2 id="configuration">配置</h2>

您可以通过以下环境变量更改 MCP 服务器的行为：

### `EXCEL_MCP_PAGING_CELLS_LIMIT`

单次分页操作中读取的最大单元格数。  
[默认：4000]

## 许可证

版权所有 (c) 2025 Kazuki Negoro

excel-mcp-server 基于 [MIT 许可证](LICENSE) 发布
