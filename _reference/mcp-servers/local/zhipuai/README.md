# Zhipu Web Search MCP Server

> <https://github.com/THUDM>

## What is Zhipu Web Search MCP Service?

Zhipu Web Search MCP Server is a search engine launched by Zhipu Open Platform (BigModel.cn), specifically designed for large models. It integrates five search engines, allowing users to flexibly compare and switch between them. Building upon the web crawling and ranking capabilities of traditional search engines, it enhances intent recognition capabilities, returning results more suitable for large model processing (such as webpage titles, URLs, summaries, site names, site icons, etc.), helping AI applications achieve "dynamic knowledge acquisition" and "precise scenario adaptation" capabilities.

## How to use Zhipu Web Search MCP?

It can be configured in clients that support the MCP protocol, such as Cursor, Cherry Studio, etc. Copy your API key from the Zhipu BigModel Open Platform, and set up the server command according to the documentation.

## Key Features of Zhipu Web Search MCP Service

**Real-time Web Search**: Retrieves real-time information and webpage links from across the internet.

**Domain-Specific Search Support**: Users can search for content within a specified domain by entering the domain name. This enhances search efficiency and result relevance, catering to personalized needs such as professional research, brand monitoring, and security management.

**Flexible Result Quantity Customization**: Users can freely set the number of search results from 1 to 50 based on their requirements. This precise control over information quantity helps avoid information redundancy or insufficiency.

**Accurate Time Filtering**: Users can filter websites published within a day, a week, a month, a year, or without any time limit. This assists in accurately retrieving the latest information and historical data, meeting users' multi-condition search needs.

**Independent Summary Adjustment**: We offer two modes for generating website summaries. The medium mode provides a summary of approximately 400-600 words, while the newly added high mode can generate up to 2500 words of extended context summary, allowing the model to produce more comprehensive answers. Users can adjust the mode according to their needs.

**Intent-Enhanced Search**: Combines Zhipu's proprietary models for vectors, semantic matching, timeliness, content quality, etc., to perform intent recognition on user queries and optimize the extraction of query search terms. It also utilizes a proprietary semantic reranker to analyze the relevance between the user's question and the given search results, assigns scores and ranks them based on timeliness, providing users with accurate and reliable search results.

**Multi-Engine Support**: Integrates Zhipu's proprietary engine and mainstream search engines (Sogou/Quark/Jina.ai). Developers can flexibly invoke them according to scenario-specific needs, leveraging the advantages of different search tools.

| Name | Description |
|------|-------------|
| Zhipu Proprietary Search Basic Edition | Provides basic search capabilities, excellent value for money. |
| Zhipu Proprietary Search Pro Edition | Supports ultra-long full text, more comprehensive search results; multiple engines reduce empty result rates and improve search efficiency; higher search result recall rate, more accurate answers. |
| Sogou Search | Comprehensive content; can crawl content from the Tencent ecosystem (Tencent News, Penguin Accounts) and Zhihu; higher answer authoritativeness in encyclopedia and medical vertical scenarios. |
| Quark Search | Supports searching within a specified scope, e.g., can specify searches within the finance and legal industries, improving recall accuracy; full-text output, provides a 3000-character long full-text service covering 95% of webpages, no need to re-parse URLs. |
| Jina AI Search | Output results are direct and concise, suitable for scenarios requiring clear and direct answers; can accurately parse complex HTML and convert it into clean Markdown or JSON format. |
| Bing Search | Possessing a vast index that encompasses billions of web pages and various types of content; featuring a rich array of search functions, including searches for web pages, images, videos, news, shopping, academic resources, maps, and more. |

**Structured output**: Returns a data format suitable for LLM processing (including title/URL/summary/site name/icon, etc.)

## Use Cases for Zhipu Web Search MCP Service

**Professional Vertical Research**: Conduct high-quality webpage retrieval and multi-source information integration, such as academic research, legal, and financial industry analysis reports.

**Business Intelligence Radar**: Real-time data tracking and analysis, such as monitoring industry dynamics, competitor information, market trends, etc.

**AI Assistant/Chatbot**: Provides real-time information search capabilities to ensure the accuracy and timeliness of answers.

**Consumer Decision-making and Planning Design**: Compare multiple options and find the optimal solution based on the latest information such as weather, news, and tickets.

**Talent Profiling and Resume Optimization**: For HR, search and compare job descriptions (JDs) in the market to help generate or optimize JDs; for job seekers, search and compare resumes with target job JDs and templates of successful resumes in related industries to help generate/optimize resumes.

## Frequently Asked Questions

**Q: How to obtain an API Key?**
A: You need to register a developer account on the [Zhipu BigModel Open Platform](https://zhipuaishengchan.datasink.sensorsdata.cn/t/ME) to obtain an API Key.

**Q: Is there a fee for using Zhipu Web Search MCP?**
A: You can receive a free quota on the Zhipu BigModel Open Platform. If the free quota is exhausted, you will need to pay for usage. If you have any questions, you can consult customer service on the Zhipu BigModel Open Platform.

**Q: Which search engines are supported?**
A: It supports Zhipu's proprietary engines and mainstream search engines. Zhipu proprietary engines: search_std (Basic Edition), search_pro (Advanced Edition). Third-party engines: search_pro_sogou (Sogou), search_pro_quark (Quark), search_pro_jina (Jina.ai), search_pro_bing (Bing).

## Installation Guide

Supports clients that run the MCP protocol, such as Cursor, Cherry Studio, etc.
Click to obtain the API Key from the [Zhipu BigModel Open Platform](https://zhipuaishengchan.datasink.sensorsdata.cn/t/ME).

**Using in Cursor**

Cursor version 0.45.6 provides MCP functionality. Cursor will act as an MCP service client to use the MCP service. The MCP service can be integrated into Cursor through simple configuration.

Operation Path: Cursor Settings -> 【Features】-> 【MCP Servers】.

**Configure MCP Server**

```json
{
    "mcpServers": {
        "zhipu-web-search-sse": {
            "url": "https://open.bigmodel.cn/api/mcp/web_search/sse?Authorization= YOUR API Key"
        }
    }
}
```
