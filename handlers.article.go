// handlers.article.go

package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"strings"
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
		Color string `form:"colour" json:"colour" binding:"required,oneof=blue yellow"`
	}

	// Testing with
	// curl -s http://localhost:3000/validate -H "Accept: application/json" -H "Content-Type: application/json" --data '{"Name":"test.coma","colour": "blue"}'

	if err := c.ShouldBind(&query); err != nil {
		var verr validator.ValidationErrors

		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": Descriptive(verr)})
			return
		}

		// need to convert this to a better error
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	//data := c.BindJSON(&query)
	c.JSON(http.StatusOK, gin.H{"status": query.Name})
}

// IsDebugging returns true if the framework is running in debug mode.
// Use SetMode(gin.ReleaseMode) to disable debug mode.
func IsDebugging() bool {
	return gin.Mode() == gin.DebugMode
}

func debugPrint(format string, values ...interface{}) {
	if IsDebugging() {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		fmt.Fprintf(gin.DefaultWriter, "[GIN-debug] "+format, values...)
	}
}

type submission struct {
	URL   string
	UUID  string
	count int
}

type submissions struct {
	data []submission
}

func callback(c *gin.Context) {
	var callbackUUID struct {
		UUID string `form:"uuid" json:"uuid" binding:"required"`
	}

	// Testing with
	// curl -s http://localhost:3000/callback -H "Accept: application/json" -H "Content-Type: application/json" --data '{"UUID": "4ea5303b-90be-42e7-bb77-ed6590e87370"}'

	if err := c.ShouldBind(&callbackUUID); err != nil {
		debugPrint("Error: %s", err.Error())
		var verr validator.ValidationErrors

		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": Descriptive(verr)})
			return
		}

		// need to convert this to a better error
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	for k, data := range ourSubmissions.data {
		if data.UUID == callbackUUID.UUID {
			debugPrint("hit")
			debugPrint("count: %d", data.count)
			ourSubmissions.data[k].count++
			if data.count > 5 {
				c.JSON(http.StatusOK, gin.H{"status": "ready"})
			} else {
				c.JSON(http.StatusOK, gin.H{"status": "not ready yet"})
			}
			return
		}
	}
	debugPrint("miss")
	c.JSON(http.StatusOK, gin.H{"Miss": ""})

}

func submittedURL(c *gin.Context) {
	// Validation
	// https://seb-nyberg.medium.com/better-validation-errors-in-go-gin-88f983564a3d

	// https://blog.depa.do/post/gin-validation-errors-handling

	var submittedURL struct {
		URL string `form:"url" json:"url" binding:"required"`
	}

	// Testing with
	// curl -s http://localhost:3000/validate -H "Accept: application/json" -H "Content-Type: application/json" --data '{"Name":"test.coma","colour": "blue"}'

	if err := c.ShouldBind(&submittedURL); err != nil {
		debugPrint("Error: %s", err.Error())
		var verr validator.ValidationErrors

		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": Descriptive(verr)})
			return
		}

		// need to convert this to a better error
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	myUUID := uuid.NewString()

	submission := submission{}
	submission.URL = submittedURL.URL
	submission.count = 0
	submission.UUID = myUUID

	ourSubmissions.data = append(ourSubmissions.data, submission)

	c.JSON(http.StatusOK, gin.H{"ID": myUUID})
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
