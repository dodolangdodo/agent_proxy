package model_test

import (
	"path/filepath"
	"testing"

	"github.com/labring/aiproxy/core/common"
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/mode"
)

func TestModelConfigShouldSummaryServiceTier(t *testing.T) {
	t.Run("default false does not record summary service tier", func(t *testing.T) {
		cfg := &model.ModelConfig{}
		if cfg.ShouldSummaryServiceTier() {
			t.Fatal("expected zero-value summary_service_tier to default to false")
		}
	})

	t.Run("explicit false disables summary service tier recording", func(t *testing.T) {
		cfg := &model.ModelConfig{SummaryServiceTier: false}
		if cfg.ShouldSummaryServiceTier() {
			t.Fatal("expected false summary_service_tier to disable recording")
		}
	})
}

func TestModelConfigLoadFromGroupModelConfigSummaryServiceTier(t *testing.T) {
	base := (&model.ModelConfig{
		SummaryServiceTier: true,
	}).LoadFromGroupModelConfig(model.GroupModelConfig{
		OverrideSummaryServiceTier: true,
		SummaryServiceTier:         false,
	})

	if base.SummaryServiceTier {
		t.Fatal("expected override to set summary_service_tier to false")
	}
}

func TestModelConfigShouldSummaryClaudeLongContext(t *testing.T) {
	t.Run("default false does not record claude long context", func(t *testing.T) {
		cfg := &model.ModelConfig{}
		if cfg.ShouldSummaryClaudeLongContext() {
			t.Fatal("expected zero-value summary_claude_long_context to default to false")
		}
	})

	t.Run("explicit true enables claude long context", func(t *testing.T) {
		cfg := &model.ModelConfig{SummaryClaudeLongContext: true}
		if !cfg.ShouldSummaryClaudeLongContext() {
			t.Fatal("expected true summary_claude_long_context to enable recording")
		}
	})
}

func TestModelConfigLoadFromGroupModelConfigSummaryClaudeLongContext(t *testing.T) {
	base := (&model.ModelConfig{
		SummaryClaudeLongContext: false,
	}).LoadFromGroupModelConfig(
		model.GroupModelConfig{
			OverrideSummaryClaudeLongContext: true,
			SummaryClaudeLongContext:         true,
		},
	)

	if !base.SummaryClaudeLongContext {
		t.Fatal("expected override to set summary_claude_long_context to true")
	}
}

func TestModelConfigLoadFromGroupModelConfigBodyStorageMaxSize(t *testing.T) {
	base := (&model.ModelConfig{
		RequestBodyStorageMaxSize:  1024,
		ResponseBodyStorageMaxSize: 2048,
	}).LoadFromGroupModelConfig(
		model.GroupModelConfig{
			OverrideRequestBodyStorageMaxSize:  true,
			RequestBodyStorageMaxSize:          512,
			OverrideResponseBodyStorageMaxSize: true,
			ResponseBodyStorageMaxSize:         1536,
		},
	)

	if base.RequestBodyStorageMaxSize != 512 {
		t.Fatalf(
			"expected request_body_storage_max_size to be overridden to 512, got %d",
			base.RequestBodyStorageMaxSize,
		)
	}

	if base.ResponseBodyStorageMaxSize != 1536 {
		t.Fatalf(
			"expected response_body_storage_max_size to be overridden to 1536, got %d",
			base.ResponseBodyStorageMaxSize,
		)
	}
}

func TestModelConfigLoadFromGroupModelConfigMaxImageGenerationCount(t *testing.T) {
	base := (&model.ModelConfig{
		MaxImageGenerationCount: 4,
	}).LoadFromGroupModelConfig(
		model.GroupModelConfig{
			OverrideMaxImageGenerationCount: true,
			MaxImageGenerationCount:         2,
		},
	)

	if base.MaxImageGenerationCount != 2 {
		t.Fatalf(
			"expected max_image_generation_count to be overridden to 2, got %d",
			base.MaxImageGenerationCount,
		)
	}
}

func TestModelConfigLoadFromGroupModelConfigMaxVideoGenerationSeconds(t *testing.T) {
	base := (&model.ModelConfig{
		MaxVideoGenerationSeconds: 10,
	}).LoadFromGroupModelConfig(
		model.GroupModelConfig{
			OverrideMaxVideoGenerationSeconds: true,
			MaxVideoGenerationSeconds:         5,
		},
	)

	if base.MaxVideoGenerationSeconds != 5 {
		t.Fatalf(
			"expected max_video_generation_seconds to be overridden to 5, got %d",
			base.MaxVideoGenerationSeconds,
		)
	}
}

