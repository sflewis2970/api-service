package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sflewis2970/trivia-service/messages"
)

// Home is a http handler that receives messages from a client
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the trivia service app\n")
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
func GetQuestion(w http.ResponseWriter, r *http.Request) {
	// Display a log message
	log.Print("data received from client...")

	// Get category from query parameter
	// For now this is not used but may be used at a later date
	r.URL.Query().Get("category")

	// Process API Get Request
	qResponse, dsTable := controller.publicAPI.GetQuestion()

	// Insert question into datastore
	insertErr := controller.triviaModel.InsertQuestion(qResponse.QuestionID, dsTable)

	// Add question to data store
	if insertErr != nil {
		errMsg := "Datastore: Insert error..."
		log.Print(errMsg, ": ", insertErr)
		qResponse.QuestionID = ""
		qResponse.Category = ""
		qResponse.Question = ""
		qResponse.Choices = []string{}
		qResponse.Error = errMsg
	}

	// Write JSON to stream
	json.NewEncoder(w).Encode(qResponse)

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
func AnswerQuestion(w http.ResponseWriter, r *http.Request) {
	var aRequest messages.AnswerRequest

	// Read JSON from stream
	json.NewDecoder(r.Body).Decode(&aRequest)

	// Check the model for the answer
	aResponse := controller.triviaModel.AnswerQuestion(aRequest)

	// Write JSON to stream
	json.NewEncoder(w).Encode(aResponse)

	// Display a log message
	log.Print("data sent back to client...")
}
