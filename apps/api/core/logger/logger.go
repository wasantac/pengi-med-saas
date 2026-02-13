package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func Init(env string) {
	var config zap.Config

	if env == "production" {
		config = zap.NewProductionConfig()
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	var err error
	Log, err = config.Build()
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	if Log == nil {
		// Fallback implementation if logger is not initialized (e.g. during tests)
		// Or panic if strict init is required. For now, fallback to default console.
		// However, safest to assume Init is called.
		// If not, we could panic or just log to stdout using fmt.
		// Let's assume Init is called as per plan.
		panic("Logger not initialized. Call logger.Init() first.")
	}
	Log.Info(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	if Log == nil {
		panic("Logger not initialized. Call logger.Init() first.")
	}
	Log.Error(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	if Log == nil {
		panic("Logger not initialized. Call logger.Init() first.")
	}
	Log.Debug(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	if Log == nil {
		panic("Logger not initialized. Call logger.Init() first.")
	}
	Log.Warn(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	if Log == nil {
		panic("Logger not initialized. Call logger.Init() first.")
	}
	Log.Fatal(message, fields...)
}
