// Package bootstrap provides application startup functionality.
package bootstrap

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	aiEngine "github.com/mzakiaklhairi/velora/internal/ai/engine"
	aiHandler "github.com/mzakiaklhairi/velora/internal/ai/handler"
	airepository "github.com/mzakiaklhairi/velora/internal/ai/repository"
	aiRoutes "github.com/mzakiaklhairi/velora/internal/ai/routes"
	aiService "github.com/mzakiaklhairi/velora/internal/ai/service"
	aiSummary "github.com/mzakiaklhairi/velora/internal/ai/summary"
	"github.com/mzakiaklhairi/velora/internal/infrastructure"
	"github.com/mzakiaklhairi/velora/internal/infrastructure/jwt"
	authHandler "github.com/mzakiaklhairi/velora/internal/modules/auth/handler"
	"github.com/mzakiaklhairi/velora/internal/modules/auth/repository"
	authRoutes "github.com/mzakiaklhairi/velora/internal/modules/auth/routes"
	authService "github.com/mzakiaklhairi/velora/internal/modules/auth/service"
	userrepo "github.com/mzakiaklhairi/velora/internal/modules/user/repository"
	"github.com/mzakiaklhairi/velora/internal/shared"
	"github.com/mzakiaklhairi/velora/internal/shared/scheduler"
)

// Router holds the Gin router
type Router struct {
	Engine *gin.Engine
}

// InitRouter initializes the Gin router
func InitRouter(debug bool) *Router {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	// Recovery middleware
	engine.Use(gin.Recovery())

	shared.Info("Router initialized",
		"mode", gin.Mode(),
	)

	return &Router{Engine: engine}
}

// RegisterRoutes registers all application routes
func (r *Router) RegisterRoutes(
	readyHandler func(*gin.Context),
	healthHandler func(*gin.Context),
	rootHandler func(*gin.Context),
	userRepo userrepo.UserRepository,
	jwtService *jwt.JWTService,
	refreshTokenRepo repository.RefreshTokenRepository,
	db *gorm.DB,
	refreshExpires time.Duration,
) {
	// Health check endpoints
	r.Engine.GET("/", rootHandler)
	r.Engine.GET("/health", healthHandler)
	r.Engine.GET("/ready", readyHandler)

	// API v1 routes
	apiV1 := r.Engine.Group("/api/v1")

	// AI health endpoint (public, no auth required)
	aiRoutes.SetupAIRoutes(apiV1)

	// Initialize auth dependencies
	authSvc := authService.NewAuthServiceImpl(userRepo, nil, refreshTokenRepo, jwtService, db, refreshExpires)
	authHdlr := authHandler.NewAuthHandler(authSvc)

	// Register auth routes
	authRoutes.RegisterRoutes(apiV1, authHdlr)

	// Protected auth endpoint
	authProtected := apiV1.Group("/auth")
	authProtected.Use(infrastructure.AuthMiddleware(jwtService))
	{
		authProtected.GET("/me", func(c *gin.Context) {
			userID := infrastructure.GetUserID(c)

			user, err := userRepo.FindByID(c.Request.Context(), userID)
			if err != nil {
				shared.ErrorResponse(c, http.StatusNotFound, "User not found")
				return
			}

			shared.Success(c, http.StatusOK, gin.H{
				"id":         user.ID,
				"name":       user.Name,
				"email":      user.Email,
				"status":     user.Status,
				"created_at": user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
				"updated_at": user.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
			})
		})
	}

	// Protected test endpoint for middleware testing
	protected := apiV1.Group("/protected")
	protected.Use(infrastructure.AuthMiddleware(jwtService))
	{
		protected.GET("/test", func(c *gin.Context) {
			shared.Success(c, http.StatusOK, gin.H{
				"user_id": infrastructure.GetUserID(c),
				"email":   infrastructure.GetUserEmail(c),
				"name":    infrastructure.GetUserName(c),
			})
		})
	}

	// Initialize workspace dependencies
	workspaceRepo := airepository.NewPostgresWorkspaceRepository(db)
	workspaceService := aiService.NewWorkspaceServiceImpl(workspaceRepo)
	workspaceHandler := aiHandler.NewWorkspaceHandler(workspaceService)

	// Register workspace routes
	aiRoutes.RegisterRoutes(apiV1, workspaceHandler, jwtService)

	// Initialize conversation and message dependencies
	conversationRepo := airepository.NewPostgresConversationRepository(db)
	messageRepo := airepository.NewPostgresMessageRepository(db)
	conversationService := aiService.NewConversationServiceImpl(conversationRepo)
	messageService := aiService.NewMessageServiceImpl(messageRepo)
	conversationHandler := aiHandler.NewConversationHandler(conversationService)

	// Initialize summary service
	summaryProvider := aiSummary.NewProviderAdapter(nil) // Will be resolved at runtime
	summarySvc := aiSummary.NewService(conversationRepo, messageRepo, summaryProvider, nil)

	// Initialize scheduler
	sched := scheduler.NewImmediateScheduler()

	// Initialize ChatService
	providerResolver := aiEngine.NewDefaultProviderResolver()
	chatEngine := aiEngine.NewChatService(
		conversationRepo,
		messageRepo,
		messageRepo,
		providerResolver,
		summarySvc,
		sched,
	)
	chatService := aiService.NewChatServiceImpl(conversationRepo, messageRepo, chatEngine)
	messageHandler := aiHandler.NewMessageHandler(messageService, chatService)

	// Register conversation and message routes with auth middleware
	workspacesGroup := apiV1.Group("/workspaces")
	workspacesGroup.Use(infrastructure.AuthMiddleware(jwtService))
	aiRoutes.RegisterConversationRoutes(workspacesGroup, conversationHandler)
	aiRoutes.RegisterMessageRoutes(workspacesGroup, messageHandler)
}
