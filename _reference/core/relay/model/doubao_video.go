package model

type DoubaoVideoTaskResponse struct {
	ID                    string            `json:"id,omitempty"`
	Model                 string            `json:"model,omitempty"`
	Status                string            `json:"status,omitempty"`
	Error                 *OpenAIError      `json:"error,omitempty"`
	Content               DoubaoVideoOutput `json:"content,omitempty"`
	Usage                 DoubaoVideoUsage  `json:"usage,omitempty"`
	Seed                  int64             `json:"seed,omitempty"`
	Resolution            string            `json:"resolution,omitempty"`
	Ratio                 string            `json:"ratio,omitempty"`
	Duration              int               `json:"duration,omitempty"`
	Frames                int               `json:"frames,omitempty"`
	FramesPerSecond       int               `json:"framespersecond,omitempty"`
	CreatedAt             int64             `json:"created_at,omitempty"`
	UpdatedAt             int64             `json:"updated_at,omitempty"`
	ServiceTier           string            `json:"service_tier,omitempty"`
	ExecutionExpiresAfter int64             `json:"execution_expires_after,omitempty"`
	GenerateAudio         *bool             `json:"generate_audio,omitempty"`
	Draft                 *bool             `json:"draft,omitempty"`
	DraftTaskID           string            `json:"draft_task_id,omitempty"`
}

type DoubaoVideoOutput struct {
	VideoURL     string `json:"video_url,omitempty"`
	LastFrameURL string `json:"last_frame_url,omitempty"`
	FileURL      string `json:"file_url,omitempty"`
}

type DoubaoVideoUsage struct {
	CompletionTokens int64                `json:"completion_tokens,omitempty"`
	TotalTokens      int64                `json:"total_tokens,omitempty"`
	ToolUsage        DoubaoVideoToolUsage `json:"tool_usage,omitempty"`
}

type DoubaoVideoToolUsage struct {
	WebSearch int64 `json:"web_search,omitempty"`
}
