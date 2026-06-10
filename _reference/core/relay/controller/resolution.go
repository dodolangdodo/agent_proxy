package controller

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	relaymodel "github.com/labring/aiproxy/core/relay/model"
)

type resolutionAliasFunc func(string) []string

const (
	openAIImageResolutionOptions = "auto, <width>x<height>"
	openAIVideoResolutionOptions = "<width>x<height>"
	geminiVideoResolutionOptions = "720p, 1080p, 4k"
	noResolutionOptions          = "none"
)

func supportedImageResolutionMatches(
	resolution string,
	supported []string,
	fuzzy bool,
) bool {
	resolution = normalizeSupportedResolutionValue(resolution)
	if resolution == "" || len(supported) == 0 {
		return true
	}

	if !imageResolutionValue(resolution) {
		return false
	}

	supported = normalizeSupportedImageResolutionValues(supported)
	if len(supported) == 0 {
		return false
	}

	if slices.Contains(supported, resolution) {
		return true
	}

	if !fuzzy {
		return false
	}

	supportedAliases := supportedResolutionAliases(supported, supportedImageResolutionAliases)
	for _, alias := range requestImageResolutionAliases(resolution) {
		if slices.Contains(supportedAliases, alias) {
			return true
		}
	}

	return false
}

func supportedResolutionMatches(
	resolution string,
	supported []string,
	supportedAliasesFunc resolutionAliasFunc,
	requestAliasesFunc resolutionAliasFunc,
	fuzzy bool,
) bool {
	resolution = normalizeSupportedResolutionValue(resolution)
	if resolution == "" || len(supported) == 0 {
		return true
	}

	supportedExact := normalizeSupportedResolutionValues(supported)
	if len(supportedExact) == 0 {
		return false
	}

	if slices.Contains(supportedExact, resolution) {
		return true
	}

	if !fuzzy {
		return false
	}

	supportedAliases := supportedResolutionAliases(supported, supportedAliasesFunc)
	for _, alias := range requestAliasesFunc(resolution) {
		if slices.Contains(supportedAliases, alias) {
			return true
		}
	}

	return false
}

func normalizeSupportedImageResolutionValues(resolutions []string) []string {
	normalized := make([]string, 0, len(resolutions))
	for _, resolution := range resolutions {
		resolution = normalizeSupportedResolutionValue(resolution)
		if resolution != "" && imageResolutionValue(resolution) {
			normalized = append(normalized, resolution)
		}
	}

	slices.Sort(normalized)

	return slices.Compact(normalized)
}

func normalizeSupportedResolutionValues(resolutions []string) []string {
	normalized := make([]string, 0, len(resolutions))
	for _, resolution := range resolutions {
		resolution = normalizeSupportedResolutionValue(resolution)
		if resolution != "" {
			normalized = append(normalized, resolution)
		}
	}

	slices.Sort(normalized)

	return slices.Compact(normalized)
}

func normalizeSupportedResolutionValue(resolution string) string {
	resolution = strings.ToLower(strings.TrimSpace(resolution))
	resolution = strings.ReplaceAll(resolution, " ", "")
	resolution = strings.ReplaceAll(resolution, "×", "x")
	resolution = strings.ReplaceAll(resolution, "*", "x")

	return resolution
}

func validateOpenAIImageResolutionFormat(
	resolution string,
	supported []string,
	fuzzy bool,
) error {
	resolution = strings.ToLower(strings.TrimSpace(resolution))
	if resolution == "" || resolution == "auto" || dimensionResolutionValue(resolution) {
		return nil
	}

	return NewBadRequestParamError(
		fmt.Sprintf(
			"invalid image resolution `%s`, supported resolutions: %s",
			resolution,
			openAIImageSupportedResolutionOptions(supported, fuzzy),
		),
	)
}

func dimensionResolutionValue(resolution string) bool {
	width, height, ok := parseOpenAIDimensions(resolution)
	return ok && width > 0 && height > 0
}

func parseOpenAIDimensions(resolution string) (int, int, bool) {
	widthText, heightText, ok := strings.Cut(strings.TrimSpace(resolution), "x")
	if !ok || widthText == "" || heightText == "" || strings.Contains(heightText, "x") {
		return 0, 0, false
	}

	width, err := strconv.Atoi(widthText)
	if err != nil {
		return 0, 0, false
	}

	height, err := strconv.Atoi(heightText)
	if err != nil {
		return 0, 0, false
	}

	return width, height, true
}

func supportedResolutionAliases(
	resolutions []string,
	aliases resolutionAliasFunc,
) []string {
	result := make([]string, 0, len(resolutions)*3)
	for _, resolution := range resolutions {
		result = append(result, aliases(normalizeSupportedResolutionValue(resolution))...)
	}

	slices.Sort(result)

	return slices.Compact(result)
}

func supportedImageResolutionAliases(resolution string) []string {
	if resolution == "" {
		return nil
	}

	result := make([]string, 0, 2)
	if imageResolutionValue(resolution) {
		result = append(result, resolution)
	}

	if imageSize := imageResolutionSizeAlias(resolution); imageSize != "" {
		result = append(result, "size:"+imageSize)
	}

	return slices.Compact(result)
}

func requestImageResolutionAliases(resolution string) []string {
	if resolution == "" {
		return nil
	}

	result := make([]string, 0, 2)
	if imageResolutionValue(resolution) {
		result = append(result, resolution)
	}

	if imageSize := imageResolutionSizeAlias(resolution); imageSize != "" {
		result = append(result, "size:"+imageSize)
	}

	return slices.Compact(result)
}

