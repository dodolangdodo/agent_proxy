package homeassistant

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
			"homeassistant",
			"Home Assistant",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Home Assistant"),
			mcpservers.WithTags([]string{"homeassistant", "home", "assistant"}),
			mcpservers.WithGitHubURL(
				"https://github.com/tevonsb/homeassistant-mcp",
			),
			mcpservers.WithDescription(
				"A MCP server for Home Assistant",
			),
			mcpservers.WithDescriptionCN(
				"Home Assistant 的 MCP 服务器",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
