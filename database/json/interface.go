package json

type IRecipes interface {
	GetByID(id string) (*Recipe, int)
}
