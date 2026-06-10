package zhipuai

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
			"zhipuai",
			"ZhipuAI",
			model.PublicMCPTypeDocs,
			mcpservers.WithTags([]string{"search"}),
			mcpservers.WithGitHubURL("https://github.com/THUDM"),
			mcpservers.WithDescription(
				"ZhipuAI MCP Server, helps you easily integrate the transaction creation, query, and refund capabilities of the ZhipuAI MCP server into your LLM application, allowing you to further develop smart tools with payment capabilities.",
			),
			mcpservers.WithDescriptionCN(
				"智谱网络搜索MCP服务器是专为大模型设计的搜索引擎，整合了四种搜索引擎，可以让用户灵活对比和切换。在传统搜索引擎的网页爬取和排序能力基础上，增强意图识别能力，返回更适用于大模型处理的结果（如网页标题、网址、摘要、站点名称、站点图标等），帮助AI应用实现“动态知识获取”和“精准场景适配”的能力。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
