package model_test

import (
	"testing"
	"time"

	"github.com/labring/aiproxy/core/model"
)

func TestPrice_ValidateConditionalPrices(t *testing.T) {
	tests := []struct {
		name    string
		price   model.Price
		wantErr bool
	}{
		{
			name: "Empty conditional prices",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{},
			},
			wantErr: false,
		},
		{
			name: "Nil conditional prices",
			price: model.Price{
				ConditionalPrices: nil,
			},
			wantErr: false,
		},
		{
			name: "Valid single condition",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin:  0,
							InputTokenMax:  32000,
							OutputTokenMin: 0,
							OutputTokenMax: 200,
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.002,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid multiple conditions - doubao-seed-1.6 example",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin:  0,
							InputTokenMax:  32000,
							OutputTokenMin: 0,
							OutputTokenMax: 200,
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.002,
						},
					},
					{
						Condition: model.PriceCondition{
							InputTokenMin:  0,
							InputTokenMax:  32000,
							OutputTokenMin: 201,
							OutputTokenMax: 16000,
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.008,
						},
					},
					{
						Condition: model.PriceCondition{
							InputTokenMin: 32001,
							InputTokenMax: 128000,
						},
						Price: model.Price{
							InputPrice:  0.0012,
							OutputPrice: 0.016,
						},
					},
					{
						Condition: model.PriceCondition{
							InputTokenMin: 128001,
							InputTokenMax: 256000,
						},
						Price: model.Price{
							InputPrice:  0.0024,
							OutputPrice: 0.024,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid input token range - min > max",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 32000,
							InputTokenMax: 1000, // min > max
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.002,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid output token range - min > max",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin:  0,
							InputTokenMax:  32000,
							OutputTokenMin: 1000,
							OutputTokenMax: 500, // min > max
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.002,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Same-specificity overlapping input and output ranges",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin:  1,
							InputTokenMax:  32000,
							OutputTokenMin: 1,
							OutputTokenMax: 500,
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.002,
						},
					},
					{
						Condition: model.PriceCondition{
							InputTokenMin:  20000, // overlaps with previous
							InputTokenMax:  50000,
							OutputTokenMin: 200, // overlaps with previous
							OutputTokenMax: 1000,
						},
						Price: model.Price{
							InputPrice:  0.0012,
							OutputPrice: 0.008,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Overlapping input ranges but non-overlapping output ranges (valid)",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin:  0,
							InputTokenMax:  32000,
							OutputTokenMin: 0,
							OutputTokenMax: 200,
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.002,
						},
					},
					{
						Condition: model.PriceCondition{
							InputTokenMin:  0,     // same input range
							InputTokenMax:  32000, // same input range
							OutputTokenMin: 201,   // non-overlapping output range
							OutputTokenMax: 16000,
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.008,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Unordered non-overlapping conditions are allowed",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 32001,
							InputTokenMax: 128000,
						},
						Price: model.Price{
							InputPrice:  0.0012,
							OutputPrice: 0.016,
						},
					},
					{
						Condition: model.PriceCondition{
							InputTokenMin: 1,
							InputTokenMax: 32000,
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.002,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid consecutive ranges",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.002,
						},
					},
					{
						Condition: model.PriceCondition{
							InputTokenMin: 32001, // consecutive with previous
							InputTokenMax: 128000,
						},
						Price: model.Price{
							InputPrice:  0.0012,
							OutputPrice: 0.016,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Gap between ranges (valid)",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.002,
						},
					},
					{
						Condition: model.PriceCondition{
							InputTokenMin: 50000, // gap between 32000 and 50000
							InputTokenMax: 128000,
						},
						Price: model.Price{
							InputPrice:  0.0012,
							OutputPrice: 0.016,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Unbounded ranges (zero values)",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0, // unbounded min
							InputTokenMax: 0, // unbounded max
						},
						Price: model.Price{
							InputPrice:  0.001,
							OutputPrice: 0.002,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Mixed bounded and unbounded ranges",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.002,
						},
					},
					{
						Condition: model.PriceCondition{
							InputTokenMin: 32001,
							InputTokenMax: 0, // unbounded max
						},
						Price: model.Price{
							InputPrice:  0.0012,
							OutputPrice: 0.016,
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.price.ValidateConditionalPrices()

			if tt.wantErr {
				if err == nil {
					t.Errorf("%s: ValidateConditionalPrices() expected error but got nil", tt.name)
				}
				return
			}

			if err != nil {
				t.Errorf("%s: ValidateConditionalPrices() unexpected error = %v", tt.name, err)
			}
		})
	}
}

