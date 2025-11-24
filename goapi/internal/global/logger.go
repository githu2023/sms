package global

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

// InitLogger initializes the global logger.
func InitLogger(mode string) error {
	var err error
	if mode == "release" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		return err
	}

	// 设置为全局logger，这样任何地方都可以使用 zap.S() 调用
	zap.ReplaceGlobals(logger)
	return nil
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

