package utils

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/donnyirianto/be-stock-app/config"
)

var logger *zap.Logger

func InitLogger() {

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.MkdirAll("logs", 0755)
	}

	// config log rotation
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    20,
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   true,
	}

	fileWriteSyncer := zapcore.AddSync(lumberjackLogger)

	consoleWriteSyncer := zapcore.AddSync(zapcore.Lock(os.Stdout))

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.LineEnding = zapcore.DefaultLineEnding

	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewJSONEncoder(encoderConfig)

	fileCore := zapcore.NewCore(fileEncoder, fileWriteSyncer, zapcore.InfoLevel)
	consoleCore := zapcore.NewCore(consoleEncoder, consoleWriteSyncer, zapcore.InfoLevel)

	core := zapcore.NewTee(fileCore, consoleCore)

	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

}

func GetLogger() *zap.Logger {
	return logger
}

func LoggerMiddleware(logger *zap.Logger, cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Log request
		logRequest(c, cfg.App.Version)

		if err := c.Next(); err != nil {
			logger.Error("Error processing request", zap.Error(err))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":    500,
				"status":  "error",
				"message": "Internal Server Error",
			})
		}

		// Log response
		logResponse(c, start, cfg.App.Version)

		return nil
	}
}

func logRequest(c *fiber.Ctx, version string) {
	logger.Info("Request received",
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
		zap.String("ip", c.IP()),
		zap.String("user_agent", string(c.Request().Header.UserAgent())),
		zap.String("version", version),
		zap.Time("timestamp", time.Now()),
	)
}

func logResponse(c *fiber.Ctx, start time.Time, version string) {
	duration := time.Since(start)
	statusCode := c.Response().StatusCode()

	logger.Info("Response sent",
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
		zap.Int("status", statusCode),
		zap.String("duration", duration.String()),
		zap.String("version", version),
		zap.Time("timestamp", time.Now()),
	)
}
