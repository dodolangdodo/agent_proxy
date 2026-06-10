package stockanalysis

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
			"stock-analysis",
			"Stock Analysis",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("股票分析"),
			mcpservers.WithGitHubURL(
				"https://github.com/giptilabs/mcp-stock-analysis",
			),
			mcpservers.WithTags([]string{"stock"}),
			mcpservers.WithDescription(
				"Use Yahoo Finance API to provide access to real-time and historical Indian stock data, allowing local LLMs to retrieve stock information through MCP-compatible agents (such as Claude Desktop and Cursor).e Yahoo Finance API to provide access to real-time and historical Indian stock data, allowing local LLMs to retrieve stock information through MCP-compatible agents (such as Claude Desktop and Cursor).",
			),
			mcpservers.WithDescriptionCN(
				"通过Yahoo Finance API提供对实时和历史印度股票数据的访问，使本地LLM能够通过与MCP兼容的代理（如Claude Desktop和Cursor）检索股票信息。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
