//nolint:testpackage
package openai

import (
	"testing"

	"github.com/labring/aiproxy/core/relay/model"
	"github.com/stretchr/testify/require"
)

func TestCalculateVideoUsageKeepsLandscapeDimensions(t *testing.T) {
	t.Parallel()

	_, usageContext := calculateVideoUsage(&model.VideoGenerationJob{
		NVariants: 1,
		NSeconds:  5,
		Width:     1280,
		Height:    720,
	})

	require.Equal(t, "1280x720", usageContext.Resolution)
}

func TestCalculateVideoUsageKeepsPortraitDimensions(t *testing.T) {
	t.Parallel()

	_, usageContext := calculateVideoUsage(&model.VideoGenerationJob{
		NVariants: 1,
		NSeconds:  5,
		Width:     720,
		Height:    1280,
	})

	require.Equal(t, "720x1280", usageContext.Resolution)
}

func TestCalculateVideoUsageKeepsEmptyResolutionWithoutDimensions(t *testing.T) {
	t.Parallel()

	_, usageContext := calculateVideoUsage(&model.VideoGenerationJob{
		NVariants: 1,
		NSeconds:  5,
	})

	require.Empty(t, usageContext.Resolution)
}
