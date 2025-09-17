package models

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/sflewis2970/trivia-service/common"
	"github.com/sflewis2970/trivia-service/config"
	"github.com/sflewis2970/trivia-service/messages"
)

var model *TriviaModel

func initialize() {
	// Initialize logging
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// set environment variables
	_ = os.Setenv(config.HOST, "")
	_ = os.Setenv(config.PORT, "8080")
	_ = os.Setenv(config.NBR_OF_RETRIES, "3")

	// Go-redis settings
	_ = os.Setenv(config.REDIS_TLS_URL, "localhost")
	_ = os.Setenv(config.REDIS_URL, "localhost")
	_ = os.Setenv(config.REDIS_PORT, "6379")

	// DB Test setting
	// For now, we will test against a real database. To do so uncomment the following line
	// Later, database mocking will be added
	_ = os.Setenv("DB_TEST", "TESTDB")

	// Create config object
	// Get config data
	cfg := config.New()
	cfg.GetData(config.REFRESH_CONFIG_DATA)

	// Create TriviaModel
	model = NewTriviaModel()
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

func AddTriviaQuestionTestCase() error {
	// Create trivia object
	trivia := createTrivia()

	// Add trivia to database
	addErr := model.AddTriviaQuestion(trivia)
	if addErr != nil {
		errMsg := fmt.Sprintf("AddTriviaQuestion(%v): unexpected error occurred, call returned %v", trivia, addErr.Error())
		log.Print(errMsg)
		return addErr
	}

	return nil
}

func TestAddAndRetrieveTriviaQuestion(t *testing.T) {
	TestAddTriviaQuestion(t)
	TestGetTriviaAnswer(t)
}

func TestAddTriviaQuestion(t *testing.T) {
	initialize()

	// Check model creation
	if model == nil {
		t.Errorf("model has invalid value: %v", model)
	}

	// Check dbtest settings
	dbTest := os.Getenv("DB_TEST")
	if len(dbTest) == 0 {
		log.Print("Environment variable not set...skipping test")
		t.Skip()
	}

	addErr := AddTriviaQuestionTestCase()
	if addErr != nil {
		t.Errorf("AddTriviaQuestionTestCase() return an error: %v", addErr.Error())
	}
}

func GetTriviaAnswerTestCase() error {
	// Simulate user responses
	responses := []string{"1", "2", "3", "4", "999"}

	for _, response := range responses {
		aRequest := createAnswerRequest(response)
		_, answerErr := model.SubmitTriviaAnswer(aRequest)
		if answerErr != nil {
			errMsg := fmt.Sprintf("GetTriviaAnswer(%v): unexpected error occurred, call returned %v", aRequest, answerErr.Error())
			return errors.New(errMsg)
		}
	}

	return nil
}

func TestGetTriviaAnswer(t *testing.T) {
	initialize()

	// Check model creation
	if model == nil {
		t.Errorf("model has invalid value: %v", model)
	}

	// Check dbtest settings
	dbTest := os.Getenv("DB_TEST")
	if len(dbTest) == 0 {
		log.Print("Environment variable not set...skipping test")
		t.Skip()
	}

	getErr := GetTriviaAnswerTestCase()
	if getErr != nil {
		t.Errorf("GetTriviaAnswerTestCase() return an error: %v", getErr.Error())
	}
}

func DeleteTriviaQuestionTestCase() error {
	// Create trivia object
	trivia := createTrivia()

	// Add trivia to database
	addErr := model.AddTriviaQuestion(trivia)
	if addErr != nil {
		errMsg := fmt.Sprintf("AddTriviaQuestion(%v): unexpected error occurred, call returned %v", trivia, addErr.Error())
		return errors.New(errMsg)
	}

	// delete trivia from database
	delErr := model.DeleteTriviaQuestion(trivia.QuestionID)
	if delErr != nil {
		errMsg := fmt.Sprintf("DeleteTriviaQuestion(%v): unexpected error occurred, call returned %v", trivia.QuestionID, delErr.Error())
		return errors.New(errMsg)
	}

	return nil
}

func TestDeleteTriviaQuestion(t *testing.T) {
	initialize()

	// Check model creation
	if model == nil {
		t.Errorf("model has invalid value: %v", model)
	}

	// Check dbtest setting
	dbTest := os.Getenv("DB_TEST")
	if len(dbTest) == 0 {
		log.Print("Environment variable not set...skipping test")
		t.Skip()
	}

	delErr := DeleteTriviaQuestionTestCase()
	if delErr != nil {
		t.Errorf("DeleteTriviaQuestionTestCase() return an error: %v", delErr.Error())
	}
}

func TestNewModel(t *testing.T) {
	// Initialize environment for testing
	initialize()

	// Check model creation
	if model == nil {
		t.Errorf("model has invalid value: %v", model)
	}
}

func BenchmarkAddTriviaQuestion(b *testing.B) {
	initialize()

	// Check model creation
	if model == nil {
		b.Errorf("model has invalid value: %v", model)
	}

	dbTest := os.Getenv("DB_TEST")
	if len(dbTest) == 0 {
		log.Print("Environment variable not set...skipping test")
		b.Skip()
	}

	for idx := 0; idx < b.N; idx++ {
		// Create Trivia Question
		tcErr := AddTriviaQuestionTestCase()
		if tcErr != nil {
			b.Errorf("%s\n", tcErr.Error())
		}
	}
}

func BenchmarkGetTriviaAnswer(b *testing.B) {
	initialize()

	// Check model creation
	if model == nil {
		b.Errorf("model has invalid value: %v", model)
	}

	dbTest := os.Getenv("DB_TEST")
	if len(dbTest) == 0 {
		log.Print("Environment variable not set...skipping test")
		b.Skip()
	}

	for idx := 0; idx < b.N; idx++ {
		// Create Trivia Question
		tcErr := GetTriviaAnswerTestCase()
		if tcErr != nil {
			b.Errorf("%s\n", tcErr.Error())
		}
	}
}

func BenchmarkDeleteTriviaQuestion(b *testing.B) {
	initialize()

	// Check model creation
	if model == nil {
		b.Errorf("model has invalid value: %v", model)
	}

	dbTest := os.Getenv("DB_TEST")
	if len(dbTest) == 0 {
		log.Print("Environment variable not set...skipping test")
		b.Skip()
	}

	for idx := 0; idx < b.N; idx++ {
		// Create Trivia Question
		tcErr := DeleteTriviaQuestionTestCase()
		if tcErr != nil {
			b.Errorf("%s\n", tcErr.Error())
		}
	}
}
