package fetch

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
			"fetch",
			"Fetch",
			model.PublicMCPTypeEmbed,
			mcpservers.WithNameCN("网页内容获取"),
			mcpservers.WithNewServerFunc(NewServer),
			mcpservers.WithListToolsFunc(ListTools),
			mcpservers.WithGitHubURL(
				"https://github.com/modelcontextprotocol/servers/tree/main/src/fetch",
			),
			mcpservers.WithConfigTemplates(configTemplates),
			mcpservers.WithTags([]string{"fetch", "web", "html", "markdown"}),
			mcpservers.WithDescription(
				"A Model Context Protocol server that provides web content fetching capabilities. This server enables LLMs to retrieve and process content from web pages, converting HTML to markdown for easier consumption.",
			),
			mcpservers.WithDescriptionCN(
				"提供网页内容获取功能的模型上下文协议服务器。此服务器使LLM能够从网页检索和处理内容，将HTML转换为markdown以便于使用。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
