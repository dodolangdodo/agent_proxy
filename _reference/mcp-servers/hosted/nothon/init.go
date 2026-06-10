package notion

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
			"notion",
			"Notion",
			model.PublicMCPTypeEmbed,
			mcpservers.WithNameCN("Notion"),
			mcpservers.WithNewServerFunc(NewServer),
			mcpservers.WithListToolsFunc(ListTools),
			mcpservers.WithConfigTemplates(configTemplates),
			mcpservers.WithGitHubURL(
				"https://github.com/suekou/mcp-notion-server",
			),
			mcpservers.WithTags([]string{"notion", "productivity", "notes", "database", "blocks"}),
			mcpservers.WithDescription(
				"A Model Context Protocol server that provides comprehensive access to Notion workspaces. This server enables LLMs to interact with Notion pages, databases, blocks, and users through the official Notion API.",
			),
			mcpservers.WithDescriptionCN(
				"提供对Notion工作区全面访问的模型上下文协议服务器。此服务器使LLM能够通过官方Notion API与Notion页面、数据库、块和用户进行交互。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
