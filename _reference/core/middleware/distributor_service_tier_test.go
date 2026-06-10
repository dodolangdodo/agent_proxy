//nolint:testpackage
package middleware

import (
	"testing"

	"github.com/labring/aiproxy/core/common"
	"github.com/labring/aiproxy/core/relay/mode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRequestServiceTierFromNode(t *testing.T) {
	t.Parallel()

	t.Run("chat completions uses service_tier", func(t *testing.T) {
		t.Parallel()

		node, err := common.GetJSONNodeNoCopy(
			[]byte(`{"service_tier":"priority","serviceTier":"flex"}`),
		)
		require.NoError(t, err)

		tier, err := getRequestServiceTierFromNode(&node, mode.ChatCompletions)
		require.NoError(t, err)
		assert.Equal(t, "priority", tier)
	})

	t.Run("anthropic uses service_tier", func(t *testing.T) {
		t.Parallel()

		node, err := common.GetJSONNodeNoCopy(
			[]byte(`{"service_tier":"scale","serviceTier":"flex"}`),
		)
		require.NoError(t, err)

		tier, err := getRequestServiceTierFromNode(&node, mode.Anthropic)
		require.NoError(t, err)
		assert.Equal(t, "scale", tier)
	})

	t.Run("gemini uses serviceTier", func(t *testing.T) {
		t.Parallel()

		node, err := common.GetJSONNodeNoCopy(
			[]byte(`{"service_tier":"priority","serviceTier":"flex"}`),
		)
		require.NoError(t, err)

		tier, err := getRequestServiceTierFromNode(&node, mode.Gemini)
		require.NoError(t, err)
		assert.Equal(t, "flex", tier)
	})
}
