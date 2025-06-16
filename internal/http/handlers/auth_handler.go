package handlers

import (
	"net/http"

	"github.com/LavaJover/shvark-admin-service/internal/grpcclients"
	"github.com/LavaJover/shvark-admin-service/internal/http/dto"
	ssopb "github.com/LavaJover/shvark-sso-service/proto/gen"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	SSOClient *grpcclients.SSOClient
}

func NewAuthHandler(ssoClient *grpcclients.SSOClient) *AuthHandler {
	return &AuthHandler{SSOClient: ssoClient}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var httpRequest dto.LoginRequest
	if err := c.ShouldBindJSON(&httpRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	response, err := h.SSOClient.Login(&ssopb.LoginRequest{
		Login: httpRequest.Login,
		Password: httpRequest.Password,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "login failed"})
		return
	}
	c.JSON(http.StatusOK, dto.LoginResponse{
		AccessToken: response.AccessToken,
		RefreshToken: response.RefreshToken,
	})
}