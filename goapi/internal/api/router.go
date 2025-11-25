package api

import (
	"sms-platform/goapi/internal/api/handler"
	"sms-platform/goapi/internal/api/middleware"
	"sms-platform/goapi/internal/config"
	"sms-platform/goapi/internal/global"
	"sms-platform/goapi/internal/repository"
	"sms-platform/goapi/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter creates and configures a new Gin router.
func NewRouter(cfg config.Config) *gin.Engine { // Pass config to router
	// Set Gin mode from config
	gin.SetMode(cfg.Server.Mode)
	router := gin.New()

	// Global Middleware
	router.Use(gin.Logger())   // Standard logger
	router.Use(gin.Recovery()) // Recover from panics
	router.Use(middleware.RequestID())
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
	}))

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	// --- Initialize Repositories ---
	db := global.GetDB() // GetDB() is safe to call here as InitDB has been called in main.go
	customerRepo := repository.NewCustomerRepository(db)
	businessTypeRepo := repository.NewBusinessTypeRepository(db)
	transactionRepo := repository.NewTransactionRepository(db, customerRepo)
	logRepo := repository.NewLogRepository(db)
	providerRepo := repository.NewProviderRepository(db)
	assignmentRepo := repository.NewPhoneAssignmentRepository(db)
	whitelistRepo := repository.NewWhitelistRepository(db)
	customerBusinessConfigRepo := repository.NewCustomerBusinessConfigRepository(db)
	platformBusinessTypeRepo := repository.NewPlatformBusinessTypeRepository(db)
	platformProviderBusinessMappingRepo := repository.NewPlatformProviderBusinessMappingRepository(db)
	providerBusinessTypeRepo := repository.NewProviderBusinessTypeRepository(db)

	// --- Initialize Services ---
	userService := service.NewUserService(customerRepo, cfg.JWT)
	businessService := service.NewBusinessService(businessTypeRepo, customerBusinessConfigRepo)
	transactionService := service.NewTransactionService(transactionRepo)

	phoneService := service.NewPhoneService(
		transactionRepo,
		logRepo,
		assignmentRepo,
		customerBusinessConfigRepo,
		businessTypeRepo,
		platformBusinessTypeRepo,
		platformProviderBusinessMappingRepo,
		providerBusinessTypeRepo,
		customerRepo,
		db,
	)
	assignmentService := service.NewAssignmentService(assignmentRepo, businessTypeRepo, providerRepo)
	whitelistService := service.NewWhitelistService(whitelistRepo)

	// Initialize and start scheduler service
	// 使用全局ProviderManager，通过ProviderID查找对应的provider
	schedulerService := service.NewSchedulerService(cfg.Scheduler, assignmentRepo, providerRepo, transactionRepo, customerRepo, db)
	schedulerService.Start() // --- Initialize Handlers ---
	userHandler := handler.NewUserHandler(userService, cfg.JWT)
	businessHandler := handler.NewBusinessHandler(businessService)
	authHandler := handler.NewAuthHandler(userService, cfg.JWT)
	balanceHandler := handler.NewBalanceHandler(transactionService)
	phoneHandler := handler.NewPhoneHandler(phoneService)
	assignmentHandler := handler.NewAssignmentHandler(assignmentService)
	whitelistHandler := handler.NewWhitelistHandler(whitelistService)

	// --- API Groups ---

	// Client API Group (for Flutter app)
	clientV1 := router.Group("/client/v1")
	{
		clientV1.POST("/register", userHandler.Register)
		clientV1.POST("/login", userHandler.Login)

		// Authenticated routes for clientV1
		clientAuth := clientV1.Group("/")
		clientAuth.Use(middleware.JWTAuthMiddleware(cfg.JWT.Secret))
		{
			clientAuth.GET("/business_types", businessHandler.GetBusinessTypes)
			clientAuth.GET("/profile", userHandler.GetProfile)
			clientAuth.POST("/change_password", userHandler.UpdatePassword) // New
			clientAuth.GET("/balance", balanceHandler.GetBalance)           // New
			clientAuth.POST("/get_phone", phoneHandler.GetPhone)            // New
			clientAuth.POST("/get_code", phoneHandler.GetCode)              // New
			clientAuth.GET("/phone_status", phoneHandler.GetPhoneStatus)    // New
			clientAuth.GET("/assignments", assignmentHandler.GetAssignments)
			clientAuth.GET("/assignments/recent", assignmentHandler.GetRecentAssignments)
			clientAuth.GET("/assignments/statistics", assignmentHandler.GetCostStatistics)
			// 白名单相关接口
			clientAuth.GET("/whitelist", whitelistHandler.ListWhitelists)
			clientAuth.POST("/whitelist", whitelistHandler.AddWhitelist)
			clientAuth.DELETE("/whitelist", whitelistHandler.DeleteWhitelist)
		}
	}

	// Programmatic API Group (for other apps/scripts)
	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/get_token", authHandler.GetAPIToken) // New

		// Authenticated routes for apiV1
		apiAuth := apiV1.Group("/")
		apiAuth.Use(middleware.APITokenAuthMiddleware(cfg.JWT.Secret)) // New
		{
			apiAuth.GET("/business_types", businessHandler.GetBusinessTypes) // This route should also be authenticated
			apiAuth.GET("/balance", balanceHandler.GetBalance)               // New
			apiAuth.POST("/get_phone", phoneHandler.GetPhone)                // New
			apiAuth.POST("/get_code", phoneHandler.GetCode)                  // New
			apiAuth.GET("/phone_status", phoneHandler.GetPhoneStatus)        // New
			apiAuth.GET("/assignments", assignmentHandler.GetAssignments)
			apiAuth.GET("/assignments/recent", assignmentHandler.GetRecentAssignments)
			apiAuth.GET("/assignments/statistics", assignmentHandler.GetCostStatistics)
			// 白名单相关接口
			apiAuth.GET("/whitelist", whitelistHandler.ListWhitelists)
			apiAuth.POST("/whitelist", whitelistHandler.AddWhitelist)
			apiAuth.DELETE("/whitelist", whitelistHandler.DeleteWhitelist)
		}
	}

	return router
}
