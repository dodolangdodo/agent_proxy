package brightdata

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
			"brightdata",
			"Bright Data",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Bright Data"),
			mcpservers.WithTags([]string{"web-scraping", "data-extraction", "proxy"}),
			mcpservers.WithGitHubURL(
				"https://github.com/brightdata/brightdata-mcp",
			),
			mcpservers.WithDescription(
				"Enhance AI Agents with Real-Time Web Data. Enables LLMs, agents and apps to access, discover and extract web data in real-time, allowing seamless web searching, website navigation, and data retrieval without getting blocked.",
			),
			mcpservers.WithDescriptionCN(
				"使用实时网络数据增强AI代理。让LLM、代理和应用程序能够实时访问、发现和提取网络数据，允许无缝的网络搜索、网站导航和数据检索而不被阻止。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
