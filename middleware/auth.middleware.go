package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stevenfamy/go-send-message-slack/config"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("authToken")
		token := config.GetConfig("API_AUTH_TOKEN")

		if authToken != token {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not Authorized!"})
			c.Abort()
		}

		c.Next()
	}
}
