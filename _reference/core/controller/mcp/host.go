package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/common"
	"github.com/labring/aiproxy/core/common/config"
	"github.com/labring/aiproxy/core/middleware"
	"github.com/labring/aiproxy/core/model"
	mcpservers "github.com/labring/aiproxy/mcp-servers"
	"github.com/mark3labs/mcp-go/mcp"
)

func routeHostMCP(
	c *gin.Context,
	publicHandler, groupHandler func(c *gin.Context, mcpID string),
) {
	log := common.GetLogger(c)
	host := c.Request.Host

	log.Debugf("route host mcp: %s", host)

	publicMCPHost := config.GetPublicMCPHost()
	groupMCPHost := config.GetGroupMCPHost()

	switch {
	case publicMCPHost != "" && strings.HasSuffix(host, publicMCPHost):
		mcpID := strings.TrimSuffix(host, "."+publicMCPHost)
		publicHandler(c, mcpID)
	case groupMCPHost != "" && strings.HasSuffix(host, groupMCPHost):
		mcpID := strings.TrimSuffix(host, "."+groupMCPHost)
		groupHandler(c, mcpID)
	default:
		http.Error(c.Writer, "invalid host", http.StatusNotFound)
	}
}

// HostMCPSSEServer godoc
//
//	@Summary	Public MCP SSE Server
//	@Security	ApiKeyAuth
//	@Router		/sse [get]
func HostMCPSSEServer(c *gin.Context) {
	routeHostMCP(c, func(c *gin.Context, mcpID string) {
		publicMcp, err := model.CacheGetPublicMCP(mcpID)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusBadRequest)
			return
		}

		if publicMcp.Status != model.PublicMCPStatusEnabled {
			http.Error(c.Writer, "mcp is not enabled", http.StatusBadRequest)
			return
		}

		group := middleware.GetGroup(c)
		paramsFunc := newGroupParams(publicMcp.ID, group.ID)

		handlePublicSSEMCP(c, publicMcp, paramsFunc, sseEndpoint)
	}, func(c *gin.Context, mcpID string) {
		group := middleware.GetGroup(c)

		groupMcp, err := model.CacheGetGroupMCP(group.ID, mcpID)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusNotFound)
			return
		}

		if groupMcp.Status != model.GroupMCPStatusEnabled {
			http.Error(c.Writer, "mcp is not enabled", http.StatusNotFound)
			return
		}

		handleGroupSSEMCPServer(c, groupMcp, sseEndpoint)
	})
}

// HostMCPStreamable godoc
//
//	@Summary	Host MCP Streamable Server
//	@Security	ApiKeyAuth
//	@Router		/mcp [get]
//	@Router		/mcp [post]
//	@Router		/mcp [delete]
func HostMCPStreamable(c *gin.Context) {
	routeHostMCP(c, func(c *gin.Context, mcpID string) {
		publicMcp, err := model.CacheGetPublicMCP(mcpID)
		if err != nil {
			c.JSON(http.StatusBadRequest, mcpservers.CreateMCPErrorResponse(
				mcp.NewRequestId(nil),
				mcp.INVALID_REQUEST,
				err.Error(),
			))

			return
		}

		if publicMcp.Status != model.PublicMCPStatusEnabled {
			c.JSON(http.StatusNotFound, mcpservers.CreateMCPErrorResponse(
				mcp.NewRequestId(nil),
				mcp.INVALID_REQUEST,
				"mcp is not enabled",
			))

			return
		}

		group := middleware.GetGroup(c)
		paramsFunc := newGroupParams(publicMcp.ID, group.ID)

		handlePublicStreamable(c, publicMcp, paramsFunc)
	}, func(c *gin.Context, mcpID string) {
		group := middleware.GetGroup(c)

		groupMcp, err := model.CacheGetGroupMCP(group.ID, mcpID)
		if err != nil {
			c.JSON(http.StatusNotFound, mcpservers.CreateMCPErrorResponse(
				mcp.NewRequestId(nil),
				mcp.INVALID_REQUEST,
				err.Error(),
			))

			return
		}

		if groupMcp.Status != model.GroupMCPStatusEnabled {
			c.JSON(http.StatusNotFound, mcpservers.CreateMCPErrorResponse(
				mcp.NewRequestId(nil),
				mcp.INVALID_REQUEST,
				"mcp is not enabled",
			))

			return
		}

		handleGroupStreamable(c, groupMcp)
	})
}
