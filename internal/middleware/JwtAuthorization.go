package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secret = []byte(getEnv("JWT_SECRET", "supersecret"))

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}

func JwtMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("JWT")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "jwt token in empty",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(token, "Bearer ")

		jwtToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method for jwt")
			}
			return secret, nil
		})

		if err != nil || !jwtToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid jwt token",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
