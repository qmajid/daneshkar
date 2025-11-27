package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuthRequiredUnauthorized_WithEmptyRequestHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//--------------- server
	router := gin.New()
	router.Use(AuthRequired("super-secret"))
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "allowed")
	})

	//------------- request -> client
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	//--------------------- server
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, resp.Code)
	}

	if !strings.Contains(resp.Body.String(), `"unauthorized"`) {
		t.Fatalf("expected unauthorized message, got %q", resp.Body.String())
	}
}

func TestAuthRequiredUnauthorized_HeaderNotEqual(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//--------------- server
	router := gin.New()
	router.Use(AuthRequired("super-secret"))
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "allowed")
	})

	//------------- request -> client
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-API-KEY", "aaaaaaaaaaaaaaaaa")
	resp := httptest.NewRecorder()

	//--------------------- server
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, resp.Code)
	}

	if !strings.Contains(resp.Body.String(), `"unauthorized"`) {
		t.Fatalf("expected unauthorized message, got %q", resp.Body.String())
	}
}

func TestAuthRequiredAuthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(AuthRequired("super-secret"))

	called := false
	router.GET("/hello", func(c *gin.Context) {
		called = true
		c.String(http.StatusOK, "allowed")
	})

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	req.Header.Set("X-API-KEY", "super-secret")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.Code)
	}

	if resp.Body.String() != "allowed" {
		t.Fatalf("expected body %q, got %q", "allowed", resp.Body.String())
	}

	if !called {
		t.Fatalf("expected next handler to be called")
	}
}
