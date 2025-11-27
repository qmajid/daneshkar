package main

import (
	"github.com/qmajid/daneshkar/cmd"
)

// @title           Daneshkar API
// @version         1.0
// @description     This is a sample server for Daneshkar API.
// @host            localhost:8081
// @BasePath        /
// @schemes         http

func main() {
	cmd.Execute()
}
