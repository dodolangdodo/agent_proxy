package applescript

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
			"applescript",
			"Apple Script",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Apple Script"),
			mcpservers.WithTags([]string{"applescript", "apple", "script"}),
			mcpservers.WithGitHubURL(
				"https://github.com/peakmojo/applescript-mcp",
			),
			mcpservers.WithDescription(
				"MCP server that execute applescript giving you full control of your Mac",
			),
			mcpservers.WithDescriptionCN(
				"执行 AppleScript 的 MCP 服务器，让您完全掌控您的 Mac",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
