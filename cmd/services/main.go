package main

import (
	"log"
	"net/http"

	"github.com/sflewis2970/trivia-service/config"
	"github.com/sflewis2970/trivia-service/routes"
)

func main() {
	// Initialize logging
	log.SetFlags(log.Ldate | log.Lshortfile)

	// Create config object and load config data into memory
	log.Print("Loading config data...")
	cfg := config.GetConfig()

	// Create App
	rs := routes.CreateRoutingServer()

	// Start Server
	log.Print("Web service server is ready...")

	addr := cfg.GetConfigData().HostName + cfg.GetConfigData().Port
	log.Fatal(http.ListenAndServe(addr, rs.Router))
}
