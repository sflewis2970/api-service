package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/sflewis2970/trivia-service/api"
	"github.com/sflewis2970/trivia-service/common"
	"github.com/sflewis2970/trivia-service/datastore"
)

// Global questions datastore
var questionsDS *datastore.QuestionDS

const (
	DELIMITER     string = "-"
	NBR_OF_GROUPS int    = 1
)

type TriviaRequest struct {
	Request  string `json:"request"`
	Category string `json:"category"`
	Limit    string `json:"limit"`
}

type TriviaResponse struct {
	Request    string `json:"request"`
	Timestamp  string `json:"timestamp"`
	Category   string `json:"category"`
	QuestionID string `json:"questionid"`
	Question   string `json:"question"`
	Choices    string `json:"choices"`
	Warning    string `json:"warning,omitempty"`
	Error      string `json:"error,omitempty"`
}

type QuestionAnswer struct {
	Request    string `json:"request"`
	QuestionID string `json:"questionid"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
}

type QuestionResponse struct {
	Request    string `json:"request"`
	Timestamp  string `json:"timestamp"`
	QuestionID string `json:"questionid"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	Correct    bool   `json:"correct"`
	Message    string `json:"message,omitempty"`
	Warning    string `json:"warning,omitempty"`
	Error      string `json:"error,omitempty"`
}

func Home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Welcome to the trivia service app\n")
}

func GetQuestion(rw http.ResponseWriter, r *http.Request) {
	var tRequest TriviaRequest

	// Display a log message
	log.Print("data received from client...")

	// Decode request into JSON format
	json.NewDecoder(r.Body).Decode(&tRequest)

	// Initialize data store when needed
	if questionsDS == nil {
		log.Print("questions data store NOT ready...")
		questionsDS = datastore.InitializeDataStore()
	}

	var tResponse TriviaResponse

	// Send request to API
	log.Print("category: ", tRequest.Category)
	log.Print("limit: ", tRequest.Limit)
	apiResponseErr, apiResponses, timestamp := api.TriviaRequest(tRequest.Category, tRequest.Limit)

	// Build API Response
	tResponse.Request = tRequest.Request
	tResponse.Timestamp = timestamp

	if apiResponseErr != nil {
		tResponse.Category = tRequest.Category
		tResponse.Error = apiResponseErr.Error()
	} else {
		// After getting a valid response from the API, generate a question ID
		tResponse.QuestionID = uuid.New().String()
		tResponse.QuestionID = common.BuildUUID(tResponse.QuestionID, DELIMITER, NBR_OF_GROUPS)
		tResponse.Category = apiResponses[0].Category
		tResponse.Question = apiResponses[0].Question

		// Add question to data store
		questionsDS.AddQuestionAndAnswer(tResponse.QuestionID, tResponse.Question, apiResponses[0].Answer)

		// Build choices string
		apiResponsesSize := len(apiResponses)
		if apiResponsesSize == 1 {
			tResponse.Choices = tResponse.Choices + apiResponses[0].Answer
		} else {
			for idx := 0; idx < apiResponsesSize-1; idx++ {
				tResponse.Choices = tResponse.Choices + apiResponses[idx].Answer + ", "
			}
		}

		tResponse.Choices = tResponse.Choices + apiResponses[apiResponsesSize-1].Answer
	}

	// Write JSON to stream
	json.NewEncoder(rw).Encode(tResponse)

	// Display a log message
	log.Print("data sent back to client...")
}

func AnswerQuestion(rw http.ResponseWriter, r *http.Request) {
	var question QuestionAnswer

	// Display a log message
	log.Print("data received from client...")

	// Decode request into JSON format
	json.NewDecoder(r.Body).Decode(&question)

	// Initialize data store when needed
	var questionResponse QuestionResponse
	if questionsDS == nil {
		log.Print("Questions data store not ready...")
	} else {
		if questionsDS.CheckAnswer(question.QuestionID, question.Answer) {
			questionResponse.Message = "Congrats! That is correct"
			questionResponse.Correct = true
		} else {
			questionResponse.Message = "Nice try! That is NOT correct"
			questionResponse.Correct = false
		}
	}

	// Write JSON to stream
	json.NewEncoder(rw).Encode(questionResponse)

	// Display a log message
	log.Print("data sent back to client...")
}

func InitializeDataStore() {
	log.Print("initializing questions data store...")
	questionsDS = datastore.InitializeDataStore()
}