func TestModelConfigLoadFromGroupModelConfigMaxVideoGenerationCount(t *testing.T) {
	base := (&model.ModelConfig{
		MaxVideoGenerationCount: 4,
	}).LoadFromGroupModelConfig(
		model.GroupModelConfig{
			OverrideMaxVideoGenerationCount: true,
			MaxVideoGenerationCount:         2,
		},
	)

	if base.MaxVideoGenerationCount != 2 {
		t.Fatalf(
			"expected max_video_generation_count to be overridden to 2, got %d",
			base.MaxVideoGenerationCount,
		)
	}
}

func TestModelConfigLoadFromGroupModelConfigTimeoutConfig(t *testing.T) {
	base := (&model.ModelConfig{
		Type: mode.ChatCompletions,
		TimeoutConfig: model.TimeoutConfig{
			RequestTimeout:       30,
			StreamRequestTimeout: 60,
		},
	}).LoadFromGroupModelConfig(
		model.GroupModelConfig{
			OverrideTimeoutConfig: true,
			TimeoutConfig: model.TimeoutConfig{
				RequestTimeout:       45,
				StreamRequestTimeout: 90,
			},
		},
	)

	if base.TimeoutConfig.RequestTimeout != 45 {
		t.Fatalf(
			"expected request timeout to be overridden to 45, got %d",
			base.TimeoutConfig.RequestTimeout,
		)
	}

	if base.TimeoutConfig.StreamRequestTimeout != 90 {
		t.Fatalf(
			"expected stream request timeout to be overridden to 90, got %d",
			base.TimeoutConfig.StreamRequestTimeout,
		)
	}
}

func TestModelConfigSupportStreamTimeout(t *testing.T) {
	t.Run("supports stream timeout for stream-capable modes", func(t *testing.T) {
		modes := []mode.Mode{
			mode.ChatCompletions,
			mode.Completions,
			mode.Anthropic,
			mode.Responses,
			mode.Gemini,
		}

		for _, m := range modes {
			cfg := &model.ModelConfig{Type: m}
			if !cfg.SupportStreamTimeout() {
				t.Fatalf("expected mode %s to support stream timeout", m.String())
			}
		}
	})

	t.Run("does not support stream timeout for non-stream modes", func(t *testing.T) {
		modes := []mode.Mode{
			mode.Embeddings,
			mode.Moderations,
			mode.ImagesGenerations,
			mode.ImagesEdits,
			mode.AudioSpeech,
			mode.AudioTranscription,
			mode.AudioTranslation,
			mode.Rerank,
			mode.ParsePdf,
			mode.VideoGenerationsJobs,
			mode.Videos,
		}

		for _, m := range modes {
			cfg := &model.ModelConfig{Type: m}
			if cfg.SupportStreamTimeout() {
				t.Fatalf("expected mode %s to not support stream timeout", m.String())
			}
		}
	})
}

func TestModelConfigBeforeSaveClearsUnsupportedStreamTimeout(t *testing.T) {
	cfg := &model.ModelConfig{
		Model: "test-embedding",
		Type:  mode.Embeddings,
		TimeoutConfig: model.TimeoutConfig{
			RequestTimeout:       30,
			StreamRequestTimeout: 120,
		},
	}

	if err := cfg.BeforeSave(nil); err != nil {
		t.Fatalf("expected BeforeSave to succeed, got error: %v", err)
	}

	if cfg.TimeoutConfig.RequestTimeout != 30 {
		t.Fatalf(
			"expected request timeout to remain unchanged, got %d",
			cfg.TimeoutConfig.RequestTimeout,
		)
	}

	if cfg.TimeoutConfig.StreamRequestTimeout != 0 {
		t.Fatalf(
			"expected unsupported stream timeout to be cleared, got %d",
			cfg.TimeoutConfig.StreamRequestTimeout,
		)
	}
}

