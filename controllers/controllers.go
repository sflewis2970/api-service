package controllers

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/sflewis2970/trivia-service/config"
	"github.com/sflewis2970/trivia-service/handlers/trivia"
)

// Controller struct definition
type Controller struct {
	cfg           *config.Config
	cfgData       *config.CfgData
	Router        *mux.Router
	triviaHandler *trivia.TriviaHandler
}

// Package controller object
var controller *Controller

func (c *Controller) setupRoutes() {
	// Display log message
	log.Print("Setting up trivia service routes")

	// Trivia routes
	c.Router.HandleFunc("/api/v1/trivia/getquestion", c.triviaHandler.GetTriviaQuestion).Methods("GET")
	c.Router.HandleFunc("/api/v1/trivia/submitanswer", c.triviaHandler.SubmitTriviaAnswer).Methods("POST")
}

// New Export functions
func New() *Controller {
	// Create controller component
	log.Print("Creating controller object...")
	controller = new(Controller)

	// Load config data
	var getCfgDataErr error
	controller.cfg = config.New()
	controller.cfgData, getCfgDataErr = controller.cfg.GetData()
	if getCfgDataErr != nil {
		log.Print("Error getting config data: ", getCfgDataErr)
		return nil
	}

	// Trivia handlers
	controller.triviaHandler = trivia.New()

	// Set controller routes
	controller.Router = mux.NewRouter()
	controller.setupRoutes()

	return controller
}
