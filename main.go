// main.go

package main

import "routers"

func main() {
	// Initialize the routes
	router := routers.InitializeRoutes()

	// Start serving the application
	router.Run(":1668")
}
