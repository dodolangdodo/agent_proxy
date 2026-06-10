//nolint:testpackage
package controller

import (
	"testing"

	"github.com/labring/aiproxy/core/common/config"
	"github.com/labring/aiproxy/core/model"
	"github.com/stretchr/testify/require"
)

func TestPublicMCPEndpointUsesDefaultHostForPathMode(t *testing.T) {
	resetMCPHostsForTest(t)

	config.SetDefaultHost("public.example.com")

	ep := NewPublicMCPEndpoint("request.example.local", model.PublicMCP{
		ID:   "search",
		Type: model.PublicMCPTypeProxySSE,
	})

	require.Equal(t, "public.example.com", ep.Host)
	require.Equal(t, "/mcp/public/search/sse", ep.SSE)
	require.Equal(t, "/mcp/public/search", ep.StreamableHTTP)
}

func TestPublicMCPEndpointKeepsSubdomainHostWhenConfigured(t *testing.T) {
	resetMCPHostsForTest(t)

	config.SetDefaultHost("public.example.com")
	config.SetPublicMCPHost("public-mcp.example.com")

	ep := NewPublicMCPEndpoint("request.example.local", model.PublicMCP{
		ID:   "search",
		Type: model.PublicMCPTypeProxySSE,
	})

	require.Equal(t, "search.public-mcp.example.com", ep.Host)
	require.Equal(t, "/sse", ep.SSE)
	require.Equal(t, "/mcp", ep.StreamableHTTP)
}

func TestGroupMCPEndpointUsesDefaultHostForPathMode(t *testing.T) {
	resetMCPHostsForTest(t)

	config.SetDefaultHost("public.example.com")

	resp := NewGroupMCPResponse("request.example.local", model.GroupMCP{
		ID:   "tools",
		Type: model.GroupMCPTypeProxySSE,
	})

	require.Equal(t, "public.example.com", resp.Endpoints.Host)
	require.Equal(t, "/mcp/group/tools/sse", resp.Endpoints.SSE)
	require.Equal(t, "/mcp/group/tools", resp.Endpoints.StreamableHTTP)
}

func TestGroupMCPEndpointKeepsSubdomainHostWhenConfigured(t *testing.T) {
	resetMCPHostsForTest(t)

	config.SetDefaultHost("public.example.com")
	config.SetGroupMCPHost("group-mcp.example.com")

	resp := NewGroupMCPResponse("request.example.local", model.GroupMCP{
		ID:   "tools",
		Type: model.GroupMCPTypeProxySSE,
	})

	require.Equal(t, "tools.group-mcp.example.com", resp.Endpoints.Host)
	require.Equal(t, "/sse", resp.Endpoints.SSE)
	require.Equal(t, "/mcp", resp.Endpoints.StreamableHTTP)
}

func resetMCPHostsForTest(t *testing.T) {
	t.Helper()

	t.Setenv("DEFAULT_HOST", "")
	t.Setenv("DEFAULT_MCP_HOST", "")
	t.Setenv("PUBLIC_MCP_HOST", "")
	t.Setenv("GROUP_MCP_HOST", "")

	oldDefaultHost := config.GetDefaultHost()
	oldDefaultMCPHost := config.GetConfiguredDefaultMCPHost()
	oldPublicMCPHost := config.GetPublicMCPHost()
	oldGroupMCPHost := config.GetGroupMCPHost()
	t.Cleanup(func() {
		config.SetDefaultHost(oldDefaultHost)
		config.SetDefaultMCPHost(oldDefaultMCPHost)
		config.SetPublicMCPHost(oldPublicMCPHost)
		config.SetGroupMCPHost(oldGroupMCPHost)
	})

	config.SetDefaultHost("")
	config.SetDefaultMCPHost("")
	config.SetPublicMCPHost("")
	config.SetGroupMCPHost("")
}
