package openai

import (
	"errors"
	"slices"
	"strings"
	"time"

	"github.com/bytedance/sonic/ast"
	"github.com/labring/aiproxy/core/relay/meta"
	relaymodel "github.com/labring/aiproxy/core/relay/model"
	"github.com/labring/aiproxy/core/relay/utils"
)

// StreamReasoningToReasoningContentPreHandler rewrites
// choices.[*].delta.reasoning -> choices.[*].delta.reasoning_content.
func StreamReasoningToReasoningContentPreHandler(_ *meta.Meta, node *ast.Node) error {
	choicesNode := node.Get("choices")

	nodes, err := choicesNode.ArrayUseNode()
	if err != nil {
		return err
	}

	for index, choice := range nodes {
		deltaNode := choice.Get("delta")

		reasoningString, err := deltaNode.Get("reasoning").String()
		if err != nil {
			if errors.Is(err, ast.ErrNotExist) {
				continue
			}
			return err
		}

		_, err = deltaNode.Set("reasoning_content", ast.NewString(reasoningString))
		if err != nil {
			return err
		}

		_, err = deltaNode.Unset("reasoning")
		if err != nil {
			return err
		}

		_, err = choicesNode.SetByIndex(index, choice)
		if err != nil {
			return err
		}
	}

	return nil
}

// ReasoningToReasoningContentPreHandler rewrites
// choices.[*].message.reasoning -> choices.[*].message.reasoning_content.
func ReasoningToReasoningContentPreHandler(_ *meta.Meta, node *ast.Node) error {
	choicesNode := node.Get("choices")

	nodes, err := choicesNode.ArrayUseNode()
	if err != nil {
		return err
	}

	for index, choice := range nodes {
		messageNode := choice.Get("message")

		reasoningString, err := messageNode.Get("reasoning").String()
		if err != nil {
			if errors.Is(err, ast.ErrNotExist) {
				continue
			}
			return err
		}

		_, err = messageNode.Set("reasoning_content", ast.NewString(reasoningString))
		if err != nil {
			return err
		}

		_, err = messageNode.Unset("reasoning")
		if err != nil {
			return err
		}

		_, err = choicesNode.SetByIndex(index, choice)
		if err != nil {
			return err
		}
	}

	return nil
}

func applyReasoningToOpenAIRequestForModel(
	m *meta.Meta,
	req *relaymodel.GeneralOpenAIRequest,
	reasoning relaymodel.NormalizedReasoning,
) {
	if req == nil || !reasoning.Specified {
		return
	}

	effort := utils.ReasoningToOpenAIEffort(reasoning)
	if effort == "" {
		return
	}

	effortString := openAIReasoningEffortForMeta(m, effort)
	req.ReasoningEffort = &effortString
	req.Thinking = nil
}

func applyReasoningToResponsesRequestForModel(
	m *meta.Meta,
	req *relaymodel.CreateResponseRequest,
	reasoning relaymodel.NormalizedReasoning,
) {
	if req == nil || !reasoning.Specified {
		return
	}

	effort := utils.ReasoningToOpenAIEffort(reasoning)
	if effort == "" {
		return
	}

	effortString := openAIReasoningEffortForMeta(m, effort)
	req.Reasoning = &relaymodel.ResponseReasoning{
		Effort: &effortString,
	}
}

func patchOpenAIReasoningEffort(m *meta.Meta) func(node *ast.Node) error {
	return func(node *ast.Node) error {
		if node == nil {
			return nil
		}

		effortNode := node.Get("reasoning_effort")
		if !effortNode.Exists() {
			return nil
		}

		effortString, err := effortNode.String()
		if err != nil {
			return nil
		}

		effort := relaymodel.NormalizeReasoningEffort(effortString)
		if effort == "" {
			return nil
		}

		_, err = node.Set(
			"reasoning_effort",
			ast.NewString(openAIReasoningEffortForMeta(m, effort)),
		)

		return err
	}
}

func patchOpenAIResponsesReasoningEffort(m *meta.Meta) func(node *ast.Node) error {
	return func(node *ast.Node) error {
		if node == nil {
			return nil
		}

		reasoningNode := node.Get("reasoning")
		if !reasoningNode.Exists() {
			return nil
		}

		effortNode := reasoningNode.Get("effort")
		if !effortNode.Exists() {
			return nil
		}

		effortString, err := effortNode.String()
		if err != nil {
			return nil
		}

		effort := relaymodel.NormalizeReasoningEffort(effortString)
		if effort == "" {
			return nil
		}

		_, err = reasoningNode.Set(
			"effort",
			ast.NewString(openAIReasoningEffortForMeta(m, effort)),
		)
		if err != nil {
			return err
		}

		_, err = node.Set("reasoning", *reasoningNode)

		return err
	}
}

func openAIReasoningEffortForMeta(
	m *meta.Meta,
	effort relaymodel.ReasoningEffort,
) relaymodel.ReasoningEffort {
	if m == nil {
		return effort
	}

	return openAIReasoningEffortForModel(m.OriginModel, m.ActualModel, effort)
}

func openAIReasoningEffortForModel(
	originModel string,
	actualModel string,
	effort relaymodel.ReasoningEffort,
) relaymodel.ReasoningEffort {
	effort = relaymodel.NormalizeReasoningEffort(effort)
	if effort == "" {
		return effort
	}

	supported, ok := openAIReasoningEffortsForModel(originModel, actualModel)
	if !ok {
		return effort
	}

	if slices.Contains(supported, effort) {
		return effort
	}

	return closestOpenAIReasoningEffort(effort, supported)
}

