package v1

import (
	"github.com/gin-gonic/gin"
	handler "github.com/qmajid/daneshkar/internal/handler"
)

func InitRoute(e *gin.Engine) {
	v1 := e.Group("/v1")
	v1.GET("/ping", handler.Pong)
	v1.GET("/recipes", handler.Recipes)
	v1.GET("/recipes/:id", handler.RecipesByID)
}
