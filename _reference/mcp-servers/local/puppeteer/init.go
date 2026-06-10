package puppeteer

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
			"puppeteer",
			"Puppeteer",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Puppeteer"),
			mcpservers.WithTags([]string{"browser"}),
			mcpservers.WithGitHubURL(
				"https://github.com/modelcontextprotocol/servers-archived/tree/main/src/puppeteer",
			),
			mcpservers.WithDescription(
				"A Model Context Protocol (MCP) server that provides browser automation capabilities using Puppeteer. This server allows LLMs to interact with web pages, take screenshots, and execute JavaScript in a real browser environment.",
			),
			mcpservers.WithDescriptionCN(
				"一个使用Puppeteer提供浏览器自动化功能的Model Context Protocol服务器。此服务器使LLM能够与网页交互、截屏以及在真实的浏览器环境中执行JavaScript。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
