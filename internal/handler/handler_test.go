package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	recipes "github.com/qmajid/daneshkar/database/json"
	mocks "github.com/qmajid/daneshkar/mocks"
	gomock "go.uber.org/mock/gomock"
)

func TestGetRecipeByID_Success(t *testing.T) {
	mockRecipe := &recipes.Recipe{
		ID:    "123",
		Title: "Mocked Pizza",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mocks.NewMockIRecipes(ctrl)
	mockService.EXPECT().GetByID(gomock.Any()).Times(1).Return(mockRecipe, http.StatusOK)

	h := NewHandler(mockService)
	// --------------------

	gin.SetMode(gin.TestMode)
	// --------------- server
	router := gin.New()
	router.GET("/recipes/:id", h.RecipesByID)

	// ------------- request -> client
	req := httptest.NewRequest(http.MethodGet, "/recipes/123", nil)
	resp := httptest.NewRecorder()

	// --------------------- server
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.Code)
	}
}

func TestGetRecipeByID_NofFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mocks.NewMockIRecipes(ctrl)
	mockService.EXPECT().GetByID(gomock.Any()).Times(1).Return(nil, http.StatusNotFound)

	h := NewHandler(mockService)
	// --------------------

	gin.SetMode(gin.TestMode)
	// --------------- server
	router := gin.New()
	router.GET("/recipes/:id", h.RecipesByID)

	// ------------- request -> client
	req := httptest.NewRequest(http.MethodGet, "/recipes/123456", nil)
	resp := httptest.NewRecorder()

	// --------------------- server
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, resp.Code)
	}
}
