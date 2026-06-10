package ai360

import (
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/mode"
)

var ModelList = []model.ModelConfig{
	{
		Model: "360GPT_S2_V9",
		Type:  mode.ChatCompletions,
		Owner: model.ModelOwnerAI360,
	},
	{
		Model: "embedding-bert-512-v1",
		Type:  mode.Embeddings,
		Owner: model.ModelOwnerAI360,
	},
	{
		Model: "embedding_s1_v1",
		Type:  mode.Embeddings,
		Owner: model.ModelOwnerAI360,
	},
	{
		Model: "semantic_similarity_s1_v1",
		Type:  mode.Embeddings,
		Owner: model.ModelOwnerAI360,
	},
}