func TestPrice_ValidateConditionalPrices_WithTime(t *testing.T) {
	now := time.Now().Unix()
	future := now + 3600 // 1 hour from now
	past := now - 3600   // 1 hour ago

	tests := []struct {
		name    string
		price   model.Price
		wantErr bool
	}{
		{
			name: "Valid time range - future time window",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
							StartTime:     now,
							EndTime:       future,
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.002,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid time range - no end time",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
							StartTime:     now,
							EndTime:       0, // no end limit
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.002,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Valid time range - no start time",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
							StartTime:     0, // no start limit
							EndTime:       future,
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.002,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid time range - start time >= end time",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
							StartTime:     future,
							EndTime:       now, // end before start
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.002,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid time range - start time equals end time",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
							StartTime:     now,
							EndTime:       now, // same time
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.002,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Multiple conditions with different time ranges",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
							StartTime:     past,
							EndTime:       now,
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.002,
						},
					},
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
							StartTime:     now,
							EndTime:       future,
						},
						Price: model.Price{
							InputPrice:  0.001,
							OutputPrice: 0.003,
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.price.ValidateConditionalPrices()

			if tt.wantErr {
				if err == nil {
					t.Errorf("%s: ValidateConditionalPrices() expected error but got nil", tt.name)
				}
				return
			}

			if err != nil {
				t.Errorf("%s: ValidateConditionalPrices() unexpected error = %v", tt.name, err)
			}
		})
	}
}

func TestPrice_SelectConditionalPrice_WithTime(t *testing.T) {
	now := time.Now().Unix()
	past := now - 3600      // 1 hour ago
	future := now + 3600    // 1 hour from now
	farFuture := now + 7200 // 2 hours from now

	tests := []struct {
		name           string
		price          model.Price
		usage          model.Usage
		expectedInput  float64
		expectedOutput float64
	}{
		{
			name: "Select price within active time range",
			price: model.Price{
				InputPrice:  0.001,
				OutputPrice: 0.002,
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
							StartTime:     past,
							EndTime:       future,
						},
						Price: model.Price{
							InputPrice:  0.0005,
							OutputPrice: 0.001,
						},
					},
				},
			},
			usage: model.Usage{
				InputTokens:  1000,
				OutputTokens: 500,
			},
			expectedInput:  0.0005,
			expectedOutput: 0.001,
		},
		{
			name: "Fallback to default price when time range not active (before start)",
			price: model.Price{
				InputPrice:  0.001,
				OutputPrice: 0.002,
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
							StartTime:     future,
							EndTime:       farFuture,
						},
						Price: model.Price{
							InputPrice:  0.0005,
							OutputPrice: 0.001,
						},
					},
				},
			},
			usage: model.Usage{
				InputTokens:  1000,
				OutputTokens: 500,
			},
			expectedInput:  0.001,
			expectedOutput: 0.002,
		},
		{
			name: "Fallback to default price when time range expired",
			price: model.Price{
				InputPrice:  0.001,
				OutputPrice: 0.002,
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
							StartTime:     past - 7200, // 3 hours ago
							EndTime:       past,        // 1 hour ago
						},
						Price: model.Price{
							InputPrice:  0.0005,
							OutputPrice: 0.001,
						},
					},
				},
			},
			usage: model.Usage{
				InputTokens:  1000,
				OutputTokens: 500,
			},
			expectedInput:  0.001,
			expectedOutput: 0.002,
		},
		{
			name: "Select first matching price with multiple time-based conditions",
			price: model.Price{
				InputPrice:  0.001,
				OutputPrice: 0.002,
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
							StartTime:     past,
							EndTime:       future,
						},
						Price: model.Price{
							InputPrice:  0.0005,
							OutputPrice: 0.001,
						},
					},
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
							StartTime:     past,
							EndTime:       farFuture, // broader time range
						},
						Price: model.Price{
							InputPrice:  0.0008,
							OutputPrice: 0.0015,
						},
					},
				},
			},
			usage: model.Usage{
				InputTokens:  1000,
				OutputTokens: 500,
			},
			expectedInput:  0.0005,
			expectedOutput: 0.001,
		},
		{
			name: "Time range with no end time (ongoing promotion)",
			price: model.Price{
				InputPrice:  0.001,
				OutputPrice: 0.002,
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
							StartTime:     past,
							EndTime:       0, // no end time
						},
						Price: model.Price{
							InputPrice:  0.0005,
							OutputPrice: 0.001,
						},
					},
				},
			},
			usage: model.Usage{
				InputTokens:  1000,
				OutputTokens: 500,
			},
			expectedInput:  0.0005,
			expectedOutput: 0.001,
		},
		{
			name: "Time range with no start time (promotion until end)",
			price: model.Price{
				InputPrice:  0.001,
				OutputPrice: 0.002,
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
							StartTime:     0, // no start time
							EndTime:       future,
						},
						Price: model.Price{
							InputPrice:  0.0005,
							OutputPrice: 0.001,
						},
					},
				},
			},
			usage: model.Usage{
				InputTokens:  1000,
				OutputTokens: 500,
			},
			expectedInput:  0.0005,
			expectedOutput: 0.001,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selectedPrice := tt.price.SelectConditionalPrice(tt.usage, model.UsageContext{})

			if float64(selectedPrice.InputPrice) != tt.expectedInput {
				t.Errorf("%s: expected input price %v, got %v",
					tt.name, tt.expectedInput, float64(selectedPrice.InputPrice))
			}

			if float64(selectedPrice.OutputPrice) != tt.expectedOutput {
				t.Errorf("%s: expected output price %v, got %v",
					tt.name, tt.expectedOutput, float64(selectedPrice.OutputPrice))
			}
		})
	}
}

