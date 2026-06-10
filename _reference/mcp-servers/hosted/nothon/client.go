package notion

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/bytedance/sonic"
)

// Client represents a Notion API client
type Client struct {
	token   string
	baseURL string
	client  *http.Client
}

// NewClient creates a new Notion API client
func NewClient(token string) *Client {
	return &Client{
		token:   token,
		baseURL: "https://api.notion.com/v1",
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// makeRequest makes an HTTP request to the Notion API
func (c *Client) makeRequest(
	ctx context.Context,
	method, endpoint string,
	body any,
	v any,
) error {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := sonic.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}

		reqBody = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+endpoint, reqBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", "2022-06-28")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error %d: %s", resp.StatusCode, resp.Status)
	}

	return sonic.ConfigDefault.NewDecoder(resp.Body).Decode(v)
}

// AppendBlockChildren appends children to a block
func (c *Client) AppendBlockChildren(
	ctx context.Context,
	blockID string,
	children []BlockResponse,
) (*BlockResponse, error) {
	body := map[string]any{
		"children": children,
	}

	var result BlockResponse

	err := c.makeRequest(ctx, http.MethodPatch, "/blocks/"+blockID+"/children", body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// RetrieveBlock retrieves a block
func (c *Client) RetrieveBlock(ctx context.Context, blockID string) (*BlockResponse, error) {
	var result BlockResponse

	err := c.makeRequest(ctx, http.MethodGet, "/blocks/"+blockID, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// RetrieveBlockChildren retrieves children of a block
func (c *Client) RetrieveBlockChildren(
	ctx context.Context,
	blockID string,
	startCursor *string,
	pageSize *int,
) (*ListResponse, error) {
	params := url.Values{}
	if startCursor != nil {
		params.Set("start_cursor", *startCursor)
	}

	if pageSize != nil {
		params.Set("page_size", strconv.Itoa(*pageSize))
	}

	endpoint := "/blocks/" + blockID + "/children"
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	var result ListResponse

	err := c.makeRequest(ctx, http.MethodGet, endpoint, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteBlock deletes a block
func (c *Client) DeleteBlock(ctx context.Context, blockID string) (*BlockResponse, error) {
	var result BlockResponse

	err := c.makeRequest(ctx, http.MethodDelete, "/blocks/"+blockID, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateBlock updates a block
func (c *Client) UpdateBlock(
	ctx context.Context,
	blockID string,
	block BlockResponse,
) (*BlockResponse, error) {
	var result BlockResponse

	err := c.makeRequest(ctx, http.MethodPatch, "/blocks/"+blockID, block, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// RetrievePage retrieves a page
func (c *Client) RetrievePage(ctx context.Context, pageID string) (*PageResponse, error) {
	var result PageResponse

	err := c.makeRequest(ctx, http.MethodGet, "/pages/"+pageID, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdatePageProperties updates page properties
func (c *Client) UpdatePageProperties(
	ctx context.Context,
	pageID string,
	properties map[string]any,
) (*PageResponse, error) {
	body := map[string]any{
		"properties": properties,
	}

	var result PageResponse

	err := c.makeRequest(ctx, http.MethodPatch, "/pages/"+pageID, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ListAllUsers lists all users
func (c *Client) ListAllUsers(
	ctx context.Context,
	startCursor *string,
	pageSize *int,
) (*ListResponse, error) {
	params := url.Values{}
	if startCursor != nil {
		params.Set("start_cursor", *startCursor)
	}

	if pageSize != nil {
		params.Set("page_size", strconv.Itoa(*pageSize))
	}

	endpoint := "/users"
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	var result ListResponse

	err := c.makeRequest(ctx, http.MethodGet, endpoint, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// RetrieveUser retrieves a user
func (c *Client) RetrieveUser(ctx context.Context, userID string) (*UserResponse, error) {
	var result UserResponse

	err := c.makeRequest(ctx, http.MethodGet, "/users/"+userID, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// RetrieveBotUser retrieves the bot user
func (c *Client) RetrieveBotUser(ctx context.Context) (*UserResponse, error) {
	var result UserResponse

	err := c.makeRequest(ctx, http.MethodGet, "/users/me", nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateDatabase creates a database
func (c *Client) CreateDatabase(
	ctx context.Context,
	parent Parent,
	properties map[string]any,
	title []RichTextItemResponse,
) (*DatabaseResponse, error) {
	body := map[string]any{
		"parent":     parent,
		"properties": properties,
	}
	if title != nil {
		body["title"] = title
	}

	var result DatabaseResponse

	err := c.makeRequest(ctx, http.MethodPost, "/databases", body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// QueryDatabase queries a database
func (c *Client) QueryDatabase(
	ctx context.Context,
	databaseID string,
	filter map[string]any,
	sorts []SortObject,
	startCursor *string,
	pageSize *int,
) (*ListResponse, error) {
	body := map[string]any{}
	if filter != nil {
		body["filter"] = filter
	}

	if sorts != nil {
		body["sorts"] = sorts
	}

	if startCursor != nil {
		body["start_cursor"] = *startCursor
	}

	if pageSize != nil {
		body["page_size"] = *pageSize
	}

	var result ListResponse

	err := c.makeRequest(ctx, http.MethodPost, "/databases/"+databaseID+"/query", body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// RetrieveDatabase retrieves a database
func (c *Client) RetrieveDatabase(
	ctx context.Context,
	databaseID string,
) (*DatabaseResponse, error) {
	var result DatabaseResponse

	err := c.makeRequest(ctx, http.MethodGet, "/databases/"+databaseID, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateDatabase updates a database
func (c *Client) UpdateDatabase(
	ctx context.Context,
	databaseID string,
	title, description []RichTextItemResponse,
	properties map[string]any,
) (*DatabaseResponse, error) {
	body := map[string]any{}
	if title != nil {
		body["title"] = title
	}

	if description != nil {
		body["description"] = description
	}

	if properties != nil {
		body["properties"] = properties
	}

	var result DatabaseResponse

	err := c.makeRequest(ctx, http.MethodPatch, "/databases/"+databaseID, body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateDatabaseItem creates a database item
func (c *Client) CreateDatabaseItem(
	ctx context.Context,
	databaseID string,
	properties map[string]any,
) (*PageResponse, error) {
	body := map[string]any{
		"parent": map[string]any{
			"database_id": databaseID,
		},
		"properties": properties,
	}

	var result PageResponse

	err := c.makeRequest(ctx, http.MethodPost, "/pages", body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateComment creates a comment
func (c *Client) CreateComment(
	ctx context.Context,
	parent *CommentParentInput,
	discussionID *string,
	richText []RichTextItemResponse,
) (*CommentResponse, error) {
	body := map[string]any{
		"rich_text": richText,
	}
	if parent != nil {
		body["parent"] = parent
	}

	if discussionID != nil {
		body["discussion_id"] = *discussionID
	}

	var result CommentResponse

	err := c.makeRequest(ctx, http.MethodPost, "/comments", body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// RetrieveComments retrieves comments
func (c *Client) RetrieveComments(
	ctx context.Context,
	blockID string,
	startCursor *string,
	pageSize *int,
) (*ListResponse, error) {
	params := url.Values{}
	params.Set("block_id", blockID)

	if startCursor != nil {
		params.Set("start_cursor", *startCursor)
	}

	if pageSize != nil {
		params.Set("page_size", strconv.Itoa(*pageSize))
	}

	endpoint := "/comments?" + params.Encode()

	var result ListResponse

	err := c.makeRequest(ctx, http.MethodGet, endpoint, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Search searches pages and databases
func (c *Client) Search(
	ctx context.Context,
	query *string,
	filter *SearchFilter,
	sort *SearchSort,
	startCursor *string,
	pageSize *int,
) (*ListResponse, error) {
	body := map[string]any{}
	if query != nil {
		body["query"] = *query
	}

	if filter != nil {
		body["filter"] = filter
	}

	if sort != nil {
		body["sort"] = sort
	}

	if startCursor != nil {
		body["start_cursor"] = *startCursor
	}

	if pageSize != nil {
		body["page_size"] = *pageSize
	}

	var result ListResponse

	err := c.makeRequest(ctx, http.MethodPost, "/search", body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
