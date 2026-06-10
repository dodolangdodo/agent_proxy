package openai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bytedance/sonic"
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/adaptor"
	"github.com/labring/aiproxy/core/relay/mode"
	relaymodel "github.com/labring/aiproxy/core/relay/model"
	relayutils "github.com/labring/aiproxy/core/relay/utils"
	log "github.com/sirupsen/logrus"
)

var _ adaptor.AsyncUsageFetcher = (*Adaptor)(nil)

func (a *Adaptor) FetchAsyncUsage(
	ctx context.Context,
	request adaptor.AsyncUsageRequest,
) (model.Usage, model.UsageContext, bool, error) {
	channel := request.Channel
	info := request.Info

	switch mode.Mode(info.Mode) {
	case mode.VideoGenerationsJobs,
		mode.Videos,
		mode.VideosRemix,
		mode.VideosEdits,
		mode.VideosExtensions:
		return a.fetchVideoUsage(ctx, channel, info)
	case mode.Responses, mode.ChatCompletions, mode.Anthropic, mode.Gemini:
		return a.fetchResponseUsage(ctx, channel, info)
	default:
		return model.Usage{}, model.UsageContext{}, false, fmt.Errorf(
			"unsupported async usage mode: %d",
			info.Mode,
		)
	}
}

func (a *Adaptor) fetchVideoUsage(
	ctx context.Context,
	channel *model.Channel,
	info *model.AsyncUsageInfo,
) (model.Usage, model.UsageContext, bool, error) {
	if info.UpstreamID == "" {
		return model.Usage{}, model.UsageContext{}, false, errors.New("upstream id is empty")
	}

	if isOpenAIVideosMode(mode.Mode(info.Mode)) {
		return a.fetchVideoObjectUsage(ctx, channel, info)
	}

	resp, err := a.fetchAsyncUsageObject(ctx, channel, info, "/video/generations/jobs")
	if err != nil {
		return model.Usage{}, model.UsageContext{}, false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Usage{}, model.UsageContext{}, false, fmt.Errorf(
			"unexpected status code: %d",
			resp.StatusCode,
		)
	}

	var job relaymodel.VideoGenerationJob
	if err := sonic.ConfigDefault.NewDecoder(resp.Body).Decode(&job); err != nil {
		return model.Usage{}, model.UsageContext{}, false, fmt.Errorf("decode video job: %w", err)
	}

	switch job.Status {
	case relaymodel.VideoGenerationJobStatusSucceeded:
		usage, usageContext := calculateVideoUsage(&job)
		return usage, usageContext, true, nil
	case relaymodel.VideoGenerationJobStatusQueued,
		relaymodel.VideoGenerationJobStatusProcessing,
		relaymodel.VideoGenerationJobStatusRunning:
		return model.Usage{}, model.UsageContext{}, false, nil
	default:
		return model.Usage{}, model.UsageContext{}, true, fmt.Errorf(
			"video job ended with status %q",
			job.Status,
		)
	}
}

func isOpenAIVideosMode(m mode.Mode) bool {
	return m == mode.Videos ||
		m == mode.VideosRemix ||
		m == mode.VideosEdits ||
		m == mode.VideosExtensions
}

func (a *Adaptor) fetchVideoObjectUsage(
	ctx context.Context,
	channel *model.Channel,
	info *model.AsyncUsageInfo,
) (model.Usage, model.UsageContext, bool, error) {
	resp, err := a.fetchAsyncUsageObject(ctx, channel, info, "/videos")
	if err != nil {
		return model.Usage{}, model.UsageContext{}, false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Usage{}, model.UsageContext{}, false, fmt.Errorf(
			"unexpected status code: %d",
			resp.StatusCode,
		)
	}

	var video relaymodel.Video
	if err := sonic.ConfigDefault.NewDecoder(resp.Body).Decode(&video); err != nil {
		return model.Usage{}, model.UsageContext{}, false, fmt.Errorf("decode video: %w", err)
	}

	switch video.Status {
	case relaymodel.VideoStatusCompleted, relaymodel.VideoStatusSucceeded:
		usage, usageContext := calculateOfficialVideoUsage(&video)
		return usage, usageContext, true, nil
	case relaymodel.VideoStatusQueued, relaymodel.VideoStatusInProgress, "":
		return model.Usage{}, model.UsageContext{}, false, nil
	default:
		return model.Usage{}, model.UsageContext{}, true, fmt.Errorf(
			"video ended with status %q",
			video.Status,
		)
	}
}

