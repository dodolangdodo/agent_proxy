package context7

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
			"context7",
			"Context7",
			model.PublicMCPTypeProxyStreamable,
			mcpservers.WithNameCN("Context7"),
			mcpservers.WithGitHubURL("https://github.com/upstash/context7"),
			mcpservers.WithProxyConfigTemplates(configTemplates),
			mcpservers.WithTags([]string{"context7", "document", "code", "copilot"}),
			mcpservers.WithDescription(
				"Context7 is a tool that can help you find the latest documents and fight against code hallucinations.",
			),
			mcpservers.WithDescriptionCN(
				"Context7 可以获取最新文档，对抗代码幻觉。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
