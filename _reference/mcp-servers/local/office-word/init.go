package officeword

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
			"office-word",
			"Office Word",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Word文档"),
			mcpservers.WithGitHubURL(
				"https://github.com/gongrzhe/office-word-mcp-server",
			),
			mcpservers.WithTags([]string{"word"}),
			mcpservers.WithDescription(
				"A Model Context Protocol (MCP) server that allows LLMs to interact with Microsoft Word documents using natural language. It provides a set of tools that can create, edit, and read Microsoft Word documents.",
			),
			mcpservers.WithDescriptionCN(
				"一种用于与Microsoft Word文档交互的模型上下文协议（MCP）服务器，允许LLM通过自然语言与Word文档进行交互。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
