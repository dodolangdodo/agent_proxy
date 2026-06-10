package howtocook

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
			"howtocook",
			"HowToCook Recipe Server",
			model.PublicMCPTypeEmbed,
			mcpservers.WithNameCN("程序员做饭指南"),
			mcpservers.WithNewServerFunc(NewServer),
			mcpservers.WithListToolsFunc(ListTools),
			mcpservers.WithDescription(
				"A recipe recommendation server based on the HowToCook project. Provides intelligent meal planning, recipe search by category, and dish recommendations based on the number of people.",
			),
			mcpservers.WithDescriptionCN("基于程序员做饭指南项目的菜谱推荐服务器。提供智能膳食计划、按分类搜索菜谱以及根据用餐人数推荐菜品的功能。"),
			mcpservers.WithGitHubURL("https://github.com/Anduin2017/HowToCook"),
			mcpservers.WithTags([]string{"recipe", "cooking", "meal", "food", "chinese"}),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
