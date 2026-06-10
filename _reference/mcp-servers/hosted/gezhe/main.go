package gezhe

import (
	"github.com/labring/aiproxy/core/model"
	mcpservers "github.com/labring/aiproxy/mcp-servers"
)

const defaultURL = "https://mcp.gezhe.com/mcp"

var configTemplates = mcpservers.ProxyConfigTemplates{
	"API_KEY": {
		ConfigTemplate: mcpservers.ConfigTemplate{
			Name:        "API_KEY",
			Required:    mcpservers.ConfigRequiredTypeInitOrReusingOnly,
			Example:     "bx7Qt1BLbxRq...",
			Description: "The key of the Gezhe MCP server: https://pro.gezhe.com/settings",
		},
		Type: model.ParamTypeQuery,
	},

	"url": {
		ConfigTemplate: mcpservers.ConfigTemplate{
			Name:        "URL",
			Required:    mcpservers.ConfigRequiredTypeInitOptional,
			Example:     defaultURL,
			Default:     defaultURL,
			Description: "The Streamable http URL of the gezhe MCP server",
		},
		Type: model.ParamTypeURL,
	},
}
