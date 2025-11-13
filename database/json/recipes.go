package json

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Recipe struct {
	//swagger:ignore
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	PersianName  string    `json:"persian_name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

var recipes []Recipe

const ReceipsFilePath string = "./database/json/recipes.json"

func init() {
	file, err := os.ReadFile(ReceipsFilePath)
	if err != nil {
		log.Fatal("failed to read recipes file: ", err)
	}

	if err := json.Unmarshal(file, &recipes); err != nil {
		log.Fatal("failed to unmarshal recipes JSON:", err)
	}
}

func GetAll() []Recipe {
	return recipes
}

func GetByID(id string) (*Recipe, int) {
	for _, recipe := range recipes {
		if recipe.ID == id {
			return &recipe, http.StatusOK
		}
	}
	return nil, http.StatusNotFound
}

func InsertRecipe(newRecipe Recipe) (int, error) {
	// Add the new recipe to the slice
	recipes = append(recipes, newRecipe)

	// Marshal the updated recipes slice to JSON
	data, err := json.MarshalIndent(recipes, "", "  ")
	if err != nil {
		return http.StatusBadRequest, err
	}

	// Write the marshaled data back to the JSON file
	err = os.WriteFile(ReceipsFilePath, data, 0644)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func UpdateRecipe(id string, updatedRecipe Recipe) (int, error) {
	updated := false
	for i, recipe := range recipes {
		if recipe.ID == id {
			// Replace the recipe at index i with updatedRecipe.
			// Ensure the ID remains unchanged.
			updatedRecipe.ID = id
			recipes[i] = updatedRecipe
			updated = true
			break
		}
	}

	if !updated {
		return http.StatusNotFound, nil
	}

	// Marshal the updated recipes slice to JSON
	data, err := json.MarshalIndent(recipes, "", "  ")
	if err != nil {
		return http.StatusBadRequest, err
	}

	// Write the marshaled data back to the JSON file
	err = os.WriteFile(ReceipsFilePath, data, 0644)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}


func DeleteRecipe(id string) (int, error) {
	found := false
	newRecipes := make([]Recipe, 0, len(recipes))
	for _, recipe := range recipes {
		if recipe.ID == id {
			found = true
			continue // skip the recipe to delete
		}
		newRecipes = append(newRecipes, recipe)
	}

	if !found {
		return http.StatusNotFound, nil
	}

	recipes = newRecipes

	// Marshal the updated recipes slice to JSON
	data, err := json.MarshalIndent(recipes, "", "  ")
	if err != nil {
		return http.StatusBadRequest, err
	}

	// Write the marshaled data back to the JSON file
	err = os.WriteFile(ReceipsFilePath, data, 0644)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func PatchRecipeTime(id string, newTime string) (int, error) {
	found := false

	for i, recipe := range recipes {
		if recipe.ID == id {
			parsedTime, err := time.Parse(time.RFC3339, newTime)
			if err != nil {
				return http.StatusBadRequest, err
			}
			recipes[i].PublishedAt = parsedTime
			found = true
			break
		}
	}

	if !found {
		return http.StatusNotFound, nil
	}

	// Marshal the updated recipes slice to JSON
	data, err := json.MarshalIndent(recipes, "", "  ")
	if err != nil {
		return http.StatusBadRequest, err
	}

	// Write the marshaled data back to the JSON file
	err = os.WriteFile(ReceipsFilePath, data, 0644)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