func imageResolutionValue(resolution string) bool {
	switch resolution {
	case "512", "1k", "2k", "4k":
		return true
	}

	return dimensionResolutionValue(resolution)
}

func imageResolutionSizeAlias(resolution string) string {
	switch resolution {
	case "512", "1k", "2k", "4k":
		return resolution
	}

	width, height, ok := relaymodel.ParseVideoDimensions(resolution)
	if !ok || width <= 0 || height <= 0 {
		return ""
	}

	longSide := max(width, height)
	switch {
	case longSide >= 3500:
		return "4k"
	case longSide >= 1500:
		return "2k"
	case longSide >= 900:
		return "1k"
	default:
		return "512"
	}
}

func openAIImageSupportedResolutionOptions(supported []string, fuzzy bool) string {
	if len(supported) == 0 {
		return openAIImageResolutionOptions
	}

	options := make([]string, 0, len(supported)+1)
	for _, resolution := range normalizeSupportedImageResolutionValues(supported) {
		switch resolution {
		case "512":
			if fuzzy {
				options = append(options, openAIImagePresetResolutionOptions("512")...)
			}
		case "1k":
			if fuzzy {
				options = append(options, openAIImagePresetResolutionOptions("1k")...)
			}
		case "2k":
			if fuzzy {
				options = append(options, openAIImagePresetResolutionOptions("2k")...)
			}
		case "4k":
			if fuzzy {
				options = append(options, openAIImagePresetResolutionOptions("4k")...)
			}
		default:
			if dimensionResolutionValue(resolution) {
				options = append(options, resolution)
			}
		}
	}

	if len(options) == 0 {
		return noResolutionOptions
	}

	sortResolutionOptions(options)

	return strings.Join(slices.Compact(options), ", ")
}

func openAIImagePresetResolutionOptions(size string) []string {
	switch size {
	case "512":
		return []string{"512x512", "768x512", "512x768"}
	case "1k":
		return []string{"1024x1024", "1536x1024", "1024x1536"}
	case "2k":
		return []string{"2048x2048", "3072x2048", "2048x3072"}
	case "4k":
		return []string{"4096x4096", "6144x4096", "4096x6144"}
	default:
		return nil
	}
}

func openAIVideoSupportedResolutionOptions(supported []string, fuzzy bool) string {
	if len(supported) == 0 {
		return openAIVideoResolutionOptions
	}

	options := make([]string, 0, len(supported))
	for _, resolution := range normalizeSupportedResolutionValues(supported) {
		if dimensionResolutionValue(resolution) {
			options = append(options, resolution)
			continue
		}

		if width, height, ok := canonicalOpenAIVideoDimensionsForTier(resolution); fuzzy && ok {
			options = append(options, fmt.Sprintf("%dx%d", width, height))
		}
	}

	if len(options) == 0 {
		return noResolutionOptions
	}

	slices.Sort(options)

	return strings.Join(slices.Compact(options), ", ")
}

func sortResolutionOptions(options []string) {
	slices.SortFunc(options, func(a, b string) int {
		aWidth, aHeight, aOK := relaymodel.ParseVideoDimensions(a)

		bWidth, bHeight, bOK := relaymodel.ParseVideoDimensions(b)
		if aOK && bOK {
			aShort, aLong := min(aWidth, aHeight), max(aWidth, aHeight)

			bShort, bLong := min(bWidth, bHeight), max(bWidth, bHeight)
			if aShort != bShort {
				return aShort - bShort
			}

			if aLong != bLong {
				return aLong - bLong
			}

			if orientationRank(aWidth, aHeight) != orientationRank(bWidth, bHeight) {
				return orientationRank(aWidth, aHeight) - orientationRank(bWidth, bHeight)
			}

			return strings.Compare(a, b)
		}

		if aOK {
			return -1
		}

		if bOK {
			return 1
		}

		return strings.Compare(a, b)
	})
}

func orientationRank(width, height int) int {
	switch {
	case width == height:
		return 0
	case width > height:
		return 1
	default:
		return 2
	}
}

func geminiVideoSupportedResolutionOptions(supported []string, fuzzy bool) string {
	if len(supported) == 0 {
		return geminiVideoResolutionOptions
	}

	options := make([]string, 0, len(supported))
	for _, resolution := range normalizeSupportedResolutionValues(supported) {
		switch resolution {
		case "720p", "1080p", "4k":
			options = append(options, resolution)
		case "480p":
			continue
		default:
			if width, height, ok := relaymodel.ParseVideoDimensions(resolution); fuzzy && ok {
				if tier := relaymodel.VideoResolutionFromDimensions(width, height); tier != "" &&
					tier != "480p" {
					options = append(options, tier)
				}
			}
		}
	}

	if len(options) == 0 {
		return noResolutionOptions
	}

	slices.Sort(options)

	return strings.Join(slices.Compact(options), ", ")
}

func canonicalOpenAIVideoDimensionsForTier(resolution string) (int, int, bool) {
	switch resolution {
	case "480p":
		return 854, 480, true
	case "720p":
		return 1280, 720, true
	case "1080p":
		return 1920, 1080, true
	case "4k":
		return 3840, 2160, true
	default:
		return 0, 0, false
	}
}

func videoResolutionAliases(resolution string) []string {
	if resolution == "" {
		return nil
	}

	result := []string{resolution}
	if width, height, ok := relaymodel.ParseVideoDimensions(resolution); ok {
		if tier := relaymodel.VideoResolutionFromDimensions(width, height); tier != "" {
			result = append(result, tier)
		}
	}

	return slices.Compact(result)
}