func TestPrice_SelectConditionalPrice_WithServiceTier(t *testing.T) {
	tests := []struct {
		name          string
		price         model.Price
		usage         model.Usage
		serviceTier   string
		expectedInput float64
	}{
		{
			name: "match specific service tier",
			price: model.Price{
				InputPrice: 0.001,
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							ServiceTier: "priority",
						},
						Price: model.Price{
							InputPrice: 0.003,
						},
					},
				},
			},
			usage: model.Usage{
				InputTokens: 1000,
			},
			serviceTier:   "priority",
			expectedInput: 0.003,
		},
		{
			name: "fallback when service tier not matched",
			price: model.Price{
				InputPrice: 0.001,
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							ServiceTier: "priority",
						},
						Price: model.Price{
							InputPrice: 0.003,
						},
					},
				},
			},
			usage: model.Usage{
				InputTokens: 1000,
			},
			serviceTier:   "default",
			expectedInput: 0.001,
		},
		{
			name: "case-insensitive service tier match",
			price: model.Price{
				InputPrice: 0.001,
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							ServiceTier: "Priority",
						},
						Price: model.Price{
							InputPrice: 0.003,
						},
					},
				},
			},
			usage: model.Usage{
				InputTokens: 1000,
			},
			serviceTier:   "PRIORITY",
			expectedInput: 0.003,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selectedPrice := tt.price.SelectConditionalPrice(
				tt.usage,
				model.UsageContext{ServiceTier: tt.serviceTier},
			)
			if float64(selectedPrice.InputPrice) != tt.expectedInput {
				t.Errorf("%s: expected input price %v, got %v",
					tt.name, tt.expectedInput, float64(selectedPrice.InputPrice))
			}
		})
	}
}

func TestPrice_SelectConditionalPrice_UsesMostSpecificMatchingCondition(t *testing.T) {
	price := model.Price{
		InputPrice: 0.001,
		ConditionalPrices: []model.ConditionalPrice{
			{
				Condition: model.PriceCondition{
					InputTokenMax: 32000,
				},
				Price: model.Price{
					InputPrice: 0.002,
				},
			},
			{
				Condition: model.PriceCondition{
					InputTokenMax: 32000,
					ServiceTier:   "priority",
				},
				Price: model.Price{
					InputPrice: 0.004,
				},
			},
		},
	}

	selectedPrice := price.SelectConditionalPrice(
		model.Usage{InputTokens: 1000},
		model.UsageContext{ServiceTier: "priority"},
	)
	if float64(selectedPrice.InputPrice) != 0.004 {
		t.Fatalf("expected more specific price 0.004, got %v", selectedPrice.InputPrice)
	}
}

