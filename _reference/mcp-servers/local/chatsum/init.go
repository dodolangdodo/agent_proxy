package chatsum

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
			"chatsum",
			"ChatSum",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("聊天摘要"),
			mcpservers.WithGitHubURL(
				"https://github.com/chatmcp/mcp-server-chatsum",
			),
			mcpservers.WithTags([]string{"chat"}),
			mcpservers.WithDescription(
				"A Model Context Protocol (MCP) server that allows LLMs to generate summaries of chat conversations using natural language.",
			),
			mcpservers.WithDescriptionCN(
				"一种用于生成聊天会话摘要的模型上下文协议（MCP）服务器。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
