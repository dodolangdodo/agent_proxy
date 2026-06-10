package notion

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/bytedance/sonic"
	mcpservers "github.com/labring/aiproxy/mcp-servers"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Server represents the Notion MCP server
type Server struct {
	*server.MCPServer
	client                   *Client
	enableMarkdownConversion bool
}

// NewServer creates a new Notion MCP server
func NewServer(config, reusingConfig map[string]string) (mcpservers.Server, error) {
	notionToken := config["notion-api-token"]
	if notionToken == "" {
		notionToken = reusingConfig["notion-api-token"]
	}

	if notionToken == "" {
		return nil, errors.New("NOTION_API_TOKEN is required")
	}

	enabledToolsStr := config["enabled-tools"]
	enabledToolsSet := ParseEnabledTools(enabledToolsStr)

	enableMarkdownStr := config["enable-markdown"]
	enableMarkdownConversion, _ := strconv.ParseBool(enableMarkdownStr)

	// Create MCP server
	mcpServer := server.NewMCPServer("notion-mcp", "1.0.0")

	// Create Notion client
	notionClient := NewClient(notionToken)

	notionServer := &Server{
		MCPServer:                mcpServer,
		client:                   notionClient,
		enableMarkdownConversion: enableMarkdownConversion,
	}

	// Add tools
	notionServer.addTools(enabledToolsSet)

	return notionServer, nil
}

func ListTools(ctx context.Context) ([]mcp.Tool, error) {
	notionServer := &Server{
		MCPServer: server.NewMCPServer("notion-mcp", "1.0.0"),
	}
	notionServer.addTools(nil)

	return mcpservers.ListServerTools(ctx, notionServer)
}

// addTools adds all tools to the server
func (s *Server) addTools(enabledToolsSet map[string]bool) {
	allTools := []mcp.Tool{
		getAppendBlockChildrenTool(),
		getRetrieveBlockTool(),
		getRetrieveBlockChildrenTool(),
		getDeleteBlockTool(),
		getUpdateBlockTool(),
		getRetrievePageTool(),
		getUpdatePagePropertiesTool(),
		getQueryDatabaseTool(),
		getSearchTool(),
		// Add other tools as needed
	}

	tools := FilterTools(allTools, enabledToolsSet)

	for _, tool := range tools {
		s.AddTool(tool, s.createToolHandler(tool.Name))
	}
}

// createToolHandler creates a handler for a specific tool
func (s *Server) createToolHandler(
	toolName string,
) func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		var (
			response any
			err      error
		)

		switch toolName {
		case "notion_append_block_children":
			response, err = s.handleAppendBlockChildren(ctx, args)
		case "notion_retrieve_block":
			response, err = s.handleRetrieveBlock(ctx, args)
		case "notion_retrieve_block_children":
			response, err = s.handleRetrieveBlockChildren(ctx, args)
		case "notion_delete_block":
			response, err = s.handleDeleteBlock(ctx, args)
		case "notion_update_block":
			response, err = s.handleUpdateBlock(ctx, args)
		case "notion_retrieve_page":
			response, err = s.handleRetrievePage(ctx, args)
		case "notion_update_page_properties":
			response, err = s.handleUpdatePageProperties(ctx, args)
		case "notion_query_database":
			response, err = s.handleQueryDatabase(ctx, args)
		case "notion_search":
			response, err = s.handleSearch(ctx, args)
		default:
			return nil, fmt.Errorf("unknown tool: %s", toolName)
		}

		if err != nil {
			return nil, err
		}

		// Check format parameter and return appropriate response
		requestedFormat := "markdown"
		if format, ok := args["format"].(string); ok {
			requestedFormat = format
		}

		// Convert to markdown if both conditions are met
		if s.enableMarkdownConversion && requestedFormat == "markdown" {
			markdown, err := ConvertToMarkdown(response)
			if err != nil {
				return nil, fmt.Errorf("failed to convert to markdown: %w", err)
			}

			return mcp.NewToolResultText(markdown), nil
		}

		// Return JSON response
		jsonResponse, err := sonic.MarshalIndent(response, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}

		return mcp.NewToolResultText(string(jsonResponse)), nil
	}
}

// Tool handlers

func (s *Server) handleAppendBlockChildren(
	ctx context.Context,
	args map[string]any,
) (any, error) {
	blockID, ok := args["block_id"].(string)
	if !ok || blockID == "" {
		return nil, errors.New("block_id is required")
	}

	childrenData, ok := args["children"]
	if !ok {
		return nil, errors.New("children is required")
	}

	// Convert children to BlockResponse slice
	childrenJSON, err := json.Marshal(childrenData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal children: %w", err)
	}

	var children []BlockResponse
	if err := json.Unmarshal(childrenJSON, &children); err != nil {
		return nil, fmt.Errorf("failed to unmarshal children: %w", err)
	}

	return s.client.AppendBlockChildren(ctx, blockID, children)
}

