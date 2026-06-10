package alibabaclouddms

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
			"alibabacloud-dms",
			"Alibaba Cloud DMS",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("阿里云DMS"),
			mcpservers.WithGitHubURL(
				"https://github.com/aliyun/alibabacloud-dms-mcp-server",
			),
			mcpservers.WithTags([]string{"database"}),
			mcpservers.WithDescription(
				"The preferred unified data access channel for AI, supporting secure access to 30+ data sources (Alibaba Cloud全系/主流数据库/数仓).",
			),
			mcpservers.WithDescriptionCN(
				"AI 首选的统一数据访问通道，支持30多种数据源(阿里云全系/主流数据库/数仓)的安全访问。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
