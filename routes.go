// routes.go

package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)
	router.GET("/robin", showRobinPage)
}
