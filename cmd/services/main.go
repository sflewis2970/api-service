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

	// Get config data
	cfgData, getCfgDataErr := config.Get().GetData(config.UPDATE_CONFIG_DATA)
	if getCfgDataErr != nil {
		log.Fatal("Error getting config data: ", getCfgDataErr)
	}

	// Create App
	rs := routes.CreateRoutingServer()

	// Start Server
	log.Print("Web service server is ready...")

	addr := cfgData.HostName + cfgData.HostPort
	log.Print("The address used by the service is: ", addr)
	log.Fatal(http.ListenAndServe(addr, rs.Router))
}
