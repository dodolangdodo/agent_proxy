//nolint:testpackage
package azure

import (
	"testing"

	relaymodel "github.com/labring/aiproxy/core/relay/model"
	"github.com/stretchr/testify/require"
)

func TestVideoResolutionFromDimensionsNormalizesLandscapeResolution(t *testing.T) {
	t.Parallel()

	size := relaymodel.VideoResolutionFromDimensions(1280, 720)

	require.Equal(t, "720p", size)
}

func TestVideoResolutionFromDimensionsNormalizesPortraitResolution(t *testing.T) {
	t.Parallel()

	size := relaymodel.VideoResolutionFromDimensions(720, 1280)

	require.Equal(t, "720p", size)
}

func TestVideoGenerationJobPriceResolutionKeepsEmptyResolutionWithoutDimensions(t *testing.T) {
	t.Parallel()

	size := videoGenerationJobPriceResolution(&relaymodel.VideoGenerationJob{})

	require.Empty(t, size)
}
