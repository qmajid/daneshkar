package json

type IRecipes interface {
	Load(path string)
	GetAll() []Recipe
	GetByID(id string) (*Recipe, int)
	InsertRecipe(newRecipe Recipe) (int, error)
	UpdateRecipe(updatedRecipe Recipe) (int, error)
	DeleteRecipe(id string) (int, error)
	PatchRecipeTime(patchRecipes Recipe) (int, error)
}
