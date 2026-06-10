package grafana

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
			"grafana",
			"Grafana",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Grafana"),
			mcpservers.WithTags([]string{"grafana"}),
			mcpservers.WithGitHubURL(
				"https://github.com/grafana/mcp-grafana",
			),
			mcpservers.WithDescription(
				"Connects Grafana to Claude AI through the Model Context Protocol (MCP), allowing Claude to interact with and control Grafana directly, enabling AI-assisted data analysis.",
			),
			mcpservers.WithDescriptionCN(
				"通过模型上下文协议将Grafana与Claude AI连接起来，通过允许Claude直接与Grafana会话交互和控制，实现人工智能辅助的数据分析。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
