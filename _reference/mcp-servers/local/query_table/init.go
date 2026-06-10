package querytable

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
			"query-table",
			"Stock Data Query",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("股票数据查询"),
			mcpservers.WithGitHubURL(
				"https://github.com/wukan1986/mcp_query_table",
			),
			mcpservers.WithTags([]string{"finance"}),
			mcpservers.WithDescription(
				"A Model Context Protocol (MCP) server that allows LLMs to query financial data from various websites using natural language.",
			),
			mcpservers.WithDescriptionCN(
				"基于playwright实现的财经网页表格爬虫，支持Model Context Protocol (MCP)，可查询同花顺问财、通达信问小达、东方财富条件选股等。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
