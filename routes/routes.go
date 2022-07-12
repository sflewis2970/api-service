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
	// Setup routes
	rs.Router.HandleFunc("/", controllers.Home)

	// Display log message
	log.Print("Setting up trivia service routes")

	// Client-side routes
	rs.Router.HandleFunc("/api/v1/trivia/question", controllers.GetQuestion)
	rs.Router.HandleFunc("/api/v1/trivia/answer", controllers.AnswerQuestion)
}

func New() *RoutingServer {
	rs := new(RoutingServer)

	// Initialize controller
	controllers.New()

	// Create mux Router and setup routes
	rs.Router = mux.NewRouter()
	rs.setupRoutes()

	return rs
}
