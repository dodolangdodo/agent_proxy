package bingcn

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type SearchResult struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Link    string `json:"link"`
	Snippet string `json:"snippet"`
}

const (
	DefaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
)

// SearchEngine represents a Bing search engine
type SearchEngine struct {
	userAgent     string
	client        *http.Client
	searchResults sync.Map // map[string]*SearchResult
}

// SearchOptions contains options for search
type SearchOptions struct {
	Query      string
	NumResults int
	Language   string
}

// NewSearchEngine creates a new search engine instance
func NewSearchEngine(userAgent string, timeout time.Duration) *SearchEngine {
	if userAgent == "" {
		userAgent = DefaultUserAgent
	}

	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
				MinVersion:         tls.VersionTLS12,
			},
		},
	}

	return &SearchEngine{
		userAgent: userAgent,
		client:    client,
	}
}

// Search performs a Bing search and returns results
func (e *SearchEngine) Search(ctx context.Context, options SearchOptions) ([]*SearchResult, error) {
	if options.Query == "" {
		return nil, errors.New("query cannot be empty")
	}

	if options.NumResults <= 0 {
		options.NumResults = 5
	}

	// Build search URL
	searchURL := e.buildSearchURL(options.Query, options.Language)

	// Create and execute request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, searchURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	e.setSearchHeaders(req)

	resp, err := e.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search request failed with status: %d", resp.StatusCode)
	}

	// Read and decode response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	content := e.decodeContent(body, resp.Header.Get("Content-Type"))

	// Parse HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Extract results
	return e.extractSearchResults(doc, options.Query, options.NumResults), nil
}

// GetSearchResult retrieves a stored search result by ID
func (e *SearchEngine) GetSearchResult(resultID string) (*SearchResult, bool) {
	value, ok := e.searchResults.Load(resultID)
	if !ok {
		return nil, false
	}

	result, ok := value.(*SearchResult)

	return result, ok
}

// buildSearchURL constructs the Bing search URL
func (e *SearchEngine) buildSearchURL(query, language string) string {
	baseURL := "https://cn.bing.com/search"
	params := url.Values{}
	params.Set("q", query)
	params.Set("setlang", "zh-CN")
	params.Set("ensearch", "0")

	if language != "" {
		params.Set("setlang", language)
	}

	return fmt.Sprintf("%s?%s", baseURL, params.Encode())
}

