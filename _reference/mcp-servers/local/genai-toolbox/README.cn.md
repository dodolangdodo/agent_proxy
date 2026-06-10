# 数据库 MCP 工具箱

[![Discord](https://img.shields.io/badge/Discord-%235865F2.svg?style=for-the-badge&logo=discord&logoColor=white)](https://discord.gg/Dmm69peqjh)
[![Go Report Card](https://goreportcard.com/badge/github.com/googleapis/genai-toolbox)](https://goreportcard.com/report/github.com/googleapis/genai-toolbox)

> [!NOTE]
> 数据库 MCP 工具箱目前处于测试阶段，在第一个稳定版本 (v1.0) 发布之前可能会有重大变更。

数据库 MCP 工具箱是一个开源的数据库 MCP 服务器。它通过处理连接池、身份验证等复杂性，使您能够更轻松、更快速、更安全地开发工具。

本 README 提供了简要概述。有关详细信息，请参阅[完整文档](https://googleapis.github.io/genai-toolbox/)。

> [!NOTE]
> 此解决方案最初名为"数据库生成式 AI 工具箱"，因为其初始开发早于 MCP，但为了与最近添加的 MCP 兼容性保持一致而重命名。

<!-- TOC ignore:true -->
## 目录

<!-- TOC -->

- [数据库 MCP 工具箱](#数据库-mcp-工具箱)
  - [目录](#目录)
  - [为什么选择工具箱？](#为什么选择工具箱)
  - [总体架构](#总体架构)
  - [入门指南](#入门指南)
    - [安装服务器](#安装服务器)
    - [运行服务器](#运行服务器)
    - [集成您的应用程序](#集成您的应用程序)
  - [配置](#配置)
    - [数据源](#数据源)
    - [工具](#工具)
    - [工具集](#工具集)
  - [版本控制](#版本控制)
  - [贡献](#贡献)

<!-- /TOC -->

## 为什么选择工具箱？

工具箱帮助您构建生成式 AI 工具，让您的智能体访问数据库中的数据。工具箱提供：

- **简化开发**：用不到 10 行代码将工具集成到您的智能体中，在多个智能体或框架之间重用工具，并更轻松地部署新版本的工具。
- **更好的性能**：连接池、身份验证等最佳实践。
- **增强的安全性**：集成身份验证，更安全地访问您的数据
- **端到端可观测性**：开箱即用的指标和跟踪，内置对 OpenTelemetry 的支持。

**⚡ 用 AI 数据库助手为您的工作流程增压 ⚡**

停止上下文切换，让您的 AI 助手成为真正的协作开发者。通过[使用 MCP 工具箱将您的 IDE 连接到数据库][connect-ide]，您可以委托复杂且耗时的数据库任务，让您构建得更快并专注于重要的事情。这不仅仅是代码补全；这是为您的 AI 提供处理整个开发生命周期所需的上下文。

以下是它如何为您节省时间：

- **用自然语言查询**：直接在您的 IDE 中使用自然语言与数据交互。询问复杂问题，如*"2024年有多少订单被交付，其中包含哪些商品？"*，无需编写任何 SQL。
- **自动化数据库管理**：简单描述您的数据需求，让 AI 助手为您管理数据库。它可以处理生成查询、创建表、添加索引等。
- **生成上下文感知代码**：让您的 AI 助手深入了解您的实时数据库架构，生成应用程序代码和测试。这通过确保生成的代码可直接使用来加速开发周期。
- **大幅减少开发开销**：大幅减少在手动设置和样板代码上花费的时间。MCP 工具箱帮助简化冗长的数据库配置、重复代码和容易出错的架构迁移。

了解[如何使用 MCP 将您的 AI 工具（IDE）连接到工具箱][connect-ide]。

[connect-ide]: https://googleapis.github.io/genai-toolbox/how-to/connect-ide/

## 总体架构

工具箱位于您的应用程序编排框架和数据库之间，提供用于修改、分发或调用工具的控制平面。它通过为您提供存储和更新工具的集中位置来简化工具管理，允许您在智能体和应用程序之间共享工具，并在不必重新部署应用程序的情况下更新这些工具。

![architecture](./docs/en/getting-started/introduction/architecture.png)

## 入门指南

### 安装服务器

有关最新版本，请查看[发布页面][releases]并按照您的操作系统和 CPU 架构使用以下说明。

[releases]: https://github.com/googleapis/genai-toolbox/releases

<details open>
<summary>二进制文件</summary>

以二进制文件方式安装工具箱：

<!-- {x-release-please-start-version} -->
```sh
# 其他版本请查看发布页面
export VERSION=0.7.0
curl -O https://storage.googleapis.com/genai-toolbox/v$VERSION/linux/amd64/toolbox
chmod +x toolbox
```

</details>

<details>
<summary>容器镜像</summary>
您也可以将工具箱安装为容器：

```sh
# 其他版本请查看发布页面
export VERSION=0.7.0
docker pull us-central1-docker.pkg.dev/database-toolbox/toolbox/toolbox:$VERSION
```

</details>

<details>
<summary>从源码编译</summary>

要从源码安装，请确保您已安装最新版本的 [Go](https://go.dev/doc/install)，然后运行以下命令：

```sh
go install github.com/googleapis/genai-toolbox@v0.7.0
```
<!-- {x-release-please-end} -->

</details>

### 运行服务器

[配置](#配置)一个 `tools.yaml` 来定义您的工具，然后执行 `toolbox` 启动服务器：

```sh
./toolbox --tools-file "tools.yaml"
```

您可以使用 `toolbox help` 查看完整的标志列表！要停止服务器，发送终止信号（大多数平台上是 `ctrl+c`）。

有关部署到不同环境的更详细文档，请查看[操作指南部分](https://googleapis.github.io/genai-toolbox/how-to/)中的资源

### 集成您的应用程序

一旦您的服务器启动并运行，您就可以将工具加载到您的应用程序中。请参阅下面用于各种框架的客户端 SDK 列表：

<details open>
<summary>Core</summary>

1. 安装 [Toolbox Core SDK][toolbox-core]：

    ```bash
    pip install toolbox-core
    ```

1. 加载工具：

    ```python
    from toolbox_core import ToolboxClient

    # 更新 url 以指向您的服务器
    async with ToolboxClient("http://127.0.0.1:5000") as client:

        # 这些工具可以传递给您的应用程序！
        tools = await client.load_toolset("toolset_name")
    ```

有关使用 Toolbox Core SDK 的更详细说明，请参阅[项目的 README][toolbox-core-readme]。

[toolbox-core]: https://pypi.org/project/toolbox-core/
[toolbox-core-readme]: https://github.com/googleapis/mcp-toolbox-sdk-python/tree/main/packages/toolbox-core/README.md

</details>
<details>
<summary>LangChain / LangGraph</summary>

1. 安装 [Toolbox LangChain SDK][toolbox-langchain]：

    ```bash
    pip install toolbox-langchain
    ```

1. 加载工具：

    ```python
    from toolbox_langchain import ToolboxClient

    # 更新 url 以指向您的服务器
    async with ToolboxClient("http://127.0.0.1:5000") as client:

        # 这些工具可以传递给您的应用程序！
        tools = client.load_toolset()
    ```

有关使用 Toolbox LangChain SDK 的更详细说明，请参阅[项目的 README][toolbox-langchain-readme]。

[toolbox-langchain]: https://pypi.org/project/toolbox-langchain/
[toolbox-langchain-readme]: https://github.com/googleapis/mcp-toolbox-sdk-python/blob/main/packages/toolbox-langchain/README.md

</details>

<details>
<summary>LlamaIndex</summary>

1. 安装 [Toolbox Llamaindex SDK][toolbox-llamaindex]：

    ```bash
    pip install toolbox-llamaindex
    ```

1. 加载工具：

    ```python
    from toolbox_llamaindex import ToolboxClient

    # 更新 url 以指向您的服务器
    async with ToolboxClient("http://127.0.0.1:5000") as client:

        # 这些工具可以传递给您的应用程序！
        tools = client.load_toolset()
    ```

有关使用 Toolbox Llamaindex SDK 的更详细说明，请参阅[项目的 README][toolbox-llamaindex-readme]。

[toolbox-llamaindex]: https://pypi.org/project/toolbox-llamaindex/
[toolbox-llamaindex-readme]: https://github.com/googleapis/genai-toolbox-llamaindex-python/blob/main/README.md

</details>

## 配置

配置工具箱的主要方式是通过 `tools.yaml` 文件。如果您有多个文件，可以使用 `--tools-file tools.yaml` 标志告诉工具箱加载哪个文件。

您可以在[资源](https://googleapis.github.io/genai-toolbox/resources/)中找到所有资源类型的更详细参考文档。

### 数据源

`tools.yaml` 的 `sources` 部分定义了您的工具箱应该访问哪些数据源。大多数工具至少有一个要执行的数据源。

```yaml
sources:
  my-pg-source:
    kind: postgres
    host: 127.0.0.1
    port: 5432
    database: toolbox_db
    user: toolbox_user
    password: my-password
```

有关配置不同类型数据源的更多详细信息，请参阅[数据源](https://googleapis.github.io/genai-toolbox/resources/sources)。

### 工具

`tools.yaml` 的 `tools` 部分定义了智能体可以执行的操作：它是什么类型的工具、影响哪些数据源、使用什么参数等。

```yaml
tools:
  search-hotels-by-name:
    kind: postgres-sql
    source: my-pg-source
    description: 根据名称搜索酒店。
    parameters:
      - name: name
        type: string
        description: 酒店的名称。
    statement: SELECT * FROM hotels WHERE name ILIKE '%' || $1 || '%';
```

有关配置不同类型工具的更多详细信息，请参阅[工具](https://googleapis.github.io/genai-toolbox/resources/tools)。

### 工具集

`tools.yaml` 的 `toolsets` 部分允许您定义要一起加载的工具组。这对于基于智能体或应用程序定义不同组很有用。

```yaml
toolsets:
    my_first_toolset:
        - my_first_tool
        - my_second_tool
    my_second_toolset:
        - my_second_tool
        - my_third_tool
```

您可以按名称加载工具集：

```python
# 这将加载所有工具
all_tools = client.load_toolset()

# 这将只加载 'my_second_toolset' 中列出的工具
my_second_toolset = client.load_toolset("my_second_toolset")
```

## 版本控制

此项目使用[语义版本控制](https://semver.org/)，包括 `MAJOR.MINOR.PATCH` 版本号，递增规则如下：

- MAJOR 版本：当我们进行不兼容的 API 更改时
- MINOR 版本：当我们以向后兼容的方式添加功能时
- PATCH 版本：当我们进行向后兼容的错误修复时

这适用的公共 API 是与工具箱关联的 CLI、与官方 SDK 的交互以及 `tools.yaml` 文件中的定义。

## 贡献

欢迎贡献。请参阅[贡献指南](CONTRIBUTING.md)开始。

请注意，此项目发布时附有贡献者行为准则。通过参与此项目，您同意遵守其条款。有关更多信息，请参阅[贡献者行为准则](CODE_OF_CONDUCT.md)。
