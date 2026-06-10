package config_test

import (
	"testing"

	"github.com/labring/aiproxy/core/common/config"
	"github.com/stretchr/testify/require"
)

func TestDefaultMCPHostFallsBackToDefaultHost(t *testing.T) {
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

	config.SetDefaultHost("public.example.com")
	config.SetDefaultMCPHost("")
	config.SetPublicMCPHost("")
	config.SetGroupMCPHost("")

	require.Equal(t, "public.example.com", config.GetDefaultHost())
	require.Empty(t, config.GetConfiguredDefaultMCPHost())
	require.Equal(t, "public.example.com", config.GetDefaultMCPHost())
	require.Empty(t, config.GetPublicMCPHost())
	require.Empty(t, config.GetGroupMCPHost())
}

func TestConfiguredDefaultMCPHostOverridesDefaultHost(t *testing.T) {
	t.Setenv("DEFAULT_HOST", "")
	t.Setenv("DEFAULT_MCP_HOST", "")

	oldDefaultHost := config.GetDefaultHost()
	oldDefaultMCPHost := config.GetConfiguredDefaultMCPHost()
	t.Cleanup(func() {
		config.SetDefaultHost(oldDefaultHost)
		config.SetDefaultMCPHost(oldDefaultMCPHost)
	})

	config.SetDefaultHost("public.example.com")
	config.SetDefaultMCPHost("mcp.example.com")

	require.Equal(t, "mcp.example.com", config.GetConfiguredDefaultMCPHost())
	require.Equal(t, "mcp.example.com", config.GetDefaultMCPHost())
}
