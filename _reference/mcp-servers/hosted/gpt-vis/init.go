package gptvis

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
			"gpt-vis",
			"GPT Vis",
			model.PublicMCPTypeEmbed,
			mcpservers.WithNameCN("可视化图表"),
			mcpservers.WithNewServerFunc(NewServer),
			mcpservers.WithListToolsFunc(ListTools),
			mcpservers.WithGitHubURL(
				"https://github.com/antvis/mcp-server-chart",
			),
			mcpservers.WithConfigTemplates(configTemplates),
			mcpservers.WithTags([]string{"chart", "visualization", "graph", "plot", "diagram"}),
			mcpservers.WithDescription(
				"Use ant GPT Vis to generate charts. By default, it uses the free service at https://antv-studio.alipay.com/api/gpt-vis. Supports various chart types including line, column, pie, and more.",
			),
			mcpservers.WithDescriptionCN(
				"使用蚂蚁GPT Vis生成图表。默认使用 https://antv-studio.alipay.com/api/gpt-vis 的免费服务。支持多种图表类型，包括折线图、柱状图、饼图等。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
