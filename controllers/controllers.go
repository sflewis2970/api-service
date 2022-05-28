package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sflewis2970/trivia-service/api"
	"github.com/sflewis2970/trivia-service/common"
	"github.com/sflewis2970/trivia-service/datastore"
)

// Global questions datastore
var questionsDS *datastore.QuestionDS

const (
	DASH          string = "-"
	COMMA         string = ","
	SPACE         string = " "
	NBR_OF_GROUPS int    = 1
	FIND_ERROR    int    = -1
)

type controllerComponents struct {
	useLocalDB bool
	muxRouter  *mux.Router
}

var controllerComponent *controllerComponents

type QuestionResponse struct {
	QuestionID string `json:"questionid,omitempty"`
	Question   string `json:"question"`
	Timestamp  string `json:"timestamp"`
	Category   string `json:"category"`
	Choices    string `json:"choices"`
	Warning    string `json:"warning,omitempty"`
	Error      string `json:"error,omitempty"`
}

type AnswerResponse struct {
	QuestionID string `json:"questionid,omitempty"`
	Question   string `json:"question"`
	Timestamp  string `json:"timestamp"`
	Category   string `json:"category"`
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
	// Extract data from request
	vars := mux.Vars(r)
	category := vars["category"]
	limit := vars["limit"]

	// Display a log message
	log.Print("data received from client...")

	// Initialize data store when needed
	if questionsDS == nil {
		log.Print("questions data store NOT ready...")
	}

	var qResponse QuestionResponse

	// Send request to API
	apiResponseErr, apiResponses, timestamp := api.TriviaRequest(category, limit)

	// set API Response size
	apiResponsesSize := len(apiResponses)

	// Build API Response
	qResponse.Timestamp = timestamp

	if apiResponseErr != nil {
		qResponse.Category = category
		qResponse.Error = apiResponseErr.Error()
	} else if apiResponsesSize == 0 {
		qResponse.Warning = "No question was returned, select another category"
	} else {
		// After getting a valid response from the API, generate a question ID
		qResponse.QuestionID = uuid.New().String()
		qResponse.QuestionID = common.BuildUUID(qResponse.QuestionID, DASH, NBR_OF_GROUPS)
		qResponse.Category = apiResponses[0].Category
		qResponse.Question = apiResponses[0].Question
		questionPos := 0

		// Build choices string
		if apiResponsesSize == 1 {
			qResponse.Choices = apiResponses[0].Answer
		} else {
			// Build choices string
			choiceList := []string{}
			for idx := 0; idx < apiResponsesSize; idx++ {
				choiceList = append(choiceList, apiResponses[idx].Answer)
			}

			// Shuttle list
			choiceList = common.ShuffleList(choiceList)

			// Find answer index position
			questionPos = common.FindPosition(choiceList, apiResponses[0].Answer)
			if questionPos > FIND_ERROR {
				questionPos++
			}

			// After the list has been shuffled build the string
			qResponse.Choices = common.BuildDelimitedStr(choiceList, ",")
		}

		// Create Q&A stuct object
		qa := datastore.QuestionAndAnswer{}
		qa.Question = qResponse.Question
		qa.Category = qResponse.Category
		qa.AnswerPos = questionPos
		qa.Answer = apiResponses[0].Answer

		// Add question to data store
		questionsDS.AddQuestionAndAnswer(qResponse.QuestionID, qa)
	}

	// Write JSON to stream
	json.NewEncoder(rw).Encode(qResponse)

	// Display a log message
	log.Print("data sent back to client...")
}

func AnswerQuestion(rw http.ResponseWriter, r *http.Request) {
	// Extract data from request
	vars := mux.Vars(r)
	questionID := vars["questionID"]
	answer := vars["answer"]

	// Initialize data store when needed
	var aResponse AnswerResponse
	if questionsDS == nil {
		log.Print("Questions data store not ready...")
		aResponse.Error = "Datastore unavailable"
	} else {
		timestamp, newQA := questionsDS.CheckAnswer(questionID, answer)
		aResponse.Question = newQA.Question
		aResponse.Category = newQA.Category
		aResponse.Answer = newQA.Answer
		aResponse.Timestamp = timestamp
		aResponse.Message = newQA.Message
		aResponse.Correct = newQA.Correct
	}

	// Write JSON to stream
	json.NewEncoder(rw).Encode(aResponse)

	// Display a log message
	log.Print("data sent back to client...")
}

func InitializeController(muxRouter *mux.Router) {
	// Check to see if datastore server
	log.Print("initializing questions data store...")
	questionsDS = datastore.InitializeDataStore()

	// Controller Components
	controllerComponent := new(controllerComponents)
	controllerComponent.muxRouter = muxRouter
}
