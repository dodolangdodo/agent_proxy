package netdata

import (
	_ "embed"

	"github.com/labring/aiproxy/core/model"
	mcpservers "github.com/labring/aiproxy/mcp-servers"
)

//go:embed README.md
var readme string

//go:embed README.cn.md
var readmeCN string

// need import in mcpregister/init.go
func init() {
	mcpservers.Register(
		mcpservers.NewMcp(
			"netdata",
			"Netdata",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Netdata"),
			mcpservers.WithTags([]string{"monitoring", "infrastructure", "devops"}),
			mcpservers.WithGitHubURL(
				"https://github.com/netdata/netdata/tree/master/src/web/mcp",
			),
			mcpservers.WithDescription(
				"Netdata Model Context Protocol (MCP) integration enables AI assistants to interact with infrastructure monitoring data, providing access to metrics, logs, alerts, and live system information for DevOps/SRE/SysAdmin assistance.",
			),
			mcpservers.WithDescriptionCN(
				"Netdata 模型上下文协议 (MCP) 集成使 AI 助手能够与基础设施监控数据进行交互，提供对指标、日志、告警和实时系统信息的访问，为 DevOps/SRE/系统管理员提供协助。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
