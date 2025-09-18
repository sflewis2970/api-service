package trivia

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/sflewis2970/trivia-service/config"
)

func initialize() {
	// Initialize logging
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// set environment variables
	_ = os.Setenv(config.HOST, "")
	_ = os.Setenv(config.PORT, "8080")
	_ = os.Setenv(config.NBR_OF_RETRIES, "3")

	// Go-redis settings
	_ = os.Setenv(config.REDIS_TLS_URL, "cache")
	_ = os.Setenv(config.REDIS_URL, "cache")
	_ = os.Setenv(config.REDIS_PORT, "6379")

	// DB Test setting
	// For now, we will test against a real database. To do so uncomment the following line
	// Later, database mocking will be added
	_ = os.Setenv("DB_TEST", "TESTDB")

	// Create config object
	// Get config data
	cfg := config.New()
	cfg.GetData(config.REFRESH_CONFIG_DATA)
}

func GetTriviaQuestionTestCase() error {
	// Create request.
	request, reqErr := http.NewRequest("GET", "/api/v1/trivia/questions", nil)
	if reqErr != nil {
		return errors.New("error creating request")
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
		errMsg := fmt.Sprintf("handler returned wrong status code: got: %d, expected: %d", gotStatus, http.StatusOK)
		return errors.New(errMsg)
	}

	return nil
}

func TestGetTriviaQuestion(t *testing.T) {
	t.Skip() // skipping unit until the decision is made whether to use miniredis or use mocking

	// Create request
	request, reqErr := http.NewRequest("GET", "/api/v1/trivia/getquestion", nil)
	if reqErr != nil {
		return
	}

	// We create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	triviaHandler := CreateNewHandlerObject()
	handler := http.HandlerFunc(triviaHandler.GetTriviaQuestion)

	// Server request
	handler.ServeHTTP(rr, request)

	// Check status code.
	gotStatus := rr.Code
	if (gotStatus != http.StatusOK) && (gotStatus != http.StatusCreated) {
		errMsg := fmt.Sprintf("handler returned wrong status code: got: %d, expected: %d", gotStatus, http.StatusOK)
		log.Print(errMsg)
		return
	}

}

func SubmitTriviaAnswerTestCase() error {
	// Create jsonData
	var jsonData = []byte(`{
		"questionid": "8cd76569",
		"response": "2"
	}`)

	// Create request.
	request, reqErr := http.NewRequest("POST", "/api/v1/trivia/questions", bytes.NewBuffer(jsonData))
	if reqErr != nil {
		return errors.New("error creating request")
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
		errMsg := fmt.Sprintf("handler returned wrong status code: got: %d, expected: %d", gotStatus, http.StatusOK)
		return errors.New(errMsg)
	}

	return nil
}

func TestSubmitTriviaAnswer(t *testing.T) {
	t.Skip() // skipping unit until the decision is made whether to use miniredis or use mocking
	initialize()

	dbTest := os.Getenv("DB_TEST")
	if len(dbTest) == 0 {
		log.Print("Environment variable not set...skipping test")
		t.Skip()
	}

	// Create Trivia Question
	tcErr := GetTriviaQuestionTestCase()
	if tcErr != nil {
		t.Errorf("%s\n", tcErr.Error())
	}

	// Submit Trivia Question Answer
	tcErr = SubmitTriviaAnswerTestCase()
	if tcErr != nil {
		t.Errorf("%s\n", tcErr.Error())
	}
}

func TestNewHandler(t *testing.T) {
	newHandler := CreateNewHandlerObject()

	if newHandler == nil {
		t.Errorf("Could nort create Handler object!")
	} else {
		fmt.Println("Created trivia handler!")
	}
}

func CreateNewHandlerObject() *TriviaHandler {
	newHandler := New()

	return newHandler

}

func BenchmarkGetTriviaQuestion(b *testing.B) {
	initialize()

	dbTest := os.Getenv("DB_TEST")
	if len(dbTest) == 0 {
		log.Print("Environment variable not set...skipping test")
		b.Skip()
	}

	for idx := 0; idx < b.N; idx++ {
		// Create Trivia Question
		tcErr := GetTriviaQuestionTestCase()
		if tcErr != nil {
			b.Errorf("%s\n", tcErr.Error())
		}
	}
}

func BenchmarkSubmitTriviaAnswer(b *testing.B) {
	initialize()

	dbTest := os.Getenv("DB_TEST")
	if len(dbTest) == 0 {
		log.Print("Environment variable not set...skipping test")
		b.Skip()
	}

	// Create Trivia Question
	tcErr := GetTriviaQuestionTestCase()
	if tcErr != nil {
		b.Errorf("%s\n", tcErr.Error())
	}

	for idx := 0; idx < b.N; idx++ {
		// Submit Trivia Question Answer
		tcErr = SubmitTriviaAnswerTestCase()
		if tcErr != nil {
			b.Errorf("%s\n", tcErr.Error())
		}
	}
}
