package notion

import (
	mcpservers "github.com/labring/aiproxy/mcp-servers"
	"github.com/mark3labs/mcp-go/mcp"
)

// Configuration templates
var configTemplates = mcpservers.ConfigTemplates{
	"notion-api-token": {
		Name:        "Notion API Token",
		Required:    mcpservers.ConfigRequiredTypeInitOrReusingOnly,
		Example:     "ntn_xxxxxxxxx",
		Description: "Your Notion API integration token",
	},
	"enabled-tools": {
		Name:        "Enabled Tools",
		Required:    mcpservers.ConfigRequiredTypeInitOptional,
		Example:     "notion_retrieve_page,notion_query_database",
		Description: "Comma-separated list of tools to enable",
	},
	"enable-markdown": {
		Name:        "Enable Markdown Conversion",
		Required:    mcpservers.ConfigRequiredTypeInitOptional,
		Example:     "true",
		Description: "Enable experimental Markdown conversion for responses",
	},
}

// Common schema definitions
const commonIDDescription = "It should be a 32-character string (excluding hyphens) formatted as 8-4-4-4-12 with hyphens (-)."

var formatParameter = map[string]any{
	"type":        "string",
	"enum":        []string{"json", "markdown"},
	"description": "Specify the response format. 'json' returns the original data structure, 'markdown' returns a more readable format. Use 'markdown' when the user only needs to read the page and isn't planning to write or modify it. Use 'json' when the user needs to read the page with the intention of writing to or modifying it.",
	"default":     "markdown",
}

// Tool schemas
func getAppendBlockChildrenTool() mcp.Tool {
	return mcp.Tool{
		Name:        "notion_append_block_children",
		Description: "Append new children blocks to a specified parent block in Notion. Requires insert content capabilities. You can optionally specify the 'after' parameter to append after a certain block.",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]any{
				"block_id": map[string]any{
					"type":        "string",
					"description": "The ID of the parent block." + commonIDDescription,
				},
				"children": map[string]any{
					"type":        "array",
					"description": "Array of block objects to append. Each block must follow the Notion block schema.",
					"items":       getBlockObjectSchema(),
				},
				"after": map[string]any{
					"type":        "string",
					"description": "The ID of the existing block that the new block should be appended after." + commonIDDescription,
				},
				"format": formatParameter,
			},
			Required: []string{"block_id", "children"},
		},
	}
}

func getRetrieveBlockTool() mcp.Tool {
	return mcp.Tool{
		Name:        "notion_retrieve_block",
		Description: "Retrieve a block from Notion",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]any{
				"block_id": map[string]any{
					"type":        "string",
					"description": "The ID of the block to retrieve." + commonIDDescription,
				},
				"format": formatParameter,
			},
			Required: []string{"block_id"},
		},
	}
}

func getRetrieveBlockChildrenTool() mcp.Tool {
	return mcp.Tool{
		Name:        "notion_retrieve_block_children",
		Description: "Retrieve the children of a block",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]any{
				"block_id": map[string]any{
					"type":        "string",
					"description": "The ID of the block." + commonIDDescription,
				},
				"start_cursor": map[string]any{
					"type":        "string",
					"description": "Pagination cursor for next page of results",
				},
				"page_size": map[string]any{
					"type":        "number",
					"description": "Number of results per page (max 100)",
				},
				"format": formatParameter,
			},
			Required: []string{"block_id"},
		},
	}
}

func getDeleteBlockTool() mcp.Tool {
	return mcp.Tool{
		Name:        "notion_delete_block",
		Description: "Delete a block in Notion",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]any{
				"block_id": map[string]any{
					"type":        "string",
					"description": "The ID of the block to delete." + commonIDDescription,
				},
				"format": formatParameter,
			},
			Required: []string{"block_id"},
		},
	}
}

func getUpdateBlockTool() mcp.Tool {
	return mcp.Tool{
		Name:        "notion_update_block",
		Description: "Update the content of a block in Notion based on its type. The update replaces the entire value for a given field.",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]any{
				"block_id": map[string]any{
					"type":        "string",
					"description": "The ID of the block to update." + commonIDDescription,
				},
				"block": map[string]any{
					"type":        "object",
					"description": "The updated content for the block. Must match the block's type schema.",
				},
				"format": formatParameter,
			},
			Required: []string{"block_id", "block"},
		},
	}
}

func getRetrievePageTool() mcp.Tool {
	return mcp.Tool{
		Name:        "notion_retrieve_page",
		Description: "Retrieve a page from Notion",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]any{
				"page_id": map[string]any{
					"type":        "string",
					"description": "The ID of the page to retrieve." + commonIDDescription,
				},
				"format": formatParameter,
			},
			Required: []string{"page_id"},
		},
	}
}

