package config

import (
	"github.com/spf13/viper"
)

// Config holds all configuration for the application.
type Config struct {
	Server    ServerConfig    `mapstructure:"server"`
	Database  DatabaseConfig  `mapstructure:"database"`
	Redis     RedisConfig     `mapstructure:"redis"`
	JWT       JWTConfig       `mapstructure:"jwt"`
	Scheduler SchedulerConfig `mapstructure:"scheduler"`
	Logging   LoggingConfig   `mapstructure:"logging"`
}

// JWTConfig holds JWT specific configuration.
type JWTConfig struct {
	Secret            string `mapstructure:"secret"`
	ClientTokenExpiry int    `mapstructure:"client_token_expiry"` // 客户端JWT过期时间(秒)
	APITokenExpiry    int    `mapstructure:"api_token_expiry"`    // API Token过期时间(秒)
}

// ServerConfig holds server specific configuration.
type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

// DatabaseConfig holds database connection configuration.
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

// RedisConfig holds redis connection configuration.
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// SchedulerConfig holds task scheduler configuration.
type SchedulerConfig struct {
	CodeCheckInterval         int `mapstructure:"code_check_interval"`         // 验证码检查间隔 (秒)
	CodeTimeout               int `mapstructure:"code_timeout"`                // 验证码获取超时时间 (秒)
	AssignmentCleanupInterval int `mapstructure:"assignment_cleanup_interval"` // 过期分配清理间隔 (秒)
}

// LoggingConfig holds logging related configuration
type LoggingConfig struct {
	FilePath   string `mapstructure:"file"`
	Level      string `mapstructure:"level"`
	MaxSize    int    `mapstructure:"max_size"`    // megabytes
	MaxBackups int    `mapstructure:"max_backups"` // number of files
	MaxAge     int    `mapstructure:"max_age"`     // days
	Compress   bool   `mapstructure:"compress"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
