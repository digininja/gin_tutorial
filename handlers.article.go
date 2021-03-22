// handlers.article.go

package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

import "github.com/go-playground/validator/v10"

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func Descriptive(verr validator.ValidationErrors) []ValidationError {
	errs := []ValidationError{}

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs = append(errs, ValidationError{Field: f.Field(), Reason: err})
	}

	return errs
}

func validateJSON(c *gin.Context) {
	var query struct {
		Name  string `form:"name" json:"name" binding:"required"`
		Color string `form:"color" json:"color" binding:"required,oneof=blue yellow"`
	}

	// Testing with
	// curl -s http://localhost:3000/validate -H "Accept: application/json" -H "Content-Type: application/json" --data '{"Name":"test.coma","coxlor": "blue"}'

	if err := c.ShouldBind(&query); err != nil {

		xType := fmt.Sprintf("%T", err)
		fmt.Println(xType)
		// prints validator.ValidationErrors

		var verr validator.ValidationErrors

		// This always returns false
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": Descriptive(verr)})
			fmt.Println("in here")
			return
		}

		// error message is
		// "Key: 'color' Error:Field validation for 'color' failed on the 'required' tag"
		// so know validator is called and failing

		fmt.Println("other")
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func submittedURL(c *gin.Context) {
	// Validation
	// https://seb-nyberg.medium.com/better-validation-errors-in-go-gin-88f983564a3d

	// https://blog.depa.do/post/gin-validation-errors-handling

	var url URL

	if err := c.ShouldBind(&url); err != nil {

		if verr, ok := err.(validator.ValidationErrors); ok {
			fmt.Printf("this is actually a validation error, %s\n", verr)
		}
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

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
