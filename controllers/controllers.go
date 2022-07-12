package controllers

import (
	"log"

	apis "github.com/sflewis2970/trivia-service/api"
	"github.com/sflewis2970/trivia-service/config"
	"github.com/sflewis2970/trivia-service/models"
)

// Controller struct definition
type Controller struct {
	cfgData     *config.ConfigData
	publicAPI   *apis.API
	triviaModel *models.Model
}

// Packge controller object
var controller *Controller

// Export functions
func New() {
	if controller == nil {
		// Create controller component
		log.Print("Creating controller object...")
		controller = new(Controller)

		// Load config data
		var getCfgDataErr error
		controller.cfgData, getCfgDataErr = config.Get().GetData()
		if getCfgDataErr != nil {
			log.Print("Error getting config data: ", getCfgDataErr)
			return
		}

		// Create trivia model
		controller.triviaModel = models.New()

		// Create trivia api
		controller.publicAPI = apis.New()
	}
}
