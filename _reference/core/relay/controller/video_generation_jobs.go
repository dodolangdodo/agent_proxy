package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/common"
	"github.com/labring/aiproxy/core/model"
	relaymodel "github.com/labring/aiproxy/core/relay/model"
	"github.com/labring/aiproxy/core/relay/utils"
)

func ValidateVideoGenerationJobRequest(c *gin.Context, mc model.ModelConfig) error {
	request, err := getVideoGenerationJobRequest(c)
	if err != nil {
		return err
	}

	return validateVideoGenerationJobRequest(request, mc)
}

func GetVideoGenerationJobRequestPrice(c *gin.Context, mc model.ModelConfig) (model.Price, error) {
	request, err := getVideoGenerationJobRequest(c)
	if err != nil {
		return model.Price{}, err
	}

	if err := validateVideoGenerationJobRequest(request, mc); err != nil {
		return model.Price{}, err
	}

	return getVideoRequestPrice(mc.Price), nil
}

func GetVideoGenerationJobRequestUsage(c *gin.Context, mc model.ModelConfig) (RequestUsage, error) {
	request, err := getVideoGenerationJobRequest(c)
	if err != nil {
		return RequestUsage{}, err
	}

	if err := validateVideoGenerationJobRequest(request, mc); err != nil {
		return RequestUsage{}, err
	}

	return RequestUsage{
		// Video providers bill with incompatible units: returned videos, seconds,
		// resolution tiers, or provider-specific async usage. Do not reserve
		// balance from a preflight estimate; response/async handlers provide the
		// final billable usage after the upstream result is known.
		Usage: model.Usage{},
		Context: model.UsageContext{
			Resolution: videoRequestPriceResolution(request),
		},
	}, nil
}

func validateOpenAIVideoJobResolutionFormat(request *relaymodel.VideoGenerationJobRequest) error {
	if request == nil {
		return nil
	}

	if (request.Width == 0) != (request.Height == 0) {
		return NewBadRequestParamError("width and height must be provided together")
	}

	return nil
}

func validateVideoGenerationJobRequest(
	request *relaymodel.VideoGenerationJobRequest,
	mc model.ModelConfig,
) error {
	fuzzy := !mc.DisableResolutionFuzzyMatch

	if err := validateOpenAIVideoJobResolutionFormat(request); err != nil {
		return err
	}

	if err := validateVideoGenerationSeconds(
		request.NSeconds,
		mc.MaxVideoGenerationSeconds,
	); err != nil {
		return err
	}

	if err := validateVideoGenerationCount(
		request.NVariants,
		mc.MaxVideoGenerationCount,
	); err != nil {
		return err
	}

	return validateSupportedVideoResolution(
		videoRequestPriceResolution(request),
		mc,
		openAIVideoSupportedResolutionOptions(mc.AllowedResolutions, fuzzy),
	)
}

func getVideoGenerationJobRequest(c *gin.Context) (*relaymodel.VideoGenerationJobRequest, error) {
	contentType := c.Request.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "multipart/form-data") {
		if err := common.ParseMultipartFormWithLimit(c.Request); err != nil {
			return nil, NewBadRequestParamError(err.Error())
		}

		return getMultipartVideoGenerationJobRequest(c)
	}

	request, err := utils.UnmarshalVideoGenerationJobRequest(c.Request)
	if err != nil {
		return nil, err
	}

	if err := validateParsedVideoGenerationJobRequest(request); err != nil {
		return nil, err
	}

	return request, nil
}

func validateParsedVideoGenerationJobRequest(request *relaymodel.VideoGenerationJobRequest) error {
	if request.NSeconds < 0 {
		return NewBadRequestParamError("invalid n_seconds: must be non-negative")
	}

	if request.NVariants < 0 {
		return NewBadRequestParamError("invalid n_variants: must be non-negative")
	}

	return nil
}

func getMultipartVideoGenerationJobRequest(
	c *gin.Context,
) (*relaymodel.VideoGenerationJobRequest, error) {
	request := &relaymodel.VideoGenerationJobRequest{
		Prompt: c.PostForm("prompt"),
		Model:  c.PostForm("model"),
	}

	var err error
	if request.Width, err = parseOptionalPositiveInt(c.PostForm("width"), "width"); err != nil {
		return nil, err
	}

	if request.Height, err = parseOptionalPositiveInt(c.PostForm("height"), "height"); err != nil {
		return nil, err
	}

	request.NVariants, err = parseOptionalPositiveInt(c.PostForm("n_variants"), "n_variants")
	if err != nil {
		return nil, err
	}

	request.NSeconds, err = parseOptionalPositiveInt(c.PostForm("n_seconds"), "n_seconds")
	if err != nil {
		return nil, err
	}

	return request, nil
}

func videoRequestPriceResolution(request *relaymodel.VideoGenerationJobRequest) string {
	if request == nil {
		return ""
	}

	if request.Width > 0 && request.Height > 0 {
		return relaymodel.VideoResolutionFromDimensions(request.Width, request.Height)
	}

	return ""
}
