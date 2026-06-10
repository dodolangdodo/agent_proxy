package figma

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
			"figma",
			"Figma Design",
			model.PublicMCPTypeEmbed,
			mcpservers.WithNameCN("Figma 设计"),
			mcpservers.WithNewServerFunc(NewServer),
			mcpservers.WithListToolsFunc(ListTools),
			mcpservers.WithGitHubURL(
				"https://github.com/GLips/Figma-Context-MCP",
			),
			mcpservers.WithConfigTemplates(configTemplates),
			mcpservers.WithTags([]string{"figma", "design", "ui", "ux", "prototyping"}),
			mcpservers.WithDescription(
				"A Model Context Protocol server that provides access to Figma design files. This server enables LLMs to retrieve and process Figma design data, converting it to simplified formats for easier consumption.",
			),
			mcpservers.WithDescriptionCN(
				"提供对Figma设计文件访问的模型上下文协议服务器。此服务器使LLM能够检索和处理Figma设计数据，将其转换为简化格式以便于使用。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
