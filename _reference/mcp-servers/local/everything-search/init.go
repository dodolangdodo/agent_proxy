package everythingsearch

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
			"everything-search",
			"Everything Search",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Everything Search"),
			mcpservers.WithGitHubURL(
				"https://github.com/mamertofabian/mcp-everything-search",
			),
			mcpservers.WithTags([]string{"search"}),
			mcpservers.WithDescription(
				"Use Everything SDK's fast file search functionality",
			),
			mcpservers.WithDescriptionCN(
				"使用Everything SDK的快速文件搜索功能",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
