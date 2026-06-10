package magic

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
			"magic",
			"Magic Component Generator",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("智能组件生成器"),
			mcpservers.WithTags([]string{"ui"}),
			mcpservers.WithGitHubURL(
				"https://github.com/21st-dev/magic-mcp",
			),
			mcpservers.WithDescription(
				"A powerful AI-driven tool that helps developers create beautiful, modern UI components instantly through natural language descriptions. It integrates seamlessly with popular IDEs and provides a streamlined workflow for UI development.",
			),
			mcpservers.WithDescriptionCN(
				"一个由人工智能驱动的工具，可以从自然语言描述生成现代UI组件，并与流行的IDE集成，以简化UI开发工作流程。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
