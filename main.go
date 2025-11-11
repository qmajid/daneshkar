package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/qmajid/daneshkar/internal/route/v1"
)

type Info struct {
	Name   string `json:"name"`
	Family string `json:"family"`
}

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()
	v1.InitRoute(r)

	// Define a simple GET endpoint
	r.GET("/ping", Pong)
	r.GET("/user/:name", PathParam)
	r.GET("/welcome", QueryParam)
	r.POST("/info", BodyParam)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func PathParam(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func QueryParam(c *gin.Context) {
	firstname := c.DefaultQuery("name", "Guest")
	lastname := c.Query("family") // shortcut for c.Request.URL.Query().Get("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func BodyParam(c *gin.Context) {
	var i Info
	err := c.ShouldBindJSON(&i)
	if err != nil {
		c.String(http.StatusBadRequest, "")
		return
	}
	c.String(http.StatusOK, "Hello %s %s", i.Name, i.Family)
}
