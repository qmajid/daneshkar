package cmd

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	v1 "github.com/qmajid/daneshkar/internal/route/v1"
	"github.com/spf13/cobra"
)

var (
	port int
)

// serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run web server",
	Run: func(cmd *cobra.Command, args []string) {
		addr := fmt.Sprintf(":%d", port)

		r := gin.Default()
		// r.Use(middleware.AuthRequired("test-key"), middleware.JwtMiddlware())

		// serve /static/*.css
		r.Static("/static", "./static")
		// load templates
		r.LoadHTMLGlob("templates/*.html")

		v1.InitRoute(r)

		if err := r.Run(addr); err != nil {
			log.Fatalf("failed to run server: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntVarP(&port, "port", "p", 8081, "server port")
}
