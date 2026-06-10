package engine

import (
	"context"
	"strings"
	"time"

	"github.com/labring/aiproxy/mcp-servers/hosted/bingcn"
)

type BingCNEngine struct {
	bingcn *bingcn.SearchEngine
}

func NewBingCNEngine() *BingCNEngine {
	return &BingCNEngine{
		bingcn: bingcn.NewSearchEngine("", 10*time.Second),
	}
}

func (e *BingCNEngine) Search(
	ctx context.Context,
	query SearchQuery,
) ([]SearchResult, error) {
	options := bingcn.SearchOptions{
		Query:      strings.Join(query.Queries, " "),
		NumResults: query.MaxResults,
		Language:   query.Language,
	}

	results, err := e.bingcn.Search(ctx, options)
	if err != nil {
		return nil, err
	}

	searchResults := make([]SearchResult, 0, len(results))
	for _, result := range results {
		searchResults = append(searchResults, SearchResult{
			Title:   result.Title,
			Link:    result.Link,
			Content: result.Snippet,
		})
	}

	return searchResults, nil
}
