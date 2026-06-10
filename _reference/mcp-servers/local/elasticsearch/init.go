package elasticsearch

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
			"elasticsearch",
			"Elasticsearch",
			model.PublicMCPTypeDocs,
			mcpservers.WithTags([]string{"database"}),
			mcpservers.WithGitHubURL("https://github.com/elastic/mcp-server-elasticsearch"),
			mcpservers.WithDescription(
				"Connects Claude and other MCP clients to Elasticsearch data, allowing users to interact with their Elasticsearch indexes through natural language conversations.",
			),
			mcpservers.WithDescriptionCN(
				"将Claude和其他MCP客户端连接到Elasticsearch数据，允许用户通过自然语言对话与其Elasticsearch索引进行交互。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
