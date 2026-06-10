package alipay

import (
	_ "embed"

	"github.com/labring/aiproxy/core/model"
	mcpservers "github.com/labring/aiproxy/mcp-servers"
)

//go:embed README.md
var readme string

// need import in mcpregister/init.go
func init() {
	mcpservers.Register(
		mcpservers.NewMcp(
			"alipay",
			"Alipay",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("支付宝"),
			mcpservers.WithTags([]string{"pay"}),
			mcpservers.WithDescription(
				"支付宝 MCP Server，让你可以轻松将支付宝开放平台提供的交易创建、查询、退款等能力集成到你的 LLM 应用中，并进一步创建具备支付能力的智能工具。",
			),
			mcpservers.WithReadme(readme),
		),
	)
}
