package main

import (
	"github.com/hawk-eye03/kafka-poc/lib/config"
	"github.com/hawk-eye03/kafka-poc/lib/consumers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initLogger() *zap.Logger {
	loggerConfig := zap.Config{
		Encoding:    "console",
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel), // Set your desired log level
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:    "level",
			TimeKey:     "time",
			MessageKey:  "msg",
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		},
	}

	// Replace the global logger instance with the configured logger
	logger, err := loggerConfig.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	return logger
}
func main() {
	// initialise logger
	logger := initLogger()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	// Load config for DB, App and Kafka
	config := config.LoadConfig()

	switch config.App.Mode {
	case "DAILY_SUMMARY_CONSUMER":
		zap.L().Info("Started Daily Summary Consumer")
		dailySummaryConsumer := consumers.NewDailySummaryConsumer(config)
		baseConsumer := consumers.NewBaseConsumer(dailySummaryConsumer, *config)
		baseConsumer.StartConsumer()

	default:
		zap.L().Info("No valid mode found. Exiting...")
	}
}
