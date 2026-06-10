package xhs

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
			"xhs",
			"小红书",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("小红书"),
			mcpservers.WithTags([]string{"docs"}),
			mcpservers.WithDescription(
				"A Model Context Protocol (MCP) server that allows LLMs to publish Red Notes on Xiaohongshu (Red Notebook) using natural language.",
			),
			mcpservers.WithDescriptionCN(
				"一种用于发布小红书（红笔记）的模型上下文协议（MCP）服务器",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
