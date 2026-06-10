package milvus

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
			"milvus",
			"Milvus",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Milvus"),
			mcpservers.WithTags([]string{"vector-database", "search"}),
			mcpservers.WithGitHubURL(
				"https://github.com/zilliztech/mcp-server-milvus",
			),
			mcpservers.WithDescription(
				"A Model Context Protocol (MCP) server that provides access to Milvus vector database functionality, enabling seamless integration between LLM applications and vector search capabilities.",
			),
			mcpservers.WithDescriptionCN(
				"一个模型上下文协议（MCP）服务器，提供对Milvus向量数据库功能的访问，实现LLM应用程序与向量搜索功能之间的无缝集成。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
