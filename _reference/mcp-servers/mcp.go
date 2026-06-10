package mcpservers

import (
	"context"
	"errors"
	"fmt"

	"github.com/labring/aiproxy/core/model"
	"github.com/mark3labs/mcp-go/mcp"
)

type ConfigValueValidator func(value string) error

type ConfigRequiredType int

const (
	ConfigRequiredTypeInitOptional ConfigRequiredType = iota
	ConfigRequiredTypeReusingOptional
	ConfigRequiredTypeInitOnly
	ConfigRequiredTypeReusingOnly
	ConfigRequiredTypeInitOrReusingOnly
)

func (c ConfigRequiredType) Validate(config, reusingConfig string) error {
	switch c {
	case ConfigRequiredTypeInitOnly:
		if config == "" {
			return errors.New("config is required")
		}
	case ConfigRequiredTypeReusingOnly:
		if reusingConfig == "" {
			return errors.New("reusing config is required")
		}
	case ConfigRequiredTypeInitOrReusingOnly:
		if config == "" && reusingConfig == "" {
			return errors.New("config or reusing config is required")
		}

		if config != "" && reusingConfig != "" {
			return errors.New(
				"config and reusing config are both provided, but only one is allowed",
			)
		}
	}

	return nil
}

type ConfigTemplate struct {
	Name        string               `json:"name"`
	Required    ConfigRequiredType   `json:"required"`
	Example     string               `json:"example,omitempty"`
	Default     string               `json:"default,omitempty"`
	Description string               `json:"description,omitempty"`
	Validator   ConfigValueValidator `json:"-"`
}

type ConfigTemplates = map[string]ConfigTemplate

type ProxyConfigTemplate struct {
	ConfigTemplate
	Type model.ProxyParamType
}

type ProxyConfigTemplates = map[string]ProxyConfigTemplate

func ValidateConfigTemplatesConfig(
	ct ConfigTemplates,
	config, reusingConfig map[string]string,
) error {
	if len(ct) == 0 {
		return nil
	}

	for key, template := range ct {
		c := config[key]

		rc := reusingConfig[key]
		if err := template.Required.Validate(c, rc); err != nil {
			return fmt.Errorf("config required %s is invalid: %w", key, err)
		}

		if template.Validator != nil {
			if c != "" {
				if err := template.Validator(c); err != nil {
					return fmt.Errorf("config %s is invalid: %w", key, err)
				}
			} else if rc != "" {
				if err := template.Validator(rc); err != nil {
					return fmt.Errorf("reusing config %s is invalid: %w", key, err)
				}
			}
		}
	}

	return nil
}

func CheckConfigTemplatesValidate(ct ConfigTemplates) error {
	for key, value := range ct {
		err := CheckConfigTemplateValidate(value)
		if err != nil {
			return fmt.Errorf("config %s validate error: %w", key, err)
		}
	}

	return nil
}

func CheckProxyConfigTemplatesValidate(ct ProxyConfigTemplates) error {
	for key, value := range ct {
		err := CheckConfigTemplateValidate(value.ConfigTemplate)
		if err != nil {
			return fmt.Errorf("config %s validate error: %w", key, err)
		}
	}

	return nil
}

func CheckConfigTemplateValidate(value ConfigTemplate) error {
	if value.Name == "" {
		return errors.New("name is required")
	}

	if value.Description == "" {
		return errors.New("description is required")
	}

	if value.Example != "" && value.Validator != nil {
		if err := value.Validator(value.Example); err != nil {
			return fmt.Errorf("example is invalid: %w", err)
		}
	}

	if value.Default != "" && value.Validator != nil {
		if err := value.Validator(value.Default); err != nil {
			return fmt.Errorf("default is invalid: %w", err)
		}
	}

	return nil
}

type (
	NewServerFunc func(config, reusingConfig map[string]string) (Server, error)
	ListToolsFunc func(ctx context.Context) ([]mcp.Tool, error)
)

type McpServer struct {
	model.PublicMCP
	ConfigTemplates      ConfigTemplates
	ProxyConfigTemplates ProxyConfigTemplates
	newServer            NewServerFunc
	listTools            ListToolsFunc
	disableCache         bool
}

type McpConfig func(*McpServer)

func WithNameCN(nameCN string) McpConfig {
	return func(e *McpServer) {
		e.NameCN = nameCN
	}
}

func WithDescription(description string) McpConfig {
	return func(e *McpServer) {
		e.Description = description
	}
}

func WithDescriptionCN(descriptionCN string) McpConfig {
	return func(e *McpServer) {
		e.DescriptionCN = descriptionCN
	}
}

func WithReadme(readme string) McpConfig {
	return func(e *McpServer) {
		e.Readme = readme
	}
}

func WithReadmeCN(readmeCN string) McpConfig {
	return func(e *McpServer) {
		e.ReadmeCN = readmeCN
	}
}

func WithReadmeURL(readmeURL string) McpConfig {
	return func(e *McpServer) {
		e.ReadmeURL = readmeURL
	}
}

func WithReadmeCNURL(readmeCNURL string) McpConfig {
	return func(e *McpServer) {
		e.ReadmeCNURL = readmeCNURL
	}
}

func WithGitHubURL(gitHubURL string) McpConfig {
	return func(e *McpServer) {
		e.GitHubURL = gitHubURL
	}
}

func WithType(t model.PublicMCPType) McpConfig {
	return func(e *McpServer) {
		e.Type = t
	}
}

func WithLogoURL(logoURL string) McpConfig {
	return func(e *McpServer) {
		e.LogoURL = logoURL
	}
}

func WithTags(tags []string) McpConfig {
	return func(e *McpServer) {
		e.Tags = tags
	}
}

func WithConfigTemplates(configTemplates ConfigTemplates) McpConfig {
	return func(e *McpServer) {
		e.ConfigTemplates = configTemplates
	}
}

func WithProxyConfigTemplates(proxyConfigTemplates ProxyConfigTemplates) McpConfig {
	return func(e *McpServer) {
		e.ProxyConfigTemplates = proxyConfigTemplates
	}
}

func WithListToolsFunc(listTools ListToolsFunc) McpConfig {
	return func(e *McpServer) {
		e.listTools = listTools
	}
}

func WithNewServerFunc(newServer NewServerFunc) McpConfig {
	return func(e *McpServer) {
		e.newServer = newServer
	}
}

func WithDisableCache(disableCache bool) McpConfig {
	return func(e *McpServer) {
		e.disableCache = disableCache
	}
}

func NewMcp(id, name string, mcpType model.PublicMCPType, opts ...McpConfig) McpServer {
	e := McpServer{
		PublicMCP: model.PublicMCP{
			ID:   id,
			Name: name,
			Type: mcpType,
		},
	}
	for _, opt := range opts {
		opt(&e)
	}

	return e
}

var (
	ErrNotImplNewServer = errors.New("not impl new server")
	ErrNotImplListTools = errors.New("not impl list tools")
)

func (e *McpServer) NewServer(config, reusingConfig map[string]string) (Server, error) {
	if e.newServer == nil {
		return nil, ErrNotImplNewServer
	}

	if err := ValidateConfigTemplatesConfig(e.ConfigTemplates, config, reusingConfig); err != nil {
		return nil, fmt.Errorf("mcp %s config is invalid: %w", e.ID, err)
	}

	return e.newServer(config, reusingConfig)
}

func (e *McpServer) ListTools(ctx context.Context) ([]mcp.Tool, error) {
	if e.listTools == nil {
		return nil, ErrNotImplListTools
	}
	return e.listTools(ctx)
}
