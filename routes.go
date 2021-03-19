// routes.go

package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)
	router.GET("/robin", showRobinPage)
	router.GET("/article/view/:article_id", getArticle)
	router.NoRoute(custom404)
}
