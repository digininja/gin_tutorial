// handlers.article.go

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func showRobinPage(c *gin.Context) {

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"robin.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "Robin's Page",
			"hello": "World",
		},
	)

}
func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)

}
