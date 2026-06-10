package weather

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
			"weather",
			"Weather",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("天气"),
			mcpservers.WithGitHubURL(
				"https://github.com/CodeByWaqas/weather-mcp-server",
			),
			mcpservers.WithTags([]string{"weather"}),
			mcpservers.WithDescription(
				"A MCP server that provides real-time weather information (including temperature, humidity, wind speed, and sunrise/sunset times) through the OpenWeatherMap API.",
			),
			mcpservers.WithDescriptionCN(
				"一个通过OpenWeatherMap API提供实时天气信息（包括温度、湿度、风速和日出/日落时间）的MCP服务器。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
