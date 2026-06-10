package forevervm

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
			"forevervm",
			"ForeverVM",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("ForeverVM"),
			mcpservers.WithTags([]string{"python", "repl", "code-execution"}),
			mcpservers.WithGitHubURL(
				"https://github.com/jamsocket/forevervm/tree/main/javascript/mcp-server",
			),
			mcpservers.WithDescription(
				"MCP Server for ForeverVM, enabling Claude to execute code in a Python REPL.",
			),
			mcpservers.WithDescriptionCN(
				"ForeverVM 的 MCP 服务器，使 Claude 能够在 Python REPL 中执行代码。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
