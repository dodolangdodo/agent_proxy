package gemini

import (
	"strings"

	relaymeta "github.com/labring/aiproxy/core/relay/meta"
	relaymodel "github.com/labring/aiproxy/core/relay/model"
)

func trackGeminiWebSearch(
	response *relaymodel.GeminiChatResponse,
	queries map[string]struct{},
	grounded *bool,
	gemini3 *bool,
) {
	if response == nil {
		return
	}

	if response.IsGemini3Model() {
		*gemini3 = true
	}

	for query := range response.WebSearchQuerySet() {
		queries[query] = struct{}{}
	}

	if !*grounded && response.GetWebSearchCount() > 0 {
		*grounded = true
	}
}

func geminiWebSearchCount(queries map[string]struct{}, grounded, gemini3 bool) int64 {
	if gemini3 && len(queries) > 0 {
		return int64(len(queries))
	}

	if grounded {
		return 1
	}

	return 0
}

func isGemini3Meta(meta *relaymeta.Meta) bool {
	if meta == nil {
		return false
	}

	return isGemini3ModelName(meta.ActualModel) || isGemini3ModelName(meta.OriginModel)
}

func isGemini3ModelName(modelName string) bool {
	modelName = strings.ToLower(strings.TrimSpace(modelName))
	return strings.HasPrefix(modelName, "gemini-3")
}