// setSearchHeaders sets appropriate headers for search requests
func (e *SearchEngine) setSearchHeaders(req *http.Request) {
	headers := map[string]string{
		"User-Agent":                e.userAgent,
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8",
		"Accept-Language":           "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":             "no-cache",
		"Pragma":                    "no-cache",
		"Sec-Fetch-Dest":            "document",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-Site":            "none",
		"Sec-Fetch-User":            "?1",
		"Upgrade-Insecure-Requests": "1",
		"Cookie":                    "SRCHHPGUSR=SRCHLANG=zh-Hans; _EDGE_S=ui=zh-cn; _EDGE_V=1",
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

// decodeContent attempts to properly decode content based on encoding
func (e *SearchEngine) decodeContent(body []byte, contentType string) string {
	content := string(body)

	// Check if content type suggests GBK encoding
	if strings.Contains(strings.ToLower(contentType), "gbk") ||
		strings.Contains(strings.ToLower(contentType), "gb2312") {
		if decoded, err := e.decodeGBK(body); err == nil {
			content = decoded
		}
	}

	return content
}

// decodeGBK decodes GBK encoded content to UTF-8
func (e *SearchEngine) decodeGBK(data []byte) (string, error) {
	reader := transform.NewReader(
		strings.NewReader(string(data)),
		simplifiedchinese.GBK.NewDecoder(),
	)

	decoded, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}

// extractSearchResults extracts search results from parsed HTML
func (e *SearchEngine) extractSearchResults(
	doc *goquery.Document,
	query string,
	numResults int,
) []*SearchResult {
	var results []*SearchResult

	// Try different selectors for Bing search results
	selectors := []string{
		"#b_results > li.b_algo",
		"#b_results > .b_ans",
		"#b_results > li",
	}

	for _, selector := range selectors {
		doc.Find(selector).Each(func(i int, element *goquery.Selection) {
			if len(results) >= numResults {
				return
			}

			result := e.parseSearchResultElement(element, i)
			if result != nil {
				// Store result for later retrieval
				e.searchResults.Store(result.ID, result)
				results = append(results, result)
			}
		})

		// If we found results with this selector, stop trying others
		if len(results) > 0 {
			break
		}
	}

	// If no results found, create a fallback result
	if len(results) == 0 {
		fallbackResult := e.createFallbackResult(query)
		e.searchResults.Store(fallbackResult.ID, fallbackResult)
		results = append(results, fallbackResult)
	}

	return results
}

// parseSearchResultElement parses a single search result element
func (e *SearchEngine) parseSearchResultElement(
	element *goquery.Selection,
	index int,
) *SearchResult {
	// Skip ads
	if element.HasClass("b_ad") {
		return nil
	}

	title, link := e.extractTitleAndLink(element)
	snippet := e.extractSnippet(element, title)

	// Fix incomplete links
	if link != "" && !strings.HasPrefix(link, "http") {
		link = e.fixIncompleteLink(link)
	}

	// Skip if no meaningful content
	if title == "" && snippet == "" {
		return nil
	}

	// Create unique ID
	id := fmt.Sprintf("result_%d_%d", time.Now().UnixNano(), index)

	return &SearchResult{
		ID:      id,
		Title:   title,
		Link:    link,
		Snippet: snippet,
	}
}

// extractTitleAndLink extracts title and link from a search result element
func (e *SearchEngine) extractTitleAndLink(element *goquery.Selection) (string, string) {
	// Try to find title and link in h2 a
	titleElement := element.Find("h2 a").First()
	if titleElement.Length() > 0 {
		title := strings.TrimSpace(titleElement.Text())
		link, _ := titleElement.Attr("href")
		return title, link
	}

	// Try alternative selectors
	altTitleElement := element.Find(".b_title a, a.tilk, a strong").First()
	if altTitleElement.Length() > 0 {
		title := strings.TrimSpace(altTitleElement.Text())
		link, _ := altTitleElement.Attr("href")
		return title, link
	}

	return "", ""
}

// extractSnippet extracts snippet from a search result element
func (e *SearchEngine) extractSnippet(element *goquery.Selection, title string) string {
	// Try to find snippet in common Bing snippet selectors
	snippetElement := element.Find(".b_caption p, .b_snippet, .b_algoSlug").First()
	if snippetElement.Length() > 0 {
		return strings.TrimSpace(snippetElement.Text())
	}

	// If no snippet found, use entire element text and clean it up
	snippet := strings.TrimSpace(element.Text())

	// Remove title from snippet
	if title != "" && strings.Contains(snippet, title) {
		snippet = strings.ReplaceAll(snippet, title, "")
		snippet = strings.TrimSpace(snippet)
	}

	// Truncate if too long
	if len(snippet) > 150 {
		snippet = snippet[:150] + "..."
	}

	return snippet
}

// fixIncompleteLink fixes incomplete URLs
func (e *SearchEngine) fixIncompleteLink(link string) string {
	if strings.HasPrefix(link, "/") {
		return "https://cn.bing.com" + link
	}
	return "https://cn.bing.com/" + link
}

// createFallbackResult creates a fallback result when no results are found
func (e *SearchEngine) createFallbackResult(query string) *SearchResult {
	id := fmt.Sprintf("result_%d_fallback", time.Now().UnixNano())

	return &SearchResult{
		ID:      id,
		Title:   "搜索结果: " + query,
		Link:    "https://cn.bing.com/search?q=" + url.QueryEscape(query),
		Snippet: fmt.Sprintf("未能解析关于 \"%s\" 的搜索结果，但您可以直接访问必应搜索页面查看。", query),
	}
}
