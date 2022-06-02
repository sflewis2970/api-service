package routes

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/sflewis2970/trivia-service/controllers"
)

type RoutingServer struct {
	Router *mux.Router
}

func (rs *RoutingServer) setupRoutes() {
	// Display log message
	log.Print("Setting up web service routes")

	// Initialize Datastore before receiving any messages
	controllers.InitializeController(rs.Router)

	// Setup routes
	rs.Router.HandleFunc("/", controllers.Home)
	rs.Router.HandleFunc("/api/v1/question", controllers.GetQuestion)
	rs.Router.HandleFunc("/api/v1/answer", controllers.AnswerQuestion)
}

func CreateRoutingServer() *RoutingServer {
	rs := new(RoutingServer)

	rs.Router = mux.NewRouter()
	rs.setupRoutes()

	return rs
}
