package markdownify

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
			"markdownify",
			"Markdownify",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Markdownify"),
			mcpservers.WithGitHubURL(
				"https://github.com/zcaceres/markdownify-mcp",
			),
			mcpservers.WithTags([]string{"markdown"}),
			mcpservers.WithDescription(
				"Convert various file types and web content to Markdown format. It provides a set of tools that can convert PDF, image, audio file, web page, etc. to readable and shareable Markdown text.",
			),
			mcpservers.WithDescriptionCN(
				"将各种文件类型和网页内容转换为Markdown格式。它提供了一套工具，可以将PDF、图像、音频文件、网页等转换为易读且易于分享的Markdown文本。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
