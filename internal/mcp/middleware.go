package mcp

import (
	"context"
	"time"

	"atlassian-dc-mcp-go/internal/utils/logging"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"go.uber.org/zap"
)

// LoggingMiddleware creates a middleware that logs method calls with their duration
func LoggingMiddleware() mcp.Middleware {
	return func(next mcp.MethodHandler) mcp.MethodHandler {
		return func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
			logger := logging.GetLogger()

			start := time.Now()

			logger.Info("Starting method",
				zap.String("method", method),
			)

			result, err := next(ctx, method, req)

			duration := time.Since(start)
			if err != nil {
				logger.Error("Method failed",
					zap.String("method", method),
					zap.Duration("duration", duration),
					zap.Error(err),
				)
			} else {
				logger.Info("Method completed",
					zap.String("method", method),
					zap.Duration("duration", duration),
				)
			}

			return result, err
		}
	}
}
