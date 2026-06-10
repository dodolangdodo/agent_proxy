package firecrawl

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
			"firecrawl",
			"Firecrawl",
			model.PublicMCPTypeEmbed,
			mcpservers.WithNewServerFunc(NewServer),
			mcpservers.WithListToolsFunc(ListTools),
			mcpservers.WithGitHubURL(
				"https://github.com/mendableai/firecrawl-mcp-server",
			),
			mcpservers.WithConfigTemplates(configTemplates),
			mcpservers.WithTags(
				[]string{"firecrawl", "web", "scraping", "crawling", "search", "extraction"},
			),
			mcpservers.WithDescription(
				"A Model Context Protocol (MCP) server implementation that integrates with Firecrawl for web scraping capabilities. Supports web scraping, crawling, discovery, search, content extraction, deep research and batch scraping.",
			),
			mcpservers.WithDescriptionCN(
				"集成Firecrawl网页抓取功能的模型上下文协议(MCP)服务器实现。支持网页抓取、爬取、发现、搜索、内容提取、深度研究和批量抓取。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
