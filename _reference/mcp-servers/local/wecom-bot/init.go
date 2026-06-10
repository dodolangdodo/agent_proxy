package wecombot

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
			"wecom-bot",
			"WeCom Bot",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("企业微信机器人"),
			mcpservers.WithGitHubURL(
				"https://github.com/mamertofabian/mcp-wecom-bot",
			),
			mcpservers.WithTags([]string{"wecom"}),
			mcpservers.WithDescription(
				"A use fastmcp to send messages via WeCom bot, supports asynchronous communication and message tracking via Webhook.",
			),
			mcpservers.WithDescriptionCN(
				"一个使用FastMCP通过企业微信机器人发送消息的服务器，支持通过Webhook进行异步通信和消息追踪。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
