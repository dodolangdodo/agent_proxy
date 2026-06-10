package hefengweather

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
			"hefeng-weather",
			"HeFeng Weather",
			model.PublicMCPTypeEmbed,
			mcpservers.WithNameCN("和风天气"),
			mcpservers.WithNewServerFunc(NewServer),
			mcpservers.WithListToolsFunc(ListTools),
			mcpservers.WithGitHubURL(
				"https://github.com/shanggqm/hefeng-mcp-weather",
			),
			mcpservers.WithConfigTemplates(configTemplates),
			mcpservers.WithTags([]string{"weather", "天气", "和风天气", "forecast", "china"}),
			mcpservers.WithDescription(
				"A Model Context Protocol server that provides weather forecast data for locations in China through HeFeng Weather API. Supports real-time weather, hourly forecasts, and daily forecasts.",
			),
			mcpservers.WithDescriptionCN("通过和风天气API为中国地区提供天气预报数据的模型上下文协议服务器。支持实时天气、小时预报和日预报。"),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
