package sqlite

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
			"sqlite",
			"SQLite",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("SQLite"),
			mcpservers.WithTags([]string{"database"}),
			mcpservers.WithGitHubURL(
				"https://github.com/modelcontextprotocol/servers-archived/tree/main/src/sqlite",
			),
			mcpservers.WithDescription(
				"A Model Context Protocol (MCP) server implementation that provides database interaction and business intelligence features through SQLite. The server supports running SQL queries, analyzing business data, and automatically generating business insight memos.",
			),
			mcpservers.WithDescriptionCN(
				"一种模型上下文协议（MCP）服务器实现，通过SQLite提供数据库交互和商业智能功能。该服务器支持运行SQL查询、分析业务数据以及自动生成业务洞察备忘录。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
