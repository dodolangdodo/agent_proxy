package kubernetes

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
			"kubernetes",
			"Kubernetes",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Kubernetes"),
			mcpservers.WithTags([]string{"kubernetes"}),
			mcpservers.WithGitHubURL(
				"https://github.com/Flux159/mcp-server-kubernetes",
			),
			mcpservers.WithDescription(
				"MCP Server for kubernetes management commands",
			),
			mcpservers.WithDescriptionCN(
				"用于 Kubernetes 管理命令的 MCP 服务器",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
