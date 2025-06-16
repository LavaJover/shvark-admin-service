package middleware

import (
	"net/http"

	"github.com/LavaJover/shvark-admin-service/internal/grpcclients"
	authzpb "github.com/LavaJover/shvark-authz-service/proto/gen"
	"github.com/gin-gonic/gin"
)

func RequirePermission(authzClient *grpcclients.AuthzClient, object, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("user_id")
		response, err := authzClient.CheckPermission(&authzpb.CheckPermissionRequest{
			UserId: userID,
			Object: object,
			Action: action,
		})
		if err != nil || !response.Allowed{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "access denied"})
			return
		}
		c.Next()
	}
}