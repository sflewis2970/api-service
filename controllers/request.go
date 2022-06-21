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

// Home is a http handler that receives messages from a client
func Home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Welcome to the trivia service app\n")
}

// GetQuestion is a http handler that receives a client request.
// Clients will send a request when they want to receive a trivia question from the trivia API.
// The format used is: 'http://<server-name>:8080/question?category=name'. category is optional
// When 'category' is supplied the trivia API returns a question related to the requested category
// When 'category' is omitted, the trivia API determines whether not the selected question is related
// to a category.
// The request returns a QuestionResponse object.
// The format for QuestionResponse is:
//       {"questionid": "<random_id>",
//        "question": "<question from trivia API>",
//        "category": "<category is not required and could be blank>",
//        "choices": "<choices are generated from API. One answer is correct, the otrhers are incorrect>",
//        "timestamp": "<formatted string of when the API returned the question>",
//        "warning": "<optional warning message>",
//        "error": "<optional error message>"}
func GetQuestion(rw http.ResponseWriter, r *http.Request) {
	// Get category from query parameter
	categoryStr := r.URL.Query().Get("category")

	// For now limit is NOT changeable
	limit := 0

	// Display a log message
	log.Print("data received from client...")

	// Question (Request) Response message
	var qResponse QuestionResponse

	// Initialize data store when needed
	requestComplete := false
	var apiResponseErr error
	var apiResponses []api.TriviaResponse
	apiResponsesSize := 0
	timestamp := ""

	// Check for duplicates for marking the request as complete
	for !requestComplete {
		// Send request to API
		apiResponses, timestamp, apiResponseErr = api.TriviaRequest(categoryStr, limit)

		// Get API Response size
		apiResponsesSize = len(apiResponses)

		if apiResponsesSize > 0 {
			// When results are returned, make sure there are no duplicate answers
			if !containsDuplicates(apiResponses) {
				log.Print("No duplicates found...")
				requestComplete = true
			} else {
				log.Print("Found duplicates...")
			}
		} else {
			// An error occurred or no results found
			requestComplete = true
		}
	}

	// Build API Response
	qResponse.Timestamp = timestamp

	if apiResponseErr != nil {
		// If an error occurs let the client know
		qResponse.Category = categoryStr
		qResponse.Error = apiResponseErr.Error()
	} else if apiResponsesSize == 0 {
		// If no error occurs but no response is returned
		// let the client know, the most likely case where this happens
		// is when the client supplies a category that does not exist
		qResponse.Warning = "No question was returned, select another category"
	} else {
		// Since the client is no longer allowed to supply a limit
		// there should be five items returned from the API
		// After getting a valid response from the API, generate a question ID
		qResponse.QuestionID = uuid.New().String()
		qResponse.QuestionID = common.BuildUUID(qResponse.QuestionID, DASH, ONE_SET)
		qResponse.Category = apiResponses[0].Category
		qResponse.Question = apiResponses[0].Question

		// Build choices string
		choiceList := []string{}
		for idx := 0; idx < apiResponsesSize; idx++ {
			choiceList = append(choiceList, apiResponses[idx].Answer)
		}

		// Shuttle list
		choiceList = common.ShuffleList(choiceList)

		// After the list has been shuffled build the string
		qResponse.Choices = common.BuildDelimitedStr(choiceList, ",")

		// Create data store table struct
		var dsTable datastore.DataStoreTable
		dsTable.Question = qResponse.Question
		dsTable.Category = qResponse.Category
		dsTable.Answer = apiResponses[0].Answer

		// Add question to data store
		cComponent.ds.AddQuestionAndAnswer(qResponse.QuestionID, dsTable)
	}

	// Write JSON to stream
	json.NewEncoder(rw).Encode(qResponse)

	// Display a log message
	log.Print("data sent back to client...")
}

// AnswerQuestion is a http handler that receives a response message from the client.
// The client is responsing to question received from the trivia API.
// The request uses the form of: 'http://<server-name>:8080/checkanswer' including a
// json object:
//        {"questionid": "<id received in the question response>",
//         "response": "<answer question from list of choices>"}
// The client will receive a response in the form of the following:
//       {"question": "<the question the client provided the answer for>",
//        "timestamp": "<formatted string of when the API returned the question>",
//        "category": "<if the question is linked to a category that information will be provided here>",
//        "response": "<the response the client provided>",
//        "answer": "<the answer to the question>",
//        "message": "<message to client whether or not question was answered correctly>",
//        "warning": "<optional warning message>",
//        "error": "<optional error message>"}
func AnswerQuestion(rw http.ResponseWriter, r *http.Request) {
	var aRequest AnswerRequest

	// Read JSON from stream
	json.NewDecoder(r.Body).Decode(&aRequest)

	// Initialize data store when needed
	var aResponse AnswerResponse
	timestamp, newQA, caErr := cComponent.ds.CheckAnswer(aRequest.QuestionID, aRequest.Response)
	if caErr != nil {
		log.Print("Error return from CheckAnswer call: ", caErr)
		aResponse.Error = caErr.Error()
	} else {
		// Build Response mesasge
		aResponse.Question = newQA.Question
		aResponse.Category = newQA.Category
		aResponse.Response = newQA.Response
		aResponse.Answer = newQA.Answer
		aResponse.Timestamp = timestamp
		aResponse.Message = newQA.Message
		aResponse.Warning = newQA.Warning
		aResponse.Error = newQA.Error
		aResponse.Correct = newQA.Correct
	}

	// Write JSON to stream
	json.NewEncoder(rw).Encode(aResponse)

	// Display a log message
	log.Print("data sent back to client...")
}
