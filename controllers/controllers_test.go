package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetQuestion(t *testing.T) {
	// Define input string
	jsonStr := []byte(`{"category":""}`)

	// Create new request
	request, reqErr := http.NewRequest("GET", "/trivia", bytes.NewBuffer(jsonStr))
	if reqErr != nil {
		t.Errorf("Could not create request.\n")
	}

	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetQuestion)
	handler.ServeHTTP(rr, request)

	// Check response code
	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned invalid status code: got %d, expected: %d\n", status, http.StatusOK)
	}

	// Check body
	expected := `{"request":"trivia"}`
	body := rr.Body.String()
	if body != expected {
		t.Errorf("Handler returned an unpected body: got %v, expected: %v", body, expected)
	}
}

func TestAnswerQuestion(t *testing.T) {
	// Define input string
	jsonStr := []byte(`{"questionid": "1e5c6450",
	                    "question": "In the song My Darling Clemantine how did Clemantine die",
						"choices": "Drowning"}`)

	// Create new request
	request, reqErr := http.NewRequest("GET", "/trivia", bytes.NewBuffer(jsonStr))
	if reqErr != nil {
		t.Errorf("Could not create request.\n")
	}

	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AnswerQuestion)
	handler.ServeHTTP(rr, request)

	// Check response code
	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned invalid status code: got %d, expected: %d\n", status, http.StatusOK)
	}

	// Check body
	expected := `{"request":"trivia"}`
	body := rr.Body.String()
	if body != expected {
		t.Errorf("Handler returned an unpected body: got %v, expected: %v", body, expected)
	}
}

func BenchmarkAnswerQuestion(b *testing.B) {
	// benchmark
	for idx := 0; idx < b.N; idx++ {
		// AnswerQuestion(timeNow, timeFormat)
	}
}
