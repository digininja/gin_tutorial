// routes.go

package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)
	router.GET("/robin", showRobinPage)
	router.GET("/article/view/:article_id", getArticle)

	apiRoutes := router.Group("/api")
	{
		apiRoutes.GET("/ver", checkAPILogin(), getVer)
	}

	router.NoRoute(custom404)
}
