package github

import (
	"github.com/labring/aiproxy/core/model"
	mcpservers "github.com/labring/aiproxy/mcp-servers"
)

const defaultURL = "https://api.githubcopilot.com/mcp/"

var configTemplates = mcpservers.ProxyConfigTemplates{
	"Authorization": {
		ConfigTemplate: mcpservers.ConfigTemplate{
			Name:        "PAT",
			Required:    mcpservers.ConfigRequiredTypeInitOrReusingOnly,
			Example:     "ghp_dvAK8Al8aWP...",
			Description: "The Personal Access Token of the GitHub MCP server: https://github.com/settings/tokens",
		},
		Type: model.ParamTypeHeader,
	},

	"url": {
		ConfigTemplate: mcpservers.ConfigTemplate{
			Name:        "URL",
			Required:    mcpservers.ConfigRequiredTypeInitOptional,
			Example:     defaultURL,
			Default:     defaultURL,
			Description: "The Streamable http URL of the GitHub MCP server",
		},
		Type: model.ParamTypeURL,
	},
}
