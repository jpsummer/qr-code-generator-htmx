package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jpsummer/qr-code-generator-htmx/internals"
)

func main() {
	router := gin.Default()

	// initialize config
	app := internals.Config{Router: router}

	// serve static files
	router.Static("/static", "./static")

	// routes
	app.Routes()

	router.Run("localhost:8080")
}
