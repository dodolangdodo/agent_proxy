package codesandbox

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
			"code-sandbox",
			"Code Sandbox",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("代码沙箱"),
			mcpservers.WithTags([]string{"code-sandbox", "code", "sandbox"}),
			mcpservers.WithGitHubURL(
				"https://github.com/Automata-Labs-team/code-sandbox-mcp",
			),
			mcpservers.WithDescription(
				"An MCP server to create secure code sandbox environment for executing code within Docker containers. This MCP server provides AI applications with a safe and isolated environment for running code while maintaining security through containerization.",
			),
			mcpservers.WithDescriptionCN(
				"一个 MCP 服务器，用于在 Docker 容器内创建安全的代码沙箱环境以执行代码。该 MCP 服务器为 AI 应用程序提供了一个安全且隔离的运行代码环境，同时通过容器化技术保障安全性。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
