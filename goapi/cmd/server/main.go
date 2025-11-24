package main

import (
	"fmt"
	"log" // Use standard log for fatal errors before zap is fully configured
	"sms-platform/goapi/internal/api"
	"sms-platform/goapi/internal/config"
	"sms-platform/goapi/internal/global"
	"sms-platform/goapi/internal/repository"

	"go.uber.org/zap"
)

func main() {
	// 1. Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// 2. Initialize Global Logger
	err = global.InitLogger(cfg.Server.Mode)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer global.SyncLogger() // Flushes buffer, if any

	zap.S().Infof("Starting SMS Platform GO API in %s mode...", cfg.Server.Mode)

	// 3. Initialize Database Connection
	err = global.InitDB(cfg.Database)
	if err != nil {
		zap.S().Fatalf("Failed to initialize database: %v", err)
	}
	zap.S().Info("Database connection established.")

	// 4. Initialize Redis Connection (optional)
	if cfg.Redis.Host != "" {
		err = global.InitRedis(cfg.Redis)
		if err != nil {
			zap.S().Warnf("Failed to initialize redis (continuing without redis): %v", err)
		} else {
			zap.S().Info("Redis connection established.")
		}
	}

	// 5. Initialize Provider Manager
	db := global.GetDB()
	providerRepo := repository.NewProviderRepository(db)
	providerBusinessTypeRepo := repository.NewProviderBusinessTypeRepository(db)
	err = global.InitProviderManager(db, providerRepo, providerBusinessTypeRepo)
	if err != nil {
		zap.S().Warnf("Failed to initialize provider manager (continuing): %v", err)
	} else {
		zap.S().Info("Provider manager initialized.")
	}

	// Auto-migrate database models (for development/testing)
	// TODO: Consider a separate migration tool for production
	// err = global.GetDB().AutoMigrate(&domain.Customer{}, &domain.Provider{}) // Add all domain models here
	// if err != nil {
	// 	zap.S().Fatalf("Failed to auto-migrate database: %v", err)
	// }
	// zap.S().Info("Database auto-migration completed.")

	// 6. Create Gin Router
	router := api.NewRouter(cfg)

	// 7. Start Server
	serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)
	zap.S().Infof("Server starting on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		zap.S().Fatalf("Server failed to start: %v", err)
	}
}