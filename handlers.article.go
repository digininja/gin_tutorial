// handlers.article.go

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func showRobinPage(c *gin.Context) {

	render(
		c,
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

	render(
		c,
		http.StatusOK,
		"index.html",
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)

}

func getArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := getArticleByID(articleID); err == nil {
			// Call the HTML method of the Context to render a template
			render(
				c,
				// Set the HTTP status to 200 (OK)
				http.StatusOK,
				// Use the index.html template
				"article.html",
				// Pass the data that the page uses
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)

		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func custom404(c *gin.Context) {
	pe := pageError{http.StatusNotFound, "Page Not Found"}
	render(
		c,
		// Set the HTTP status to 400 (not found)
		http.StatusNotFound,
		// Use the 404.html template
		"404.html",
		gin.H{
			"title":   "Page Not Found",
			"payload": pe,
		},
	)
}
