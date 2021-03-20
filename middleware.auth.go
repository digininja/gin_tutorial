// middleware.auth.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIAuthRequired(c *gin.Context) {
	apiKey := c.PostForm("apiKey")
	headerKey := c.GetHeader("apiKey")

	if headerKey != "123" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized missing header key"})
		return
	}
	if apiKey != "123" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized missing POST key"})
		return
	}
	/*
	   session := sessions.Default(c)
	   	user := session.Get(userkey)
	   	if user == nil {
	   		// Abort the request with the appropriate error code
	   		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	   		return
	   	}
	*/
	// Continue down the chain to handler etc
	c.Next()
}
