package trivia

import (
	"log"
	"testing"
)

func TestTriviaRequest(t *testing.T) {
	// Create api object
	api := New()

	// Test cases
	testCases := []struct {
		category string
	}{
		{category: ""},
		{category: "mathematics"},
		{category: "sfl"},
	}

	for _, tt := range testCases {
		limitVal := 0
		gotVals, gotTimestamp, gotError := api.triviaRequest(tt.category, limitVal)

		gotValsSize := len(gotVals)
		categoryValSize := len(tt.category)

		// Category and limit have empty values
		if categoryValSize == 0 {
			if gotError != nil {
				t.Errorf("TriviaRequest(%v, %v): error not expected, got %v", tt.category, limitVal, gotError)
			}

			if gotValsSize != TriviaMaxRecordCount {
				t.Errorf("TriviaRequest(%v, %v): did not return the correct number of records, got %d - expected: %d", tt.category, limitVal, gotValsSize, TriviaMaxRecordCount)
			}

			if len(gotTimestamp) == 0 {
				t.Errorf("TriviaRequest(%v, %v): did not return a valid time stamp, got %s", tt.category, limitVal, gotTimestamp)
			}
		}

		// Category has a non-empty value
		if categoryValSize > 0 && isItemInCategoryList(tt.category) {
			if gotError != nil {
				t.Errorf("TriviaRequest(%v, %v): error not expected, got %v", tt.category, limitVal, gotError)
			}

			if gotValsSize != TriviaMaxRecordCount {
				t.Errorf("TriviaRequest(%v, %v): did not return the correct number of records, got %d - expected: %d", tt.category, limitVal, gotValsSize, TriviaMaxRecordCount)
			}

			if len(gotTimestamp) == 0 {
				t.Errorf("TriviaRequest(%v, %v): did not return a valid time stamp, got %s", tt.category, limitVal, gotTimestamp)
			}
		}

		// Category has a non-empty value and the category is not list in the category list
		if categoryValSize > 0 && !isItemInCategoryList(tt.category) {
			if gotError != nil {
				t.Errorf("TriviaRequest(%v, %v): error not expected, got %v", tt.category, limitVal, gotError)
			}

			if gotValsSize != EmptyRecordCount {
				t.Errorf("TriviaRequest(%v, %v): did not return the correct number of records, got %d - expected: %d", tt.category, limitVal, gotValsSize, EmptyRecordCount)
			}

			if len(gotTimestamp) == 0 {
				t.Errorf("TriviaRequest(%v, %v): did not return a valid time stamp, got %s", tt.category, limitVal, gotTimestamp)
			}
		}
	}
}

func BenchmarkTriviaRequest(b *testing.B) {
	api := New()

	// benchmark
	category := ""
	limit := 0

	for idx := 0; idx < b.N; idx++ {
		_, _, requestErr := api.triviaRequest(category, limit)
		if requestErr != nil {
			log.Print("Error processing trivia request...", requestErr)
		}
	}
}
