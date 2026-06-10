package main

import (
	"context"
	"errors"
	"flag"
	"log"

	"github.com/labring/aiproxy/openapi-mcp/convert"
	"github.com/mark3labs/mcp-go/server"
)

var (
	sse  string
	file string
	v2   bool
)

func init() {
	flag.StringVar(&sse, "sse", "", "it will use sse protocol, example: :3000")
	flag.StringVar(&file, "file", "", "openapi file path")
	flag.BoolVar(&v2, "v2", false, "openapi v2 version")
}

func newServer() (*server.MCPServer, error) {
	parser := convert.NewParser()

	var err error
	if v2 {
		err = parser.ParseFileV2(file)
	} else {
		err = parser.ParseFile(file)
	}

	if err != nil {
		return nil, err
	}

	converter := convert.NewConverter(parser, convert.Options{
		OpenAPIFrom: file,
	})

	return converter.Convert()
}

func main() {
	flag.Parse()

	if file == "" {
		log.Fatal("Not provied openapi file")
	}

	s, err := newServer()
	if err != nil {
		log.Fatalf("Failed to new mcp server: %v", err)
	}

	if sse != "" {
		log.Printf("SSE MCP Server Starting")

		err = server.NewSSEServer(s).Start(sse)
	} else {
		err = server.ServeStdio(s)
	}

	if err != nil && !errors.Is(err, context.Canceled) {
		log.Fatalf("Failed to serve MCP: %v", err)
	}
}
