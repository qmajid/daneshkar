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

func RecipesByID(c *gin.Context) {
	id := c.Param("id")
	recipe, status := recipes.GetByID(id)

	if recipe == nil {
		c.JSON(status, gin.H{
			"error": "recipe not found",
		})
		return
	}

	c.JSON(status, recipe)
}

func InsertRecipe(c *gin.Context) {
	var newRecipe recipes.Recipe

	if err := c.ShouldBindJSON(&newRecipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}



	// Optionally, initialize ID/timestamp here if needed
	// e.g. generate unique ID or set PublishedAt

	status, err := recipes.InsertRecipe(newRecipe)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "recipe inserted successfully",
		"recipe":  newRecipe,
	})
}
