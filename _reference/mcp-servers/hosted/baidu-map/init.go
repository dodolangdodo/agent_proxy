package baidumap

import (
	_ "embed"

	"github.com/labring/aiproxy/core/model"
	mcpservers "github.com/labring/aiproxy/mcp-servers"
)

//go:embed README.md
var readme string

func init() {
	mcpservers.Register(
		mcpservers.NewMcp(
			"baidu-map",
			"Baidu Map",
			model.PublicMCPTypeEmbed,
			mcpservers.WithNameCN("百度地图"),
			mcpservers.WithNewServerFunc(NewServer),
			mcpservers.WithListToolsFunc(ListTools),
			mcpservers.WithGitHubURL(
				"https://github.com/baidu-maps/mcp",
			),
			mcpservers.WithConfigTemplates(configTemplates),
			mcpservers.WithTags(
				[]string{"baidu", "map", "geocoding", "search", "weather", "traffic"},
			),
			mcpservers.WithDescription(
				"Baidu Map API now fully supports MCP protocol, the first map service provider in China to be compatible with MCP protocol. Includes 10 API interfaces that comply with MCP protocol standards, covering reverse geocoding, place search, route planning, etc.",
			),
			mcpservers.WithDescriptionCN(
				"百度地图API现已全面兼容MCP协议，是国内首家兼容MCP协议的地图服务商。包含10个符合MCP协议标准的API接口，涵盖逆地理编码、地点检索、路线规划等。",
			),
			mcpservers.WithReadme(readme),
		),
	)
}
