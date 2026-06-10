# Figma MCP Server

> <https://github.com/GLips/Figma-Context-MCP>

A Model Context Protocol server that provides access to Figma design files. This server enables LLMs to retrieve and process Figma design data, converting it to simplified formats for easier consumption.

## Features

- **Design File Access**: Retrieve complete Figma design files or specific nodes
- **Simplified Data Format**: Convert complex Figma data to simplified, AI-friendly formats
- **Image Download**: Download SVG and PNG images from Figma designs
- **Flexible Authentication**: Support for both Personal Access Tokens and OAuth
- **Multiple Output Formats**: Support for both YAML and JSON output formats

## Setup

### Prerequisites

1. Create a Figma Personal Access Token at <https://www.figma.com/developers/api#access-tokens>
2. Or set up OAuth authentication for your application

### Configuration

The server requires the following configuration:

- `figma-api-key` (required if not using OAuth): Your Figma Personal Access Token
- `figma-oauth-token` (optional): Your Figma OAuth Bearer token (takes precedence over API key)
- `output-format` (optional): Output format for design data (yaml or json, default: yaml)

## Available Tools

### get_figma_data

Retrieve layout information about a Figma file or specific node.

**Parameters:**

- `fileKey` (required): The key of the Figma file to fetch
- `nodeId` (optional): The ID of a specific node to fetch
- `depth` (optional): How many levels deep to traverse the node tree

### download_figma_images

Download SVG and PNG images from Figma designs.

**Parameters:**

- `fileKey` (required): The key of the Figma file containing the nodes
- `nodes` (required): Array of nodes to fetch as images
- `localPath` (required): Directory path where images will be saved
- `pngScale` (optional): Export scale for PNG images (default: 2)
- `svgOptions` (optional): Options for SVG export

## Authentication

The server supports two authentication methods:

1. **Personal Access Token**: Set the `figma-api-key` configuration
2. **OAuth Bearer Token**: Set the `figma-oauth-token` configuration (takes precedence)

## Output Formats

The server can output data in two formats:

- **YAML** (default): Human-readable format
- **JSON**: Machine-readable format

Set the `output-format` configuration to choose your preferred format.
