package main

import (
	"fmt"
	"log"

	"github.com/LavaJover/shvark-admin-service/internal/config"
	"github.com/LavaJover/shvark-admin-service/internal/grpcclients"
	"github.com/LavaJover/shvark-admin-service/internal/http/handlers"
	"github.com/LavaJover/shvark-admin-service/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// init config
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// init authz-client
	authzAddr := "localhost:50054"
	authzClient, err := grpcclients.NewAuthzClient(authzAddr)
	if err != nil {
		log.Fatalf("failed to init authz client\n")
	}

	// init sso client
	ssoAddr := "localhost:50051"
	ssoClient, err := grpcclients.NewSSOClient(ssoAddr)
	if err != nil {
		log.Fatalf("failed to init sso client\n")
	}
	authHandler := handlers.NewAuthHandler(ssoClient)

	r := gin.Default()

	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.RequirePermission(authzClient, "admin:panel", "access"))
	{
		adminGroup.POST("/login", authHandler.Login)
	}

	r.Run("9090")
}