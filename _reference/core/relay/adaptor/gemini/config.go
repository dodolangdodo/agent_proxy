package gemini

import "github.com/labring/aiproxy/core/relay/meta"

type Config struct {
	Safety                         string `json:"safety"`
	DisableAutoImageURLToBase64    bool   `json:"disable_auto_image_url_to_base64"`
	DisableAutoAudioURLToBase64    bool   `json:"disable_auto_audio_url_to_base64"`
	DisableAutoVideoURLToBase64    bool   `json:"disable_auto_video_url_to_base64"`
	EnablePersonGenerationAllowAll bool   `json:"enable_person_generation_allow_all"`
}

func loadConfig(meta *meta.Meta) (Config, error) {
	return (&Adaptor{}).loadConfig(meta)
}

func (a *Adaptor) loadConfig(meta *meta.Meta) (Config, error) {
	return a.configCache.Load(meta, Config{})
}
