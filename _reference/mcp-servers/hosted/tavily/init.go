package tavily

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
			"tavily",
			"Tavily AI Search",
			model.PublicMCPTypeEmbed,
			mcpservers.WithNameCN("Tavily AI 搜索"),
			mcpservers.WithNewServerFunc(NewServer),
			mcpservers.WithListToolsFunc(ListTools),
			mcpservers.WithDescription(
				"A powerful web search MCP server powered by Tavily's AI search engine. Provides real-time web search, content extraction, web crawling, and site mapping capabilities with advanced filtering and customization options.",
			),
			mcpservers.WithDescriptionCN(
				"基于Tavily AI搜索引擎的强大网络搜索MCP服务器。提供实时网络搜索、内容提取、网页爬取和站点映射功能，具有高级过滤和自定义选项。",
			),
			mcpservers.WithGitHubURL("https://github.com/tavily-ai/tavily-mcp"),
			mcpservers.WithConfigTemplates(configTemplates),
			mcpservers.WithTags(
				[]string{"search", "ai", "web", "crawling", "extraction", "tavily"},
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
