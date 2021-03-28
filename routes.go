// routes.go

package main

func initializeRoutes() {

	// router.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))

	// Handle the index route
	router.GET("/", showIndexPage)
	router.POST("/validate", validateJSON) // not used
	router.GET("/submitURL", showSubmitURL)
	router.POST("/submitURL", submitURL) // Send in the URL
	router.POST("/callback", callback)   // Check on state of UUID
	router.GET("/article/view/:article_id", getArticle)

	router.Static("/resources", "./resources")

	apiRoutes := router.Group("/api")
	apiRoutes.Use(APIAuthRequired)
	{
		apiRoutes.GET("/ver", getVer)
		apiRoutes.POST("/ver", getVer)
	}

	router.NoRoute(custom404)
}
