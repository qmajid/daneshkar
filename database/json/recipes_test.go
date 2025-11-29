package json

import "testing"

func BenchmarkGetRecipeByID_NotFound(b *testing.B) {
	s := NewJsonService()
	s.Load("recipes.json")

	for i := 0; i < b.N; i++ {
		_, _ = s.GetByID("")
	}
}

func BenchmarkGetRecipeByID_FoundAtFirst(b *testing.B) {
	s := NewJsonService()
	s.Load("recipes.json")

	for i := 0; i < b.N; i++ {
		_, _ = s.GetByID("p001")
	}
}

func BenchmarkGetRecipeByID_FoundAtMiddle(b *testing.B) {
	s := NewJsonService()
	s.Load("recipes.json")

	for i := 0; i < b.N; i++ {
		_, _ = s.GetByID("p010")
	}
}
