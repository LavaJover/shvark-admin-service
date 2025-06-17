package handlers

import (
	"net/http"

	"github.com/LavaJover/shvark-admin-service/internal/grpcclients"
	"github.com/LavaJover/shvark-admin-service/internal/http/dto"
	authzpb "github.com/LavaJover/shvark-authz-service/proto/gen"
	ssopb "github.com/LavaJover/shvark-sso-service/proto/gen"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	SSOClient *grpcclients.SSOClient
	AuthzClient *grpcclients.AuthzClient
}

func NewAuthHandler(ssoClient *grpcclients.SSOClient, authzClient *grpcclients.AuthzClient) *AuthHandler {
	return &AuthHandler{SSOClient: ssoClient, AuthzClient: authzClient}
}

// @Summary Login user
// @Description Authenticate user and receive JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.LoginRequest true "user creds"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /auth/login [post]
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

	validationResponse, err := h.SSOClient.ValidateToken(&ssopb.ValidateTokenRequest{
		AccessToken: response.AccessToken,
	})
	if err != nil || !validationResponse.Valid{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid access token"})
		return
	}
	userID := validationResponse.UserId

	authzResponse, err := h.AuthzClient.CheckPermission(&authzpb.CheckPermissionRequest{
		UserId: userID,
		Object: "*",
		Action: "*",
	})
	if err != nil || !authzResponse.Allowed {
		c.JSON(http.StatusForbidden, gin.H{"error": "not enough rights"})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		AccessToken: response.AccessToken,
		RefreshToken: response.RefreshToken,
	})
}