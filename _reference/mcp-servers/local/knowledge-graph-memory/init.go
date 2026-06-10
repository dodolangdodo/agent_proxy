package knowledgegraphmemory

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
			"knowledge-graph-memory",
			"Knowledge Graph Memory",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("知识图谱记忆"),
			mcpservers.WithTags([]string{"knowledge-graph"}),
			mcpservers.WithGitHubURL(
				"https://github.com/modelcontextprotocol/servers/tree/main/src/memory",
			),
			mcpservers.WithDescription(
				"A basic method for implementing persistent memory using local knowledge graphs. This allows Claude to remember user-related information across multiple conversations.",
			),
			mcpservers.WithDescriptionCN(
				"使用本地知识图谱实现持久内存的基本方法。这使得Claude能够在多次聊天中记住用户的相关信息。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
