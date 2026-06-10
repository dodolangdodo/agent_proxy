package vertexai

import (
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/mode"
)

var ModelList = []model.ModelConfig{
	{
		Model: "claude-3-haiku@20240307",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
	},
	{
		Model: "claude-3-sonnet@20240229",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
	},
	{
		Model: "claude-3-opus@20240229",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
	},
	{
		Model: "claude-3-5-sonnet@20240620",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
	},
	{
		Model: "claude-3-5-sonnet-v2@20241022",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
	},
	{
		Model: "claude-3-5-haiku@20241022",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAnthropic,
	},
}
