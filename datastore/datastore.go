package datastore

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/sflewis2970/trivia-service/config"
)

var qds *QuestionDataStore

type QuestionDataStore struct {
	cfgData      *config.ConfigData
	serverStatus StatusCode
}

// AddQuestionAndAnswer sends a request to the DataStore server to add a question to the datastore
func (q *QuestionDataStore) Insert(questionID string, dst DataStoreTable) error {
	// Create AddQuestionRequest message
	var aqRequest AddQuestionRequest

	// Build AddQuestionRequest
	aqRequest.QuestionID = questionID
	aqRequest.Question = dst.Question
	aqRequest.Category = dst.Category
	aqRequest.Answer = dst.Answer

	// Build request body from AddQuestionRequest
	requestBody, marshalErr := json.Marshal(aqRequest)
	if marshalErr != nil {
		log.Print("marshaling error: ", marshalErr)
		return marshalErr
	}

	// Post request
	url := q.cfgData.DataStoreName + q.cfgData.DataStorePort + DS_INSERT_PATH
	response, postErr := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if postErr != nil {
		return postErr
	}
	defer response.Body.Close()

	// Handle add question response
	var aqResponse AddQuestionResponse

	// Read response stream into JSON
	json.NewDecoder(response.Body).Decode(&aqResponse)

	return nil
}

// CheckAnswer sends a request to the DataStore server to determine if the question was answered correctly
func (q *QuestionDataStore) Get(questionID string) (string, *QuestionAndAnswer, error) {
	timestamp := ""
	var caRequest CheckAnswerRequest

	// Build add question request
	caRequest.QuestionID = questionID

	// Convert struct to byte array
	requestBody, marshalErr := json.Marshal(caRequest)
	if marshalErr != nil {
		log.Print("marshaling error: ", marshalErr)
		return "", nil, marshalErr
	}

	// Create a http request
	url := q.cfgData.DataStoreName + q.cfgData.DataStorePort + DS_GET_PATH
	response, postErr := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if postErr != nil {
		return "", nil, postErr
	}
	defer response.Body.Close()

	// Handle add question response
	var caResponse CheckAnswerResponse

	// Read response stream into JSON
	json.NewDecoder(response.Body).Decode(&caResponse)

	// Update QuestionAndAnswer struct
	timestamp = caResponse.Timestamp

	var newQA *QuestionAndAnswer
	newQA = new(QuestionAndAnswer)
	newQA.Question = caResponse.Question
	newQA.Category = caResponse.Category
	newQA.Answer = caResponse.Answer
	newQA.Message = caResponse.Message
	newQA.Warning = caResponse.Warning
	newQA.Error = caResponse.Error

	return timestamp, newQA, nil
}

// SendStatusRequest sends a request for the status of the datastore server
func (q *QuestionDataStore) sendStatusRequest() StatusCode {
	url := q.cfgData.DataStoreName + q.cfgData.DataStorePort + DS_STATUS_PATH
	log.Print("sending status request to ", url)

	// http request
	response, getErr := http.Get(url)
	if getErr != nil {
		log.Print("A response error has occurred...")
		return StatusCode(DS_REQUEST_ERROR)
	}
	defer response.Body.Close()

	// Status (Request) Response
	var sResponse StatusResponse

	// Read JSON from stream
	json.NewDecoder(response.Body).Decode(&sResponse)

	return sResponse.Status
}

// CreateDataStore prepares the datastore component waits for the datastore server before allowing messages to be sent
func Initialize() *QuestionDataStore {
	if qds == nil {
		// Create QuestionDataStore object
		log.Print("Creating QuestionDataStore object")
		qds = new(QuestionDataStore)

		// Update QuestionDataStore fields
		var getCfgDataErr error
		qds.cfgData, getCfgDataErr = config.Get().GetData()
		if getCfgDataErr != nil {
			log.Print("error getting config data...")
			return nil
		}

		qds.serverStatus = StatusCode(DS_NOT_STARTED)

		// Wait for DataStore server to become available
		for qds.serverStatus != StatusCode(DS_RUNNING) {
			// Get datastore server status
			qds.serverStatus = qds.sendStatusRequest()

			// Once the datastore is up and running get out!
			if qds.serverStatus == StatusCode(DS_RUNNING) {
				break
			} else {
				log.Print("waiting for Datastore server...")
			}

			// Sleep for 3 seconds
			time.Sleep(time.Second * 3)
		}
	}

	return qds
}
