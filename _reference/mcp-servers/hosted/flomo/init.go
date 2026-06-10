package flomo

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
			"flomo",
			"Flomo",
			model.PublicMCPTypeEmbed,
			mcpservers.WithNewServerFunc(NewServer),
			mcpservers.WithListToolsFunc(ListTools),
			mcpservers.WithGitHubURL(
				"https://github.com/chatmcp/mcp-server-flomo",
			),
			mcpservers.WithConfigTemplates(configTemplates),
			mcpservers.WithTags([]string{"flomo", "notes", "writing", "productivity"}),
			mcpservers.WithDescription(
				"Write notes to Flomo. A TypeScript-based MCP server that helps you write notes to Flomo with markdown format support.",
			),
			mcpservers.WithDescriptionCN(
				"向Flomo写入笔记。基于TypeScript的MCP服务器，帮助您向Flomo写入笔记，支持markdown格式。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
