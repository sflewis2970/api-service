package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sflewis2970/trivia-service/config"
	"github.com/sflewis2970/trivia-service/routes"
)

func setCfgEnv() {
	// Set hostname environment variable
	os.Setenv(config.HOSTNAME, "")

	// Set hostport environment variable
	os.Setenv(config.HOSTPORT, ":8080")

	// Set Datastore Server address or DNS environment variable
	os.Setenv(config.DSNAME, "http://ds-service")
	os.Setenv(config.DSPORT, ":9090")

	// Response messages
	os.Setenv(config.CONGRATS, "Congrats! That is correct")
	os.Setenv(config.TRYAGAIN, "Nice try! Better luck on the next question")
}

func main() {
	// Initialize logging
	log.SetFlags(log.Ldate | log.Lshortfile)

	useCfgFile := os.Getenv("USECONFIGFILE")
	if len(useCfgFile) == 0 {
		setCfgEnv()
	}

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
