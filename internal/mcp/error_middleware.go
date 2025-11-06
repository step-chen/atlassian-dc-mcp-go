package mcp

import (
	"context"

	"atlassian-dc-mcp-go/internal/utils/logging"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"go.uber.org/zap"
)

// ErrorMiddleware creates a lightweight error handling middleware
// This middleware focuses on capturing and logging errors without changing error propagation
func ErrorMiddleware() mcp.Middleware {
	return func(next mcp.MethodHandler) mcp.MethodHandler {
		return func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
			logger := logging.GetLogger()

			// Execute the next handler
			result, err := next(ctx, method, req)

			// If there's an error, log detailed information
			if err != nil {
				// Log structured error information
				logger.Error("Method execution failed",
					zap.String("method", method),
					zap.String("error", err.Error()),
					zap.Any("request", req),
				)

				// In debug mode, log more detailed stack information
				logger.Debug("Full error details",
					zap.String("method", method),
					zap.Error(err),
				)
			}

			// Don't change error propagation, let the MCP SDK handle the error
			return result, err
		}
	}
}
