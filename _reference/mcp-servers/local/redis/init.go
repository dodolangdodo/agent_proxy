package redis

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
			"redis",
			"Redis",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Redis"),
			mcpservers.WithTags([]string{"database"}),
			mcpservers.WithGitHubURL(
				"https://github.com/redis/mcp-redis",
			),
			mcpservers.WithDescription(
				"A natural language interface designed for agentic applications to efficiently manage and search data in Redis.",
			),
			mcpservers.WithDescriptionCN(
				"为智能应用程序设计的自然语言接口，用于高效管理和搜索 Redis 中的数据。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
