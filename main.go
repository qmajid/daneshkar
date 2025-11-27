package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qmajid/daneshkar/cmd"
	v1 "github.com/qmajid/daneshkar/internal/route/v1"
)

// @title           Daneshkar API
// @version         1.0
// @description     This is a sample server for Daneshkar API.
// @host            localhost:8081
// @BasePath        /
// @schemes         http

type Info struct {
	Name   string `json:"name"`
	Family string `json:"family"`
}

const asciiArt = ` ______   _______  _        _______  _______           _        _______  _______ 
(  __  \ (  ___  )( (    /|(  ____ \(  ____ \|\     /|| \    /\(  ___  )(  ____ )
| (  \  )| (   ) ||  \  ( || (    \/| (    \/| )   ( ||  \  / /| (   ) || (    )|
| |   ) || (___) ||   \ | || (__    | (_____ | (___) ||  (_/ / | (___) || (____)|
| |   | ||  ___  || (\ \) ||  __)   (_____  )|  ___  ||   _ (  |  ___  ||     __)
| |   ) || (   ) || | \   || (            ) || (   ) ||  ( \ \ | (   ) || (\ (   
| (__/  )| )   ( || )  \  || (____/\/\____) || )   ( ||  /  \ \| )   ( || ) \ \__
(______/ |/     \||/    )_)(_______/\_______)|/     \||_/    \/|/     \||/   \__/
`

func main() {

	cmd.Execute()
	return

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()
	// r.Use(middleware.AuthRequired("test-key"), middleware.JwtMiddlware())

	// serve /static/*.css
	r.Static("/static", "./static")
	// load templates
	r.LoadHTMLGlob("templates/*.html")

	v1.InitRoute(r)

	// Define a simple GET endpoint
	r.GET("", AsciiArt)
	r.GET("/hello", HelloWorld)
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

func AsciiArt(c *gin.Context) {
	c.String(http.StatusOK, asciiArt)
}

func HelloWorld(c *gin.Context) {
	c.File("static/index.html")
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
