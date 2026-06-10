package filesystem

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
			"filesystem",
			"Filesystem",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("文件系统"),
			mcpservers.WithTags([]string{"filesystem"}),
			mcpservers.WithGitHubURL(
				"https://github.com/modelcontextprotocol/servers/tree/main/src/filesystem",
			),
			mcpservers.WithDescription(
				"A Node.js MCP server for file system operations.",
			),
			mcpservers.WithDescriptionCN(
				"实现用于文件系统操作的模型上下文协议（MCP）的 Node.js 服务器。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
