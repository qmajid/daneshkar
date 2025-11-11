package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	recipes "github.com/qmajid/daneshkar/database/json"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong-from-v1",
	})
}

func Recipes(c *gin.Context) {
	c.JSON(http.StatusOK, recipes.GetAll())
}