func TestPrice_SelectConditionalPrice_MostSpecificConditionIsOrderIndependent(t *testing.T) {
	price := model.Price{
		InputPrice: 0.001,
		ConditionalPrices: []model.ConditionalPrice{
			{
				Condition: model.PriceCondition{
					InputTokenMax: 32000,
				},
				Price: model.Price{
					InputPrice: 0.002,
				},
			},
			{
				Condition: model.PriceCondition{
					InputTokenMax: 32000,
					ServiceTier:   "priority",
				},
				Price: model.Price{
					InputPrice: 0.004,
				},
			},
		},
	}

	priorityPrice := price.SelectConditionalPrice(
		model.Usage{InputTokens: 1000},
		model.UsageContext{ServiceTier: "priority"},
	)
	if float64(priorityPrice.InputPrice) != 0.004 {
		t.Fatalf("expected priority price 0.004, got %v", priorityPrice.InputPrice)
	}

	defaultPrice := price.SelectConditionalPrice(
		model.Usage{InputTokens: 1000},
		model.UsageContext{ServiceTier: "default"},
	)
	if float64(defaultPrice.InputPrice) != 0.002 {
		t.Fatalf("expected default price 0.002, got %v", defaultPrice.InputPrice)
	}
}

func TestPrice_SelectConditionalPrice_TokenAndTimeBoundsIncreaseSpecificity(t *testing.T) {
	now := time.Now().Unix()
	price := model.Price{
		InputPrice: 0.001,
		ConditionalPrices: []model.ConditionalPrice{
			{
				Condition: model.PriceCondition{
					InputTokenMax: 32000,
				},
				Price: model.Price{
					InputPrice: 0.002,
				},
			},
			{
				Condition: model.PriceCondition{
					InputTokenMin: 100,
					InputTokenMax: 2000,
				},
				Price: model.Price{
					InputPrice: 0.003,
				},
			},
			{
				Condition: model.PriceCondition{
					InputTokenMin: 100,
					InputTokenMax: 2000,
					StartTime:     now - 60,
					EndTime:       now + 60,
				},
				Price: model.Price{
					InputPrice: 0.004,
				},
			},
		},
	}

	selectedPrice := price.SelectConditionalPrice(
		model.Usage{InputTokens: 1000},
		model.UsageContext{},
	)
	if float64(selectedPrice.InputPrice) != 0.004 {
		t.Fatalf("expected token+time bounded price 0.004, got %v", selectedPrice.InputPrice)
	}
}

func TestPrice_SelectConditionalPrice_SameSpecificityKeepsFirstMatch(t *testing.T) {
	price := model.Price{
		InputPrice: 0.001,
		ConditionalPrices: []model.ConditionalPrice{
			{
				Condition: model.PriceCondition{
					InputTokenMax: 32000,
				},
				Price: model.Price{
					InputPrice: 0.002,
				},
			},
			{
				Condition: model.PriceCondition{
					OutputTokenMax: 32000,
				},
				Price: model.Price{
					InputPrice: 0.003,
				},
			},
		},
	}

	selectedPrice := price.SelectConditionalPrice(
		model.Usage{InputTokens: 1000, OutputTokens: 1000},
		model.UsageContext{},
	)
	if float64(selectedPrice.InputPrice) != 0.002 {
		t.Fatalf("expected first same-specificity price 0.002, got %v", selectedPrice.InputPrice)
	}
}

func TestUsageContextWithFallbackPreservesRequestServiceTier(t *testing.T) {
	resultContext := model.UsageContext{}
	requestContext := model.UsageContext{ServiceTier: "priority"}

	got := resultContext.WithFallback(requestContext)
	if got.ServiceTier != "priority" {
		t.Fatalf("expected fallback service tier priority, got %q", got.ServiceTier)
	}
}

func TestUsageContextWithFallbackPreservesNativeResolution(t *testing.T) {
	resultContext := model.UsageContext{Resolution: "1920x1080"}
	requestContext := model.UsageContext{NativeResolution: "1080p"}

	got := resultContext.WithFallback(requestContext)
	if got.Resolution != "1920x1080" {
		t.Fatalf("expected protocol resolution 1920x1080, got %q", got.Resolution)
	}

	if got.NativeResolution != "1080p" {
		t.Fatalf("expected native resolution 1080p, got %q", got.NativeResolution)
	}
}

