//nolint:testpackage
package model

import "testing"

func TestGeminiChatResponseWebSearchCountUsesGroundedPromptBeforeGemini3(t *testing.T) {
	t.Parallel()

	response := GeminiChatResponse{
		ModelVersion: "gemini-2.5-flash",
		Candidates: []*GeminiChatCandidate{
			{
				GroundingMetadata: &GeminiGroundingMetadata{
					WebSearchQueries: []string{"query one", "query two"},
				},
			},
		},
	}

	if got := response.GetWebSearchCount(); got != 1 {
		t.Fatalf("expected one grounded prompt, got %d", got)
	}
}

func TestGeminiChatResponseWebSearchCountUsesSearchQueriesForGemini3(t *testing.T) {
	t.Parallel()

	response := GeminiChatResponse{
		ModelVersion: "gemini-3-pro-preview",
		Candidates: []*GeminiChatCandidate{
			{
				GroundingMetadata: &GeminiGroundingMetadata{
					WebSearchQueries: []string{"query one", "query two", "query one"},
				},
			},
		},
	}

	if got := response.GetWebSearchCount(); got != 2 {
		t.Fatalf("expected two search queries, got %d", got)
	}
}
