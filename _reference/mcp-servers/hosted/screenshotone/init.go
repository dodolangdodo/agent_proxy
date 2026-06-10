package screenshotone

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
			"screenshotone",
			"ScreenshotOne",
			model.PublicMCPTypeEmbed,
			mcpservers.WithNameCN("网页截图"),
			mcpservers.WithNewServerFunc(NewServer),
			mcpservers.WithListToolsFunc(ListTools),
			mcpservers.WithGitHubURL(
				"https://github.com/screenshotone/mcp",
			),
			mcpservers.WithConfigTemplates(configTemplates),
			mcpservers.WithTags([]string{"screenshot", "web", "image", "capture", "website"}),
			mcpservers.WithDescription(
				"A Model Context Protocol server that provides website screenshot capabilities using the ScreenshotOne API. This server enables LLMs to capture screenshots of websites and return them as images.",
			),
			mcpservers.WithDescriptionCN(
				"使用ScreenshotOne API提供网页截图功能的模型上下文协议服务器。此服务器使LLM能够捕获网站截图并将其作为图像返回。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