func TestGetModelConfigLoadsFastJSONFields(t *testing.T) {
	prevDB := model.DB
	prevUsingSQLite := common.UsingSQLite

	dbPath := filepath.Join(t.TempDir(), "model-config.db")

	testDB, err := model.OpenSQLite(dbPath)
	if err != nil {
		t.Fatalf("failed to open sqlite db: %v", err)
	}

	model.DB = testDB
	common.UsingSQLite = true
	t.Cleanup(func() {
		model.DB = prevDB
		common.UsingSQLite = prevUsingSQLite
	})

	if err := testDB.AutoMigrate(&model.ModelConfig{}); err != nil {
		t.Fatalf("failed to migrate model config: %v", err)
	}

	expected := model.ModelConfig{
		Model:                       "provider:model:v1",
		Owner:                       "owner",
		Type:                        mode.ChatCompletions,
		DisableResolutionFuzzyMatch: true,
		Config: map[model.ModelConfigKey]any{
			model.ModelConfigSupportFormatsKey: []string{"json", "text"},
		},
		Plugin: map[string]map[string]any{
			"cache": {
				"enable": true,
			},
		},
		Price: model.Price{
			OutputPrice: 0.12,
			ConditionalPrices: []model.ConditionalPrice{
				{
					Condition: model.PriceCondition{
						Resolution: []string{"1024x1024"},
						Quality:    []string{"hd"},
					},
					Price: model.Price{
						OutputPrice: 0.34,
					},
				},
			},
		},
	}

	if err := testDB.Create(&expected).Error; err != nil {
		t.Fatalf("failed to create model config: %v", err)
	}

	got, err := model.GetModelConfig(expected.Model)
	if err != nil {
		t.Fatalf("expected GetModelConfig to succeed, got error: %v", err)
	}

	if got.Model != expected.Model {
		t.Fatalf("expected model %q, got %q", expected.Model, got.Model)
	}

	if got.Plugin["cache"]["enable"] != true {
		t.Fatalf("expected plugin cache.enable to be true, got %#v", got.Plugin)
	}

	if !got.DisableResolutionFuzzyMatch {
		t.Fatal("expected disable resolution fuzzy match")
	}

	if got.Price.OutputPrice != 0.12 {
		t.Fatalf("expected output price 0.12, got %#v", got.Price.OutputPrice)
	}

	if got.Price.ConditionalPrices[0].Price.OutputPrice != 0.34 {
		t.Fatalf("expected conditional output price 0.34, got %#v", got.Price.ConditionalPrices)
	}
}

func TestUpdateGroupModelConfigClearsMaxImageGenerationCount(t *testing.T) {
	prevDB := model.DB
	prevUsingSQLite := common.UsingSQLite

	dbPath := filepath.Join(t.TempDir(), "group-model-config.db")

	testDB, err := model.OpenSQLite(dbPath)
	if err != nil {
		t.Fatalf("failed to open sqlite db: %v", err)
	}

	model.DB = testDB
	common.UsingSQLite = true
	t.Cleanup(func() {
		model.DB = prevDB
		common.UsingSQLite = prevUsingSQLite
	})

	if err := testDB.AutoMigrate(&model.GroupModelConfig{}); err != nil {
		t.Fatalf("failed to migrate group model config: %v", err)
	}

	initial := model.GroupModelConfig{
		GroupID:                         "test-group",
		Model:                           "gpt-image-1",
		OverrideMaxImageGenerationCount: true,
		MaxImageGenerationCount:         5,
	}
	if err := model.SaveGroupModelConfig(initial); err != nil {
		t.Fatalf("failed to save group model config: %v", err)
	}

	updated := model.GroupModelConfig{
		GroupID:                         initial.GroupID,
		Model:                           initial.Model,
		OverrideMaxImageGenerationCount: false,
		MaxImageGenerationCount:         0,
	}
	if err := model.UpdateGroupModelConfig(updated); err != nil {
		t.Fatalf("failed to update group model config: %v", err)
	}

	got, err := model.GetGroupModelConfig(initial.GroupID, initial.Model)
	if err != nil {
		t.Fatalf("failed to get group model config: %v", err)
	}

	if got.OverrideMaxImageGenerationCount {
		t.Fatal("expected override_max_image_generation_count to be cleared")
	}

	if got.MaxImageGenerationCount != 0 {
		t.Fatalf(
			"expected max_image_generation_count to be cleared, got %d",
			got.MaxImageGenerationCount,
		)
	}
}

func TestUpdateGroupModelConfigsClearsMaxImageGenerationCount(t *testing.T) {
	prevDB := model.DB
	prevUsingSQLite := common.UsingSQLite

	dbPath := filepath.Join(t.TempDir(), "group-model-configs.db")

	testDB, err := model.OpenSQLite(dbPath)
	if err != nil {
		t.Fatalf("failed to open sqlite db: %v", err)
	}

	model.DB = testDB
	common.UsingSQLite = true
	t.Cleanup(func() {
		model.DB = prevDB
		common.UsingSQLite = prevUsingSQLite
	})

	if err := testDB.AutoMigrate(&model.GroupModelConfig{}); err != nil {
		t.Fatalf("failed to migrate group model config: %v", err)
	}

	groupID := "test-group"

	initial := []model.GroupModelConfig{
		{
			Model:                           "gpt-image-1",
			OverrideMaxImageGenerationCount: true,
			MaxImageGenerationCount:         5,
		},
	}
	if err := model.SaveGroupModelConfigs(groupID, initial); err != nil {
		t.Fatalf("failed to save group model configs: %v", err)
	}

	updated := []model.GroupModelConfig{
		{
			Model:                           initial[0].Model,
			OverrideMaxImageGenerationCount: false,
			MaxImageGenerationCount:         0,
		},
	}
	if err := model.UpdateGroupModelConfigs(groupID, updated); err != nil {
		t.Fatalf("failed to update group model configs: %v", err)
	}

	got, err := model.GetGroupModelConfig(groupID, initial[0].Model)
	if err != nil {
		t.Fatalf("failed to get group model config: %v", err)
	}

	if got.OverrideMaxImageGenerationCount {
		t.Fatal("expected override_max_image_generation_count to be cleared")
	}

	if got.MaxImageGenerationCount != 0 {
		t.Fatalf(
			"expected max_image_generation_count to be cleared, got %d",
			got.MaxImageGenerationCount,
		)
	}
}

