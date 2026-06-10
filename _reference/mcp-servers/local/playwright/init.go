package playwright

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
			"playwright",
			"Microsoft Playwright",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("微软Playwright"),
			mcpservers.WithTags([]string{"browser"}),
			mcpservers.WithGitHubURL(
				"https://github.com/microsoft/playwright-mcp",
			),
			mcpservers.WithDescription(
				"A Model Context Protocol (MCP) server implementation that allows large language models to interact with web pages through structured accessibility snapshots, without using visual models or screenshots.",
			),
			mcpservers.WithDescriptionCN(
				"使大型语言模型能够通过结构化的可访问性快照与网页交互，而无需使用视觉模型或截图。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
