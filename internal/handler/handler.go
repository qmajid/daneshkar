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

// @Summary      Get all recipes
// @Description  Retrieve all recipes.
// @Tags         recipes
// @Produce      json
// @Success      200  {array}   recipes.Recipe
// @Router       /v1/recipes [get]
func Recipes(c *gin.Context) {
	c.JSON(http.StatusOK, recipes.GetAll())

	// c.HTML(http.StatusOK, "index.html", gin.H{
	// 	"Title":   "Recipe List",
	// 	"recipes": recipes.GetAll(),
	// })
}

// @Summary      Get recipe by ID
// @Description  Get a single recipe by its ID.
// @Tags         recipes
// @Produce      json
// @Param        id   path      string  true  "Recipe ID"
// @Success      200  {object}  recipes.Recipe
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /v1/recipes/{id} [get]
func RecipesByID(c *gin.Context) {
	j := recipes.JsonRecipes{}
	id := c.Param("id")
	rcp, status := j.GetByID(id)

	if rcp == nil {
		c.JSON(status, gin.H{
			"error": "recipe not found",
		})
		return
	}

	// c.HTML(http.StatusOK, "email.html", rcp)
	c.JSON(status, rcp)
}

// @Summary      Get recipe by ID and return email format
// @Description  Get a single recipe by its ID.
// @Tags         recipes
// @Produce      json
// @Param        id   path      string  true  "Recipe ID"
// @Success      200  {object}  recipes.Recipe
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /v1/recipes/{id}/email [get]
func GenerateEmailTemplate(c *gin.Context) {
	j := recipes.JsonRecipes{}
	id := c.Param("id")
	rcp, status := j.GetByID(id)

	if rcp == nil {
		c.JSON(status, gin.H{
			"error": "recipe not found",
		})
		return
	}

	c.HTML(http.StatusOK, "email.html", rcp)
	// c.JSON(status, recipe)
}

// @Summary      Insert new recipe
// @Description  Add a new recipe.
// @Tags         recipes
// @Accept       json
// @Produce      json
// @Param        recipe  body     recipes.Recipe  true  "Recipe object"
// @Success      200     {object} map[string]string
// @Failure      400     {object} map[string]string
// @Router       /v1/recipes [post]
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

// @Summary      Update recipe
// @Description  Update an existing recipe.
// @Tags         recipes
// @Accept       json
// @Produce      json
// @Param        recipe  body     recipes.Recipe  true  "Recipe object"
// @Success      200     {object} map[string]string
// @Failure      400     {object} map[string]string
// @Router       /v1/recipes [put]
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

// @Summary      Delete recipe
// @Description  Delete a recipe by its ID.
// @Tags         recipes
// @Produce      json
// @Param        id   path      string  true  "Recipe ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /v1/recipes/{id} [delete]
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

// @Summary      Patch recipe publish time
// @Description  Update the publish time of an existing recipe.
// @Tags         recipes
// @Accept       json
// @Produce      json
// @Param        recipe  body     recipes.Recipe  true  "Recipe object (only ID is used)"
// @Success      200     {object} map[string]string
// @Failure      400     {object} map[string]string
// @Router       /v1/recipes [patch]
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
