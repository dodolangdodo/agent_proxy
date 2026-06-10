package semgrep

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
			"semgrep",
			"Semgrep",
			model.PublicMCPTypeProxyStreamable,
			mcpservers.WithNameCN("Semgrep"),
			mcpservers.WithGitHubURL("https://github.com/semgrep/semgrep-mcp"),
			mcpservers.WithProxyConfigTemplates(configTemplates),
			mcpservers.WithTags([]string{"semgrep", "security", "code", "scan"}),
			mcpservers.WithDescription(
				"A Model Context Protocol (MCP) server for using Semgrep to scan code for security vulnerabilities.",
			),
			mcpservers.WithDescriptionCN(
				"一个用于使用 Semgrep 扫描代码安全漏洞的模型上下文协议 (MCP) 服务器",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
