package time

import (
	_ "embed"

	"github.com/labring/aiproxy/core/model"
	mcpservers "github.com/labring/aiproxy/mcp-servers"
)

//go:embed README.md
var readme string

//go:embed README.cn.md
var readmeCN string

func init() {
	mcpservers.Register(
		mcpservers.NewMcp(
			"time",
			"Time",
			model.PublicMCPTypeEmbed,
			mcpservers.WithNameCN("时间"),
			mcpservers.WithNewServerFunc(NewServer),
			mcpservers.WithGitHubURL(
				"https://github.com/modelcontextprotocol/servers/tree/main/src/time",
			),
			mcpservers.WithConfigTemplates(configTemplates),
			mcpservers.WithTags([]string{"time", "timezone", "conversion", "datetime"}),
			mcpservers.WithDescription(
				"A Model Context Protocol server that provides time and timezone conversion capabilities. This server enables LLMs to get current time information and perform timezone conversions using IANA timezone names.",
			),
			mcpservers.WithDescriptionCN(
				"提供时间和时区转换功能的模型上下文协议服务器。此服务器使LLM能够获取当前时间信息并使用IANA时区名称执行时区转换。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
