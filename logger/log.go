// logger/logger.go
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Global logger variable
var Logger *zap.Logger

// Init initializes the logger
func Init() error {
	var err error
	cfg := zap.NewProductionConfig()

	// Customize configuration for local logging
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)       // Log everything including debug logs
	cfg.OutputPaths = []string{"stdout", "logs/local.log"} // Log to console and a local file
	cfg.ErrorOutputPaths = []string{"stderr"}              // Log errors to standard error
	cfg.Encoding = "console"                               // Use console-friendly formatting instead of JSON for local debugging

	// Add additional customizations if needed
	cfg.EncoderConfig.TimeKey = "timestamp"                   // Customize the timestamp field name
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Use human-readable timestamps
	cfg.EncoderConfig.StacktraceKey = "stacktrace"            // Enable stack traces for errors

	// Build the logger with the customized config
	logger, err := cfg.Build()
	if err != nil {
		return err
	}
	Logger = logger

	return nil
}

// Sync flushes any buffered log entries
func Sync() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}
