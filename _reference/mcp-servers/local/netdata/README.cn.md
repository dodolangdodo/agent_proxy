# Netdata 模型上下文协议 (MCP) 集成

Netdata 代理（以及即将推出的 Netdata Cloud）提供模型上下文协议 (MCP) 服务器，使 Claude 或 Cursor 等 AI 助手能够与您的基础设施监控数据进行交互。此集成允许 AI 助手访问指标、日志、告警和实时系统信息（进程、服务、容器、虚拟机、网络连接等），充当功能强大的 DevOps/SRE/系统管理员助手。

## 概述

AI 助手对您基础设施的可见性取决于它们在 Netdata 层次结构中的连接位置：

- **Netdata Cloud**：（尚未提供）连接到 Netdata Cloud 的 AI 助手将对您基础设施中的所有节点拥有完全可见性。
- **Netdata 父节点**：连接到 Netdata 父节点的 AI 助手将对连接到该父节点的所有子节点拥有可见性。
- **Netdata 子节点**：连接到 Netdata 子节点的 AI 助手将只对该特定节点拥有可见性。
- **Netdata 独立节点**：连接到独立 Netdata 节点的 AI 助手将只对该特定节点拥有可见性。

## 支持的 AI 助手

您可以将 Netdata 与以下 AI 助手一起使用：

