package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/LavaJover/shvark-admin-service/internal/config"
	"github.com/LavaJover/shvark-admin-service/internal/grpcclients"
	"github.com/LavaJover/shvark-admin-service/internal/http/handlers"
	"github.com/LavaJover/shvark-admin-service/internal/http/middleware"
	"github.com/LavaJover/shvark-admin-service/internal/httpclients"
	"github.com/LavaJover/shvark-admin-service/internal/usecase"
	_ "github.com/LavaJover/shvark-admin-service/pkg/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Admin Service API
// @version 1.0
// @description API for internal admin panel
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// init config
	cfg := config.MustLoad()
	fmt.Println(cfg)

	///////// Clients /////////
	// init authz-client
	slog.Info("connecting authz service...")
	authzAddr := "localhost:50054"
	authzClient, err := grpcclients.NewAuthzClient(authzAddr)
	if err != nil {
		log.Fatalf("failed to init authz client\n")
	}
	// init sso client
	slog.Info("connecting sso service...")
	ssoAddr := "localhost:50051"
	ssoClient, err := grpcclients.NewSSOClient(ssoAddr)
	if err != nil {
		log.Fatalf("failed to init sso client\n")
	}
	// init wallet client
	slog.Info("connecting to wallet service...")
	walletClient, err := httpclients.NewWalletHTTPClient()
	if err != nil {
		log.Fatalf("failed to init wallet client")
	}

	//////// Handlers ////////
	authHandler := handlers.NewAuthHandler(ssoClient, authzClient)

	traderUsecase := usecase.NewTraderUsecase(ssoClient, walletClient)
	traderhandler := handlers.NewTraderHandler(traderUsecase)

	/////// Routing //////////
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", authHandler.Login)
	}

	traderGroup := r.Group("/admin/traders")
	traderGroup.Use(middleware.AuthMiddleware("my-secret-word"), middleware.RequirePermission(authzClient, "*", "*"))
	{
		traderGroup.POST("/register", traderhandler.RegisterTrader)
	}

	r.Run(":9090")
}