package engine

import (
	"context"
)

// SearchResult represents a single search result
type SearchResult struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Content string `json:"content"`
}

// SearchQuery represents search parameters
type SearchQuery struct {
	Queries       []string
	MaxResults    int
	Language      string
	ArxivCategory string
}

// Engine interface for search engines
type Engine interface {
	Search(ctx context.Context, query SearchQuery) ([]SearchResult, error)
}
