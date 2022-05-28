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
	rs := routes.CreateRoutingServer()

	// Start Server
	log.Print("Web service server is ready...")
	log.Fatal(http.ListenAndServe(":8080", rs.Router))
}
