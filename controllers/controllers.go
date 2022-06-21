package controllers

import (
	"github.com/sflewis2970/trivia-service/api"
	"github.com/sflewis2970/trivia-service/datastore"
)

const (
	DASH             string = "-"
	COMMA            string = ","
	SPACE            string = " "
	ONE_SET          int    = 1
	FIND_ERROR       int    = -1
	EMPTY_QUESTIONID int    = 0
	EMPTY_QUESTION   int    = 0
	EMPTY_ANSWER     int    = 0
	EMPTY_CATEGORY   int    = 0
	EMPTY_CHOICES    int    = 0
	EMPTY_TIMESTAMP  int    = 0
	EMPTY_WARNING    int    = 0
)

// Unexported package functions
// containsDuplicates checks the slice for any duplicate items
func containsDuplicates(items []api.TriviaResponse) bool {
	// Initialize the map for usage
	itemsMap := make(map[string]int)

	// Since maps uses unique keys, use the string value of answer to be the key
	for idx, item := range items {
		itemsMap[item.Answer] = idx + 1
	}

	// If the size of the map is the same size of the slice, then there are no duplicates
	if len(itemsMap) != len(items) {
		return true
	}

	// Otherwise return false
	return false
}

// Global controller component
var cComponent *controllerComponent

type controllerComponent struct {
	ds *datastore.QuestionDataStore
}

func (c *controllerComponent) initializeDataStore() *datastore.QuestionDataStore {
	// Create datastore component
	return datastore.GetDataStore()
}

// Export functions
func InitializeController() {
	// Create controller component
	cComponent = new(controllerComponent)
	cComponent.ds = cComponent.initializeDataStore()
}
