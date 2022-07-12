package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/sflewis2970/trivia-service/config"
	"github.com/sflewis2970/trivia-service/routes"
)

func main() {
	// Initialize logging
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Get config data
	cfgData, getCfgDataErr := config.Get().GetData(config.UPDATE_CONFIG_DATA)
	if getCfgDataErr != nil {
		log.Fatal("Error getting config data: ", getCfgDataErr)
	}

	// Create App
	rs := routes.New()

	// setup Cors
	log.Print("Setting up CORS...")
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodPost, http.MethodGet},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	corsHandler := cors.Handler(rs.Router)

	// Server Address info
	addr := cfgData.HostName + cfgData.HostPort
	log.Print("The address used by the service is: ", addr)

	// Start Server
	log.Print("Web service server is ready...")

	// Listen and Serve
	log.Fatal(http.ListenAndServe(addr, corsHandler))
}
