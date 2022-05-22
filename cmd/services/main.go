package main

import (
	"log"
	"net/http"

	"github.com/sflewis2970/trivia-service/routes"
)

func main() {
	// Initialize logging
	log.SetFlags(log.Ldate | log.Lshortfile)

	// Create App
	routingServer := routes.CreateRoutingServer()

	// Setup routes
	routingServer.SetupRoutes()

	// Start Server
	log.Print("Web service server is ready...")
	log.Fatal(http.ListenAndServe(":8080", routingServer.Router))
}
