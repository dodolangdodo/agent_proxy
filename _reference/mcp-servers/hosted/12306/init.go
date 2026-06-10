// File: mcp-servers/12306/init.go
package train12306

import (
	_ "embed"

	"github.com/labring/aiproxy/core/model"
	mcpservers "github.com/labring/aiproxy/mcp-servers"
)

//go:embed README.md
var readme string

//go:embed README.cn.md
var readmeCN string

func init() {
	mcpservers.Register(
		mcpservers.NewMcp(
			"12306",
			"12306 Train Ticket Query",
			model.PublicMCPTypeEmbed,
			mcpservers.WithNameCN("12306 购票搜索"),
			mcpservers.WithNewServerFunc(NewServer),
			mcpservers.WithGitHubURL(
				"https://github.com/Joooook/12306-mcp",
			),
			mcpservers.WithConfigTemplates(configTemplates),
			mcpservers.WithTags([]string{"train", "12306", "ticket", "china", "railway"}),
			mcpservers.WithDescription(
				"A 12306 ticket search server based on the Model Context Protocol (MCP). Provides simple API interface for searching 12306 train tickets, filtering train information, route queries, and transfer queries.",
			),
			mcpservers.WithDescriptionCN(
				"基于模型上下文协议(MCP)的12306购票搜索服务器。提供简单的API接口，允许搜索12306购票信息、过滤列车信息、过站查询和中转查询。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
