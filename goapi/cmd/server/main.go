package main

import (
	"fmt"
	"log" // Use standard log for fatal errors before zap is fully configured
	"sms-platform/goapi/internal/api"
	"sms-platform/goapi/internal/config"
	"sms-platform/goapi/internal/repository"

	"go.uber.org/zap"
)

func main() {
	// 1. Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// 2. Initialize Logger
	var logger *zap.Logger
	if cfg.Server.Mode == "release" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync() // Flushes buffer, if any
	zap.ReplaceGlobals(logger) // Set as global logger

	zap.S().Infof("Starting SMS Platform GO API in %s mode...", cfg.Server.Mode)

	// 3. Initialize Database Connection
	err = repository.InitDB(cfg.Database)
	if err != nil {
		zap.S().Fatalf("Failed to initialize database: %v", err)
	}
	zap.S().Info("Database connection established.")

	// Auto-migrate database models (for development/testing)
	// TODO: Consider a separate migration tool for production
	// err = repository.GetDB().AutoMigrate(&domain.Customer{}, &domain.Provider{}) // Add all domain models here
	// if err != nil {
	// 	zap.S().Fatalf("Failed to auto-migrate database: %v", err)
	// }
	// zap.S().Info("Database auto-migration completed.")

	// 4. Create Gin Router
	router := api.NewRouter(cfg)

	// 5. Start Server
	serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)
	zap.S().Infof("Server starting on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		zap.S().Fatalf("Server failed to start: %v", err)
	}
}