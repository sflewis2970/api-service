package models

import (
	"github.com/sflewis2970/trivia-service/common"
	"github.com/sflewis2970/trivia-service/config"
	"github.com/sflewis2970/trivia-service/messages"
	"log"
	"os"
	"testing"
	"time"
)

var model *TriviaModel

func initialize(t *testing.T) {
	// Initialize logging
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// set environment variables
	_ = os.Setenv(config.HOST, "")
	_ = os.Setenv(config.PORT, "8080")

	// Go-redis settings
	_ = os.Setenv(config.REDIS_TLS_URL, "localhost")
	_ = os.Setenv(config.REDIS_URL, "localhost")
	_ = os.Setenv(config.REDIS_PORT, "6379")

	// Create config object
	// Get config data
	_, cfgDataErr := config.Get().GetData(config.UPDATE_CONFIG_DATA)
	if cfgDataErr != nil {
		t.Errorf("GetData(%v): error not expected, got %v", config.UPDATE_CONFIG_DATA, cfgDataErr)
	}

	// Create TriviaModel
	model = NewTriviaModel()
	if model == nil {
		t.Errorf("NewTriviaModel(): unexpected error occurred, call returned %v", model)
	}
}

func createTrivia() messages.Trivia {
	var trivia messages.Trivia

	trivia.QuestionID = "8cd76569"
	trivia.Question = "What is 1 + 1?"
	trivia.Category = "mathematics"
	trivia.Answer = "2"
	trivia.Choices = []string{"Make Selection from list...", "1", "2", "3", "4", "846"}
	trivia.Timestamp = common.GetFormattedTime(time.Now(), "Mon Jan 2 15:04:05 2006")

	return trivia
}

func createAnswerRequest(response string) messages.AnswerRequest {
	var aRequest messages.AnswerRequest

	aRequest.QuestionID = "8cd76569"
	aRequest.Response = response

	return aRequest
}

func TestAddTriviaQuestion(t *testing.T) {
	initialize(t)
	trivia := createTrivia()
	_ = model.AddTriviaQuestion(trivia)
	/*
		if addErr != nil {
			t.Errorf("AddTriviaQuestion(%v): unexpected error occurred, call returned %v", trivia, addErr.Error())
		}
	*/
}

func TestGetTriviaQuestion(t *testing.T) {
	initialize(t)
	responses := []string{"1", "2", "3", "4", "999"}

	for _, response := range responses {
		aRequest := createAnswerRequest(response)
		_, answerErr := model.GetTriviaAnswer(aRequest)
		if answerErr != nil {
			t.Errorf("GetTriviaAnswer(%v): unexpected error occurred, call returned %v", aRequest, answerErr.Error())
		}
	}
}

func TestDeleteTriviaQuestion(t *testing.T) {
	initialize(t)
	trivia := createTrivia()
	addErr := model.AddTriviaQuestion(trivia)
	if addErr != nil {
		t.Errorf("AddTriviaQuestion(%v): unexpected error occurred, call returned %v", trivia, addErr.Error())
	}

	delErr := model.DeleteTriviaQuestion(trivia.QuestionID)
	if delErr != nil {
		t.Errorf("DeleteTriviaQuestion(%v): unexpected error occurred, call returned %v", trivia.QuestionID, delErr.Error())
	}
}

func TestNewModel(t *testing.T) {
	// Initialize environment for testing
	initialize(t)
}
