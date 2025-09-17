package trivia

import (
	"fmt"
	"log"
	"testing"

	"github.com/sflewis2970/trivia-service/config"
)

func TestGetTrivia(t *testing.T) {
	api := CreateNewAPI(t)

	if api == nil {
		t.Errorf("Could not create api object with no error was returned")
	} else {
		triviaMsg, triviaErr := api.GetTrivia()
		if triviaErr != nil {
			t.Errorf("GetTrivia() returned an error: %s\n", triviaErr.Error())
			if len(triviaMsg.Question) == 0 {
				t.Error("Trivia msg doesn't contain a question")
			}
		} else {
			fmt.Println("Trivia category: ", triviaMsg.Category)
			fmt.Println("Trivia question: ", triviaMsg.Question)
		}
	}
}

func TestTriviaRequest(t *testing.T) {
	// Create api object
	api := New()

	// Test cases
	testCases := []struct {
		timestamp string
		error     error
	}{}

	for _, value := range testCases {
		fmt.Println("value: ", value.timestamp)
		gotResponses, gotTimestamp, gotError := api.triviaRequest()

		ResponseSize := len(gotResponses)
		if ResponseSize == 0 {
			t.Errorf("triviaRequest(): did not return a response.")
		}

		if gotError != nil {
			t.Errorf("TriviaRequest(): error not expected, got %v", gotError)
		}

		if len(gotTimestamp) == 0 {
			t.Errorf("TriviaRequest(): did not return a valid time stamp, got %v", gotTimestamp)
		}

		if gotError != nil {
			t.Errorf("TriviaRequest(): error not expected, got %v", gotError)
		}

		if len(gotTimestamp) == 0 {
			t.Errorf("TriviaRequest(): did not return a valid time stamp, got %s", gotTimestamp)
		}

		if gotError != nil {
			t.Errorf("TriviaRequest(): error not expected, got %v", gotError)
		}

		if len(gotTimestamp) == 0 {
			t.Errorf("TriviaRequest(): did not return a valid time stamp, got %s", gotTimestamp)
		}
	}
}

func TestReturnMultipleAnswers(t *testing.T) {
	api := CreateNewAPI(t)

	if api == nil {
		t.Errorf("API object not created!")
	} else {
		answers, answersErr := api.returnMultipleAnswers(AnswerCount)
		if answersErr != nil {
			t.Errorf("error returned an error: %s\n", answersErr.Error())
		} else {
			nbrOfAnswers := len(answers)
			for idx := 0; idx < nbrOfAnswers; idx++ {
				fmt.Println("", answers[idx])
			}
		}
	}
}

func TestNewAPI(t *testing.T) {
	api := CreateNewAPI(t)

	if api == nil {
		t.Errorf("api object is nil, cannot continue...")
	} else {
		fmt.Println("api object created")
	}
}

func BenchmarkTriviaRequest(b *testing.B) {
	api := New()

	if api == nil {
		b.Errorf("error api object is nil")
	} else {
		// benchmark
		for idx := 0; idx < b.N; idx++ {
			_, _, requestErr := api.triviaRequest()
			if requestErr != nil {
				log.Print("Error processing trivia request...", requestErr)
			}
		}
	}
}

func CreateNewAPI(t *testing.T) *API {
	envEnvErr := config.SetEnvVars()
	if envEnvErr != nil {
		t.Errorf("error attempting to set env vars, with error: %s", envEnvErr.Error())
	} else {
		api := New()
		if api == nil {
			t.Errorf("api object is nil, cannot continue...")
		} else {
			return api
		}
	}

	return nil
}
