# Grafana MCP 服务器

一个用于 Grafana 的[模型上下文协议][mcp] (MCP) 服务器。

这提供了对您的 Grafana 实例和周围生态系统的访问。

## 功能特性

_以下功能目前在 MCP 服务器中可用。此列表仅供参考，不代表未来功能的路线图或承诺。_

### 仪表板

- **搜索仪表板：** 按标题或其他元数据查找仪表板
- **通过 UID 获取仪表板：** 使用唯一标识符检索完整的仪表板详情
- **更新或创建仪表板：** 修改现有仪表板或创建新仪表板。_注意：由于上下文窗口限制，请谨慎使用；参见 [issue #101](https://github.com/grafana/mcp-grafana/issues/101)_
- **获取面板查询和数据源信息：** 从仪表板的每个面板获取标题、查询字符串和数据源信息（包括 UID 和类型，如果可用）

### 数据源

- **列出和获取数据源信息：** 查看所有配置的数据源并检索每个数据源的详细信息。
  - _支持的数据源类型：Prometheus、Loki。_

### Prometheus 查询

- **查询 Prometheus：** 对 Prometheus 数据源执行 PromQL 查询（支持即时和范围指标查询）。
- **查询 Prometheus 元数据：** 从 Prometheus 数据源检索指标元数据、指标名称、标签名称和标签值。

### Loki 查询

- **查询 Loki 日志和指标：** 使用 LogQL 对 Loki 数据源运行日志查询和指标查询。
- **查询 Loki 元数据：** 从 Loki 数据源检索标签名称、标签值和流统计信息。

### 事件

- **搜索、创建、更新和关闭事件：** 在 Grafana Incident 中管理事件，包括搜索、创建、更新和解决事件。

### Sift 调查

- **创建 Sift 调查：** 启动新的 Sift 调查来分析日志或跟踪。
- **列出 Sift 调查：** 检索 Sift 调查列表，支持限制参数。
- **获取 Sift 调查：** 通过 UUID 检索特定 Sift 调查的详情。
- **获取 Sift 分析：** 从 Sift 调查中检索特定分析。
- **在日志中查找错误模式：** 使用 Sift 检测 Loki 日志中的高级错误模式。
- **查找慢请求：** 使用 Sift (Tempo) 检测慢请求。

### 告警

- **列出和获取告警规则信息：** 在 Grafana 中查看告警规则及其状态（触发/正常/错误等）。
- **列出联系点：** 查看 Grafana 中配置的通知联系点。

### Grafana OnCall

- **列出和管理排班表：** 在 Grafana OnCall 中查看和管理值班排班表。
- **获取班次详情：** 检索特定值班班次的详细信息。
- **获取当前值班用户：** 查看哪些用户当前在排班表中值班。
- **列出团队和用户：** 查看所有 OnCall 团队和用户。

### 管理

- **列出团队：** 查看 Grafana 中所有配置的团队。

工具列表是可配置的，因此您可以选择要向 MCP 客户端提供哪些工具。
如果您不使用某些功能或不想占用太多上下文窗口，这很有用。
要禁用某类工具，请在启动服务器时使用 `--disable-<category>` 标志。例如，要禁用
OnCall 工具，请使用 `--disable-oncall`。

### 工具

| 工具                              | 类别        | 描述                                                        |
| --------------------------------- | ----------- | ------------------------------------------------------------------ |
| `list_teams`                      | Admin       | 列出所有团队                                                     |
| `search_dashboards`               | Search      | 搜索仪表板                                              |
| `get_dashboard_by_uid`            | Dashboard   | 通过 uid 获取仪表板                                             |
| `update_dashboard`                | Dashboard   | 更新或创建新仪表板                                   |
| `get_dashboard_panel_queries`     | Dashboard   | 从仪表板获取面板标题、查询、数据源 UID 和类型 |
| `list_datasources`                | Datasources | 列出数据源                                                   |
| `get_datasource_by_uid`           | Datasources | 通过 uid 获取数据源                                            |
| `get_datasource_by_name`          | Datasources | 通过名称获取数据源                                           |
| `query_prometheus`                | Prometheus  | 对 Prometheus 数据源执行查询                    |
| `list_prometheus_metric_metadata` | Prometheus  | 列出指标元数据                                               |
| `list_prometheus_metric_names`    | Prometheus  | 列出可用的指标名称                                        |
| `list_prometheus_label_names`     | Prometheus  | 列出匹配选择器的标签名称                               |
| `list_prometheus_label_values`    | Prometheus  | 列出特定标签的值                                   |
| `list_incidents`                  | Incident    | 列出 Grafana Incident 中的事件                                 |
| `create_incident`                 | Incident    | 在 Grafana Incident 中创建事件                             |
| `add_activity_to_incident`        | Incident    | 向 Grafana Incident 中的事件添加活动项            |
| `resolve_incident`                | Incident    | 解决 Grafana Incident 中的事件                            |
| `query_loki_logs`                 | Loki        | 使用 LogQL 查询和检索日志（日志或指标查询） |
| `list_loki_label_names`           | Loki        | 列出日志中所有可用的标签名称                             |
| `list_loki_label_values`          | Loki        | 列出特定日志标签的值                               |
| `query_loki_stats`                | Loki        | 获取日志流的统计信息                                   |
| `list_alert_rules`                | Alerting    | 列出告警规则                                                   |
| `get_alert_rule_by_uid`           | Alerting    | 通过 UID 获取告警规则                                              |
| `list_oncall_schedules`           | OnCall      | 列出 Grafana OnCall 的排班表                                 |
| `get_oncall_shift`                | OnCall      | 获取特定 OnCall 班次的详情                            |
| `get_current_oncall_users`        | OnCall      | 获取特定排班表当前值班的用户                |
| `list_oncall_teams`               | OnCall      | 列出 Grafana OnCall 的团队                                     |
| `list_oncall_users`               | OnCall      | 列出 Grafana OnCall 的用户                                     |
| `get_investigation`               | Sift        | 通过 UUID 检索现有的 Sift 调查                |
| `get_analysis`                    | Sift        | 从 Sift 调查中检索特定分析             |
| `list_investigations`             | Sift        | 检索带有可选限制的 Sift 调查列表      |
| `find_error_pattern_logs`         | Sift        | 在 Loki 日志中查找高级错误模式。                        |
| `find_slow_requests`              | Sift        | 从相关的 tempo 数据源查找慢请求。           |
| `list_pyroscope_label_names`      | Pyroscope   | 列出匹配选择器的标签名称                               |
| `list_pyroscope_label_values`     | Pyroscope   | 列出标签名称匹配选择器的标签值             |
| `list_pyroscope_profile_types`    | Pyroscope   | 列出可用的配置文件类型                                       |
| `fetch_pyroscope_profile`         | Pyroscope   | 获取 DOT 格式的配置文件进行分析                       |

## 使用方法

1. 在 Grafana 中创建一个具有足够权限的服务账户来使用您想要使用的工具，
   生成服务账户令牌，并将其复制到剪贴板以在配置文件中使用。
   详情请参考 [Grafana 文档][service-account]。

2. 您有多种选项来安装 `mcp-grafana`：

   - **Docker 镜像**：使用 Docker Hub 的预构建 Docker 镜像。

     **重要**：Docker 镜像的入口点默认配置为在 SSE 模式下运行 MCP 服务器，但大多数用户会希望使用 STDIO 模式直接与 AI 助手（如 Claude Desktop）集成：

     1. **STDIO 模式**：对于 stdio 模式，您必须明确使用 `-t stdio` 覆盖默认设置，并包含 `-i` 标志以保持 stdin 打开：

     ```bash
     docker pull mcp/grafana
     docker run --rm -i -e GRAFANA_URL=http://localhost:3000 -e GRAFANA_API_KEY=<your service account token> mcp/grafana -t stdio
     ```

     2. **SSE 模式**：在此模式下，服务器作为 HTTP 服务器运行，客户端连接到它。您必须使用 `-p` 标志暴露端口 8000：

     ```bash
     docker pull mcp/grafana
     docker run --rm -p 8000:8000 -e GRAFANA_URL=http://localhost:3000 -e GRAFANA_API_KEY=<your service account token> mcp/grafana
     ```

     3. **可流式 HTTP 模式**：在此模式下，服务器作为独立进程运行，可以处理多个客户端连接。您必须使用 `-p` 标志暴露端口 8000：对于此模式，您必须明确使用 `-t streamable-http` 覆盖默认设置

     ```bash
     docker pull mcp/grafana
     docker run --rm -p 8000:8000 -e GRAFANA_URL=http://localhost:3000 -e GRAFANA_API_KEY=<your service account token> mcp/grafana -t streamable-http
     ```

   - **下载二进制文件**：从 [发布页面](https://github.com/grafana/mcp-grafana/releases) 下载最新版本的 `mcp-grafana` 并将其放在您的 `$PATH` 中。

   - **从源码构建**：如果您安装了 Go 工具链，您也可以从源码构建和安装，使用 `GOBIN` 环境变量
     指定二进制文件应安装的目录。这也应该在您的 `PATH` 中。

     ```bash
     GOBIN="$HOME/go/bin" go install github.com/grafana/mcp-grafana/cmd/mcp-grafana@latest
     ```

3. 将服务器配置添加到您的客户端配置文件中。例如，对于 Claude Desktop：

   **如果使用二进制文件：**

   ```json
   {
     "mcpServers": {
       "grafana": {
         "command": "mcp-grafana",
         "args": [],
         "env": {
           "GRAFANA_URL": "http://localhost:3000",
           "GRAFANA_API_KEY": "<your service account token>"
         }
       }
     }
   }
   ```

> 注意：如果您在 Claude Desktop 中看到 `Error: spawn mcp-grafana ENOENT`，您需要指定 `mcp-grafana` 的完整路径。

   **如果使用 Docker：**

   ```json
   {
     "mcpServers": {
       "grafana": {
         "command": "docker",
         "args": [
           "run",
           "--rm",
           "-i",
           "-e",
           "GRAFANA_URL",
           "-e",
           "GRAFANA_API_KEY",
           "mcp/grafana",
           "-t",
           "stdio"
         ],
         "env": {
           "GRAFANA_URL": "http://localhost:3000",
           "GRAFANA_API_KEY": "<your service account token>"
         }
       }
     }
   }
   ```

   > 注意：`-t stdio` 参数在这里是必需的，因为它覆盖了 Docker 镜像中的默认 SSE 模式。

**在 VSCode 中使用远程 MCP 服务器**

如果您使用 VSCode 并在 SSE 模式下运行 MCP 服务器（这是使用 Docker 镜像而不覆盖传输时的默认模式），请确保您的 `.vscode/settings.json` 包含以下内容：

```json
"mcp": {
  "servers": {
    "grafana": {
      "type": "sse",
      "url": "http://localhost:8000/sse"
    }
  }
}
```

### 调试模式

您可以通过添加 `-debug` 标志来启用 Grafana 传输的调试模式。这将提供 MCP 服务器和 Grafana API 之间 HTTP 请求和响应的详细日志记录，这对故障排除很有帮助。

要在 Claude Desktop 配置中使用调试模式，请按如下方式更新您的配置：

**如果使用二进制文件：**

```json
{
  "mcpServers": {
    "grafana": {
      "command": "mcp-grafana",
      "args": ["-debug"],
      "env": {
        "GRAFANA_URL": "http://localhost:3000",
        "GRAFANA_API_KEY": "<your service account token>"
      }
    }
  }
}
```

**如果使用 Docker：**

```json
{
  "mcpServers": {
    "grafana": {
      "command": "docker",
      "args": [
        "run",
        "--rm",
        "-i",
        "-e",
        "GRAFANA_URL",
        "-e",
        "GRAFANA_API_KEY",
        "mcp/grafana",
        "-t",
        "stdio",
        "-debug"
      ],
      "env": {
        "GRAFANA_URL": "http://localhost:3000",
        "GRAFANA_API_KEY": "<your service account token>"
      }
    }
  }
}
```

> 注意：与标准配置一样，需要 `-t stdio` 参数来覆盖 Docker 镜像中的默认 SSE 模式。

### TLS 配置

如果您的 Grafana 实例在 mTLS 后面或需要自定义 TLS 证书，您可以配置 MCP 服务器使用自定义证书。服务器支持以下 TLS 配置选项：

- `--tls-cert-file`：用于客户端身份验证的 TLS 证书文件路径
- `--tls-key-file`：用于客户端身份验证的 TLS 私钥文件路径
- `--tls-ca-file`：用于服务器验证的 TLS CA 证书文件路径
- `--tls-skip-verify`：跳过 TLS 证书验证（不安全，仅用于测试）

**客户端证书身份验证示例：**

```json
{
  "mcpServers": {
    "grafana": {
      "command": "mcp-grafana",
      "args": [
        "--tls-cert-file", "/path/to/client.crt",
        "--tls-key-file", "/path/to/client.key",
        "--tls-ca-file", "/path/to/ca.crt"
      ],
      "env": {
        "GRAFANA_URL": "https://secure-grafana.example.com",
        "GRAFANA_API_KEY": "<your service account token>"
      }
    }
  }
}
```

**Docker 示例：**

```json
{
  "mcpServers": {
    "grafana": {
      "command": "docker",
      "args": [
        "run",
        "--rm",
        "-i",
        "-v", "/path/to/certs:/certs:ro",
        "-e", "GRAFANA_URL",
        "-e", "GRAFANA_API_KEY",
        "mcp/grafana",
        "-t", "stdio",
        "--tls-cert-file", "/certs/client.crt",
        "--tls-key-file", "/certs/client.key",
        "--tls-ca-file", "/certs/ca.crt"
      ],
      "env": {
        "GRAFANA_URL": "https://secure-grafana.example.com",
        "GRAFANA_API_KEY": "<your service account token>"
      }
    }
  }
}
```

TLS 配置应用于 MCP 服务器使用的所有 HTTP 客户端，包括：

- 主要的 Grafana OpenAPI 客户端
- Prometheus 数据源客户端
- Loki 数据源客户端
- 事件管理客户端
- Sift 调查客户端
- 告警客户端
- Asserts 客户端

**直接 CLI 使用示例：**

用于测试自签名证书：

```bash
./mcp-grafana --tls-skip-verify -debug
```

使用客户端证书身份验证：

```bash
./mcp-grafana \
  --tls-cert-file /path/to/client.crt \
  --tls-key-file /path/to/client.key \
  --tls-ca-file /path/to/ca.crt \
  -debug
```

仅使用自定义 CA 证书：

```bash
./mcp-grafana --tls-ca-file /path/to/ca.crt
```

**编程使用：**

如果您以编程方式使用此库，您也可以创建启用 TLS 的上下文函数：

```go
// 使用结构体字面量
tlsConfig := &mcpgrafana.TLSConfig{
    CertFile: "/path/to/client.crt",
    KeyFile:  "/path/to/client.key",
    CAFile:   "/path/to/ca.crt",
}
grafanaConfig := mcpgrafana.GrafanaConfig{
    Debug:     true,
    TLSConfig: tlsConfig,
}
contextFunc := mcpgrafana.ComposedStdioContextFunc(grafanaConfig)

// 或内联
grafanaConfig := mcpgrafana.GrafanaConfig{
    Debug: true,
    TLSConfig: &mcpgrafana.TLSConfig{
        CertFile: "/path/to/client.crt",
        KeyFile:  "/path/to/client.key",
        CAFile:   "/path/to/ca.crt",
    },
}
contextFunc := mcpgrafana.ComposedStdioContextFunc(grafanaConfig)
```

## 开发

欢迎贡献！如果您有任何建议或改进，请开启 issue 或提交 pull request。

此项目使用 Go 编写。按照您平台的说明安装 Go。

要在本地以 STDIO 模式运行服务器（这是本地开发的默认模式），请使用：

```bash
make run
```

要在本地以 SSE 模式运行服务器，请使用：

```bash
go run ./cmd/mcp-grafana --transport sse
```

您也可以在自定义构建的 Docker 镜像中使用 SSE 传输运行服务器。就像发布的 Docker 镜像一样，此自定义镜像的入口点默认为 SSE 模式。要构建镜像，请使用：

```
make build-image
```

要在 SSE 模式（默认）下运行镜像，请使用：

```
docker run -it --rm -p 8000:8000 mcp-grafana:latest
```

如果您需要在 STDIO 模式下运行，请覆盖传输设置：

```
docker run -it --rm mcp-grafana:latest -t stdio
```

### 测试

有三种类型的测试可用：

1. 单元测试（不需要外部依赖）：

```bash
make test-unit
```

您也可以运行单元测试：

```bash
make test
```

2. 集成测试（需要 docker 容器运行）：

```bash
make test-integration
```

3. 云测试（需要云 Grafana 实例和凭据）：

```bash
make test-cloud
```

> 注意：云测试在 CI 中自动配置。对于本地开发，您需要设置自己的 Grafana Cloud 实例和凭据。

更全面的集成测试需要 Grafana 实例在本地端口 3000 上运行；您可以使用 Docker Compose 启动一个：

```bash
docker-compose up -d
```

集成测试可以运行：

```bash
make test-all
```

如果您要添加更多工具，请为它们添加集成测试。现有测试应该是一个很好的起点。

### 代码检查

要检查代码，请运行：

```bash
make lint
```

这包括一个自定义检查器，用于检查 `jsonschema` 结构体标签中未转义的逗号。`description` 字段中的逗号必须用 `\\,` 转义以防止静默截断。您可以仅运行此检查器：

```bash
make lint-jsonschema
```

有关更多详细信息，请参阅 [JSONSchema 检查器文档](internal/linter/jsonschema/README.md)。

## 许可证

此项目根据 [Apache License, Version 2.0](LICENSE) 许可。

[mcp]: https://modelcontextprotocol.io/
[service-account]: https://grafana.com/docs/grafana/latest/administration/service-accounts/
