package figma

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/bytedance/sonic"
)

// AuthOptions represents authentication options for Figma API
type AuthOptions struct {
	AuthType  string
	AuthToken string
}

// Client represents a Figma API client
type Client struct {
	auth       AuthOptions
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new Figma API client
func NewClient(auth AuthOptions) *Client {
	return &Client{
		auth:    auth,
		baseURL: "https://api.figma.com/v1",
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// makeRequest makes an HTTP request to the Figma API
func (c *Client) makeRequest(
	ctx context.Context,
	endpoint string,
	result any,
) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set authentication headers
	if c.auth.AuthType == "oauth" {
		req.Header.Set("Authorization", "Bearer "+c.auth.AuthToken)
	} else {
		req.Header.Set("X-Figma-Token", c.auth.AuthToken)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	if err := sonic.ConfigDefault.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}

// GetFile retrieves a Figma file
func (c *Client) GetFile(
	ctx context.Context,
	fileKey string,
	depth *int,
) (*SimplifiedDesign, error) {
	endpoint := "/files/" + fileKey
	if depth != nil {
		endpoint += "?depth=" + strconv.Itoa(*depth)
	}

	var response FileResponse
	if err := c.makeRequest(ctx, endpoint, &response); err != nil {
		return nil, err
	}

	return ParseFigmaResponse(response), nil
}

// GetNode retrieves a specific node from a Figma file
func (c *Client) GetNode(
	ctx context.Context,
	fileKey, nodeID string,
	depth *int,
) (*SimplifiedDesign, error) {
	endpoint := "/files/" + fileKey + "/nodes?ids=" + nodeID
	if depth != nil {
		endpoint += "&depth=" + strconv.Itoa(*depth)
	}

	var response NodesResponse
	if err := c.makeRequest(ctx, endpoint, &response); err != nil {
		return nil, err
	}

	return ParseFigmaNodesResponse(response), nil
}

// DownloadImages downloads images from Figma
func (c *Client) DownloadImages(
	ctx context.Context,
	fileKey string,
	nodes []ImageNode,
	pngScale float64,
	svgOptions SVGOptions,
) ([][]byte, error) {
	// Separate image fills and render requests
	var (
		imageFills     []ImageNode
		renderRequests []ImageNode
	)

	for _, node := range nodes {
		if node.ImageRef != "" {
			imageFills = append(imageFills, node)
		} else {
			renderRequests = append(renderRequests, node)
		}
	}

	var results [][]byte

	// Download image fills
	if len(imageFills) > 0 {
		fillResults, err := c.downloadImageFills(ctx, fileKey, imageFills)
		if err != nil {
			return nil, err
		}

		results = append(results, fillResults...)
	}

	// Download rendered images
	if len(renderRequests) > 0 {
		renderResults, err := c.downloadRenderedImages(
			ctx,
			fileKey,
			renderRequests,
			pngScale,
			svgOptions,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, renderResults...)
	}

	return results, nil
}

// downloadImageFills downloads image fills
func (c *Client) downloadImageFills(
	ctx context.Context,
	fileKey string,
	nodes []ImageNode,
) ([][]byte, error) {
	endpoint := "/files/" + fileKey + "/images"

	var response ImageFillsResponse
	if err := c.makeRequest(ctx, endpoint, &response); err != nil {
		return nil, err
	}

	results := make([][]byte, len(nodes))
	for i, node := range nodes {
		imageURL, exists := response.Meta.Images[node.ImageRef]
		if !exists {
			continue
		}

		image, err := c.downloadImage(ctx, imageURL)
		if err != nil {
			continue
		}

		results[i] = image
	}

	return results, nil
}

// downloadRenderedImages downloads rendered images
func (c *Client) downloadRenderedImages(
	ctx context.Context,
	fileKey string,
	nodes []ImageNode,
	pngScale float64,
	svgOptions SVGOptions,
) ([][]byte, error) {
	// Separate PNG and SVG nodes
	var pngNodes, svgNodes []ImageNode
	for _, node := range nodes {
		if strings.HasSuffix(strings.ToLower(node.FileName), ".svg") {
			svgNodes = append(svgNodes, node)
		} else {
			pngNodes = append(pngNodes, node)
		}
	}

	var allResults [][]byte

	// Download PNG images
	if len(pngNodes) > 0 {
		pngResults, err := c.downloadPNGImages(ctx, fileKey, pngNodes, pngScale)
		if err != nil {
			return nil, err
		}

		allResults = append(allResults, pngResults...)
	}

	// Download SVG images
	if len(svgNodes) > 0 {
		svgResults, err := c.downloadSVGImages(ctx, fileKey, svgNodes, svgOptions)
		if err != nil {
			return nil, err
		}

		allResults = append(allResults, svgResults...)
	}

	return allResults, nil
}

// downloadPNGImages downloads PNG images
func (c *Client) downloadPNGImages(
	ctx context.Context,
	fileKey string,
	nodes []ImageNode,
	scale float64,
) ([][]byte, error) {
	nodeIDs := make([]string, len(nodes))
	for i, node := range nodes {
		nodeIDs[i] = node.NodeID
	}

	params := url.Values{
		"ids":    {strings.Join(nodeIDs, ",")},
		"format": {"png"},
		"scale":  {strconv.FormatFloat(scale, 'f', -1, 64)},
	}

	endpoint := "/images/" + fileKey + "?" + params.Encode()

	var response ImagesResponse
	if err := c.makeRequest(ctx, endpoint, &response); err != nil {
		return nil, err
	}

	results := make([][]byte, len(nodes))
	for i, node := range nodes {
		imageURL, exists := response.Images[node.NodeID]
		if !exists {
			results[i] = nil
			continue
		}

		image, err := c.downloadImage(ctx, imageURL)
		if err != nil {
			continue
		}

		results[i] = image
	}

	return results, nil
}

// downloadSVGImages downloads SVG images
func (c *Client) downloadSVGImages(
	ctx context.Context,
	fileKey string,
	nodes []ImageNode,
	options SVGOptions,
) ([][]byte, error) {
	nodeIDs := make([]string, len(nodes))
	for i, node := range nodes {
		nodeIDs[i] = node.NodeID
	}

	params := url.Values{
		"ids":                 {strings.Join(nodeIDs, ",")},
		"format":              {"svg"},
		"svg_outline_text":    {strconv.FormatBool(options.OutlineText)},
		"svg_include_id":      {strconv.FormatBool(options.IncludeID)},
		"svg_simplify_stroke": {strconv.FormatBool(options.SimplifyStroke)},
	}

	endpoint := "/images/" + fileKey + "?" + params.Encode()

	var response ImagesResponse
	if err := c.makeRequest(ctx, endpoint, &response); err != nil {
		return nil, err
	}

	results := make([][]byte, len(nodes))
	for i, node := range nodes {
		imageURL, exists := response.Images[node.NodeID]
		if !exists {
			results[i] = nil
			continue
		}

		image, err := c.downloadImage(ctx, imageURL)
		if err != nil {
			results[i] = nil
			continue
		}

		results[i] = image
	}

	return results, nil
}

// downloadImage downloads an image from URL to local path
func (c *Client) downloadImage(ctx context.Context, imageURL string) ([]byte, error) {
	// Download image
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, imageURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download image: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
