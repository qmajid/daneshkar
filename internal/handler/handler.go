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

func InsertNewRecipes(c *gin.Context) {
	var r recipes.Recipe

	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body request"})
		return
	}

	statusCode, err := recipes.InsertRecipe(r)
	if err != nil {
		c.JSON(statusCode, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully insert new recipes."})
}

func UpdateRecipes(c *gin.Context) {
	var r recipes.Recipe

	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body request"})
		return
	}

	statusCode, err := recipes.UpdateRecipe(r)
	if err != nil {
		c.JSON(statusCode, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully update current recipes."})
}

func DeleteRecipes(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	code, err := recipes.DeleteRecipe(id)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully delete"})
}

func PatchRecipesTime(c *gin.Context) {
	var r recipes.Recipe

	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body request"})
		return
	}

	statusCode, err := recipes.PatchRecipeTime(r)
	if err != nil {
		c.JSON(statusCode, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully patched recipe publish time."})
}
