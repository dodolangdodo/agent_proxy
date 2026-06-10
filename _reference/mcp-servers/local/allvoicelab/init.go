package allvoicelab

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
			"allvoicelab",
			"AllVoiceLab",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("趣丸千音"),
			mcpservers.WithTags([]string{"voice"}),
			mcpservers.WithGitHubURL("https://github.com/allvoicelab/AllVoiceLab-MCP"),
			mcpservers.WithDescription(
				"Official AllVoiceLab MCP Server, supports interaction with powerful text-to-speech and video translation APIs. Allows MCP clients like Claude Desktop, Cursor, Windsurf, OpenAI Agents, etc. to generate speech, translate videos, and smart voice changing.",
			),
			mcpservers.WithDescriptionCN(
				"官方 AllVoiceLab 模型上下文协议(MCP)服务器，支持与强大的文本转语音和视频翻译API交互。允许MCP客户端如Claude Desktop、Cursor、Windsurf、OpenAI Agents等生成语音、翻译视频、智能变声等功能。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
