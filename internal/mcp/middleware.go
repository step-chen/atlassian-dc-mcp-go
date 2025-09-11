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
func LoggingMiddleware() mcp.Middleware {
	// Threshold for logging - only log methods that take longer than this
	const timeThreshold = 100 * time.Millisecond
	
	return func(next mcp.MethodHandler) mcp.MethodHandler {
		return func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
			logger := logging.GetLogger()

			start := time.Now()
			result, err := next(ctx, method, req)
			duration := time.Since(start)

			// Only log if duration exceeds threshold or an error occurred
			if err != nil {
				logger.Error("Method failed",
					zap.String("method", method),
					zap.Duration("duration", duration),
					zap.Error(err),
				)
			} else if duration >= timeThreshold {
				logger.Info("Method completed slowly",
					zap.String("method", method),
					zap.Duration("duration", duration),
				)
			}

			return result, err
		}
	}
}