func (a *Adaptor) fetchResponseUsage(
	ctx context.Context,
	channel *model.Channel,
	info *model.AsyncUsageInfo,
) (model.Usage, model.UsageContext, bool, error) {
	if info.UpstreamID == "" {
		return model.Usage{}, model.UsageContext{}, false, errors.New("upstream id is empty")
	}

	resp, err := a.fetchAsyncUsageObject(ctx, channel, info, "/responses")
	if err != nil {
		return model.Usage{}, model.UsageContext{}, false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Usage{}, model.UsageContext{}, false, fmt.Errorf(
			"unexpected status code: %d",
			resp.StatusCode,
		)
	}

	var response relaymodel.Response
	if err := sonic.ConfigDefault.NewDecoder(resp.Body).Decode(&response); err != nil {
		return model.Usage{}, model.UsageContext{}, false, fmt.Errorf("decode response: %w", err)
	}

	switch response.Status {
	case relaymodel.ResponseStatusInProgress, relaymodel.ResponseStatusQueued:
		return model.Usage{}, model.UsageContext{}, false, nil
	case relaymodel.ResponseStatusFailed,
		relaymodel.ResponseStatusIncomplete,
		relaymodel.ResponseStatusCancelled:
		return model.Usage{}, model.UsageContext{}, true, fmt.Errorf(
			"response ended with status %q",
			response.Status,
		)
	default:
		return response.ToModelUsage(), model.UsageContext{}, true, nil
	}
}

func (a *Adaptor) fetchAsyncUsageObject(
	ctx context.Context,
	channel *model.Channel,
	info *model.AsyncUsageInfo,
	path string,
) (*http.Response, error) {
	baseURL := asyncUsageBaseURL(channel, info, a.DefaultBaseURL())

	requestURL, err := url.JoinPath(baseURL, path, info.UpstreamID)
	if err != nil {
		return nil, fmt.Errorf("build async usage url: %w", err)
	}

	log.Debugf(
		"async usage fetch url: id=%d request_id=%s upstream_id=%s mode=%d channel_id=%d url=%s",
		info.ID,
		info.RequestID,
		info.UpstreamID,
		info.Mode,
		info.ChannelID,
		requestURL,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new async usage request: %w", err)
	}

	setupOpenAIAsyncUsageRequestHeader(channel, req)

	client, err := relayutils.LoadHTTPClientWithTLSConfigE(
		0,
		channel.ProxyURL,
		channel.SkipTLSVerify,
	)
	if err != nil {
		return nil, fmt.Errorf("load http client: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do async usage request: %w", err)
	}

	return resp, nil
}

func asyncUsageBaseURL(
	channel *model.Channel,
	info *model.AsyncUsageInfo,
	defaultBaseURL string,
) string {
	if info != nil && info.BaseURL != "" {
		return info.BaseURL
	}

	if channel != nil && channel.BaseURL != "" {
		return channel.BaseURL
	}

	return defaultBaseURL
}

func setupOpenAIAsyncUsageRequestHeader(channel *model.Channel, req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+channel.Key)
	req.Header.Set("Content-Type", "application/json")
}

func calculateVideoUsage(job *relaymodel.VideoGenerationJob) (model.Usage, model.UsageContext) {
	totalSeconds := job.NSeconds * job.NVariants

	return model.Usage{
			OutputTokens: model.ZeroNullInt64(totalSeconds),
			TotalTokens:  model.ZeroNullInt64(totalSeconds),
		}, model.UsageContext{
			Resolution: videoGenerationJobPriceResolution(job),
		}
}

func calculateOfficialVideoUsage(video *relaymodel.Video) (model.Usage, model.UsageContext) {
	return model.Usage{
		OutputTokens: model.ZeroNullInt64(video.Seconds),
		TotalTokens:  model.ZeroNullInt64(video.Seconds),
	}, model.UsageContext{Resolution: video.Size}
}

func videoGenerationJobPriceResolution(job *relaymodel.VideoGenerationJob) string {
	if job == nil {
		return ""
	}

	if job.Width <= 0 || job.Height <= 0 {
		return ""
	}

	return fmt.Sprintf("%dx%d", job.Width, job.Height)
}
