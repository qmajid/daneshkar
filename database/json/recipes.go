package json

import (
	"encoding/json"
	"log"
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
