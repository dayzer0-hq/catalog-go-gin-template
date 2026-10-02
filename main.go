// Command dayzer0-gin-starter is a minimal Gin HTTP service you build your project on top of.
// It exposes a single health route so you can confirm the server runs before writing anything.
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := newRouter()

	addr := ":" + port()
	log.Printf("listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}

// newRouter builds the Gin engine and registers the routes. Keep route registration here so
// handlers and their tests can share one definition of the service.
func newRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health", health)
	return r
}

// health answers a liveness probe. Add your own routes beside it in newRouter.
func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// port returns the PORT environment variable, defaulting to 8080 for local runs.
func port() string {
	if p := os.Getenv("PORT"); p != "" {
		return p
	}
	return "8080"
}
