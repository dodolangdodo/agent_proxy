package controller

import (
	"fmt"
	"net/url"

	"github.com/labring/aiproxy/core/model"
)

type ParamsFunc interface {
	GetParams() (map[string]string, error)
}

type groupParams struct {
	mcpID   string
	groupID string
}

func (g *groupParams) GetParams() (map[string]string, error) {
	param, err := model.CacheGetPublicMCPReusingParam(g.mcpID, g.groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to get reusing params: %w", err)
	}

	return param.Params, nil
}

func newGroupParams(mcpID, groupID string) ParamsFunc {
	return &groupParams{
		mcpID:   mcpID,
		groupID: groupID,
	}
}

type staticParams map[string]string

func (s staticParams) GetParams() (map[string]string, error) {
	return s, nil
}

// ReusingParamProcessor 统一处理reusing参数
type ReusingParamProcessor struct {
	mcpID      string
	paramsFunc ParamsFunc
}

func NewReusingParamProcessor(mcpID string, paramsFunc ParamsFunc) *ReusingParamProcessor {
	return &ReusingParamProcessor{
		mcpID:      mcpID,
		paramsFunc: paramsFunc,
	}
}

// ProcessProxyReusingParams 处理代理类型的reusing参数
func (p *ReusingParamProcessor) ProcessProxyReusingParams(
	reusingParams map[string]model.PublicMCPProxyReusingParam,
	headers map[string]string,
	backendQuery *url.Values,
) error {
	if len(reusingParams) == 0 {
		return nil
	}

	param, err := p.paramsFunc.GetParams()
	if err != nil {
		return err
	}

	for key, config := range reusingParams {
		value, exists := param[key]
		if !exists {
			if config.Required {
				return fmt.Errorf("required reusing parameter %s is missing", key)
			}
			continue
		}

		if err := p.applyProxyParam(key, value, config.Type, headers, backendQuery); err != nil {
			return err
		}
	}

	return nil
}

// ProcessEmbedReusingParams 处理嵌入类型的reusing参数
func (p *ReusingParamProcessor) ProcessEmbedReusingParams(
	reusingParams map[string]model.ReusingParam,
) (map[string]string, error) {
	if len(reusingParams) == 0 {
		return nil, nil
	}

	param, err := p.paramsFunc.GetParams()
	if err != nil {
		return nil, fmt.Errorf("failed to get reusing params: %w", err)
	}

	reusingConfig := make(map[string]string)
	for key, config := range reusingParams {
		value, exists := param[key]
		if !exists {
			if config.Required {
				return nil, fmt.Errorf("required reusing parameter %s is missing", key)
			}
			continue
		}

		reusingConfig[key] = value
	}

	return reusingConfig, nil
}

// applyProxyParam 应用代理参数到相应位置
func (p *ReusingParamProcessor) applyProxyParam(
	key, value string,
	paramType model.ProxyParamType,
	headers map[string]string,
	backendQuery *url.Values,
) error {
	switch paramType {
	case model.ParamTypeHeader:
		headers[key] = value
	case model.ParamTypeQuery:
		backendQuery.Set(key, value)
	case model.ParamTypeURL:
		return fmt.Errorf("URL parameter %s cannot be set via reusing", key)
	default:
		return fmt.Errorf("unknown param type: %s", paramType)
	}

	return nil
}
