package gezhe

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
			"gezhe",
			"Gezhe",
			model.PublicMCPTypeProxyStreamable,
			mcpservers.WithNameCN("歌者"),
			mcpservers.WithProxyConfigTemplates(configTemplates),
			mcpservers.WithTags([]string{"map"}),
			mcpservers.WithDescription(
				"Gezhe MCP Server provides comprehensive Gezhe services including topic generation, PPT generation, and more through Gezhe's APIs.",
			),
			mcpservers.WithDescriptionCN(
				"歌者 MCP 服务器通过歌者 API 提供全面的歌者服务，包括话题生成 PPT 等。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
