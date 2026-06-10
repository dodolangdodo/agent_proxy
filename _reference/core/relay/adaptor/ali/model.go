package ali

import model "github.com/labring/aiproxy/core/relay/model"

type ImageRequest struct {
	Input struct {
		Prompt         string `json:"prompt"`
		NegativePrompt string `json:"negative_prompt,omitempty"`
		RefImage       string `json:"ref_image,omitempty"`
	} `json:"input"`
	Model          string `json:"model"`
	ResponseFormat string `json:"response_format,omitempty"`
	Parameters     struct {
		Size         string   `json:"size,omitempty"`
		Steps        string   `json:"steps,omitempty"`
		Scale        string   `json:"scale,omitempty"`
		N            int      `json:"n,omitempty"`
		PromptExtend *bool    `json:"prompt_extend,omitempty"`
		Watermark    *bool    `json:"watermark,omitempty"`
		Seed         *int64   `json:"seed,omitempty"`
		Style        string   `json:"style,omitempty"`
		RefStrength  *float64 `json:"ref_strength,omitempty"`
		RefMode      string   `json:"ref_mode,omitempty"`
	} `json:"parameters,omitempty"`
}

type TaskResponse struct {
	RequestID string `json:"request_id,omitempty"`
	Code      string `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	Output    struct {
		TaskID     string `json:"task_id,omitempty"`
		TaskStatus string `json:"task_status,omitempty"`
		Code       string `json:"code,omitempty"`
		Message    string `json:"message,omitempty"`
		Results    []struct {
			B64Image string `json:"b64_image,omitempty"`
			URL      string `json:"url,omitempty"`
			Code     string `json:"code,omitempty"`
			Message  string `json:"message,omitempty"`
		} `json:"results,omitempty"`
		ImageURL    string `json:"image_url,omitempty"`
		B64Image    string `json:"b64_image,omitempty"`
		TaskMetrics struct {
			Total     int `json:"TOTAL,omitempty"`
			Succeeded int `json:"SUCCEEDED,omitempty"`
			Failed    int `json:"FAILED,omitempty"`
		} `json:"task_metrics,omitempty"`
	} `json:"output,omitempty"`
	Usage      AliImageUsage `json:"usage"`
	StatusCode int           `json:"status_code,omitempty"`
}

type AliImageUsage struct {
	ImageCount   int64  `json:"image_count,omitempty"`
	Width        int64  `json:"width,omitempty"`
	Height       int64  `json:"height,omitempty"`
	Size         string `json:"size,omitempty"`
	InputTokens  int64  `json:"input_tokens,omitempty"`
	OutputTokens int64  `json:"output_tokens,omitempty"`
	TotalTokens  int64  `json:"total_tokens,omitempty"`
}

type MultimodalImageRequest struct {
	Input struct {
		Messages []MultimodalImageMessage `json:"messages"`
	} `json:"input"`
	Model      string         `json:"model"`
	Parameters map[string]any `json:"parameters,omitempty"`
}

type MultimodalImageMessage struct {
	Role    string           `json:"role"`
	Content []map[string]any `json:"content"`
}

type MultimodalImageResponse struct {
	RequestID string `json:"request_id,omitempty"`
	Code      string `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	Output    struct {
		Choices []struct {
			FinishReason string `json:"finish_reason,omitempty"`
			Message      struct {
				Role    string `json:"role,omitempty"`
				Content []struct {
					Image        string `json:"image,omitempty"`
					URL          string `json:"url,omitempty"`
					B64Image     string `json:"b64_image,omitempty"`
					Text         string `json:"text,omitempty"`
					OrigPrompt   string `json:"orig_prompt,omitempty"`
					ActualPrompt string `json:"actual_prompt,omitempty"`
				} `json:"content,omitempty"`
			} `json:"message,omitempty"`
		} `json:"choices,omitempty"`
		Results []struct {
			B64Image     string `json:"b64_image,omitempty"`
			URL          string `json:"url,omitempty"`
			Image        string `json:"image,omitempty"`
			OrigPrompt   string `json:"orig_prompt,omitempty"`
			ActualPrompt string `json:"actual_prompt,omitempty"`
			Code         string `json:"code,omitempty"`
			Message      string `json:"message,omitempty"`
		} `json:"results,omitempty"`
	} `json:"output,omitempty"`
	Usage AliImageUsage `json:"usage,omitempty"`
}

type EmbeddingRequest struct {
	Parameters *struct {
		TextType string `json:"text_type,omitempty"`
	} `json:"parameters,omitempty"`
	Model string `json:"model"`
	Input struct {
		Texts []string `json:"texts"`
	} `json:"input"`
}

type Embedding struct {
	Embedding []float64 `json:"embedding"`
	TextIndex int       `json:"text_index"`
}

type EmbeddingResponse struct {
	Error
	Output struct {
		Embeddings []Embedding `json:"embeddings"`
	} `json:"output"`
	Usage model.EmbeddingUsage `json:"usage"`
}

type Error struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}
