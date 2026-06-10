package microsoftdocs

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
			"microsoft-docs",
			"Microsoft Docs",
			model.PublicMCPTypeProxyStreamable,
			mcpservers.WithNameCN("Microsoft 文档"),
			mcpservers.WithProxyConfigTemplates(configTemplates),
			mcpservers.WithTags([]string{"documentation", "microsoft"}),
			mcpservers.WithGitHubURL(
				"https://github.com/MicrosoftDocs/mcp",
			),
			mcpservers.WithDescription(
				"A cloud-hosted MCP server that provides AI assistants with real-time access to official Microsoft documentation through semantic search and content retrieval.",
			),
			mcpservers.WithDescriptionCN(
				"一个云托管的MCP服务器，通过语义搜索和内容检索为AI助手提供对官方Microsoft文档的实时访问。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
