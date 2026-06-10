package ableton

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
			"ableton",
			"Ableton",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Ableton"),
			mcpservers.WithTags([]string{"music"}),
			mcpservers.WithGitHubURL(
				"https://github.com/ahujasid/ableton-mcp",
			),
			mcpservers.WithDescription(
				"Connects Ableton Live to Claude AI through the Model Context Protocol (MCP), allowing Claude to interact with and control Ableton Live directly, enabling AI-assisted music production.",
			),
			mcpservers.WithDescriptionCN(
				"通过模型上下文协议将Ableton Live与Claude AI连接起来，通过允许Claude直接与Ableton Live会话交互和控制，实现人工智能辅助的音乐制作。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
