package controllers

import (
	"fmt"
	"testing"

	"github.com/sflewis2970/trivia-service/config"
)

func TestNewController(t *testing.T) {
	// Set environment variables
	controller := CreateNewController()

	if controller == nil {
		t.Errorf("could not create controller object")
	} else {
		controller.cfg = config.New()
		var cfgDataErr error
		controller.cfgData, cfgDataErr = controller.cfg.GetData()
		if cfgDataErr == nil {
			fmt.Printf("collected config data!")
		} else {
			t.Errorf("Error getting config data: %s", cfgDataErr.Error())
		}
	}
}

func CreateNewController() *Controller {
	controller := New()

	return controller
}
