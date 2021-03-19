// handlers.article.go

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getVer(c *gin.Context) {
	ver := version{"1.2.3"}

	render(
		c,
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"version.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title":   "Version",
			"payload": ver,
		},
	)

}