func TestUsageContextWithFallbackPreservesMediaFlags(t *testing.T) {
	resultContext := model.UsageContext{OutputAudio: new(false)}
	requestContext := model.UsageContext{
		InputVideo:  new(true),
		OutputAudio: new(true),
	}

	got := resultContext.WithFallback(requestContext)
	if got.InputVideo == nil || !*got.InputVideo {
		t.Fatalf("expected input video fallback true, got %#v", got.InputVideo)
	}

	if got.OutputAudio == nil || *got.OutputAudio {
		t.Fatalf("expected existing output audio false, got %#v", got.OutputAudio)
	}
}

func TestPrice_SelectConditionalPrice_WithMediaConditions(t *testing.T) {
	price := model.Price{
		OutputPrice: 0.08,
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
			{
				Condition: model.PriceCondition{
					Resolution: []string{"720p"},
				},
				Price: model.Price{
					OutputPrice: 0.40,
				},
			},
		},
	}

	imagePrice := price.SelectConditionalPrice(model.Usage{}, model.UsageContext{
		Resolution: "1024*1024",
		Quality:    "HD",
	})
	if float64(imagePrice.OutputPrice) != 0.34 {
		t.Fatalf("expected image conditional price 0.34, got %v", imagePrice.OutputPrice)
	}

	videoPrice := price.SelectConditionalPrice(model.Usage{}, model.UsageContext{
		Resolution: "720P",
	})
	if float64(videoPrice.OutputPrice) != 0.40 {
		t.Fatalf("expected video conditional price 0.40, got %v", videoPrice.OutputPrice)
	}

	videoDimensionPrice := price.SelectConditionalPrice(model.Usage{}, model.UsageContext{
		Resolution: "1280x720",
	})
	if float64(videoDimensionPrice.OutputPrice) != 0.40 {
		t.Fatalf(
			"expected fuzzy video dimension conditional price 0.40, got %v",
			videoDimensionPrice.OutputPrice,
		)
	}

	exactProtocolResolutionPrice := model.Price{
		OutputPrice: 0.08,
		ConditionalPrices: []model.ConditionalPrice{
			{
				Condition: model.PriceCondition{Resolution: []string{"720p"}},
				Price:     model.Price{OutputPrice: 0.40},
			},
			{
				Condition: model.PriceCondition{Resolution: []string{"1280*720"}},
				Price:     model.Price{OutputPrice: 0.50},
			},
		},
	}

	selectedExactProtocolResolutionPrice := exactProtocolResolutionPrice.SelectConditionalPrice(
		model.Usage{},
		model.UsageContext{
			Resolution:       "1280 X 720",
			NativeResolution: "720P",
		},
	)
	if float64(selectedExactProtocolResolutionPrice.OutputPrice) != 0.50 {
		t.Fatalf(
			"expected exact protocol resolution conditional price 0.50, got %v",
			selectedExactProtocolResolutionPrice.OutputPrice,
		)
	}

	disabledFuzzyPrice := price.SelectConditionalPriceWithOptions(
		model.Usage{},
		model.UsageContext{Resolution: "1280x720"},
		model.PriceSelectionOptions{DisableResolutionFuzzyMatch: true},
	)
	if float64(disabledFuzzyPrice.OutputPrice) != 0.08 {
		t.Fatalf(
			"expected disabled fuzzy resolution fallback price 0.08, got %v",
			disabledFuzzyPrice.OutputPrice,
		)
	}

	nativeResolutionPrice := price.SelectConditionalPriceWithOptions(
		model.Usage{},
		model.UsageContext{
			Resolution:       "1280x720",
			NativeResolution: "720p",
		},
		model.PriceSelectionOptions{DisableResolutionFuzzyMatch: true},
	)
	if float64(nativeResolutionPrice.OutputPrice) != 0.40 {
		t.Fatalf(
			"expected native resolution conditional price 0.40, got %v",
			nativeResolutionPrice.OutputPrice,
		)
	}

	protocolFallbackPrice := model.Price{
		OutputPrice: 0.08,
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
	}

	protocolResolutionFallbackPrice := protocolFallbackPrice.SelectConditionalPriceWithOptions(
		model.Usage{},
		model.UsageContext{
			Resolution:       "1024x1024",
			NativeResolution: "720p",
			Quality:          "hd",
		},
		model.PriceSelectionOptions{DisableResolutionFuzzyMatch: true},
	)
	if float64(protocolResolutionFallbackPrice.OutputPrice) != 0.34 {
		t.Fatalf(
			"expected protocol resolution fallback conditional price 0.34, got %v",
			protocolResolutionFallbackPrice.OutputPrice,
		)
	}

	fallbackPrice := price.SelectConditionalPrice(model.Usage{}, model.UsageContext{
		Resolution: "1024x1024",
		Quality:    "standard",
	})
	if float64(fallbackPrice.OutputPrice) != 0.08 {
		t.Fatalf("expected fallback price 0.08, got %v", fallbackPrice.OutputPrice)
	}
}

