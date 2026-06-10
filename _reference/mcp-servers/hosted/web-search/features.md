# Web Search MCP Server

This MCP server provides web search capabilities through multiple search engines including Google, Bing, Bing CN(Free), and Arxiv.

## Features

- **Multiple Search Engines**: Support for Google, Bing, Bing CN(Free), and Arxiv
- **Multi-Engine Search**: Search across multiple engines simultaneously
- **Smart Search**: Intelligent query optimization and result aggregation
- **Academic Search**: Specialized support for academic papers via Arxiv
- **Language Support**: Search in different languages
- **Configurable Results**: Control the number of results returned

## Configuration

Configure the search engines you want to use by providing their API keys:

- **Google**: Requires both API key and Custom Search Engine ID
- **Bing**: Requires Bing Search API key
- **Bing CN**: Free, no API key required
- **Arxiv**: No API key required (free to use)
- **SearchXNG**: No API key required (free to use)
- **SearchXNG Base URL**: Base URL for SearchXNG

## Available Tools

### web_search

Basic web search using a single search engine.

### multi_search

Search across multiple engines simultaneously for comprehensive results.

### smart_search

Intelligently optimize queries and aggregate results for better answers.

## Usage Examples

1. Basic search:
   - Query: "latest AI developments"
   - Engine: "google"
   - Max results: 10

2. Academic search:
   - Query: "transformer architecture"
   - Engine: "arxiv"
   - Category: "cs.AI"

3. Multi-engine search:
   - Query: "climate change impacts"
   - Engines: ["google", "bing", "arxiv", "searchxng", "bingcn"]
   - Max results per engine: 5
