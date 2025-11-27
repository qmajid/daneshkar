package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	v1 "github.com/qmajid/daneshkar/internal/route/v1"
	"github.com/spf13/cobra"
)

const asciiArt = ` ______   _______  _        _______  _______           _        _______  _______ 
(  __  \ (  ___  )( (    /|(  ____ \(  ____ \|\     /|| \    /\(  ___  )(  ____ )
| (  \  )| (   ) ||  \  ( || (    \/| (    \/| )   ( ||  \  / /| (   ) || (    )|
| |   ) || (___) ||   \ | || (__    | (_____ | (___) ||  (_/ / | (___) || (____)|
| |   | ||  ___  || (\ \) ||  __)   (_____  )|  ___  ||   _ (  |  ___  ||     __)
| |   ) || (   ) || | \   || (            ) || (   ) ||  ( \ \ | (   ) || (\ (   
| (__/  )| )   ( || )  \  || (____/\/\____) || )   ( ||  /  \ \| )   ( || ) \ \__
(______/ |/     \||/    )_)(_______/\_______)|/     \||_/    \/|/     \||/   \__/
`

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
		r.GET("", AsciiArt)
		// r.Use(middleware.AuthRequired("test-key"), middleware.JwtMiddlware())

		// serve /static/*.css
		r.Static("/static", "./static")
		// load templates
		r.LoadHTMLGlob("templates/*.html")

		v1.InitRoute(r)

		srv := &http.Server{
			Addr:    addr,
			Handler: r.Handler(),
		}

		go func() {
			// service connections
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 5 seconds.
		quit := make(chan os.Signal, 1)
		// kill (no params) by default sends syscall.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Gracefull Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Println("Server Shutdown:", err)
		}
		log.Println("Server exiting")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntVarP(&port, "port", "p", 8081, "server port")
}

func AsciiArt(c *gin.Context) {
	c.String(http.StatusOK, asciiArt)
}
