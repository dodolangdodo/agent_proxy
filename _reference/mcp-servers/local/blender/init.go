package blender

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
			"blender",
			"Blender",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Blender"),
			mcpservers.WithTags([]string{"3d"}),
			mcpservers.WithGitHubURL(
				"https://github.com/ahujasid/blender-mcp",
			),
			mcpservers.WithDescription(
				"Connects Blender to Claude AI through the Model Context Protocol (MCP), allowing Claude to interact with and control Blender directly, enabling AI-assisted 3D modeling, scene operations, and rendering.",
			),
			mcpservers.WithDescriptionCN(
				"通过模型上下文协议（MCP）将Blender连接到Claude AI，使Claude能够直接与Blender交互并对其进行控制，从而实现AI辅助的3D建模、场景操作和渲染。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
