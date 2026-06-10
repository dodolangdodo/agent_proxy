package unionpay

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
			"unionpay",
			"UnionPay",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("银联"),
			mcpservers.WithTags([]string{"pay"}),
			mcpservers.WithGitHubURL("https://github.com/allvoicelab/AllVoiceLab-MCP"),
			mcpservers.WithDescription(
				"UnionPay MCP Server, helps you easily integrate the transaction creation, query, and refund capabilities of the UnionPay MCP server into your LLM application, allowing you to further develop smart tools with payment capabilities.",
			),
			mcpservers.WithDescriptionCN(
				"银联开放平台提供的MCP服务器，可以帮助您轻松地将银联开放平台的交易创建、查询、退款等能力集成到您的LLM应用中，让您进一步开发带有支付功能的智能工具。",
			),
			mcpservers.WithReadme(readme),
		),
	)
}
