package redbooksearchcomment

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
			"redbook-search-comment",
			"Redbook Search Comment",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("小红书搜索评论"),
			mcpservers.WithGitHubURL(
				"https://github.com/chenningling/Redbook-Search-Comment-MCP2.0",
			),
			mcpservers.WithTags([]string{"search"}),
			mcpservers.WithDescription(
				"Help users automatically complete login to Xiaohongshu, search keywords, get note content, and publish AI-generated comments",
			),
			mcpservers.WithDescriptionCN(
				"帮助用户自动完成登录小红书、搜索关键词、获取笔记内容及发布AI生成评论等操作",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
