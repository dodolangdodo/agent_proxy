package xiyan

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
			"xiyan",
			"Xiyan",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("析言"),
			mcpservers.WithTags([]string{"database"}),
			mcpservers.WithGitHubURL("https://github.com/XGenerationLab/xiyan_mcp_server"),
			mcpservers.WithDescription(
				"A database query MCP service with natural language interaction, using a specialized SOTA text-to-sql model for higher query accuracy",
			),
			mcpservers.WithDescriptionCN(
				"一种自然语言交互的数据库查询MCP服务，嵌入专用text-to-sql的SOTA模型获得更高的查询精度",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
