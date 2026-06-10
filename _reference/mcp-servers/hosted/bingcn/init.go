package bingcn

import (
	_ "embed"

	"github.com/labring/aiproxy/core/model"
	mcpservers "github.com/labring/aiproxy/mcp-servers"
)

//go:embed README.md
var readme string

func init() {
	mcpservers.Register(
		mcpservers.NewMcp(
			"bing-cn-search",
			"Bing CN Search",
			model.PublicMCPTypeEmbed,
			mcpservers.WithNameCN("必应中国搜索"),
			mcpservers.WithNewServerFunc(NewServer),
			mcpservers.WithListToolsFunc(ListTools),
			mcpservers.WithGitHubURL(
				"https://github.com/yan5236/bing-cn-mcp-server",
			),
			mcpservers.WithConfigTemplates(configTemplates),
			mcpservers.WithTags([]string{"search", "bing", "web", "scraping"}),
			mcpservers.WithDescription(
				"A Model Context Protocol server that provides access to Bing CN search results. This server enables LLMs to search Bing CN and retrieve web page content.",
			),
			mcpservers.WithDescriptionCN(
				"一个提供必应中国搜索结果的模型上下文协议服务器。此服务器使LLM能够搜索必应中国并检索网页内容。",
			),
			mcpservers.WithReadme(readme),
		),
	)
}
