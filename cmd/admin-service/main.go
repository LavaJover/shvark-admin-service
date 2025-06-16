package main

import (
	"fmt"
	"log"

	"github.com/LavaJover/shvark-admin-service/internal/config"
	"github.com/LavaJover/shvark-admin-service/internal/grpcclients"
	"github.com/LavaJover/shvark-admin-service/internal/http/handlers"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/LavaJover/shvark-admin-service/pkg/docs"
	"github.com/gin-gonic/gin"
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

	// init authz-client
	// authzAddr := "localhost:50054"
	// authzClient, err := grpcclients.NewAuthzClient(authzAddr)
	// if err != nil {
	// 	log.Fatalf("failed to init authz client\n")
	// }

	// init sso client
	ssoAddr := "localhost:50051"
	ssoClient, err := grpcclients.NewSSOClient(ssoAddr)
	if err != nil {
		log.Fatalf("failed to init sso client\n")
	}
	authHandler := handlers.NewAuthHandler(ssoClient)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", authHandler.Login)
	}

	r.Run(":9090")
}