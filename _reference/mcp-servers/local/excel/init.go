package excel

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
			"excel",
			"Excel",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Excel"),
			mcpservers.WithGitHubURL(
				"https://github.com/negokaz/excel-mcp-server",
			),
			mcpservers.WithTags([]string{"excel"}),
			mcpservers.WithDescription(
				"A Model Context Protocol (MCP) server that allows LLMs to read and write Excel files using natural language. Supports xlsx, xlsm, xltx, and xltm formats.",
			),
			mcpservers.WithDescriptionCN(
				"一种用于读写Excel文件的模型上下文协议（MCP）服务器，支持xlsx、xlsm、xltx和xltm等格式。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
