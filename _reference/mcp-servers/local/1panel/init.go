package onepanel

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
			"1panel",
			"1Panel",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("1Panel"),
			mcpservers.WithTags([]string{"1panel", "panel", "server", "management"}),
			mcpservers.WithGitHubURL(
				"https://github.com/1Panel-dev/mcp-1panel",
			),
			mcpservers.WithDescription(
				"1Panel's Model Context Protocol (MCP) server implementation.",
			),
			mcpservers.WithDescriptionCN(
				"1Panel 的 Model Context Protocol (MCP) 协议服务端实现。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
