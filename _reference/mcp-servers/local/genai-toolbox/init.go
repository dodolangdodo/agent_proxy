package genaitoolbox

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
			"genai-toolbox",
			"MCP Toolbox for Databases",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("数据库 MCP 工具箱"),
			mcpservers.WithTags([]string{"database", "toolbox", "ai"}),
			mcpservers.WithGitHubURL(
				"https://github.com/googleapis/genai-toolbox",
			),
			mcpservers.WithDescription(
				"An open source MCP server for databases that enables you to develop tools easier, faster, and more securely by handling complexities such as connection pooling, authentication, and more.",
			),
			mcpservers.WithDescriptionCN(
				"一个开源的数据库 MCP 服务器，通过处理连接池、身份验证等复杂性，使您能够更轻松、更快速、更安全地开发工具。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