func getUpdatePagePropertiesTool() mcp.Tool {
	return mcp.Tool{
		Name:        "notion_update_page_properties",
		Description: "Update properties of a page or an item in a Notion database",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]any{
				"page_id": map[string]any{
					"type":        "string",
					"description": "The ID of the page or database item to update." + commonIDDescription,
				},
				"properties": map[string]any{
					"type":        "object",
					"description": "Properties to update. These correspond to the columns or fields in the database.",
				},
				"format": formatParameter,
			},
			Required: []string{"page_id", "properties"},
		},
	}
}

func getQueryDatabaseTool() mcp.Tool {
	return mcp.Tool{
		Name:        "notion_query_database",
		Description: "Query a database in Notion",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]any{
				"database_id": map[string]any{
					"type":        "string",
					"description": "The ID of the database to query." + commonIDDescription,
				},
				"filter": map[string]any{
					"type":        "object",
					"description": "Filter conditions",
				},
				"sorts": map[string]any{
					"type":        "array",
					"description": "Sort conditions",
					"items": map[string]any{
						"type": "object",
						"properties": map[string]any{
							"property":  map[string]any{"type": "string"},
							"timestamp": map[string]any{"type": "string"},
							"direction": map[string]any{
								"type": "string",
								"enum": []string{"ascending", "descending"},
							},
						},
						"required": []string{"direction"},
					},
				},
				"start_cursor": map[string]any{
					"type":        "string",
					"description": "Pagination cursor for next page of results",
				},
				"page_size": map[string]any{
					"type":        "number",
					"description": "Number of results per page (max 100)",
				},
				"format": formatParameter,
			},
			Required: []string{"database_id"},
		},
	}
}

func getSearchTool() mcp.Tool {
	return mcp.Tool{
		Name:        "notion_search",
		Description: "Search pages or databases by title in Notion",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]any{
				"query": map[string]any{
					"type":        "string",
					"description": "Text to search for in page or database titles",
				},
				"filter": map[string]any{
					"type":        "object",
					"description": "Filter results by object type (page or database)",
					"properties": map[string]any{
						"property": map[string]any{
							"type":        "string",
							"description": "Must be 'object'",
						},
						"value": map[string]any{
							"type":        "string",
							"description": "Either 'page' or 'database'",
						},
					},
				},
				"sort": map[string]any{
					"type":        "object",
					"description": "Sort order of results",
					"properties": map[string]any{
						"direction": map[string]any{
							"type": "string",
							"enum": []string{"ascending", "descending"},
						},
						"timestamp": map[string]any{
							"type": "string",
							"enum": []string{"last_edited_time"},
						},
					},
				},
				"start_cursor": map[string]any{
					"type":        "string",
					"description": "Pagination start cursor",
				},
				"page_size": map[string]any{
					"type":        "number",
					"description": "Number of results to return (max 100)",
				},
				"format": formatParameter,
			},
		},
	}
}

// Helper function to get block object schema
func getBlockObjectSchema() map[string]any {
	return map[string]any{
		"type":        "object",
		"description": "A Notion block object.",
		"properties": map[string]any{
			"object": map[string]any{
				"type":        "string",
				"description": "Should be 'block'.",
				"enum":        []string{"block"},
			},
			"type": map[string]any{
				"type":        "string",
				"description": "Type of the block. Possible values include 'paragraph', 'heading_1', 'heading_2', 'heading_3', 'bulleted_list_item', 'numbered_list_item', 'to_do', 'toggle', etc.",
			},
			"paragraph": map[string]any{
				"type":        "object",
				"description": "Paragraph block object.",
				"properties": map[string]any{
					"rich_text": map[string]any{
						"type":        "array",
						"description": "Array of rich text objects representing the content.",
						"items":       getRichTextObjectSchema(),
					},
				},
			},
		},
		"required": []string{"object", "type"},
	}
}

// Helper function to get rich text object schema
func getRichTextObjectSchema() map[string]any {
	return map[string]any{
		"type":        "object",
		"description": "A rich text object.",
		"properties": map[string]any{
			"type": map[string]any{
				"type":        "string",
				"description": "The type of this rich text object. Possible values: text, mention, equation.",
				"enum":        []string{"text", "mention", "equation"},
			},
			"text": map[string]any{
				"type":        "object",
				"description": "Object containing text content and optional link info. Required if type is 'text'.",
				"properties": map[string]any{
					"content": map[string]any{
						"type":        "string",
						"description": "The actual text content.",
					},
					"link": map[string]any{
						"type":        "object",
						"description": "Optional link object with a 'url' field.",
						"properties": map[string]any{
							"url": map[string]any{
								"type":        "string",
								"description": "The URL the text links to.",
							},
						},
					},
				},
			},
			"plain_text": map[string]any{
				"type":        "string",
				"description": "The plain text without annotations.",
			},
		},
		"required": []string{"type"},
	}
}