func openAIReasoningEffortsForModel(
	originModel string,
	actualModel string,
) ([]relaymodel.ReasoningEffort, bool) {
	return openAIReasoningEffortsForName(utils.FirstMatchingModelName(
		originModel,
		actualModel,
		func(modelName string) bool {
			_, ok := openAIReasoningEffortsForName(modelName)
			return ok
		},
	))
}

func openAIReasoningEffortsForName(modelName string) ([]relaymodel.ReasoningEffort, bool) {
	modelName = strings.ToLower(strings.TrimSpace(modelName))
	if modelName == "" {
		return nil, false
	}

	switch {
	case matchesOpenAIModelFamily(modelName, "gpt-5.4-pro") ||
		matchesOpenAIModelFamily(modelName, "gpt-5.2-pro"):
		return []relaymodel.ReasoningEffort{
			relaymodel.ReasoningEffortMedium,
			relaymodel.ReasoningEffortHigh,
			relaymodel.ReasoningEffortXHigh,
		}, true
	case matchesOpenAIModelFamily(modelName, "gpt-5.5") ||
		matchesOpenAIModelFamily(modelName, "gpt-5.4") ||
		matchesOpenAIModelFamily(modelName, "gpt-5.2"):
		return []relaymodel.ReasoningEffort{
			relaymodel.ReasoningEffortNone,
			relaymodel.ReasoningEffortLow,
			relaymodel.ReasoningEffortMedium,
			relaymodel.ReasoningEffortHigh,
			relaymodel.ReasoningEffortXHigh,
		}, true
	case matchesOpenAIModelFamily(modelName, "gpt-5-pro"):
		return []relaymodel.ReasoningEffort{
			relaymodel.ReasoningEffortHigh,
		}, true
	case matchesOpenAIModelFamily(modelName, "gpt-5.1"):
		return []relaymodel.ReasoningEffort{
			relaymodel.ReasoningEffortNone,
			relaymodel.ReasoningEffortLow,
			relaymodel.ReasoningEffortMedium,
			relaymodel.ReasoningEffortHigh,
		}, true
	case matchesOpenAIModelFamily(modelName, "gpt-5"):
		return []relaymodel.ReasoningEffort{
			relaymodel.ReasoningEffortMinimal,
			relaymodel.ReasoningEffortLow,
			relaymodel.ReasoningEffortMedium,
			relaymodel.ReasoningEffortHigh,
		}, true
	default:
		return nil, false
	}
}

func matchesOpenAIModelFamily(modelName, family string) bool {
	if modelName == family {
		return true
	}

	if hasKnownOpenAIModelSuffix(modelName, family) {
		return true
	}

	for _, separator := range []string{"/", ":"} {
		if strings.Contains(modelName, separator+family) {
			suffix := strings.TrimPrefix(
				modelName[strings.LastIndex(modelName, separator)+1:],
				family,
			)

			return suffix == "" || isKnownOpenAIModelSuffix(suffix)
		}
	}

	return false
}

func hasKnownOpenAIModelSuffix(modelName, family string) bool {
	if !strings.HasPrefix(modelName, family) {
		return false
	}

	return isKnownOpenAIModelSuffix(strings.TrimPrefix(modelName, family))
}

func isKnownOpenAIModelSuffix(suffix string) bool {
	if suffix == "" {
		return true
	}

	if matched, ok := strings.CutPrefix(suffix, "-"); ok {
		if isDateSuffix(matched) {
			return true
		}

		for _, variant := range []string{"mini", "nano", "chat-latest"} {
			if matched == variant {
				return true
			}

			if dateSuffix, ok := strings.CutPrefix(matched, variant+"-"); ok {
				return isDateSuffix(dateSuffix)
			}
		}
	}

	return false
}

func isDateSuffix(value string) bool {
	if len(value) != len(time.DateOnly) {
		return false
	}

	for index, char := range value {
		switch index {
		case 4, 7:
			if char != '-' {
				return false
			}
		default:
			if char < '0' || char > '9' {
				return false
			}
		}
	}

	return true
}

func closestOpenAIReasoningEffort(
	effort relaymodel.ReasoningEffort,
	supported []relaymodel.ReasoningEffort,
) relaymodel.ReasoningEffort {
	if len(supported) == 0 {
		return effort
	}

	target, ok := openAIReasoningEffortRank(effort)
	if !ok {
		return effort
	}

	best := supported[0]

	bestRank, ok := openAIReasoningEffortRank(best)
	if !ok {
		return best
	}

	bestDistance := absInt(bestRank - target)
	for _, candidate := range supported[1:] {
		candidateRank, ok := openAIReasoningEffortRank(candidate)
		if !ok {
			continue
		}

		distance := absInt(candidateRank - target)
		if distance < bestDistance || (distance == bestDistance && candidateRank > bestRank) {
			best = candidate
			bestRank = candidateRank
			bestDistance = distance
		}
	}

	return best
}

func openAIReasoningEffortRank(effort relaymodel.ReasoningEffort) (int, bool) {
	switch effort {
	case relaymodel.ReasoningEffortNone:
		return 0, true
	case relaymodel.ReasoningEffortMinimal:
		return 1, true
	case relaymodel.ReasoningEffortLow:
		return 2, true
	case relaymodel.ReasoningEffortMedium:
		return 3, true
	case relaymodel.ReasoningEffortHigh:
		return 4, true
	case relaymodel.ReasoningEffortXHigh:
		return 5, true
	default:
		return 0, false
	}
}

func absInt(value int) int {
	if value < 0 {
		return -value
	}

	return value
}
