package engine

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/bytedance/sonic"
)

// SearchXNGEngine implements search using SearXNG instance
type SearchXNGEngine struct {
	baseURL string
	client  *http.Client
}

// NewSearchXNGEngine creates a new SearchXNG engine instance
func NewSearchXNGEngine(baseURL string) *SearchXNGEngine {
	// Ensure baseURL doesn't have trailing slash
	baseURL = strings.TrimRight(baseURL, "/")

	return &SearchXNGEngine{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 15 * time.Second, // Slightly longer timeout as it aggregates from multiple sources
		},
	}
}

// searchXNGResponse represents the JSON response from SearXNG
type searchXNGResponse struct {
	Results []searchXNGResult `json:"results"`
}

type searchXNGResult struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Content string `json:"content"`
	Engine  string `json:"engine"`
}

// Search performs a search using the SearXNG instance
func (s *SearchXNGEngine) Search(ctx context.Context, query SearchQuery) ([]SearchResult, error) {
	// Build query parameters
	params := url.Values{}
	params.Set("q", strings.Join(query.Queries, " "))
	params.Set("format", "json")

	// Set number of results per page
	if query.MaxResults > 0 {
		// SearXNG typically returns paginated results, we'll request more to ensure we get enough
		params.Set("pageno", "1")
		// Most SearXNG instances don't directly support result count, but we can filter later
	}

	// Set language if specified
	if query.Language != "" {
		params.Set("language", query.Language)
	}

	searchURL := fmt.Sprintf("%s/search?%s", s.baseURL, params.Encode())

	// Create the request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, searchURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers to appear as a regular browser request
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; SearchClient/1.0)")

	// Execute the request
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute search request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("searchxng search failed: %d - %s", resp.StatusCode, string(body))
	}

	// Parse the response
	var response searchXNGResponse
	if err := sonic.ConfigDefault.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Convert to SearchResult format
	results := make([]SearchResult, 0, len(response.Results))
	for i, item := range response.Results {
		// Limit results if MaxResults is specified
		if query.MaxResults > 0 && i >= query.MaxResults {
			break
		}

		// Clean up content - sometimes it might be empty
		content := item.Content
		if content == "" {
			content = "Result from " + item.Engine
		}

		results = append(results, SearchResult{
			Title:   item.Title,
			Link:    item.URL,
			Content: content,
		})
	}

	return results, nil
}
