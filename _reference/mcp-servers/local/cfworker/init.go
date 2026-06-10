package cfworker

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
			"cfworker",
			"Cloudflare Workers",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Cloudflare Workers"),
			mcpservers.WithTags([]string{"cloud"}),
			mcpservers.WithGitHubURL(
				"https://github.com/cloudflare/workers-mcp",
			),
			mcpservers.WithDescription(
				"A Model Context Protocol (MCP) server for Cloudflare Workers, enabling custom functions to be accessed through natural language.",
			),
			mcpservers.WithDescriptionCN(
				"一个模型上下文协议（MCP）服务器，用于Cloudflare Workers，通过模型上下文协议（Model Context Protocol）， enables 自定义功能可以通过自然语言访问。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
