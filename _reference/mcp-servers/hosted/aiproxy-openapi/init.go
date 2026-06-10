package aiproxyopenapi

import (
	"github.com/labring/aiproxy/core/model"
	mcpservers "github.com/labring/aiproxy/mcp-servers"
)

// need import in mcpregister/init.go
func init() {
	mcpservers.Register(
		mcpservers.NewMcp(
			"aiproxy-openapi",
			"AI Proxy OpenAPI",
			model.PublicMCPTypeEmbed,
			mcpservers.WithNewServerFunc(NewServer),
			mcpservers.WithConfigTemplates(configTemplates),
			mcpservers.WithDescription(
				"AI Proxy OpenAPI MCP Server provides access to AI Proxy's administrative and management APIs through the Model Context Protocol, enabling automated management of AI services.",
			),
			mcpservers.WithDescriptionCN(
				"AI Proxy OpenAPI MCP服务器通过模型上下文协议提供对AI Proxy管理和运营API的访问，实现AI服务的自动化管理。",
			),
		),
	)
}
