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

