package controller

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bytedance/sonic/ast"
	"github.com/labring/aiproxy/core/model"
)

func getVideoRequestPrice(price model.Price) model.Price {
	setVideoOutputPriceUnit(&price, false)
	return price
}

func validateVideoGenerationSeconds(seconds, maxSeconds int) error {
	if maxSeconds <= 0 || seconds <= maxSeconds {
		return nil
	}

	return NewBadRequestParamError(
		fmt.Sprintf("seconds must be less than or equal to %d", maxSeconds),
	)
}

func validateVideoGenerationCount(count, maxCount int) error {
	if maxCount <= 0 || count <= maxCount {
		return nil
	}

	return NewBadRequestParamError(
		fmt.Sprintf("video count must be less than or equal to %d", maxCount),
	)
}

func validateOpenAIVideoSizeFormat(size string, supported []string, fuzzy bool) error {
	size = strings.ToLower(strings.TrimSpace(size))
	if size == "" || dimensionResolutionValue(size) {
		return nil
	}

	return NewBadRequestParamError(
		fmt.Sprintf(
			"invalid video size `%s`, supported resolutions: %s",
			size,
			openAIVideoSupportedResolutionOptions(supported, fuzzy),
		),
	)
}

func validateSupportedVideoResolution(
	resolution string,
	mc model.ModelConfig,
	supportedOptions string,
) error {
	if supportedResolutionMatches(
		resolution,
		mc.AllowedResolutions,
		videoResolutionAliases,
		videoResolutionAliases,
		!mc.DisableResolutionFuzzyMatch,
	) {
		return nil
	}

	return NewBadRequestParamError(
		fmt.Sprintf(
			"unsupported video resolution `%s`, supported resolutions: %s",
			resolution,
			supportedOptions,
		),
	)
}

func setVideoOutputPriceUnit(price *model.Price, force bool) {
	if price == nil {
		return
	}

	if (force || len(price.ConditionalPrices) != 0) && price.OutputPriceUnit == 0 {
		price.OutputPriceUnit = 1
	}

	for i := range price.ConditionalPrices {
		setVideoOutputPriceUnit(&price.ConditionalPrices[i].Price, true)
	}
}

func intValueFromNode(node *ast.Node, name string) (int, bool, error) {
	if node == nil || !node.Exists() || node.TypeSafe() == ast.V_NULL {
		return 0, false, nil
	}

	valueNode := node.Get(name)
	if valueNode == nil || !valueNode.Exists() || valueNode.TypeSafe() == ast.V_NULL {
		return 0, false, nil
	}

	if valueNode.TypeSafe() == ast.V_STRING {
		value, err := valueNode.String()
		if err != nil {
			return 0, true, NewBadRequestParamError(
				fmt.Sprintf("invalid %s: %s", name, err.Error()),
			)
		}

		parsed, err := parseOptionalPositiveInt(value, name)
		if err != nil {
			return 0, true, err
		}

		return parsed, true, nil
	}

	value, err := valueNode.Int64()
	if err != nil {
		return 0, true, NewBadRequestParamError(
			fmt.Sprintf("invalid %s: %s", name, err.Error()),
		)
	}

	if value < 0 {
		return 0, true, NewBadRequestParamError(
			fmt.Sprintf("invalid %s: must be non-negative", name),
		)
	}

	return int(value), true, nil
}

func stringValueFromNode(node *ast.Node, name string) string {
	if node == nil || !node.Exists() || node.TypeSafe() == ast.V_NULL {
		return ""
	}

	valueNode := node.Get(name)
	if valueNode == nil || !valueNode.Exists() || valueNode.TypeSafe() == ast.V_NULL {
		return ""
	}

	value, err := valueNode.String()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(value)
}

func firstNonEmptyStringValueFromNode(node *ast.Node, names ...string) string {
	for _, name := range names {
		if value := stringValueFromNode(node, name); value != "" {
			return value
		}
	}

	return ""
}

func parseOptionalPositiveInt(value, name string) (int, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, NewBadRequestParamError(
			fmt.Sprintf("invalid %s: %s", name, err.Error()),
		)
	}

	if parsed < 0 {
		return 0, NewBadRequestParamError(fmt.Sprintf("invalid %s: must be non-negative", name))
	}

	return parsed, nil
}