func TestPrice_SelectConditionalPrice_WithMediaFlags(t *testing.T) {
	price := model.Price{
		OutputPrice: 0.20,
		ConditionalPrices: []model.ConditionalPrice{
			{
				Condition: model.PriceCondition{
					InputMedia: new(false),
				},
				Price: model.Price{OutputPrice: 0.012},
			},
			{
				Condition: model.PriceCondition{
					Resolution: []string{"720p"},
					InputVideo: new(false),
				},
				Price: model.Price{OutputPrice: 0.046},
			},
			{
				Condition: model.PriceCondition{
					Resolution: []string{"720p"},
					InputVideo: new(true),
				},
				Price: model.Price{OutputPrice: 0.028},
			},
			{
				Condition: model.PriceCondition{
					ServiceTier: "flex",
					OutputAudio: new(false),
				},
				Price: model.Price{OutputPrice: 0.004},
			},
			{
				Condition: model.PriceCondition{ServiceTier: "flex"},
				Price:     model.Price{OutputPrice: 0.008},
			},
		},
	}

	inputVideoPrice := price.SelectConditionalPrice(model.Usage{}, model.UsageContext{
		Resolution: "1280x720",
		InputVideo: new(true),
	})
	if float64(inputVideoPrice.OutputPrice) != 0.028 {
		t.Fatalf("expected input video price 0.028, got %v", inputVideoPrice.OutputPrice)
	}

	textOnlyPrice := price.SelectConditionalPrice(model.Usage{}, model.UsageContext{
		Resolution: "720p",
		InputVideo: new(false),
	})
	if float64(textOnlyPrice.OutputPrice) != 0.046 {
		t.Fatalf("expected text-only price 0.046, got %v", textOnlyPrice.OutputPrice)
	}

	pureTextPrice := price.SelectConditionalPrice(model.Usage{}, model.UsageContext{
		InputMedia: new(false),
	})
	if float64(pureTextPrice.OutputPrice) != 0.012 {
		t.Fatalf("expected pure text price 0.012, got %v", pureTextPrice.OutputPrice)
	}

	unknownInputVideoPrice := price.SelectConditionalPrice(model.Usage{}, model.UsageContext{
		Resolution: "720p",
	})
	if float64(unknownInputVideoPrice.OutputPrice) != 0.20 {
		t.Fatalf(
			"expected base price when input video is unknown, got %v",
			unknownInputVideoPrice.OutputPrice,
		)
	}

	silentFlexPrice := price.SelectConditionalPrice(model.Usage{}, model.UsageContext{
		ServiceTier: "flex",
		OutputAudio: new(false),
	})
	if float64(silentFlexPrice.OutputPrice) != 0.004 {
		t.Fatalf("expected specific silent flex price 0.004, got %v", silentFlexPrice.OutputPrice)
	}
}

func TestPrice_SelectConditionalPrice_ResolutionAndQualityNormalizationAreIndependent(
	t *testing.T,
) {
	price := model.Price{
		OutputPrice: 0.08,
		ConditionalPrices: []model.ConditionalPrice{
			{
				Condition: model.PriceCondition{
					Resolution: []string{"1024*1024"},
					Quality:    []string{"h*d"},
				},
				Price: model.Price{
					OutputPrice: 0.12,
				},
			},
			{
				Condition: model.PriceCondition{
					Resolution: []string{"1024*1024"},
					Quality:    []string{"hxd"},
				},
				Price: model.Price{
					OutputPrice: 0.34,
				},
			},
		},
	}

	selectedPrice := price.SelectConditionalPrice(model.Usage{}, model.UsageContext{
		Resolution: "1024x1024",
		Quality:    "H*D",
	})
	if float64(selectedPrice.OutputPrice) != 0.12 {
		t.Fatalf(
			"expected quality '*' to stay distinct from 'x', got %v",
			selectedPrice.OutputPrice,
		)
	}
}

