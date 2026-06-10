package openmemory

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
			"openmemory",
			"OpenMemory",
			model.PublicMCPTypeDocs,
			mcpservers.WithTags([]string{"search"}),
			mcpservers.WithGitHubURL("https://github.com/mem0ai/mem0/tree/main/openmemory"),
			mcpservers.WithDescription(
				"OpenMemory is your personal memory layer for large language models (LLMs) - private, portable, and open-source. Your memory data is stored locally, giving you full control over your data. While building AI applications with personalized memory, ensure your data remains secure.",
			),
			mcpservers.WithDescriptionCN(
				"OpenMemory 是您的个人记忆层，适用于大型语言模型（LLM）——私有、便携且开源。您的记忆数据存储在本地，让您完全掌控自己的数据。在构建具有个性化记忆的AI应用程序的同时，确保您的数据安全。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
