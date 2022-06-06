package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetQuestionNoParameters(t *testing.T) {
	// Initialize logging
	log.SetFlags(log.Ldate | log.Lshortfile)

	// Make sure that the datastore is ready
	initializeDS()

	// Create new request
	request, reqErr := http.NewRequest("GET", "/api/v1/question", nil)
	if reqErr != nil {
		t.Errorf("Could not create request.\n")
	}

	// Setup recoder
	rRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetQuestion)
	handler.ServeHTTP(rRecorder, request)

	// Check response code
	status := rRecorder.Code
	if status != http.StatusOK {
		t.Errorf("handler returned invalid status code: got %d, expected: %d\n", status, http.StatusOK)
	}

	// Unmarshal JSON
	bodyBytes := rRecorder.Body.Bytes()

	var qResponse QuestionResponse
	unmarshalErr := json.Unmarshal(bodyBytes, &qResponse)
	if unmarshalErr != nil {
		t.Errorf(unmarshalErr.Error())
	}

	// Check question ID field
	if len(qResponse.QuestionID) == EMPTY_QUESTIONID {
		t.Errorf("Handler returned an unexpected question ID field: got %s", qResponse.QuestionID)
	}

	// Check question field
	if len(qResponse.Question) == EMPTY_QUESTION {
		t.Errorf("Handler returned an unexpected question field: got %s", qResponse.Question)
	}

	// Check timestamp field
	if len(qResponse.Timestamp) == EMPTY_TIMESTAMP {
		t.Errorf("Handler returned an unexpected timestamp field: got %s", qResponse.Timestamp)
	}

	// Check choices field
	if len(qResponse.Choices) == EMPTY_CHOICES {
		t.Errorf("Handler returned an unexpected choices field: got %s", qResponse.Choices)
	}
}

func TestGetQuestionValidCategory(t *testing.T) {
	// Initialize logging
	log.SetFlags(log.Ldate | log.Lshortfile)

	// Make sure that the datastore is ready
	initializeDS()

	// Create new request
	request, reqErr := http.NewRequest("GET", "/api/v1/question?category=sciencenature", nil)
	if reqErr != nil {
		t.Errorf("Could not create request.\n")
	}

	// Setup recoder
	rRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetQuestion)
	handler.ServeHTTP(rRecorder, request)

	// Check response code
	status := rRecorder.Code
	if status != http.StatusOK {
		t.Errorf("handler returned invalid status code: got %d, expected: %d\n", status, http.StatusOK)
	}

	// Unmarshal JSON
	bodyBytes := rRecorder.Body.Bytes()
	var qResponse QuestionResponse
	unmarshalErr := json.Unmarshal(bodyBytes, &qResponse)
	if unmarshalErr != nil {
		t.Errorf(unmarshalErr.Error())
	}

	// Check question ID field
	if len(qResponse.QuestionID) == EMPTY_QUESTIONID {
		t.Errorf("Handler returned an unexpected question ID field: got %s", qResponse.QuestionID)
	}

	// Check question field
	if len(qResponse.Question) == EMPTY_QUESTION {
		t.Errorf("Handler returned an unexpected question field: got %s", qResponse.Question)
	}

	// Check category field
	if len(qResponse.Category) == EMPTY_CATEGORY {
		t.Errorf("Handler returned an unexpected category field: got %s", qResponse.Category)
	}

	// Check timestamp field
	if len(qResponse.Timestamp) == EMPTY_TIMESTAMP {
		t.Errorf("Handler returned an unexpected timestamp field: got %s", qResponse.Timestamp)
	}

	// Check choices field
	if len(qResponse.Choices) == EMPTY_CHOICES {
		t.Errorf("Handler returned an unexpected choices field: got %s", qResponse.Choices)
	}
}

func TestGetQuestionInValidCategory(t *testing.T) {
	// Initialize logging
	log.SetFlags(log.Ldate | log.Lshortfile)

	// Make sure that the datastore is ready
	initializeDS()

	// Create new request
	request, reqErr := http.NewRequest("GET", "/api/v1/question?category=apple", nil)
	if reqErr != nil {
		t.Errorf("Could not create request.\n")
	}

	// Setup recoder
	rRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetQuestion)
	handler.ServeHTTP(rRecorder, request)

	// Check response code
	status := rRecorder.Code
	if status != http.StatusOK {
		t.Errorf("handler returned invalid status code: got %d, expected: %d\n", status, http.StatusOK)
	}

	// Unmarshal JSON
	bodyBytes := rRecorder.Body.Bytes()
	var qResponse QuestionResponse
	unmarshalErr := json.Unmarshal(bodyBytes, &qResponse)
	if unmarshalErr != nil {
		t.Errorf(unmarshalErr.Error())
	}

	// Check question ID field
	if len(qResponse.QuestionID) != EMPTY_QUESTIONID {
		t.Errorf("Handler returned an unexpected question ID field: got %s", qResponse.QuestionID)
	}

	// Check question field
	if len(qResponse.Question) != EMPTY_QUESTION {
		t.Errorf("Handler returned an unexpected question field: got %s", qResponse.Question)
	}

	// Check warning field
	if len(qResponse.Warning) == EMPTY_WARNING {
		t.Errorf("Handler returned an unexpected warning field: got %s", qResponse.Warning)
	}
}

func TestAnswerQuestion(t *testing.T) {
	// Initialize logging
	log.SetFlags(log.Ldate | log.Lshortfile)

	// TestGetQuestionNoParameters(t)
	values := map[string]string{"foo": "baz"}
	jsonData, marshalErr := json.Marshal(values)

	if marshalErr != nil {
		t.Errorf("New request error: %s", marshalErr.Error())
	}

	// Create new request
	request, reqErr := http.NewRequest("GET", "/answer", bytes.NewBuffer(jsonData))
	if reqErr != nil {
		t.Errorf("New request error: %s", reqErr.Error())
	}

	rRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(AnswerQuestion)
	handler.ServeHTTP(rRecorder, request)

	// Check response code
	status := rRecorder.Code
	if status != http.StatusOK {
		t.Errorf("handler returned invalid status code: got %d, expected: %d\n", status, http.StatusOK)
	}

	// Unmarshal JSON
	bodyBytes := rRecorder.Body.Bytes()
	var aResponse AnswerResponse
	unmarshalErr := json.Unmarshal(bodyBytes, &aResponse)
	if unmarshalErr != nil {
		t.Errorf(unmarshalErr.Error())
	}
}