func TestPrice_SelectConditionalPrice_ResolutionMultiplicationSignsMatchOnlyResolution(
	t *testing.T,
) {
	price := model.Price{
		OutputPrice: 0.08,
		ConditionalPrices: []model.ConditionalPrice{
			{
				Condition: model.PriceCondition{
					Resolution: []string{"1024*1024"},
					Quality:    []string{"hd"},
				},
				Price: model.Price{
					OutputPrice: 0.34,
				},
			},
		},
	}

	selectedPrice := price.SelectConditionalPrice(model.Usage{}, model.UsageContext{
		Resolution: "1024×1024",
		Quality:    "HD",
	})
	if float64(selectedPrice.OutputPrice) != 0.34 {
		t.Fatalf(
			"expected resolution multiplication signs to match, got %v",
			selectedPrice.OutputPrice,
		)
	}
}

func TestPrice_SelectConditionalPrice_AutoMediaConditionMatchesOnlyAuto(t *testing.T) {
	price := model.Price{
		OutputPrice: 0.08,
		ConditionalPrices: []model.ConditionalPrice{
			{
				Condition: model.PriceCondition{
					Resolution: []string{"auto"},
					Quality:    []string{"auto"},
				},
				Price: model.Price{
					OutputPrice: 0.12,
				},
			},
		},
	}

	selectedPrice := price.SelectConditionalPrice(model.Usage{}, model.UsageContext{
		Resolution: "1024x1024",
		Quality:    "standard",
	})
	if float64(selectedPrice.OutputPrice) != 0.08 {
		t.Fatalf("expected fallback price 0.08, got %v", selectedPrice.OutputPrice)
	}

	autoPrice := price.SelectConditionalPrice(model.Usage{}, model.UsageContext{
		Resolution: "auto",
		Quality:    "auto",
	})
	if float64(autoPrice.OutputPrice) != 0.12 {
		t.Fatalf("expected auto condition price 0.12, got %v", autoPrice.OutputPrice)
	}
}

func TestPrice_SelectConditionalPrice_WithMultipleMediaConditionValues(t *testing.T) {
	price := model.Price{
		OutputPrice: 0.08,
		ConditionalPrices: []model.ConditionalPrice{
			{
				Condition: model.PriceCondition{
					Resolution: []string{"1024x1024", "1024x1536"},
					Quality:    []string{"standard", "medium"},
				},
				Price: model.Price{
					OutputPrice: 0.12,
				},
			},
		},
	}

	selectedPrice := price.SelectConditionalPrice(model.Usage{}, model.UsageContext{
		Resolution: "1024x1536",
		Quality:    "medium",
	})
	if float64(selectedPrice.OutputPrice) != 0.12 {
		t.Fatalf("expected multi-value condition price 0.12, got %v", selectedPrice.OutputPrice)
	}
}

func TestPrice_ValidateConditionalPrices_WithServiceTier(t *testing.T) {
	tests := []struct {
		name    string
		price   model.Price
		wantErr bool
	}{
		{
			name: "valid service tier",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							ServiceTier: "priority",
						},
						Price: model.Price{
							InputPrice: 0.003,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid service tier",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							ServiceTier: "premium",
						},
						Price: model.Price{
							InputPrice: 0.003,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "same token range but different service tiers are allowed",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
							ServiceTier:   "default",
						},
						Price: model.Price{
							InputPrice: 0.001,
						},
					},
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
							ServiceTier:   "priority",
						},
						Price: model.Price{
							InputPrice: 0.003,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "same token range with wildcard and specific tier is allowed",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
						},
						Price: model.Price{
							InputPrice: 0.001,
						},
					},
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
							ServiceTier:   "priority",
						},
						Price: model.Price{
							InputPrice: 0.003,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "specific tier before broad tier with same token range is allowed",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							InputTokenMin: 0,
							InputTokenMax: 32000,
							ServiceTier:   "priority",
						},
						Price: model.Price{
							InputPrice: 0.003,
						},
					},
					{
						Condition: model.PriceCondition{
							ServiceTier: "priority",
						},
						Price: model.Price{
							InputPrice: 0.002,
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.price.ValidateConditionalPrices()
			if tt.wantErr && err == nil {
				t.Errorf("%s: ValidateConditionalPrices() expected error but got nil", tt.name)
			}

			if !tt.wantErr && err != nil {
				t.Errorf("%s: ValidateConditionalPrices() unexpected error = %v", tt.name, err)
			}
		})
	}
}

