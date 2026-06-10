package flomo

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/bytedance/sonic"
)

// Client represents the Flomo API client
type Client struct {
	apiURL     string
	httpClient *http.Client
}

// WriteNoteRequest represents the request structure for writing a note
type WriteNoteRequest struct {
	Content string `json:"content"`
}

// Memo represents a memo in the response
type Memo struct {
	Slug string `json:"slug"`
	URL  string `json:"url,omitempty"`
}

// Response represents the response from Flomo API
type Response struct {
	Memo    *Memo  `json:"memo,omitempty"`
	Message string `json:"message,omitempty"`
}

// NewClient creates a new Flomo client
func NewClient(apiURL string) *Client {
	return &Client{
		apiURL: apiURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// WriteNote writes a note to Flomo
func (c *Client) WriteNote(ctx context.Context, content string) (*Response, error) {
	if content == "" {
		return nil, errors.New("invalid content")
	}

	req := WriteNoteRequest{
		Content: content,
	}

	jsonData, err := sonic.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.apiURL,
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status %s", resp.Status)
	}

	var result Response
	if err := sonic.ConfigDefault.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Add memo URL if slug is present
	if result.Memo != nil && result.Memo.Slug != "" {
		result.Memo.URL = "https://v.flomoapp.com/mine/?memo_id=" + result.Memo.Slug
	}

	return &result, nil
}
