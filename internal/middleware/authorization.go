package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequired(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-API-KEY")

		if token == "" || token != apiKey {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
