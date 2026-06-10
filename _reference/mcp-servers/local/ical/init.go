package ical

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
			"ical",
			"Macos Calendar",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Macos Calendar"),
			mcpservers.WithTags([]string{"ical", "calendar", "macos"}),
			mcpservers.WithGitHubURL(
				"https://github.com/Omar-v2/mcp-ical",
			),
			mcpservers.WithDescription(
				"A Model Context Protocol Server that allows you to interact with your MacOS Calendar through natural language.",
			),
			mcpservers.WithDescriptionCN(
				"一个模型上下文协议服务器，让你能够通过自然语言与你的 MacOS 日历进行交互。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
