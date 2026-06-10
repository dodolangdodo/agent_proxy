package context7

import (
	"github.com/labring/aiproxy/core/model"
	mcpservers "github.com/labring/aiproxy/mcp-servers"
)

const defaultURL = "https://mcp.context7.com/mcp"

var configTemplates = mcpservers.ProxyConfigTemplates{
	"url": {
		ConfigTemplate: mcpservers.ConfigTemplate{
			Name:        "URL",
			Required:    mcpservers.ConfigRequiredTypeInitOptional,
			Example:     defaultURL,
			Default:     defaultURL,
			Description: "The Streamable http URL of the Context7 MCP server",
		},
		Type: model.ParamTypeURL,
	},
}
