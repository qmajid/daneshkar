package v1

import (
	"github.com/gin-gonic/gin"
	service "github.com/qmajid/daneshkar/database/json"
	handler "github.com/qmajid/daneshkar/internal/handler"
)

func InitRoute(e *gin.Engine) {
	v1 := e.Group("/v1")
	s := service.NewJsonService(nil)
	s.Load("database/json/recipes.json")
	h := handler.RecipesHandler{Service: s}

	v1.GET("/ping", h.Pong)
	v1.GET("/recipes", h.Recipes)
	v1.GET("/recipes/:id", h.RecipesByID)
	v1.GET("/recipes/:id/email", h.GenerateEmailTemplate)
	v1.POST("/recipes", h.InsertNewRecipes)
	v1.PUT("/recipes", h.UpdateRecipes)
	v1.PATCH("/recipes", h.PatchRecipesTime)
	v1.DELETE("/recipes/:id", h.DeleteRecipes)
}
