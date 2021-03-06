// routes.go

package main

func initializeRoutes() {

	// router.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))

	// Handle the index route
	router.GET("/", showIndexPage)
	router.GET("/robin", showRobinPage)
	router.POST("/robin", submittedURL)
	router.GET("/article/view/:article_id", getArticle)

	apiRoutes := router.Group("/api")
	apiRoutes.Use(APIAuthRequired)
	{
		apiRoutes.GET("/ver", getVer)
		apiRoutes.POST("/ver", getVer)
	}

	router.NoRoute(custom404)
}