func (s *Server) handleRetrieveBlock(ctx context.Context, args map[string]any) (any, error) {
	blockID, ok := args["block_id"].(string)
	if !ok || blockID == "" {
		return nil, errors.New("block_id is required")
	}

	return s.client.RetrieveBlock(ctx, blockID)
}

func (s *Server) handleRetrieveBlockChildren(
	ctx context.Context,
	args map[string]any,
) (any, error) {
	blockID, ok := args["block_id"].(string)
	if !ok || blockID == "" {
		return nil, errors.New("block_id is required")
	}

	var startCursor *string
	if sc, ok := args["start_cursor"].(string); ok && sc != "" {
		startCursor = &sc
	}

	var pageSize *int
	if ps, ok := args["page_size"].(float64); ok {
		pageSizeInt := int(ps)
		pageSize = &pageSizeInt
	}

	return s.client.RetrieveBlockChildren(ctx, blockID, startCursor, pageSize)
}

func (s *Server) handleDeleteBlock(ctx context.Context, args map[string]any) (any, error) {
	blockID, ok := args["block_id"].(string)
	if !ok || blockID == "" {
		return nil, errors.New("block_id is required")
	}

	return s.client.DeleteBlock(ctx, blockID)
}

func (s *Server) handleUpdateBlock(ctx context.Context, args map[string]any) (any, error) {
	blockID, ok := args["block_id"].(string)
	if !ok || blockID == "" {
		return nil, errors.New("block_id is required")
	}

	blockData, ok := args["block"]
	if !ok {
		return nil, errors.New("block is required")
	}

	// Convert block to BlockResponse
	blockJSON, err := json.Marshal(blockData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal block: %w", err)
	}

	var block BlockResponse
	if err := json.Unmarshal(blockJSON, &block); err != nil {
		return nil, fmt.Errorf("failed to unmarshal block: %w", err)
	}

	return s.client.UpdateBlock(ctx, blockID, block)
}

func (s *Server) handleRetrievePage(ctx context.Context, args map[string]any) (any, error) {
	pageID, ok := args["page_id"].(string)
	if !ok || pageID == "" {
		return nil, errors.New("page_id is required")
	}

	return s.client.RetrievePage(ctx, pageID)
}

func (s *Server) handleUpdatePageProperties(
	ctx context.Context,
	args map[string]any,
) (any, error) {
	pageID, ok := args["page_id"].(string)
	if !ok || pageID == "" {
		return nil, errors.New("page_id is required")
	}

	properties, ok := args["properties"].(map[string]any)
	if !ok {
		return nil, errors.New("properties is required")
	}

	return s.client.UpdatePageProperties(ctx, pageID, properties)
}

func (s *Server) handleQueryDatabase(ctx context.Context, args map[string]any) (any, error) {
	databaseID, ok := args["database_id"].(string)
	if !ok || databaseID == "" {
		return nil, errors.New("database_id is required")
	}

	var filter map[string]any
	if f, ok := args["filter"].(map[string]any); ok {
		filter = f
	}

	var sorts []SortObject
	if s, ok := args["sorts"]; ok {
		sortsJSON, err := json.Marshal(s)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal sorts: %w", err)
		}

		if err := json.Unmarshal(sortsJSON, &sorts); err != nil {
			return nil, fmt.Errorf("failed to unmarshal sorts: %w", err)
		}
	}

	var startCursor *string
	if sc, ok := args["start_cursor"].(string); ok && sc != "" {
		startCursor = &sc
	}

	var pageSize *int
	if ps, ok := args["page_size"].(float64); ok {
		pageSizeInt := int(ps)
		pageSize = &pageSizeInt
	}

	return s.client.QueryDatabase(ctx, databaseID, filter, sorts, startCursor, pageSize)
}

func (s *Server) handleSearch(ctx context.Context, args map[string]any) (any, error) {
	var query *string
	if q, ok := args["query"].(string); ok && q != "" {
		query = &q
	}

	var filter *SearchFilter
	if f, ok := args["filter"].(map[string]any); ok {
		filterJSON, err := json.Marshal(f)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal filter: %w", err)
		}

		var searchFilter SearchFilter
		if err := json.Unmarshal(filterJSON, &searchFilter); err != nil {
			return nil, fmt.Errorf("failed to unmarshal filter: %w", err)
		}

		filter = &searchFilter
	}

	var sort *SearchSort
	if s, ok := args["sort"].(map[string]any); ok {
		sortJSON, err := json.Marshal(s)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal sort: %w", err)
		}

		var searchSort SearchSort
		if err := json.Unmarshal(sortJSON, &searchSort); err != nil {
			return nil, fmt.Errorf("failed to unmarshal sort: %w", err)
		}

		sort = &searchSort
	}

	var startCursor *string
	if sc, ok := args["start_cursor"].(string); ok && sc != "" {
		startCursor = &sc
	}

	var pageSize *int
	if ps, ok := args["page_size"].(float64); ok {
		pageSizeInt := int(ps)
		pageSize = &pageSizeInt
	}

	return s.client.Search(ctx, query, filter, sort, startCursor, pageSize)
}
