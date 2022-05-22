package routes

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/sflewis2970/trivia-service/controllers"
)

type RoutingServer struct {
	Router *mux.Router
}

func (rs *RoutingServer) SetupRoutes() {
	// Display log message
	log.Print("Setting up web service routes")

	// Initialize Datastore before receiving any messages
	controllers.InitializeDataStore()

	// Setup routes
	rs.Router.HandleFunc("/", controllers.Home)
	rs.Router.HandleFunc("/trivia", controllers.GetQuestion).Methods("GET")
	rs.Router.HandleFunc("/answer", controllers.AnswerQuestion).Methods("GET")
}

func CreateRoutingServer() *RoutingServer {
	rs := new(RoutingServer)

	// Create router
	rs.Router = mux.NewRouter()

	return rs
}
