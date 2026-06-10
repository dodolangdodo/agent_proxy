package model

type GeminiVideoOperation struct {
	Name     string                       `json:"name,omitempty"`
	Done     bool                         `json:"done,omitempty"`
	Metadata map[string]any               `json:"metadata,omitempty"`
	Error    *GeminiError                 `json:"error,omitempty"`
	Response GeminiVideoOperationResponse `json:"response,omitempty"`
}

type GeminiVideoOperationResponse struct {
	GenerateVideoResponse GeminiGenerateVideoResponse `json:"generateVideoResponse,omitempty"`
	UsageMetadata         *GeminiUsageMetadata        `json:"usageMetadata,omitempty"`
}

type GeminiGenerateVideoResponse struct {
	GeneratedSamples        []GeminiGeneratedSample `json:"generatedSamples,omitempty"`
	RAIMediaFilteredCount   int                     `json:"raiMediaFilteredCount,omitempty"`
	RAIMediaFilteredReasons []string                `json:"raiMediaFilteredReasons,omitempty"`
}

type GeminiGeneratedSample struct {
	Video GeminiGeneratedVideo `json:"video,omitempty"`
}

type GeminiGeneratedVideo struct {
	URI      string `json:"uri,omitempty"`
	MimeType string `json:"mimeType,omitempty"`
}
