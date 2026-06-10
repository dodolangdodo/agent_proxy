package browserbase

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
			"browserbase",
			"BrowserBase Cloud Browser Automation",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("BrowserBase云浏览器自动化"),
			mcpservers.WithTags([]string{"browser"}),
			mcpservers.WithGitHubURL(
				"https://github.com/browserbase/mcp-server-browserbase/tree/main/browserbase",
			),
			mcpservers.WithDescription(
				"This server uses Browserbase, Puppeteer, and Stagehand to provide cloud browser automation features. This server allows large language models (LLMs) to interact with web pages, take screenshots, and execute JavaScript in a cloud browser environment.",
			),
			mcpservers.WithDescriptionCN(
				"该服务器使用 Browserbase、Puppeteer 和 Stagehand 提供云浏览器自动化功能。此服务器使大型语言模型（LLMs）能够与网页交互、截屏以及在云浏览器环境中执行 JavaScript。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
