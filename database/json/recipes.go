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
	Title        string    `json:"title"`
	PersianName  string    `json:"persian_name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

type JsonService struct {
	filePath string
	data     []Recipe
}

func NewJsonService() *JsonService {
	return &JsonService{}
}

func (s *JsonService) Load(path string) {
	s.filePath = path
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("failed to read recipes file: ", err)
	}

	if err := json.Unmarshal(file, &s.data); err != nil {
		log.Fatal("failed to unmarshal recipes JSON:", err)
	}
}

func (s *JsonService) InjectData(data []Recipe) {
	s.data = data
}

func (s *JsonService) GetAll() []Recipe {
	return s.data
}

func (s JsonService) GetByID(id string) (*Recipe, int) {
	for _, recipe := range s.data {
		if recipe.ID == id {
			return &recipe, http.StatusOK
		}
	}
	return nil, http.StatusNotFound
}

func (s *JsonService) InsertRecipe(newRecipe Recipe) (int, error) {
	// Add the new recipe to the slice
	s.data = append(s.data, newRecipe)

	return s.persistRecipes()
}

func (s *JsonService) UpdateRecipe(updatedRecipe Recipe) (int, error) {
	updated := false
	for i, recipe := range s.data {
		if recipe.ID == updatedRecipe.ID {
			// Replace the recipe at index i with updatedRecipe.
			// Ensure the ID remains unchanged.
			s.data[i] = updatedRecipe
			updated = true
			break
		}
	}

	if !updated {
		return http.StatusNotFound, nil
	}

	return s.persistRecipes()
}

func (s *JsonService) DeleteRecipe(id string) (int, error) {
	found := false
	newRecipes := make([]Recipe, 0, len(s.data))
	for _, recipe := range s.data {
		if recipe.ID == id {
			found = true
			continue // skip the recipe to delete
		}
		newRecipes = append(newRecipes, recipe)
	}

	if !found {
		return http.StatusNotFound, nil
	}

	s.data = newRecipes

	return s.persistRecipes()
}

func (s *JsonService) PatchRecipeTime(patchRecipes Recipe) (int, error) {
	found := false

	for i, r := range s.data {
		if r.ID == patchRecipes.ID {
			s.data[i].PublishedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		return http.StatusNotFound, nil
	}

	return s.persistRecipes()
}

func (s *JsonService) persistRecipes() (int, error) {
	// Marshal the updated recipes slice to JSON
	data, err := json.MarshalIndent(s.data, "", "  ")
	if err != nil {
		return http.StatusBadRequest, err
	}

	// Write the marshaled data back to the JSON file
	if err := os.WriteFile(s.filePath, data, 0644); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
