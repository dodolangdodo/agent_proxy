package mcpservers

import (
	"context"
	"errors"

	"github.com/bytedance/sonic"
	"github.com/mark3labs/mcp-go/mcp"
)

func ListServerTools(ctx context.Context, server Server) ([]mcp.Tool, error) {
	requestBytes, err := sonic.Marshal(mcp.JSONRPCRequest{
		JSONRPC: mcp.JSONRPC_VERSION,
		ID:      mcp.NewRequestId(1),
		Request: mcp.Request{
			Method: string(mcp.MethodToolsList),
			Params: mcp.RequestParams{},
		},
	})
	if err != nil {
		return nil, err
	}

	response := server.HandleMessage(ctx, requestBytes)
	if response == nil {
		return nil, errors.New("no response from server")
	}

	responseBytes, err := sonic.Marshal(response)
	if err != nil {
		return nil, err
	}

	var jsonRPCResponse struct {
		Result *mcp.ListToolsResult `json:"result"`
		Error  *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	if err := sonic.Unmarshal(responseBytes, &jsonRPCResponse); err != nil {
		return nil, err
	}

	if jsonRPCResponse.Error != nil {
		return nil, errors.New(jsonRPCResponse.Error.Message)
	}

	if jsonRPCResponse.Result == nil {
		return nil, nil
	}

	return jsonRPCResponse.Result.Tools, nil
}
