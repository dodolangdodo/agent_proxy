## Playwright MCP

> <https://github.com/microsoft/playwright-mcp>

一个模型上下文协议 (MCP) 服务器，使用 [Playwright](https://playwright.dev) 提供浏览器自动化功能。该服务器使 LLM 能够通过结构化的无障碍快照与网页交互，无需截图或视觉调优模型。

### 主要特性

- **快速且轻量级**。使用 Playwright 的无障碍树，而非基于像素的输入。
- **LLM 友好**。无需视觉模型，纯粹基于结构化数据操作。
- **确定性工具应用**。避免了基于截图方法常见的模糊性。

### 系统要求

- Node.js 18 或更新版本
- VS Code、Cursor、Windsurf、Claude Desktop 或任何其他 MCP 客户端

### 快速开始

首先，在您的客户端中安装 Playwright MCP 服务器。典型配置如下：

```js
{
  "mcpServers": {
    "playwright": {
      "command": "npx",
      "args": [
        "@playwright/mcp@latest"
      ]
    }
  }
}
```

[<img src="https://img.shields.io/badge/VS_Code-VS_Code?style=flat-square&label=Install%20Server&color=0098FF" alt="在 VS Code 中安装">](https://insiders.vscode.dev/redirect?url=vscode%3Amcp%2Finstall%3F%257B%2522name%2522%253A%2522playwright%2522%252C%2522command%2522%253A%2522npx%2522%252C%2522args%2522%253A%255B%2522%2540playwright%252Fmcp%2540latest%2522%255D%257D) [<img alt="在 VS Code Insiders 中安装" src="https://img.shields.io/badge/VS_Code_Insiders-VS_Code_Insiders?style=flat-square&label=Install%20Server&color=24bfa5">](https://insiders.vscode.dev/redirect?url=vscode-insiders%3Amcp%2Finstall%3F%257B%2522name%2522%253A%2522playwright%2522%252C%2522command%2522%253A%2522npx%2522%252C%2522args%2522%253A%255B%2522%2540playwright%252Fmcp%2540latest%2522%255D%257D)

<details><summary><b>在 VS Code 中安装</b></summary>

您也可以使用 VS Code CLI 安装 Playwright MCP 服务器：

```bash
# 对于 VS Code
code --add-mcp '{"name":"playwright","command":"npx","args":["@playwright/mcp@latest"]}'
```

安装后，Playwright MCP 服务器将可在 VS Code 中与您的 GitHub Copilot 代理一起使用。
</details>

<details>
<summary><b>在 Cursor 中安装</b></summary>

#### 点击按钮安装

[![安装 MCP 服务器](https://cursor.com/deeplink/mcp-install-dark.svg)](https://cursor.com/install-mcp?name=playwright&config=eyJjb21tYW5kIjoibnB4IEBwbGF5d3JpZ2h0L21jcEBsYXRlc3QifQ%3D%3D)

#### 或手动安装

前往 `Cursor Settings` -> `MCP` -> `Add new MCP Server`。随意命名，使用 `command` 类型，命令为 `npx @playwright/mcp`。您也可以通过点击 `Edit` 验证配置或添加命令行参数。

```js
{
  "mcpServers": {
    "playwright": {
      "command": "npx",
      "args": [
        "@playwright/mcp@latest"
      ]
    }
  }
}
```

</details>

<details>
<summary><b>在 Windsurf 中安装</b></summary>

参考 Windsurf MCP [文档](https://docs.windsurf.com/windsurf/cascade/mcp)。使用以下配置：

```js
{
  "mcpServers": {
    "playwright": {
      "command": "npx",
      "args": [
        "@playwright/mcp@latest"
      ]
    }
  }
}
```

</details>

<details>
<summary><b>在 Claude Desktop 中安装</b></summary>

参考 MCP 安装[指南](https://modelcontextprotocol.io/quickstart/user)，使用以下配置：

```js
{
  "mcpServers": {
    "playwright": {
      "command": "npx",
      "args": [
        "@playwright/mcp@latest"
      ]
    }
  }
}
```

</details>

<details>
<summary><b>在 Qodo Gen 中安装</b></summary>

在 VSCode 或 IntelliJ 中打开 [Qodo Gen](https://docs.qodo.ai/qodo-documentation/qodo-gen) 聊天面板 → Connect more tools → + Add new MCP → 粘贴以下配置：

```js
{
  "mcpServers": {
    "playwright": {
      "command": "npx",
      "args": [
        "@playwright/mcp@latest"
      ]
    }
  }
}
```

点击 <code>Save</code>。
</details>

### 配置

Playwright MCP 服务器支持以下参数。它们可以在上述 JSON 配置中作为 `"args"` 列表的一部分提供：

```
> npx @playwright/mcp@latest --help
  --allowed-origins <origins>  允许浏览器请求的源列表，用分号分隔。默认允许所有。
  --blocked-origins <origins>  阻止浏览器请求的源列表，用分号分隔。阻止列表在允许列表之前评估。如果在没有允许列表的情况下使用，不匹配阻止列表的请求仍然被允许。
  --block-service-workers      阻止 service workers
  --browser <browser>          要使用的浏览器或 chrome 频道，可能的值：chrome、firefox、webkit、msedge。
  --browser-agent <endpoint>   使用浏览器代理（实验性）。
  --caps <caps>                要启用的功能列表，用逗号分隔，可能的值：tabs、pdf、history、wait、files、install。默认是全部。
  --cdp-endpoint <endpoint>    要连接的 CDP 端点。
  --config <path>              配置文件的路径。
  --device <device>            要模拟的设备，例如："iPhone 15"
  --executable-path <path>     浏览器可执行文件的路径。
  --headless                   在无头模式下运行浏览器，默认为有头模式
  --host <host>                服务器绑定的主机。默认是 localhost。使用 0.0.0.0 绑定到所有接口。
  --ignore-https-errors        忽略 https 错误
  --isolated                   将浏览器配置文件保存在内存中，不保存到磁盘。
  --image-responses <mode>     是否向客户端发送图像响应。可以是 "allow"、"omit" 或 "auto"。默认为 "auto"，如果客户端可以显示图像则发送。
  --no-sandbox                 为通常沙盒化的所有进程类型禁用沙盒。
  --output-dir <path>          输出文件的目录路径。
  --port <port>                SSE 传输监听的端口。
  --proxy-bypass <bypass>      绕过代理的域列表，用逗号分隔，例如 ".com,chromium.org,.domain.com"
  --proxy-server <proxy>       指定代理服务器，例如 "http://myproxy:3128" 或 "socks5://myproxy:8080"
  --save-trace                 是否将会话的 Playwright Trace 保存到输出目录。
  --storage-state <path>       隔离会话的存储状态文件路径。
  --user-agent <ua string>     指定用户代理字符串
  --user-data-dir <path>       用户数据目录的路径。如果未指定，将创建临时目录。
  --viewport-size <size>       指定浏览器视口大小（像素），例如 "1280, 720"
  --vision                     运行使用截图的服务器（默认使用 Aria 快照）
```

### 用户配置文件

您可以使用持久配置文件（默认）像常规浏览器一样运行 Playwright MCP，或者在隔离上下文中用于测试会话。

**持久配置文件**

所有登录信息都将存储在持久配置文件中，如果您想清除离线状态，可以在会话之间删除它。
持久配置文件位于以下位置，您可以使用 `--user-data-dir` 参数覆盖它。

```bash
# Windows
%USERPROFILE%\AppData\Local\ms-playwright\mcp-{channel}-profile

# macOS
- ~/Library/Caches/ms-playwright/mcp-{channel}-profile

# Linux
- ~/.cache/ms-playwright/mcp-{channel}-profile
```

**隔离模式**

在隔离模式下，每个会话都在隔离配置文件中启动。每次您要求 MCP 关闭浏览器时，
会话都会关闭，该会话的所有存储状态都会丢失。您可以通过配置的 `contextOptions` 或通过 `--storage-state` 参数向浏览器提供初始存储状态。在[这里](https://playwright.dev/docs/auth)了解更多关于存储状态的信息。

```js
{
  "mcpServers": {
    "playwright": {
      "command": "npx",
      "args": [
        "@playwright/mcp@latest",
        "--isolated",
        "--storage-state={path/to/storage.json}"
      ]
    }
  }
}
```

### 配置文件

Playwright MCP 服务器可以使用 JSON 配置文件进行配置。您可以使用 `--config` 命令行选项指定配置文件：

```bash
npx @playwright/mcp@latest --config path/to/config.json
```

<details>
<summary>配置文件架构</summary>

```typescript
{
  // 浏览器配置
  browser?: {
    // 要使用的浏览器类型（chromium、firefox 或 webkit）
    browserName?: 'chromium' | 'firefox' | 'webkit';

    // 将浏览器配置文件保存在内存中，不保存到磁盘。
    isolated?: boolean;

    // 浏览器配置文件持久化的用户数据目录路径
    userDataDir?: string;

    // 浏览器启动选项（参见 Playwright 文档）
    // @see https://playwright.dev/docs/api/class-browsertype#browser-type-launch
    launchOptions?: {
      channel?: string;        // 浏览器频道（例如 'chrome'）
      headless?: boolean;      // 在无头模式下运行
      executablePath?: string; // 浏览器可执行文件路径
      // ... 其他 Playwright 启动选项
    };

    // 浏览器上下文选项
    // @see https://playwright.dev/docs/api/class-browser#browser-new-context
    contextOptions?: {
      viewport?: { width: number, height: number };
      // ... 其他 Playwright 上下文选项
    };

    // 用于连接到现有浏览器的 CDP 端点
    cdpEndpoint?: string;

    // 远程 Playwright 服务器端点
    remoteEndpoint?: string;
  },

  // 服务器配置
  server?: {
    port?: number;  // 监听端口
    host?: string;  // 绑定主机（默认：localhost）
  },

  // 启用的功能列表
  capabilities?: Array<
    'core' |    // 核心浏览器自动化
    'tabs' |    // 标签页管理
    'pdf' |     // PDF 生成
    'history' | // 浏览器历史
    'wait' |    // 等待工具
    'files' |   // 文件处理
    'install' | // 浏览器安装
    'testing'   // 测试
  >;

  // 启用视觉模式（截图而非无障碍快照）
  vision?: boolean;

  // 输出文件目录
  outputDir?: string;

  // 网络配置
  network?: {
    // 允许浏览器请求的源列表。默认允许所有。同时匹配 `allowedOrigins` 和 `blockedOrigins` 的源将被阻止。
    allowedOrigins?: string[];

    // 阻止浏览器请求的源列表。同时匹配 `allowedOrigins` 和 `blockedOrigins` 的源将被阻止。
    blockedOrigins?: string[];
  };
 
  /**
   * 不向客户端发送图像响应。
   */
  noImageResponses?: boolean;
}
```

</details>

### 独立 MCP 服务器

当在没有显示器的系统上运行有头浏览器或从 IDE 的工作进程运行时，
从具有 DISPLAY 的环境运行 MCP 服务器并传递 `--port` 标志以启用 SSE 传输。

```bash
npx @playwright/mcp@latest --port 8931
```

然后在 MCP 客户端配置中，将 `url` 设置为 SSE 端点：

```js
{
  "mcpServers": {
    "playwright": {
      "url": "http://localhost:8931/sse"
    }
  }
}
```

<details>
<summary><b>Docker</b></summary>

**注意：** Docker 实现目前仅支持无头 chromium。

```js
{
  "mcpServers": {
    "playwright": {
      "command": "docker",
      "args": ["run", "-i", "--rm", "--init", "--pull=always", "mcr.microsoft.com/playwright/mcp"]
    }
  }
}
```

您可以自己构建 Docker 镜像。

```
docker build -t mcr.microsoft.com/playwright/mcp .
```

</details>

<details>
<summary><b>编程使用</b></summary>

```js
import http from 'http';

import { createConnection } from '@playwright/mcp';
import { SSEServerTransport } from '@modelcontextprotocol/sdk/server/sse.js';

http.createServer(async (req, res) => {
  // ...

  // 创建一个无头 Playwright MCP 服务器与 SSE 传输
  const connection = await createConnection({ browser: { launchOptions: { headless: true } } });
  const transport = new SSEServerTransport('/messages', res);
  await connection.sever.connect(transport);

  // ...
});
```

</details>

### 工具

工具有两种模式可用：

1. **快照模式**（默认）：使用无障碍快照以获得更好的性能和可靠性
2. **视觉模式**：使用截图进行基于视觉的交互

要使用视觉模式，在启动服务器时添加 `--vision` 标志：

```js
{
  "mcpServers": {
    "playwright": {
      "command": "npx",
      "args": [
        "@playwright/mcp@latest",
        "--vision"
      ]
    }
  }
}
```

视觉模式最适合能够基于提供的截图使用 X Y 坐标空间与元素交互的计算机使用模型。

<details>
<summary><b>交互</b></summary>

- **browser_snapshot**
  - 标题：页面快照
  - 描述：捕获当前页面的无障碍快照，这比截图更好
  - 参数：无
  - 只读：**是**

- **browser_click**
  - 标题：点击
  - 描述：在网页上执行点击
  - 参数：
    - `element` (string)：用于获得与元素交互权限的人类可读元素描述
    - `ref` (string)：来自页面快照的确切目标元素引用
  - 只读：**否**

- **browser_drag**
  - 标题：拖拽鼠标
  - 描述：在两个元素之间执行拖放
  - 参数：
    - `startElement` (string)：用于获得与元素交互权限的人类可读源元素描述
    - `startRef` (string)：来自页面快照的确切源元素引用
    - `endElement` (string)：用于获得与元素交互权限的人类可读目标元素描述
    - `endRef` (string)：来自页面快照的确切目标元素引用
  - 只读：**否**

- **browser_hover**
  - 标题：悬停鼠标
  - 描述：在页面元素上悬停
  - 参数：
    - `element` (string)：用于获得与元素交互权限的人类可读元素描述
    - `ref` (string)：来自页面快照的确切目标元素引用
  - 只读：**是**

- **browser_type**
  - 标题：输入文本
  - 描述：在可编辑元素中输入文本
  - 参数：
    - `element` (string)：用于获得与元素交互权限的人类可读元素描述
    - `ref` (string)：来自页面快照的确切目标元素引用
    - `text` (string)：要输入到元素中的文本
    - `submit` (boolean, 可选)：是否提交输入的文本（之后按 Enter）
    - `slowly` (boolean, 可选)：是否一次输入一个字符。对于触发页面中的键处理程序很有用。默认情况下一次填入整个文本。
  - 只读：**否**

- **browser_select_option**
  - 标题：选择选项
  - 描述：在下拉菜单中选择选项
  - 参数：
    - `element` (string)：用于获得与元素交互权限的人类可读元素描述
    - `ref` (string)：来自页面快照的确切目标元素引用
    - `values` (array)：在下拉菜单中选择的值数组。可以是单个值或多个值。
  - 只读：**否**

- **browser_press_key**
  - 标题：按键
  - 描述：在键盘上按键
  - 参数：
    - `key` (string)：要按的键名或要生成的字符，如 `ArrowLeft` 或 `a`
  - 只读：**否**

- **browser_wait_for**
  - 标题：等待
  - 描述：等待文本出现或消失或指定时间过去
  - 参数：
    - `time` (number, 可选)：等待时间（秒）
    - `text` (string, 可选)：要等待的文本
    - `textGone` (string, 可选)：要等待消失的文本
  - 只读：**是**

- **browser_file_upload**
  - 标题：上传文件
  - 描述：上传一个或多个文件
  - 参数：
    - `paths` (array)：要上传的文件的绝对路径。可以是单个文件或多个文件。
  - 只读：**否**

- **browser_handle_dialog**
  - 标题：处理对话框
  - 描述：处理对话框
  - 参数：
    - `accept` (boolean)：是否接受对话框。
    - `promptText` (string, 可选)：提示对话框的文本。
  - 只读：**否**

</details>

<details>
<summary><b>导航</b></summary>

- **browser_navigate**
  - 标题：导航到 URL
  - 描述：导航到 URL
  - 参数：
    - `url` (string)：要导航到的 URL
  - 只读：**否**

- **browser_navigate_back**
  - 标题：后退
  - 描述：返回到上一页
  - 参数：无
  - 只读：**是**

- **browser_navigate_forward**
  - 标题：前进
  - 描述：前进到下一页
  - 参数：无
  - 只读：**是**

</details>

<details>
<summary><b>资源</b></summary>

- **browser_take_screenshot**
  - 标题：截图
  - 描述：截取当前页面的屏幕截图。您不能基于截图执行操作，请使用 browser_snapshot 进行操作。
  - 参数：
    - `raw` (boolean, 可选)：是否返回无压缩（PNG 格式）。默认为 false，返回 JPEG 图像。
    - `filename` (string, 可选)：保存截图的文件名。如果未指定，默认为 `page-{timestamp}.{png|jpeg}`。
    - `element` (string, 可选)：用于获得截图元素权限的人类可读元素描述。如果未提供，将截取视口的截图。如果提供了 element，也必须提供 ref。
    - `ref` (string, 可选)：来自页面快照的确切目标元素引用。如果未提供，将截取视口的截图。如果提供了 ref，也必须提供 element。
  - 只读：**是**

- **browser_pdf_save**
  - 标题：保存为 PDF
  - 描述：将页面保存为 PDF
  - 参数：
    - `filename` (string, 可选)：保存 PDF 的文件名。如果未指定，默认为 `page-{timestamp}.pdf`。
  - 只读：**是**

- **browser_network_requests**
  - 标题：列出网络请求
  - 描述：返回自加载页面以来的所有网络请求
  - 参数：无
  - 只读：**是**

- **browser_console_messages**
  - 标题：获取控制台消息
  - 描述：返回所有控制台消息
  - 参数：无
  - 只读：**是**

</details>

<details>
<summary><b>工具</b></summary>

- **browser_install**
  - 标题：安装配置中指定的浏览器
  - 描述：安装配置中指定的浏览器。如果您收到关于浏览器未安装的错误，请调用此功能。
  - 参数：无
  - 只读：**否**

- **browser_close**
  - 标题：关闭浏览器
  - 描述：关闭页面
  - 参数：无
  - 只读：**是**

- **browser_resize**
  - 标题：调整浏览器窗口大小
  - 描述：调整浏览器窗口大小
  - 参数：
    - `width` (number)：浏览器窗口宽度
    - `height` (number)：浏览器窗口高度
  - 只读：**是**

</details>

<details>
<summary><b>标签页</b></summary>

- **browser_tab_list**
  - 标题：列出标签页
  - 描述：列出浏览器标签页
  - 参数：无
  - 只读：**是**

- **browser_tab_new**
  - 标题：打开新标签页
  - 描述：打开新标签页
  - 参数：
    - `url` (string, 可选)：在新标签页中导航到的 URL。如果未提供，新标签页将为空白。
  - 只读：**是**

- **browser_tab_select**
  - 标题：选择标签页
  - 描述：按索引选择标签页
  - 参数：
    - `index` (number)：要选择的标签页索引
  - 只读：**是**

- **browser_tab_close**
  - 标题：关闭标签页
  - 描述：关闭标签页
  - 参数：
    - `index` (number, 可选)：要关闭的标签页索引。如果未提供则关闭当前标签页。
  - 只读：**否**

</details>

<details>
<summary><b>测试</b></summary>

- **browser_generate_playwright_test**
  - 标题：生成 Playwright 测试
  - 描述：为给定场景生成 Playwright 测试
  - 参数：
    - `name` (string)：测试名称
    - `description` (string)：测试描述
    - `steps` (array)：测试步骤
  - 只读：**是**

</details>

<details>
<summary><b>视觉模式</b></summary>

- **browser_screen_capture**
  - 标题：截图
  - 描述：截取当前页面的屏幕截图
  - 参数：无
  - 只读：**是**

- **browser_screen_move_mouse**
  - 标题：移动鼠标
  - 描述：将鼠标移动到给定位置
  - 参数：
    - `element` (string)：用于获得与元素交互权限的人类可读元素描述
    - `x` (number)：X 坐标
    - `y` (number)：Y 坐标
  - 只读：**是**

- **browser_screen_click**
  - 标题：点击
  - 描述：点击鼠标左键
  - 参数：
    - `element` (string)：用于获得与元素交互权限的人类可读元素描述
    - `x` (number)：X 坐标
    - `y` (number)：Y 坐标
  - 只读：**否**

- **browser_screen_drag**
  - 标题：拖拽鼠标
  - 描述：拖拽鼠标左键
  - 参数：
    - `element` (string)：用于获得与元素交互权限的人类可读元素描述
    - `startX` (number)：起始 X 坐标
    - `startY` (number)：起始 Y 坐标
    - `endX` (number)：结束 X 坐标
    - `endY` (number)：结束 Y 坐标
  - 只读：**否**

- **browser_screen_type**
  - 标题：输入文本
  - 描述：输入文本
  - 参数：
    - `text` (string)：要输入到元素中的文本
    - `submit` (boolean, 可选)：是否提交输入的文本（之后按 Enter）
  - 只读：**否**

- **browser_press_key**
  - 标题：按键
  - 描述：在键盘上按键
  - 参数：
    - `key` (string)：要按的键名或要生成的字符，如 `ArrowLeft` 或 `a`
  - 只读：**否**

- **browser_wait_for**
  - 标题：等待
  - 描述：等待文本出现或消失或指定时间过去
  - 参数：
    - `time` (number, 可选)：等待时间（秒）
    - `text` (string, 可选)：要等待的文本
    - `textGone` (string, 可选)：要等待消失的文本
  - 只读：**是**

- **browser_file_upload**
  - 标题：上传文件
  - 描述：上传一个或多个文件
  - 参数：
    - `paths` (array)：要上传的文件的绝对路径。可以是单个文件或多个文件。
  - 只读：**否**

- **browser_handle_dialog**
  - 标题：处理对话框
  - 描述：处理对话框
  - 参数：
    - `accept` (boolean)：是否接受对话框。
    - `promptText` (string, 可选)：提示对话框的文本。
  - 只读：**否**

</details>
