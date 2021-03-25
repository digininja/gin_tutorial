// main.go

package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

var ourSubmissions submissions

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, status int, templateName string, data gin.H) {

	/*
		loggedInInterface, _ := c.Get("is_api_logged_in")
		data["is_api_logged_in"] = loggedInInterface.(bool)
	*/
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(status, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(status, data["payload"])
	default:
		// Respond with HTML
		c.HTML(status, templateName, data)
	}
	return

}

func main() {
	// ConnectDatabase()

	// Set the router as the default one provided by Gin
	router = gin.Default()

	ourSubmissions = submissions{}

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*html")

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	router.Run(":3000")

}
