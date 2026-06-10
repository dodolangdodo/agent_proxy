package gitee

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
			"gitee",
			"Gitee",
			model.PublicMCPTypeProxyStreamable,
			mcpservers.WithNameCN("Gitee"),
			mcpservers.WithGitHubURL("https://github.com/oschina/mcp-gitee"),
			mcpservers.WithProxyConfigTemplates(configTemplates),
			mcpservers.WithTags([]string{"git", "gitee", "copilot"}),
			mcpservers.WithDescription(
				"Gitee MCP Server provides comprehensive Gitee services including repository search, issue tracking, and more through Gitee's APIs.",
			),
			mcpservers.WithDescriptionCN(
				"Gitee MCP 服务器通过 Gitee API 提供全面的 Gitee 服务，包括仓库搜索、问题跟踪等。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
