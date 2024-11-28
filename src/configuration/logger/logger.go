package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger

	LOG_LEVEL  = "LOG_LEVEL"
	LOG_OUTPUT = "LOG_OUTPUT"
)

func init() {

	logConfig := zap.Config{
		OutputPaths: []string{getOutputLogs()},
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build()
}

func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Info(message, tags...)
	log.Sync()
}

func getOutputLogs() string {

	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL)))

	if output == "" {
		return "stdout"
	}

	return output
}

func getLevelLogs() zapcore.Level {

	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT))) {
	case "info":
		return zap.InfoLevel

	case "debug":
		return zap.DebugLevel

	case "error":
		return zap.ErrorLevel

	default:
		return zap.InfoLevel
	}
}
