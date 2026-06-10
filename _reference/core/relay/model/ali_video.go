package model

type AliVideoTaskResponse struct {
	RequestID  string        `json:"request_id,omitempty"`
	Code       string        `json:"code,omitempty"`
	Message    string        `json:"message,omitempty"`
	Output     AliVideoTask  `json:"output,omitempty"`
	Usage      AliVideoUsage `json:"usage,omitempty"`
	StatusCode int           `json:"status_code,omitempty"`
}

type AliVideoTask struct {
	TaskID         string `json:"task_id,omitempty"`
	TaskStatus     string `json:"task_status,omitempty"`
	SubmitTime     string `json:"submit_time,omitempty"`
	ScheduledTime  string `json:"scheduled_time,omitempty"`
	EndTime        string `json:"end_time,omitempty"`
	VideoURL       string `json:"video_url,omitempty"`
	OutputVideoURL string `json:"output_video_url,omitempty"`
	OrigPrompt     string `json:"orig_prompt,omitempty"`
	Code           string `json:"code,omitempty"`
	Message        string `json:"message,omitempty"`
}

type AliVideoUsage struct {
	Duration            int64  `json:"duration,omitempty"`
	InputVideoDuration  int64  `json:"input_video_duration,omitempty"`
	OutputVideoDuration int64  `json:"output_video_duration,omitempty"`
	VideoDuration       int64  `json:"video_duration,omitempty"`
	VideoCount          int64  `json:"video_count,omitempty"`
	SR                  any    `json:"SR,omitempty"`
	Ratio               string `json:"ratio,omitempty"`
	Audio               *bool  `json:"audio,omitempty"`
}
