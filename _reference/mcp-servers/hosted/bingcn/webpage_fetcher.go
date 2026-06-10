package bingcn

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// WebpageFetcher handles fetching and extracting content from webpages
type WebpageFetcher struct {
	userAgent string
	client    *http.Client
}

// FetchOptions contains options for webpage fetching
type FetchOptions struct {
	URL         string
	MaxLength   int
	ExtractText bool
}

// WebpageContent represents extracted webpage content
type WebpageContent struct {
	Title   string
	Content string
	URL     string
	Length  int
}

// NewWebpageFetcher creates a new webpage fetcher
func NewWebpageFetcher(userAgent string, client *http.Client) *WebpageFetcher {
	if userAgent == "" {
		userAgent = DefaultUserAgent
	}

	return &WebpageFetcher{
		userAgent: userAgent,
		client:    client,
	}
}

// FetchWebpage fetches and extracts content from a webpage
func (f *WebpageFetcher) FetchWebpage(
	ctx context.Context,
	options FetchOptions,
) (*WebpageContent, error) {
	if options.URL == "" {
		return nil, errors.New("URL cannot be empty")
	}

	if options.MaxLength <= 0 {
		options.MaxLength = 8000
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, options.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	f.setWebpageHeaders(req)

	// Execute request
	resp, err := f.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch webpage: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("webpage request failed with status: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Decode content
	content := f.decodeContent(body, resp.Header.Get("Content-Type"))

	// Extract webpage content
	webpageContent, err := f.extractContent(content, options)
	if err != nil {
		return nil, fmt.Errorf("failed to extract content: %w", err)
	}

	webpageContent.URL = options.URL

	return webpageContent, nil
}

// FetchWebpageByResult fetches webpage content using a search result
func (f *WebpageFetcher) FetchWebpageByResult(
	ctx context.Context,
	result *SearchResult,
	maxLength int,
) (*WebpageContent, error) {
	options := FetchOptions{
		URL:         result.Link,
		MaxLength:   maxLength,
		ExtractText: true,
	}

	return f.FetchWebpage(ctx, options)
}

// setWebpageHeaders sets appropriate headers for webpage requests
func (f *WebpageFetcher) setWebpageHeaders(req *http.Request) {
	headers := map[string]string{
		"User-Agent":      f.userAgent,
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":   "no-cache",
		"Pragma":          "no-cache",
		"Referer":         "https://cn.bing.com/",
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

// decodeContent decodes content with proper encoding handling
func (f *WebpageFetcher) decodeContent(body []byte, contentType string) string {
	content := string(body)

	// Check if content type suggests GBK encoding
	if strings.Contains(strings.ToLower(contentType), "gbk") ||
		strings.Contains(strings.ToLower(contentType), "gb2312") {
		// Use the same decoding logic as SearchEngine
		engine := &SearchEngine{}
		if decoded, err := engine.decodeGBK(body); err == nil {
			content = decoded
		}
	}

	return content
}

// extractContent extracts and cleans the main content from HTML
func (f *WebpageFetcher) extractContent(
	htmlContent string,
	options FetchOptions,
) (*WebpageContent, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Remove unwanted elements
	f.removeUnwantedElements(doc)

	// Extract title
	title := strings.TrimSpace(doc.Find("title").Text())

	var content string
	if options.ExtractText {
		content = f.extractMainContent(doc)
	} else {
		content = htmlContent
	}

	// Clean up the text
	content = f.cleanText(content)

	// Add title if available and extracting text
	if title != "" && options.ExtractText {
		content = fmt.Sprintf("标题: %s\n\n%s", title, content)
	}

	// Truncate if too long
	if len(content) > options.MaxLength {
		content = content[:options.MaxLength] + "... (内容已截断)"
	}

	return &WebpageContent{
		Title:   title,
		Content: content,
		Length:  len(content),
	}, nil
}

// removeUnwantedElements removes unwanted HTML elements
func (f *WebpageFetcher) removeUnwantedElements(doc *goquery.Document) {
	unwantedSelectors := []string{
		"script", "style", "iframe", "noscript", "nav", "header", "footer",
		".header", ".footer", ".nav", ".sidebar", ".ad", ".advertisement",
		"#header", "#footer", "#nav", "#sidebar",
	}

	for _, selector := range unwantedSelectors {
		doc.Find(selector).Remove()
	}
}

// extractMainContent extracts the main content from the document
func (f *WebpageFetcher) extractMainContent(doc *goquery.Document) string {
	var content string

	// Try to find main content areas
	mainSelectors := []string{
		"main", "article", ".article", ".post", ".content", "#content",
		".main", "#main", ".body", "#body", ".entry", ".entry-content",
		".post-content", ".article-content", ".text", ".detail",
	}

	for _, selector := range mainSelectors {
		mainElement := doc.Find(selector)
		if mainElement.Length() > 0 {
			content = strings.TrimSpace(mainElement.Text())
			if len(content) > 100 {
				return content
			}
		}
	}

	// If no main content found, try paragraphs
	if content == "" || len(content) < 100 {
		content = f.extractParagraphs(doc)
	}

	// If still no content, get body content
	if content == "" || len(content) < 100 {
		content = strings.TrimSpace(doc.Find("body").Text())
	}

	return content
}

// extractParagraphs extracts meaningful paragraphs from the document
func (f *WebpageFetcher) extractParagraphs(doc *goquery.Document) string {
	var paragraphs []string
	doc.Find("p").Each(func(_ int, element *goquery.Selection) {
		text := strings.TrimSpace(element.Text())
		if len(text) > 20 {
			paragraphs = append(paragraphs, text)
		}
	})

	if len(paragraphs) > 0 {
		return strings.Join(paragraphs, "\n\n")
	}

	return ""
}

// cleanText cleans and normalizes extracted text
func (f *WebpageFetcher) cleanText(text string) string {
	// Replace newlines and tabs with spaces
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\t", " ")

	// Replace multiple spaces with single space
	for strings.Contains(text, "  ") {
		text = strings.ReplaceAll(text, "  ", " ")
	}

	return strings.TrimSpace(text)
}
