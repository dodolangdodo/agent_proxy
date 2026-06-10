package oceanbase

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
			"oceanbase",
			"OceanBase",
			model.PublicMCPTypeDocs,
			mcpservers.WithTags([]string{"database"}),
			mcpservers.WithGitHubURL("https://github.com/oceanbase/mcp-oceanbase"),
			mcpservers.WithDescription(
				"A model context protocol (MCP) server for secure interaction with OceanBase databases. This server allows AI assistants to list tables, read data, and execute SQL queries through controlled interfaces, making database exploration and analysis safer and more structured.",
			),
			mcpservers.WithDescriptionCN(
				"一个模型上下文协议 (MCP) 服务器，用于实现与 OceanBase 数据库的安全交互。该服务器允许 AI 助手通过受控接口列出表格、读取数据并执行 SQL 查询，从而使数据库的探索和分析更加安全、结构化。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
