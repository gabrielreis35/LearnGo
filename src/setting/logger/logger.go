package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var (
	log        *zap.Logger
	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL  = "LOG_LEVEL"
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
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	log, err := logConfig.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	log.Info("Initializing the system")
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
	log.Sync()
}

func Warn(msg string, fields ...zap.Field) {
	log.Warn(msg, fields...)
	log.Sync()
}

func Error(msg string, err error, fields ...zap.Field) {
	fields = append(fields, zap.NamedError("error", err))
	log.Error(msg, fields...)
	log.Sync()
}

func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))

	if output == "" {
		return "stdout"
	}
	return output
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "debug":
		return zapcore.DebugLevel
	case "error":
		return zapcore.ErrorLevel
	}
	return zapcore.InfoLevel
}
