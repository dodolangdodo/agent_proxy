package hfspace

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
			"hfspace",
			"HuggingFace Space",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("HuggingFace Space"),
			mcpservers.WithTags([]string{"huggingface"}),
			mcpservers.WithGitHubURL(
				"https://github.com/evalstate/mcp-hfspace",
			),
			mcpservers.WithDescription(
				"MCP Server to Use HuggingFace spaces, easy configuration and Claude Desktop mode.",
			),
			mcpservers.WithDescriptionCN(
				"MCP 服务器支持使用 HuggingFace 空间，配置简单且具备 Claude 桌面模式。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
