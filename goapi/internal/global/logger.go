package global

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"sms-platform/goapi/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

// InitLogger initializes the global logger.
func InitLogger(mode string, logCfg config.LoggingConfig) error {
	level := resolveLogLevel(mode, logCfg.Level)

	encoderCfg := zap.NewProductionEncoderConfig()
	if mode != "release" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	}
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level)

	cores := []zapcore.Core{consoleCore}

	if logCfg.FilePath != "" {
		if err := ensureLogDir(logCfg.FilePath); err != nil {
			return fmt.Errorf("create log directory: %w", err)
		}

		fileEncoder := zapcore.NewJSONEncoder(encoderCfg)
		fileCore := zapcore.NewCore(
			fileEncoder,
			zapcore.AddSync(newRollingLogger(logCfg)),
			level,
		)
		cores = append(cores, fileCore)
	}

	logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	zap.ReplaceGlobals(logger)
	return nil
}

func resolveLogLevel(mode, cfgLevel string) zapcore.Level {
	switch strings.ToLower(cfgLevel) {
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	case "info":
		return zapcore.InfoLevel
	}

	if strings.ToLower(mode) == "debug" {
		return zapcore.DebugLevel
	}
	return zapcore.InfoLevel
}

func ensureLogDir(path string) error {
	dir := filepath.Dir(path)
	return os.MkdirAll(dir, 0o755)
}

func newRollingLogger(cfg config.LoggingConfig) *lumberjack.Logger {
	maxSize := cfg.MaxSize
	if maxSize == 0 {
		maxSize = 20
	}
	maxBackups := cfg.MaxBackups
	if maxBackups == 0 {
		maxBackups = 5
	}
	maxAge := cfg.MaxAge
	if maxAge == 0 {
		maxAge = 7
	}

	return &lumberjack.Logger{
		Filename:   cfg.FilePath,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   cfg.Compress,
	}
}

// GetLogger returns the global logger instance.
// 任何地方都可以直接调用 global.GetLogger() 获取logger
// 或者直接使用 zap.S() 调用全局logger
func GetLogger() *zap.Logger {
	return logger
}

// SyncLogger flushes any buffered log entries.
func SyncLogger() {
	if logger != nil {
		logger.Sync()
	}
}

// LogInfo 记录信息日志
func LogInfo(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Info(msg, fields...)
	}
}

// LogWarn 记录警告日志
func LogWarn(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Warn(msg, fields...)
	}
}

// LogError 记录错误日志
func LogError(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Error(msg, fields...)
	}
}

// LogDebug 记录调试日志
func LogDebug(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Debug(msg, fields...)
	}
}

// LogFatal 记录致命错误日志并退出
func LogFatal(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Fatal(msg, fields...)
	}
}

// S 返回全局的SugaredLogger（便捷方法，兼容zap.S()的使用方式）
func S() *zap.SugaredLogger {
	if logger != nil {
		return logger.Sugar()
	}
	// 如果logger未初始化，返回一个no-op logger
	return zap.NewNop().Sugar()
}