- [Claude Desktop](https://claude.ai/download)：支持无限制访问的固定费用使用
- [Claude Code](https://claude.ai/code)：支持无限制访问的固定费用使用
- [Cursor](https://www.cursor.com/)：支持无限制访问的固定费用使用。支持与多个 AI 助手一起使用 Netdata，包括 Claude、ChatGPT 和 Gemini。

可能还有更多：查看 [MCP 文档](https://modelcontextprotocol.io/clients) 获取支持的 AI 助手完整列表。

所有这些 AI 助手都需要本地访问 MCP 服务器。这意味着您在计算机上本地运行的应用程序（Claude Desktop、Cursor 等）需要能够使用 `stdio` 通信连接到 Netdata。但是，由于您的 Netdata 在服务器上远程运行，您需要一个桥接器将 `stdio` 通信转换为 `WebSocket` 通信。Netdata 提供多种语言（Node.js、Python、Go）的桥接器来促进此转换。

一旦 MCP 集成到 Netdata Cloud 中，也将支持基于 Web 的 AI 助手。对于基于 Web 的 AI 助手，助手的后端连接到可公开访问的 MCP 服务器（即 Netdata Cloud）以访问基础设施可观测性数据，无需桥接器。

## 安装

MCP 服务器内置于 Netdata 中，无需额外安装。只需确保您安装了最新版本的 Netdata。

要将 Netdata 的 MCP 集成与 AI 客户端一起使用，您需要配置它们并将它们桥接到 Netdata MCP 服务器。

## AI 助手配置

大多数 AI 助手的配置都是通过配置文件完成的，对于所有助手来说几乎相同。

```json
{
  "mcpServers": {
    "netdata": {
      "command": "/usr/bin/nd-mcp",
      "args": [
        "ws://IP_OF_YOUR_NETDATA:19999/mcp?api_key=YOUR_API_KEY"
      ]
    }
  }
}
```

程序 `nd-mcp` 是将 `stdio` 通信转换为 `WebSocket` 通信的桥接程序。此程序是所有 Netdata 安装的一部分，因此通过在您的个人计算机（Linux、MacOS、Windows）上安装 Netdata，您就可以使用它。

根据您安装 Netdata 的方式，可能有不同的路径：

- `/usr/bin/nd-mcp` 或 `/usr/sbin/nd-mcp`：Linux 原生包（与 `netdata` 和 `netdatacli` 命令一起）
- `/opt/netdata/usr/bin/nd-mcp`：Linux 静态 Netdata 安装
- `/usr/local/netdata/usr/bin/nd-mcp`：MacOS 从源码安装
- `C:\\Program Files\\Netdata\\usr\\bin\\nd-mcp.exe`：Windows 安装

您还需要：

`IP_OF_YOUR_NETDATA` 是您要连接的 Netdata 实例的 IP 地址或主机名。这最终将被 Netdata Cloud URL 替换。对于此开发预览版，请使用任何 Netdata，最好是您的父节点之一。请记住，AI 助手将只"看到"连接到该 Netdata 实例的节点。

`YOUR_API_KEY` 是允许 AI 助手访问敏感功能（如日志和实时系统信息）的 API 密钥。只需启动 Netdata，它将自动为您生成一个随机 UUID。您可以在以下位置找到它：

```
/var/lib/netdata/mcp_dev_preview_api_key
```

或者，如果您安装了静态 Netdata 包，它可能位于：

```
/opt/netdata/var/lib/netdata/mcp_dev_preview_api_key
```

要查看您的 API 密钥：

```bash
sudo cat /var/lib/netdata/mcp_dev_preview_api_key
```

或

```bash
sudo cat /opt/netdata/var/lib/netdata/mcp_dev_preview_api_key
```

### Claude Desktop

要将 Netdata MCP 添加到 Claude Desktop：

1. 打开 Claude Desktop
2. 导航到开发者设置：

- **Windows/Linux**：文件 → 设置 → 开发者（或使用 Ctrl+,）
- **macOS**：Claude → 设置 → 开发者（或使用 Cmd+,）

3. 点击"编辑配置"按钮（在服务器列表下方）
4. 这将打开或显示确切的配置文件位置
5. 将上述配置添加到该文件中。

**Linux 用户**：Claude Desktop 可通过社区项目获得 (<https://github.com/fsoft72/claude-desktop-to-appimage)。它与> <https://github.com/TheAssassin/AppImageLauncher> 配合使用效果最佳。

正确配置后，您需要重启 Claude Desktop。
重启后，您应该看到"netdata"出现在 Claude Desktop 中：

- 点击"搜索和工具"按钮（就在提示下方）
- 您应该看到"netdata"列在可用工具中
- 如果您没有看到它，请检查您的配置并确保桥接器可访问

### Claude Code

对于 [Claude Code](https://claude.ai/code)，在您项目的根目录添加文件 `.mcp.json`，内容如上所述。此文件将在 Claude Code 下次在该目录中启动时自动检测到。

正确配置后，向您的 Claude Code 发出命令 `/mcp`。它应该显示可用的 MCP 服务器，包括"netdata"。

### Cursor

对于 [Cursor](https://www.cursor.com/)，将配置添加到 MCP 设置中。

## 替代的 `stdio` 到 `websocket` 桥接器

我们为您提供 3 种不同的桥接器，您可以选择最适合您环境的一种：

1. **Go 桥接器**：位于 `src/web/mcp/bridges/stdio-golang/`
2. **Node.js 桥接器**：位于 `src/web/mcp/bridges/stdio-nodejs/`
3. **Python 桥接器**：位于 `src/web/mcp/bridges/stdio-python/`

所有这些桥接器都应该提供完全相同的功能，因此您可以选择最适合您环境的一种。

每个目录都包含 `build.sh` 脚本来安装依赖项并准备桥接器。
Go 桥接器还为 Windows 用户提供了 `build.bat` 脚本。

## 功能

MCP 集成为 AI 助手提供以下访问权限：

### 基础设施发现

- **节点信息**：对基础设施中所有连接节点的完全可见性
  - 硬件规格、操作系统详情、虚拟化信息
  - 流配置和父子关系
  - 连接状态和数据收集能力
- **指标发现**：Netdata 安装收集的所有指标
  - 系统指标：CPU、内存、磁盘、网络接口
  - 应用程序指标：数据库、Web 服务器、容器
  - 硬件指标：IPMI 传感器、GPU、温度传感器
  - 自定义指标：StatsD、基于日志的指标

### 指标和分析

- **时间序列查询**：强大的数据聚合和分析
  - 多种分组选项：按维度、实例、节点或标签
  - 聚合方法：求和、平均值、最小值、最大值、百分比
  - 时间聚合：平均值、最小值、最大值、中位数、百分位数等
- **异常检测**：基于机器学习的所有指标异常检测
  - 实时异常率（0-100% 的异常时间）
  - 按指标和按维度的异常跟踪
- **相关性分析**：查找在事件期间发生变化的指标
  - 比较问题期间与基线期间
  - 统计和基于容量的相关方法
- **变异性分析**：识别不稳定或波动的指标

### 实时系统信息（需要 API 密钥和已声明的代理）

- **进程**：详细的进程信息，包括：
  - CPU 使用率、内存消耗、I/O 统计
  - 文件描述符、页面错误、父子关系
  - 容器感知的进程跟踪
- **网络连接**：活动连接，包含：
  - 协议详情、状态、地址、端口
  - 每个连接的性能指标
- **Systemd 服务和单元**：服务健康状况和资源使用情况
- **挂载点**：文件系统使用情况、容量和 inode 统计
- **块设备**：I/O 性能、延迟和利用率
- **容器和虚拟机**：容器化工作负载的资源使用情况
- **网络接口**：流量速率、数据包、丢包、链路状态
- **流状态**：实时复制和机器学习同步

### 日志访问（需要 API 密钥和已声明的代理）

- **systemd-journal**：全面的日志访问，包括：
  - 本地系统日志、用户日志和命名空间
  - 来自连接节点的远程系统日志
  - 高级过滤和搜索功能
  - 基于保留期的历史日志数据
- **Windows 事件**：查询 Windows 事件日志（在 Windows 系统上）

### 告警和监控

- **活动告警**：当前触发的警告和严重告警
  - 详细的告警信息，包括值、时间戳和上下文
  - 按类型、组件和严重性分类告警
- **告警历史**：完整的告警状态跟踪
  - 所有状态：严重、警告、清除、未定义、未初始化
  - 带有时间戳和值的告警转换
- **告警元数据**：接收者、配置和阈值

### 可用指标类别

集成提供对 Netdata 收集的所有指标类别的访问，包括：

- 核心系统：CPU、内存、磁盘、网络、进程
- 容器：Docker、cgroups、systemd 服务
- 数据库：MySQL、PostgreSQL、Redis、MongoDB
- Web 服务器：Apache、Nginx、LiteSpeed
- 硬件：IPMI、GPU、温度传感器、SMART
- 网络服务：DNS、DHCP、VPN、防火墙
- 应用程序：自定义 StatsD 指标、基于日志的指标
- 以及您的 Netdata 安装收集的任何其他指标

## 安全考虑

- MCP 集成目前提供对 Netdata 的**只读**访问
- 不暴露动态配置 - AI 助手无法读取或修改 Netdata 设置
- 访问敏感功能（日志和实时数据）需要 API 密钥
- 对于生产使用，请确保您的 Netdata 代理已声明到 Netdata Cloud

## 使用示例

配置完成后，您可以提出以下问题：

### 基础设施概览

- "显示所有连接的节点及其状态"
- "我的数据库服务器有哪些可用指标？"
- "为我的基础设施提供可观测性覆盖报告"
- "哪些节点离线或有连接问题？"

### 性能分析

- "在我所有服务器上消耗 CPU 最多的进程是什么？"
- "显示所有节点的网络接口利用率"
- "我的数据库服务器的内存使用趋势是什么？"
- "列出所有块设备及其 I/O 性能"
- "我的哪些节点有磁盘积压问题？"
- "显示容器资源使用统计"

### 异常检测和故障排除

- "在过去一小时内哪些指标显示异常行为？"
- "查找在下午 2 点故障期间发生显著变化的指标"
- "我基础设施中最不稳定的指标是什么？"
- "分析磁盘 I/O 和应用程序响应时间之间的相关性"

### 告警和监控

- "目前有任何严重告警处于活动状态吗？"
- "显示过去 24 小时内的所有告警转换"
- "哪些系统有磁盘空间警告？"
- "夜间触发和清除了哪些告警？"

### 系统日志和事件

- "显示失败服务的 systemd 日志"
- "搜索过去一小时内的身份验证失败"
- "显示所有节点的内核错误"
- "查找与内存不足条件相关的所有日志"

### 实时系统状态

- "列出所有 systemd 服务及其状态"
- "显示 Web 服务器上的活动网络连接"
- "当前的流复制状态是什么？"
- "显示可用空间不足的挂载点"

## 故障排除

1. **连接被拒绝**：确保 Netdata 正在运行并且可以在指定的 URL 访问
2. **找不到桥接器**：验证桥接器路径正确且依赖项已安装
3. **身份验证错误**：验证 API 密钥正确且代理已声明
4. **数据缺失**：检查 Netdata 代理是否启用了所需的收集器
5. **访问受限**：没有 API 密钥或未声明的代理将无法使用功能和日志

## 常见问题

- **问：我可以将 MCP 与其他 AI 助手一起使用吗？**
  - 答：是的，MCP 支持多个 AI 助手。查看 [MCP 文档](https://modelcontextprotocol.io/clients) 获取完整列表。

- **问：我需要在本地机器上运行桥接器吗？**
- 答：是的，桥接器将 `stdio` 通信转换为 `WebSocket` 以远程访问 Netdata。桥接器在您的本地机器（个人计算机）上运行以连接到 Netdata 实例。

- **问：如何找到我的 API 密钥？**
  - 答：API 密钥由 Netdata 自动生成并存储在您将连接的 Netdata 代理上的 `/var/lib/netdata/mcp_dev_preview_api_key` 或 `/opt/netdata/var/lib/netdata/mcp_dev_preview_api_key` 中。使用 `sudo cat` 查看它。

- **问：我可以将 MCP 与 Netdata Cloud 一起使用吗？**
  - 答：是的，一旦 MCP 集成到 Netdata Cloud 中，您将能够与基于 Web 的 AI 助手一起使用它，无需桥接器。

- **问：我可以使用 MCP 访问哪些数据？**
  - 答：您可以访问指标、日志、告警、实时系统信息（进程、服务、容器、网络连接）等。

- **问：我可以将 MCP 与现有的 Netdata 安装一起使用吗？**
  - 答：是的，只要您安装了最新版本的 Netdata，您就可以使用 MCP 集成，无需任何额外安装。

- **问：MCP 安全吗？**
  - 答：是的，MCP 目前提供只读访问。日志和实时系统信息等敏感功能需要 API 密钥，代理应声明到 Netdata Cloud 以供生产使用。

- **问：我的可观测性数据会暴露给 AI 公司吗？**
  - 答：是的，但这取决于您使用的 AI 助手和您拥有的订阅。例如，Claude 承诺对于某些订阅，您的数据不会用于训练他们的模型，Cursor 允许您使用多个 AI 助手。始终检查您选择的 AI 助手的隐私政策。

- **问：AI 助手的响应准确吗？**
  - 答：像 Claude 这样的 AI 助手旨在根据他们可以访问的数据提供准确和相关的响应。但是，它们可能并不总是完美的，或者在给出答案之前可能没有检查所有方面。验证关键信息很重要。

## 最佳实践

### AI 助手采样数据

有时，当您询问关于基础设施的一般性问题时，AI 助手会对基础设施的少数节点进行简单采样，而不是查询所有节点。在 Netdata 中，我们提供了正确执行此操作的工具，但 AI 助手可能不会使用它们。

示例：

问："我的服务器上运行的顶级进程/容器/虚拟机/服务是什么？"

AI 助手可能会响应来自少数节点的进程/容器/虚拟机/服务列表，而不是查询所有节点。

在 Netdata 中的正确方法是查询：

- `app.*` 图表/上下文用于 `processes`，这将返回按类别分组的所有节点上运行的进程。
- `systemd.*` 获取所有节点上运行的服务。
- `cgroup.*` 获取所有节点上的所有容器和虚拟机。

对于所有此类查询，Netdata 响应返回基数信息（很像您 Netdata 仪表板上的 NIDL 图表），因此 AI 助手可以获得更好的图片而不是采样数据。当您注意到这一点时，您可以要求 AI 助手使用更通用的查询来找到答案。

### AI 助手缺少较新的 Netdata 功能

有时您询问 AI 助手关于最近添加到 Netdata 的功能（例如日志或 Windows 功能），AI 助手不是检查通过其 MCP 连接可用的内容，而是说 Netdata 不支持该功能。回答"检查您的 MCP 工具、功能、函数"通常足以让 AI 助手检查可用功能并开始使用它们。

### AI 助手根本不使用 MCP

有时您需要指示它们使用其 MCP 连接。因此，与其说"检查我的生产数据库的性能"，您可以说"使用 netdata 检查我的生产数据库的性能"。这样，AI 助手将使用其 MCP 连接查询 Netdata 实例并为您提供相关信息。

### 使用 AI 助手做您的 DevOps/SRE/系统管理员"洗衣"

我们的建议是使用 AI 助手做"您的洗衣"：给它们具体的任务，检查它们获取该信息所做的查询，并在可能的情况下要求它们使用不同的工具/源交叉检查其答案。AI 助手通常急于得出结论，所以**挑战它们**，它们会更深入并自我纠正。请记住，您始终需要验证它们的答案，特别是对于关键任务。

### 单个 AI 助手的多个 Netdata MCP 服务器

如果您需要配置多个 MCP 服务器，您可以在 `mcpServers` 部分下使用不同的名称添加它们。示例：

```json
{
  "mcpServers": {
    "netdata-production": {
      "command": "/usr/bin/nd-mcp",
      "args": [
        "ws://IP_OF_YOUR_NETDATA:19999/mcp?api_key=YOUR_API_KEY"
      ]
    },
    "netdata-testing": {
      "command": "/usr/bin/nd-mcp",
      "args": [
        "ws://IP_OF_YOUR_NETDATA:19999/mcp?api_key=YOUR_API_KEY"
      ]
    }
  }
}
```

但是，当配置多个 netdata MCP 服务器时，所有 AI 助手都难以确定使用哪一个：

- **Claude Desktop**：似乎没有办法指示它使用正确的服务器。每个 MCP 服务器都有一个启用/禁用切换，但是它不能正常工作。因此，最好一次只配置一个 MCP 服务器。
- **Cursor**：同样，不可能指示它使用正确的服务器。但是，有一个切换并且它工作正常，但您仍然需要确保您要使用的 Netdata 服务器是唯一启用的服务器。
- **Claude Code**：此项目有不同的理念：您可以在每个项目目录（运行它的当前目录）中有不同的 `.mcp.json` 文件，因此您可以为每个项目/目录有不同的配置。由于 **Claude Code** 还支持带有 AI 助手默认指令的 `Claude.md` 文件，您可以有不同的目录，具有不同的指令和配置，因此您可以通过在不同目录中生成多个 Claude Code 实例来使用多个 Netdata MCP 服务器。

有关 Netdata 的更多信息，请访问 [netdata.cloud](https://netdata.cloud)
