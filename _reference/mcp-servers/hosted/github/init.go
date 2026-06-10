package github

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
			"github",
			"GitHub",
			model.PublicMCPTypeProxyStreamable,
			mcpservers.WithNameCN("GitHub"),
			mcpservers.WithGitHubURL("https://github.com/github/github-mcp-server"),
			mcpservers.WithProxyConfigTemplates(configTemplates),
			mcpservers.WithTags([]string{"git", "github", "copilot"}),
			mcpservers.WithDescription(
				"GitHub MCP Server provides comprehensive GitHub services including repository search, issue tracking, and more through GitHub's APIs.",
			),
			mcpservers.WithDescriptionCN(
				"GitHub MCP 服务器通过 GitHub API 提供全面的 GitHub 服务，包括仓库搜索、问题跟踪等。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
