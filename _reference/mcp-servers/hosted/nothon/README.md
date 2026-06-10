# Notion MCP Server

> <https://github.com/suekou/mcp-notion-server>

A Model Context Protocol server that provides comprehensive access to Notion workspaces. This server enables LLMs to interact with Notion pages, databases, blocks, and users through the official Notion API.

## Features

- **Page Management**: Create, read, update, and delete Notion pages
- **Database Operations**: Query databases, create database items, and manage database properties
- **Block Manipulation**: Add, retrieve, update, and delete blocks within pages
- **Search Functionality**: Search across pages and databases
- **User Management**: Retrieve user information and workspace details
- **Comment System**: Create and retrieve comments on pages and blocks
- **Rich Text Support**: Full support for Notion's rich text formatting
- **Markdown Conversion**: Optional conversion of responses to readable Markdown format

## Setup

### Prerequisites

1. Create a Notion integration at <https://www.notion.so/profile/integrations>
2. Click "New Integration".
3. Name your integration and select appropriate permissions (e.g., "Read content", "Update content").
4. Copy the "Internal Integration Token" (starts with `secret_`)
5. Share your Notion pages/databases with your integration

### Configuration

The server requires the following configuration:

- `notion-api-token` (required): Your Notion API integration token
- `enabled-tools` (optional): Comma-separated list of specific tools to enable
- `enable-markdown` (optional): Enable experimental Markdown conversion for responses

## Available Tools

### Blocks

- `notion_append_block_children` - Add new blocks to a page or block
- `notion_retrieve_block` - Get a specific block by ID
- `notion_retrieve_block_children` - Get child blocks of a block
- `notion_update_block` - Update block content
- `notion_delete_block` - Delete a block

### Pages

- `notion_retrieve_page` - Get page content and properties
- `notion_update_page_properties` - Update page properties

### Databases

- `notion_query_database` - Query database with filters and sorting
- `notion_retrieve_database` - Get database schema and properties
- `notion_create_database_item` - Create new database entries

### Search & Users

- `notion_search` - Search pages and databases
- `notion_list_all_users` - List workspace users (requires Enterprise plan)
- `notion_retrieve_user` - Get specific user details

### Comments

- `notion_create_comment` - Add comments to pages
- `notion_retrieve_comments` - Get comments from pages

## Response Formats

The server supports two response formats:

- **JSON** (default): Raw Notion API responses for programmatic use
- **Markdown**: Human-readable format for content consumption

Use the `format` parameter in tool calls to specify your preferred format.

## Security & Permissions

This server requires appropriate Notion integration permissions:

- Read content
- Update content  
- Insert content
- Read comments
- Insert comments
- Read user information (for user-related tools)

## Error Handling

The server provides detailed error messages for:

- Missing or invalid API tokens
- Insufficient permissions
- Invalid block/page/database IDs
- Malformed requests
- API rate limits

## Rate Limits

Notion API has rate limits. The server will return appropriate error messages when limits are exceeded. Consider implementing retry logic in your applications.
