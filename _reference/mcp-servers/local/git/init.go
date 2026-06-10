package git

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
			"git",
			"Git",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Git"),
			mcpservers.WithTags([]string{"git"}),
			mcpservers.WithGitHubURL(
				"https://github.com/modelcontextprotocol/servers/tree/main/src/git",
			),
			mcpservers.WithDescription(
				"A Node.js MCP server for Git operations.",
			),
			mcpservers.WithDescriptionCN(
				"一个用于Git操作的Node.js MCP服务器。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
