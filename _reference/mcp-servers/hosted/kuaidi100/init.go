package kuaidi100

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
			"kuaidi100",
			"Kuaidi100",
			model.PublicMCPTypeProxyStreamable,
			mcpservers.WithNameCN("快递100"),
			mcpservers.WithProxyConfigTemplates(configTemplates),
			mcpservers.WithTags([]string{"kuaidi", "快递"}),
			mcpservers.WithDescription(
				"Kuaidi100 MCP Server provides comprehensive Kuaidi100 services including tracking, and more through Kuaidi100's APIs.",
			),
			mcpservers.WithDescriptionCN(
				"Kuaidi100 MCP 服务器通过 Kuaidi100 API 提供全面的 Kuaidi100 服务，包括快递跟踪等。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