func TestUpdateGroupModelConfigClearsMaxVideoGenerationSeconds(t *testing.T) {
	prevDB := model.DB
	prevUsingSQLite := common.UsingSQLite

	dbPath := filepath.Join(t.TempDir(), "group-model-config-video.db")

	testDB, err := model.OpenSQLite(dbPath)
	if err != nil {
		t.Fatalf("failed to open sqlite db: %v", err)
	}

	model.DB = testDB
	common.UsingSQLite = true
	t.Cleanup(func() {
		model.DB = prevDB
		common.UsingSQLite = prevUsingSQLite
	})

	if err := testDB.AutoMigrate(&model.GroupModelConfig{}); err != nil {
		t.Fatalf("failed to migrate group model config: %v", err)
	}

	initial := model.GroupModelConfig{
		GroupID:                           "test-group",
		Model:                             "sora-2",
		OverrideMaxVideoGenerationSeconds: true,
		MaxVideoGenerationSeconds:         10,
	}
	if err := model.SaveGroupModelConfig(initial); err != nil {
		t.Fatalf("failed to save group model config: %v", err)
	}

	updated := model.GroupModelConfig{
		GroupID:                           initial.GroupID,
		Model:                             initial.Model,
		OverrideMaxVideoGenerationSeconds: false,
		MaxVideoGenerationSeconds:         0,
	}
	if err := model.UpdateGroupModelConfig(updated); err != nil {
		t.Fatalf("failed to update group model config: %v", err)
	}

	got, err := model.GetGroupModelConfig(initial.GroupID, initial.Model)
	if err != nil {
		t.Fatalf("failed to get group model config: %v", err)
	}

	if got.OverrideMaxVideoGenerationSeconds {
		t.Fatal("expected override_max_video_generation_seconds to be cleared")
	}

	if got.MaxVideoGenerationSeconds != 0 {
		t.Fatalf(
			"expected max_video_generation_seconds to be cleared, got %d",
			got.MaxVideoGenerationSeconds,
		)
	}
}

func TestUpdateGroupModelConfigClearsMaxVideoGenerationCount(t *testing.T) {
	prevDB := model.DB
	prevUsingSQLite := common.UsingSQLite

	dbPath := filepath.Join(t.TempDir(), "group-model-config-video-count.db")

	testDB, err := model.OpenSQLite(dbPath)
	if err != nil {
		t.Fatalf("failed to open sqlite db: %v", err)
	}

	model.DB = testDB
	common.UsingSQLite = true
	t.Cleanup(func() {
		model.DB = prevDB
		common.UsingSQLite = prevUsingSQLite
	})

	if err := testDB.AutoMigrate(&model.GroupModelConfig{}); err != nil {
		t.Fatalf("failed to migrate group model config: %v", err)
	}

	initial := model.GroupModelConfig{
		GroupID:                         "test-group",
		Model:                           "veo",
		OverrideMaxVideoGenerationCount: true,
		MaxVideoGenerationCount:         3,
	}
	if err := model.SaveGroupModelConfig(initial); err != nil {
		t.Fatalf("failed to save group model config: %v", err)
	}

	updated := model.GroupModelConfig{
		GroupID:                         initial.GroupID,
		Model:                           initial.Model,
		OverrideMaxVideoGenerationCount: false,
		MaxVideoGenerationCount:         0,
	}
	if err := model.UpdateGroupModelConfig(updated); err != nil {
		t.Fatalf("failed to update group model config: %v", err)
	}

	got, err := model.GetGroupModelConfig(initial.GroupID, initial.Model)
	if err != nil {
		t.Fatalf("failed to get group model config: %v", err)
	}

	if got.OverrideMaxVideoGenerationCount {
		t.Fatal("expected override_max_video_generation_count to be cleared")
	}

	if got.MaxVideoGenerationCount != 0 {
		t.Fatalf(
			"expected max_video_generation_count to be cleared, got %d",
			got.MaxVideoGenerationCount,
		)
	}
}
