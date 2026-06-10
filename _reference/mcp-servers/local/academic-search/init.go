package academicsearch

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
			"academic-search",
			"Academic Search",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("学术搜索"),
			mcpservers.WithGitHubURL(
				"https://github.com/afrise/academic-search-mcp-server",
			),
			mcpservers.WithTags([]string{"academic", "search"}),
			mcpservers.WithDescription(
				"A Model Context Protocol (MCP) server that allows LLMs to search for academic papers using natural language.",
			),
			mcpservers.WithDescriptionCN(
				"一个通过自然语言搜索学术论文的MCP服务器。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