func TestPrice_ValidateConditionalPrices_WithMediaConditions(t *testing.T) {
	tests := []struct {
		name    string
		price   model.Price
		wantErr bool
	}{
		{
			name: "same ranges with different resolutions are allowed",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{Resolution: []string{"720p"}},
						Price:     model.Price{OutputPrice: 0.4},
					},
					{
						Condition: model.PriceCondition{Resolution: []string{"1080p"}},
						Price:     model.Price{OutputPrice: 0.8},
					},
				},
			},
		},
		{
			name: "equivalent video dimension and tier do not overlap in validation",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{Resolution: []string{"720p"}},
						Price:     model.Price{OutputPrice: 0.4},
					},
					{
						Condition: model.PriceCondition{Resolution: []string{"1280x720"}},
						Price:     model.Price{OutputPrice: 0.5},
					},
				},
			},
		},
		{
			name: "same ranges with different qualities are allowed",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							Resolution: []string{"1024x1024"},
							Quality:    []string{"standard"},
						},
						Price: model.Price{OutputPrice: 0.08},
					},
					{
						Condition: model.PriceCondition{
							Resolution: []string{"1024*1024"},
							Quality:    []string{"hd"},
						},
						Price: model.Price{OutputPrice: 0.34},
					},
				},
			},
		},
		{
			name: "same ranges with disjoint size lists are allowed",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							Resolution: []string{"720p", "1080p"},
						},
						Price: model.Price{OutputPrice: 0.08},
					},
					{
						Condition: model.PriceCondition{
							Resolution: []string{"480p"},
						},
						Price: model.Price{OutputPrice: 0.04},
					},
				},
			},
		},
		{
			name: "same ranges with overlapping size lists fail",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{
							Resolution: []string{"720p", "1080p"},
						},
						Price: model.Price{OutputPrice: 0.08},
					},
					{
						Condition: model.PriceCondition{
							Resolution: []string{"1080P"},
						},
						Price: model.Price{OutputPrice: 0.12},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "normalized same size overlaps",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{Resolution: []string{"1024x1024"}},
						Price:     model.Price{OutputPrice: 0.08},
					},
					{
						Condition: model.PriceCondition{Resolution: []string{"1024*1024"}},
						Price:     model.Price{OutputPrice: 0.12},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "wildcard fallback with specific size is allowed",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{},
						Price:     model.Price{OutputPrice: 0.08},
					},
					{
						Condition: model.PriceCondition{Resolution: []string{"720p"}},
						Price:     model.Price{OutputPrice: 0.4},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "same ranges with different input video flags are allowed",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{InputVideo: new(false)},
						Price:     model.Price{OutputPrice: 0.08},
					},
					{
						Condition: model.PriceCondition{InputVideo: new(true)},
						Price:     model.Price{OutputPrice: 0.04},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "same ranges with different input media flags are allowed",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{InputMedia: new(false)},
						Price:     model.Price{OutputPrice: 0.08},
					},
					{
						Condition: model.PriceCondition{InputMedia: new(true)},
						Price:     model.Price{OutputPrice: 0.04},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "same ranges with different output audio flags are allowed",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{OutputAudio: new(false)},
						Price:     model.Price{OutputPrice: 0.08},
					},
					{
						Condition: model.PriceCondition{OutputAudio: new(true)},
						Price:     model.Price{OutputPrice: 0.04},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "same media flag overlaps",
			price: model.Price{
				ConditionalPrices: []model.ConditionalPrice{
					{
						Condition: model.PriceCondition{InputVideo: new(true)},
						Price:     model.Price{OutputPrice: 0.08},
					},
					{
						Condition: model.PriceCondition{InputVideo: new(true)},
						Price:     model.Price{OutputPrice: 0.04},
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.price.ValidateConditionalPrices()
			if tt.wantErr && err == nil {
				t.Fatalf("expected error")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}
