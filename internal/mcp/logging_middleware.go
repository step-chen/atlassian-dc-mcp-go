package mcp

import (
	"context"
	"time"

	"atlassian-dc-mcp-go/internal/utils/logging"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"go.uber.org/zap"
)

// LoggingMiddleware creates a middleware that logs method calls with their duration
// Only logs when execution time exceeds threshold or when an error occurs
func LoggingMiddleware(cfg *logging.Config) mcp.Middleware {
	return func(next mcp.MethodHandler) mcp.MethodHandler {
		return func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
			logger := logging.GetLogger()

			start := time.Now()
			result, err := next(ctx, method, req)
			duration := time.Since(start)

			// Get threshold from config, default to 100ms if not set
			timeThreshold := 100 * time.Millisecond
			if cfg != nil && cfg.LogThreshold > 0 {
				timeThreshold = time.Duration(cfg.LogThreshold) * time.Millisecond
			}

			// For tool calls, extract the actual tool name from the request parameters
			logMethod := method
			if method == "tools/call" {
				// Try to extract tool name from CallToolRequest
				if callToolReq, ok := req.(*mcp.CallToolRequest); ok && callToolReq.Params != nil {
					logMethod = callToolReq.Params.Name
				}
			}

			// Only log if duration exceeds threshold or an error occurred
			if err != nil {
				// Errors are handled by a dedicated error middleware, here we only record execution time information
				logger.Debug("Method failed",
					zap.String("method", logMethod),
					zap.Duration("duration", duration),
					zap.Error(err),
				)
			} else if duration >= timeThreshold {
				logger.Info("Method completed slowly",
					zap.String("method", logMethod),
					zap.Duration("duration", duration),
				)
			}

			return result, err
		}
	}
}
