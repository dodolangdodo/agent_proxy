package websearch

import (
	_ "embed"

	"github.com/labring/aiproxy/core/model"
	mcpservers "github.com/labring/aiproxy/mcp-servers"
)

//go:embed features.md
var readme string

//go:embed README.cn.md
var readmeCN string

// Register the server
func init() {
	mcpservers.Register(
		mcpservers.NewMcp(
			"web-search",
			"Web Search",
			model.PublicMCPTypeEmbed,
			mcpservers.WithNameCN("网络搜索"),
			mcpservers.WithNewServerFunc(NewServer),
			mcpservers.WithListToolsFunc(ListTools),
			mcpservers.WithConfigTemplates(configTemplates),
			mcpservers.WithTags([]string{"search", "web", "google", "bing", "arxiv", "searchxng"}),
			mcpservers.WithDescription(
				"A comprehensive web search MCP server that provides access to multiple search engines including Google, Bing, Bing CN(Free), and Arxiv.",
			),
			mcpservers.WithDescriptionCN(
				"一个综合的网络搜索MCP服务器，提供对Google、Bing、Bing CN(免费)和Arxiv等多个搜索引擎的访问。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
