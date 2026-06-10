package terraform

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
			"terraform",
			"Terraform",
			model.PublicMCPTypeDocs,
			mcpservers.WithNameCN("Terraform"),
			mcpservers.WithTags([]string{"terraform", "iac", "infrastructure", "cloud"}),
			mcpservers.WithGitHubURL(
				"https://github.com/hashicorp/terraform-mcp-server",
			),
			mcpservers.WithDescription(
				"Providing seamless integration with Terraform Registry APIs, enabling advanced automation and interaction capabilities for Infrastructure as Code (IaC) development.",
			),
			mcpservers.WithDescriptionCN(
				"提供与 Terraform Registry APIs 的无缝集成，为基础设施即代码 (IaC) 开发启用高级自动化和交互功能。",
			),
			mcpservers.WithReadme(readme),
			mcpservers.WithReadmeCN(readmeCN),
		),
	)
}
