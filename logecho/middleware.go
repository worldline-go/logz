package logecho

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

// RequestLoggerConfig returns a middleware.RequestLoggerConfig with the given options.
func RequestLoggerConfig(opts ...Option) middleware.RequestLoggerConfig {
	var options options
	for _, opt := range opts {
		opt(&options)
	}

	logger := log.Logger
	if options.loggerUse {
		logger = options.logger
	}

	return middleware.RequestLoggerConfig{
		Skipper: func(c echo.Context) bool {
			path := c.Request().URL.Path

			// Skip ping, health and metrics endpoints for less noise.
			for _, p := range []string{"/ping", "/health", "/metrics"} {
				if strings.HasSuffix(path, p) {
					return true
				}
			}

			return false
		},
		LogRequestID:     true,
		LogRemoteIP:      true,
		LogHost:          true,
		LogMethod:        true,
		LogURI:           true,
		LogUserAgent:     true,
		LogStatus:        true,
		LogError:         true,
		LogLatency:       true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.WithLevel(options.level).
				Str("request_id", v.RequestID).
				Str("remote_ip", v.RemoteIP).
				Str("host", v.Host).
				Str("method", v.Method).
				Str("uri", v.URI).
				Str("user_agent", v.UserAgent).
				Int("status", v.Status).
				Err(v.Error).
				Int64("latency", v.Latency.Nanoseconds()).
				Str("latency_human", v.Latency.String()).
				Str("bytes_in", v.ContentLength).
				Int64("bytes_out", v.ResponseSize).
				Msg("request")

			return nil
		},
	}
}

// ZerologLogger returns a middleware that adds a zerolog logger to the context with request_id.
//
// For request_id, it uses the echo.HeaderXRequestID header (X-Request-Id).
// WithLevel not meaningful here, as the level is set in the log message.
//
// Use after middleware.RequestID().
func ZerologLogger(opts ...Option) echo.MiddlewareFunc {
	var options options
	for _, opt := range opts {
		opt(&options)
	}

	logger := log.Logger
	if options.loggerUse {
		logger = options.logger
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()

			requestLogger := logger.With().Str("request_id", c.Response().Header().Get(echo.HeaderXRequestID)).Logger()

			c.SetRequest(c.Request().WithContext(requestLogger.WithContext(ctx)))

			return next(c)
		}
	}
}
