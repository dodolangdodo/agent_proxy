package mysql

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
			"mysql",
			"MySQL",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("MySQL"),
			mcpservers.WithGitHubURL(
				"https://github.com/designcomputer/mysql_mcp_server",
			),
			mcpservers.WithTags([]string{"database"}),
			mcpservers.WithDescription(
				"Allows AI assistants to list tables, read data, and execute SQL queries through a controlled interface, making database exploration and analysis safer and more structured.",
			),
			mcpservers.WithDescriptionCN(
				"允许AI助手通过受控接口列出表、读取数据和执行SQL查询，使数据库的探索和分析更加安全和有结构。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
