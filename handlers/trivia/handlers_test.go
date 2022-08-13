package trivia

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTriviaQuestion(t *testing.T) {
	// Create request.
	request, reqErr := http.NewRequest("GET", "/api/v1/trivia/questions", nil)
	if reqErr != nil {
		t.Error("Error creating request")
	}

	// We create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	triviaHandler := New()
	handler := http.HandlerFunc(triviaHandler.GetTriviaQuestion)

	// Server request
	handler.ServeHTTP(rr, request)

	// Check status code.
	gotStatus := rr.Code
	if gotStatus != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got: %v, expected: %v", gotStatus, http.StatusOK)
	}
}

func TestSubmitTriviaAnswer(t *testing.T) {
	// Create jsonData
	var jsonData = []byte(`{
		"questionid": "8cd76569",
		"response": "2"
	}`)

	// Create request.
	request, reqErr := http.NewRequest("POST", "/api/v1/trivia/questions", bytes.NewBuffer(jsonData))
	if reqErr != nil {
		t.Error("Error creating request")
	}

	// We create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	triviaHandler := New()
	handler := http.HandlerFunc(triviaHandler.SubmitTriviaAnswer)

	// Server request
	handler.ServeHTTP(rr, request)

	// Check status code.
	gotStatus := rr.Code
	if gotStatus != http.StatusOK {
		t.Errorf("handler returned wrong status code: got: %v, expected: %v", gotStatus, http.StatusOK)
	}
}

func BenchmarkGetTriviaQuestion(b *testing.B) {
	for idx := 0; idx < b.N; idx++ {
	}
}

func BenchmarkSubmitTriviaAnswer(b *testing.B) {
	for idx := 0; idx < b.N; idx++ {
	}
}
