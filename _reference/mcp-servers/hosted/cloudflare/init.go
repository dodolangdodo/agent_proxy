package cloudflare

import (
	_ "embed"

	"github.com/labring/aiproxy/core/model"
	mcpservers "github.com/labring/aiproxy/mcp-servers"
)

//go:embed README.md
var readme string

//go:embed README.cn.md
var readmeCN string

// need import in mcpregister/init.go
func init() {
	mcpservers.Register(
		mcpservers.NewMcp(
			"cloudflare-docs",
			"Cloudflare Docs",
			model.PublicMCPTypeProxySSE,
			mcpservers.WithNameCN("Cloudflare 文档"),
			mcpservers.WithProxyConfigTemplates(mcpservers.ProxyConfigTemplates{
				"url": {
					ConfigTemplate: mcpservers.ConfigTemplate{
						Name:        "URL",
						Required:    mcpservers.ConfigRequiredTypeInitOptional,
						Default:     "https://docs.mcp.cloudflare.com/sse",
						Description: "The Streamable http URL of the Cloudflare Docs MCP server",
					},
					Type: model.ParamTypeURL,
				},
			}),
			mcpservers.WithTags([]string{"cloudflare", "documentation"}),
			mcpservers.WithGitHubURL(
				"https://github.com/cloudflare/mcp-server-cloudflare",
			),
			mcpservers.WithDescription(
				"Get up to date reference information on Cloudflare.",
			),
			mcpservers.WithDescriptionCN(
				"获取 Cloudflare 最新参考信息。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)

	mcpservers.Register(
		mcpservers.NewMcp(
			"cloudflare-bindings",
			"Cloudflare Bindings",
			model.PublicMCPTypeProxySSE,
			mcpservers.WithNameCN("Cloudflare Bindings"),
			mcpservers.WithProxyConfigTemplates(mcpservers.ProxyConfigTemplates{
				"url": {
					ConfigTemplate: mcpservers.ConfigTemplate{
						Name:        "URL",
						Required:    mcpservers.ConfigRequiredTypeInitOptional,
						Default:     "https://bindings.mcp.cloudflare.com/sse",
						Description: "The Streamable http URL of the Cloudflare Bindings MCP server",
					},
					Type: model.ParamTypeURL,
				},
			}),
			mcpservers.WithTags([]string{"cloudflare", "workers", "development"}),
			mcpservers.WithGitHubURL(
				"https://github.com/cloudflare/mcp-server-cloudflare",
			),
			mcpservers.WithDescription(
				"Build Workers applications with storage, AI, and compute primitives.",
			),
			mcpservers.WithDescriptionCN(
				"使用存储、AI 和计算原语构建 Workers 应用。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)

	mcpservers.Register(
		mcpservers.NewMcp(
			"cloudflare-builds",
			"Cloudflare Builds",
			model.PublicMCPTypeProxySSE,
			mcpservers.WithNameCN("Cloudflare 构建"),
			mcpservers.WithProxyConfigTemplates(mcpservers.ProxyConfigTemplates{
				"url": {
					ConfigTemplate: mcpservers.ConfigTemplate{
						Name:        "URL",
						Required:    mcpservers.ConfigRequiredTypeInitOptional,
						Default:     "https://builds.mcp.cloudflare.com/sse",
						Description: "The Streamable http URL of the Cloudflare Builds MCP server",
					},
					Type: model.ParamTypeURL,
				},
			}),
			mcpservers.WithTags([]string{"cloudflare", "workers", "builds"}),
			mcpservers.WithGitHubURL(
				"https://github.com/cloudflare/mcp-server-cloudflare",
			),
			mcpservers.WithDescription(
				"Get insights and manage your Cloudflare Workers Builds.",
			),
			mcpservers.WithDescriptionCN(
				"获取洞察并管理您的 Cloudflare Workers 构建。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)

	mcpservers.Register(
		mcpservers.NewMcp(
			"cloudflare-observability",
			"Cloudflare Observability",
			model.PublicMCPTypeProxySSE,
			mcpservers.WithNameCN("Cloudflare 可观测性"),
			mcpservers.WithProxyConfigTemplates(mcpservers.ProxyConfigTemplates{
				"url": {
					ConfigTemplate: mcpservers.ConfigTemplate{
						Name:        "URL",
						Required:    mcpservers.ConfigRequiredTypeInitOptional,
						Default:     "https://observability.mcp.cloudflare.com/sse",
						Description: "The Streamable http URL of the Cloudflare Observability MCP server",
					},
					Type: model.ParamTypeURL,
				},
			}),
			mcpservers.WithTags([]string{"cloudflare", "observability", "monitoring"}),
			mcpservers.WithGitHubURL(
				"https://github.com/cloudflare/mcp-server-cloudflare",
			),
			mcpservers.WithDescription(
				"Debug and get insight into your application's logs and analytics.",
			),
			mcpservers.WithDescriptionCN(
				"调试并深入了解您应用程序的日志和分析。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)

	mcpservers.Register(
		mcpservers.NewMcp(
			"cloudflare-radar",
			"Cloudflare Radar",
			model.PublicMCPTypeProxySSE,
			mcpservers.WithNameCN("Cloudflare Radar"),
			mcpservers.WithProxyConfigTemplates(mcpservers.ProxyConfigTemplates{
				"url": {
					ConfigTemplate: mcpservers.ConfigTemplate{
						Name:        "URL",
						Required:    mcpservers.ConfigRequiredTypeInitOptional,
						Default:     "https://radar.mcp.cloudflare.com/sse",
						Description: "The Streamable http URL of the Cloudflare Radar MCP server",
					},
					Type: model.ParamTypeURL,
				},
			}),
			mcpservers.WithTags([]string{"cloudflare", "radar", "analytics"}),
			mcpservers.WithGitHubURL(
				"https://github.com/cloudflare/mcp-server-cloudflare",
			),
			mcpservers.WithDescription(
				"Get global Internet traffic insights, trends, URL scans, and other utilities.",
			),
			mcpservers.WithDescriptionCN(
				"获取全球互联网流量洞察、趋势、URL 扫描和其他实用工具。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)

	mcpservers.Register(
		mcpservers.NewMcp(
			"cloudflare-container",
			"Cloudflare Container",
			model.PublicMCPTypeProxySSE,
			mcpservers.WithNameCN("Cloudflare 容器"),
			mcpservers.WithProxyConfigTemplates(mcpservers.ProxyConfigTemplates{
				"url": {
					ConfigTemplate: mcpservers.ConfigTemplate{
						Name:        "URL",
						Required:    mcpservers.ConfigRequiredTypeInitOptional,
						Default:     "https://containers.mcp.cloudflare.com/sse",
						Description: "The Streamable http URL of the Cloudflare Container MCP server",
					},
					Type: model.ParamTypeURL,
				},
			}),
			mcpservers.WithTags([]string{"cloudflare", "container", "sandbox"}),
			mcpservers.WithGitHubURL(
				"https://github.com/cloudflare/mcp-server-cloudflare",
			),
			mcpservers.WithDescription(
				"Spin up a sandbox development environment.",
			),
			mcpservers.WithDescriptionCN(
				"启动沙盒开发环境。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)

	mcpservers.Register(
		mcpservers.NewMcp(
			"cloudflare-browser",
			"Cloudflare Browser",
			model.PublicMCPTypeProxySSE,
			mcpservers.WithNameCN("Cloudflare 浏览器"),
			mcpservers.WithProxyConfigTemplates(mcpservers.ProxyConfigTemplates{
				"url": {
					ConfigTemplate: mcpservers.ConfigTemplate{
						Name:        "URL",
						Required:    mcpservers.ConfigRequiredTypeInitOptional,
						Default:     "https://browser.mcp.cloudflare.com/sse",
						Description: "The Streamable http URL of the Cloudflare Browser Rendering MCP server",
					},
					Type: model.ParamTypeURL,
				},
			}),
			mcpservers.WithTags([]string{"cloudflare", "browser", "rendering"}),
			mcpservers.WithGitHubURL(
				"https://github.com/cloudflare/mcp-server-cloudflare",
			),
			mcpservers.WithDescription(
				"Fetch web pages, convert them to markdown and take screenshots.",
			),
			mcpservers.WithDescriptionCN(
				"获取网页、将其转换为 markdown 并截图。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)

	mcpservers.Register(
		mcpservers.NewMcp(
			"cloudflare-logpush",
			"Cloudflare Logpush",
			model.PublicMCPTypeProxySSE,
			mcpservers.WithNameCN("Cloudflare Logpush"),
			mcpservers.WithProxyConfigTemplates(mcpservers.ProxyConfigTemplates{
				"url": {
					ConfigTemplate: mcpservers.ConfigTemplate{
						Name:        "URL",
						Required:    mcpservers.ConfigRequiredTypeInitOptional,
						Default:     "https://logs.mcp.cloudflare.com/sse",
						Description: "The Streamable http URL of the Cloudflare Logpush MCP server",
					},
					Type: model.ParamTypeURL,
				},
			}),
			mcpservers.WithTags([]string{"cloudflare", "logpush", "logs"}),
			mcpservers.WithGitHubURL(
				"https://github.com/cloudflare/mcp-server-cloudflare",
			),
			mcpservers.WithDescription(
				"Get quick summaries for Logpush job health.",
			),
			mcpservers.WithDescriptionCN(
				"获取 Logpush 作业健康状况的快速摘要。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)

	mcpservers.Register(
		mcpservers.NewMcp(
			"cloudflare-ai-gateway",
			"Cloudflare AI Gateway",
			model.PublicMCPTypeProxySSE,
			mcpservers.WithNameCN("Cloudflare AI Gateway"),
			mcpservers.WithProxyConfigTemplates(mcpservers.ProxyConfigTemplates{
				"url": {
					ConfigTemplate: mcpservers.ConfigTemplate{
						Name:        "URL",
						Required:    mcpservers.ConfigRequiredTypeInitOptional,
						Default:     "https://ai-gateway.mcp.cloudflare.com/sse",
						Description: "The Streamable http URL of the Cloudflare AI Gateway MCP server",
					},
					Type: model.ParamTypeURL,
				},
			}),
			mcpservers.WithTags([]string{"cloudflare", "ai", "gateway"}),
			mcpservers.WithGitHubURL(
				"https://github.com/cloudflare/mcp-server-cloudflare",
			),
			mcpservers.WithDescription(
				"Search your logs, get details about the prompts and responses.",
			),
			mcpservers.WithDescriptionCN(
				"搜索您的日志，获取提示和响应的详细信息。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)

	mcpservers.Register(
		mcpservers.NewMcp(
			"cloudflare-autorag",
			"Cloudflare AutoRAG",
			model.PublicMCPTypeProxySSE,
			mcpservers.WithNameCN("Cloudflare AutoRAG"),
			mcpservers.WithProxyConfigTemplates(mcpservers.ProxyConfigTemplates{
				"url": {
					ConfigTemplate: mcpservers.ConfigTemplate{
						Name:        "URL",
						Required:    mcpservers.ConfigRequiredTypeInitOptional,
						Default:     "https://autorag.mcp.cloudflare.com/sse",
						Description: "The Streamable http URL of the Cloudflare AutoRAG MCP server",
					},
					Type: model.ParamTypeURL,
				},
			}),
			mcpservers.WithTags([]string{"cloudflare", "autorag", "ai"}),
			mcpservers.WithGitHubURL(
				"https://github.com/cloudflare/mcp-server-cloudflare",
			),
			mcpservers.WithDescription(
				"List and search documents on your AutoRAGs.",
			),
			mcpservers.WithDescriptionCN(
				"列出并搜索您的 AutoRAG 上的文档。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)

	mcpservers.Register(
		mcpservers.NewMcp(
			"cloudflare-auditlogs",
			"Cloudflare Audit Logs",
			model.PublicMCPTypeProxySSE,
			mcpservers.WithNameCN("Cloudflare 审计日志"),
			mcpservers.WithProxyConfigTemplates(mcpservers.ProxyConfigTemplates{
				"url": {
					ConfigTemplate: mcpservers.ConfigTemplate{
						Name:        "URL",
						Required:    mcpservers.ConfigRequiredTypeInitOptional,
						Default:     "https://auditlogs.mcp.cloudflare.com/sse",
						Description: "The Streamable http URL of the Cloudflare Audit Logs MCP server",
					},
					Type: model.ParamTypeURL,
				},
			}),
			mcpservers.WithTags([]string{"cloudflare", "audit", "logs"}),
			mcpservers.WithGitHubURL(
				"https://github.com/cloudflare/mcp-server-cloudflare",
			),
			mcpservers.WithDescription(
				"Query audit logs and generate reports for review.",
			),
			mcpservers.WithDescriptionCN(
				"查询审计日志并生成报告供审查。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)

	mcpservers.Register(
		mcpservers.NewMcp(
			"cloudflare-dns-analytics",
			"Cloudflare DNS Analytics",
			model.PublicMCPTypeProxySSE,
			mcpservers.WithNameCN("Cloudflare DNS 分析"),
			mcpservers.WithProxyConfigTemplates(mcpservers.ProxyConfigTemplates{
				"url": {
					ConfigTemplate: mcpservers.ConfigTemplate{
						Name:        "URL",
						Required:    mcpservers.ConfigRequiredTypeInitOptional,
						Default:     "https://dns-analytics.mcp.cloudflare.com/sse",
						Description: "The Streamable http URL of the Cloudflare DNS Analytics MCP server",
					},
					Type: model.ParamTypeURL,
				},
			}),
			mcpservers.WithTags([]string{"cloudflare", "dns", "analytics"}),
			mcpservers.WithGitHubURL(
				"https://github.com/cloudflare/mcp-server-cloudflare",
			),
			mcpservers.WithDescription(
				"Optimize DNS performance and debug issues based on current set up.",
			),
			mcpservers.WithDescriptionCN(
				"基于当前设置优化 DNS 性能并调试问题。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)

	mcpservers.Register(
		mcpservers.NewMcp(
			"cloudflare-dex",
			"Cloudflare DEX",
			model.PublicMCPTypeProxySSE,
			mcpservers.WithNameCN("Cloudflare DEX"),
			mcpservers.WithProxyConfigTemplates(mcpservers.ProxyConfigTemplates{
				"url": {
					ConfigTemplate: mcpservers.ConfigTemplate{
						Name:        "URL",
						Required:    mcpservers.ConfigRequiredTypeInitOptional,
						Default:     "https://dex.mcp.cloudflare.com/sse",
						Description: "The Streamable http URL of the Cloudflare Digital Experience Monitoring MCP server",
					},
					Type: model.ParamTypeURL,
				},
			}),
			mcpservers.WithTags([]string{"cloudflare", "dex", "monitoring"}),
			mcpservers.WithGitHubURL(
				"https://github.com/cloudflare/mcp-server-cloudflare",
			),
			mcpservers.WithDescription(
				"Get quick insight on critical applications for your organization.",
			),
			mcpservers.WithDescriptionCN(
				"快速洞察您组织的关键应用程序。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)

	mcpservers.Register(
		mcpservers.NewMcp(
			"cloudflare-casb",
			"Cloudflare CASB",
			model.PublicMCPTypeProxySSE,
			mcpservers.WithNameCN("Cloudflare CASB"),
			mcpservers.WithProxyConfigTemplates(mcpservers.ProxyConfigTemplates{
				"url": {
					ConfigTemplate: mcpservers.ConfigTemplate{
						Name:        "URL",
						Required:    mcpservers.ConfigRequiredTypeInitOptional,
						Default:     "https://casb.mcp.cloudflare.com/sse",
						Description: "The Streamable http URL of the Cloudflare One CASB MCP server",
					},
					Type: model.ParamTypeURL,
				},
			}),
			mcpservers.WithTags([]string{"cloudflare", "casb", "security"}),
			mcpservers.WithGitHubURL(
				"https://github.com/cloudflare/mcp-server-cloudflare",
			),
			mcpservers.WithDescription(
				"Quickly identify any security misconfigurations for SaaS applications to safeguard users & data.",
			),
			mcpservers.WithDescriptionCN(
				"快速识别 SaaS 应用程序的任何安全配置错误，以保护用户和数据。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)

	mcpservers.Register(
		mcpservers.NewMcp(
			"cloudflare-graphql",
			"Cloudflare GraphQL",
			model.PublicMCPTypeProxySSE,
			mcpservers.WithNameCN("Cloudflare GraphQL"),
			mcpservers.WithProxyConfigTemplates(mcpservers.ProxyConfigTemplates{
				"url": {
					ConfigTemplate: mcpservers.ConfigTemplate{
						Name:        "URL",
						Required:    mcpservers.ConfigRequiredTypeInitOptional,
						Default:     "https://graphql.mcp.cloudflare.com/sse",
						Description: "The Streamable http URL of the Cloudflare GraphQL MCP server",
					},
					Type: model.ParamTypeURL,
				},
			}),
			mcpservers.WithTags([]string{"cloudflare", "graphql", "api"}),
			mcpservers.WithGitHubURL(
				"https://github.com/cloudflare/mcp-server-cloudflare",
			),
			mcpservers.WithDescription(
				"Get analytics data using Cloudflare's GraphQL API.",
			),
			mcpservers.WithDescriptionCN(
				"使用 Cloudflare 的 GraphQL API 获取分析数据。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
