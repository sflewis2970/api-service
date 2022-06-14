package api

import (
	"testing"
)

func TestTriviaRequest(t *testing.T) {
	// Test cases
	testCases := []struct {
		categoryVal string
	}{
		{categoryVal: ""},
		{categoryVal: "mathematics"},
		{categoryVal: "sfl"},
	}

	for _, tt := range testCases {
		limitVal := 0
		gotVals, gotTimestamp, gotError := TriviaRequest(tt.categoryVal, limitVal)

		gotValsSize := len(gotVals)
		categoryValSize := len(tt.categoryVal)

		// Category and limit have empty values
		if categoryValSize == 0 {
			if gotError != nil {
				t.Errorf("TriviaRequest(%v, %v): error not expected, got %v", tt.categoryVal, limitVal, gotError)
			}

			if gotValsSize != TriviaMaxRecordCount {
				t.Errorf("TriviaRequest(%v, %v): did not return the correct number of records, got %d - expected: %d", tt.categoryVal, limitVal, gotValsSize, TriviaMaxRecordCount)
			}

			if len(gotTimestamp) == 0 {
				t.Errorf("TriviaRequest(%v, %v): did not return a valid time stamp, got %s", tt.categoryVal, limitVal, gotTimestamp)
			}
		}

		// Category has a non-empty value while limit has an empty
		if categoryValSize > 0 && isItemInCategoryList(tt.categoryVal) {
			if gotError != nil {
				t.Errorf("TriviaRequest(%v, %v): error not expected, got %v", tt.categoryVal, limitVal, gotError)
			}

			if gotValsSize != TriviaMaxRecordCount {
				t.Errorf("TriviaRequest(%v, %v): did not return the correct number of records, got %d - expected: %d", tt.categoryVal, limitVal, gotValsSize, TriviaMaxRecordCount)
			}

			if len(gotTimestamp) == 0 {
				t.Errorf("TriviaRequest(%v, %v): did not return a valid time stamp, got %s", tt.categoryVal, limitVal, gotTimestamp)
			}
		}

		// Category has a non-empty value while and the category is not list in the category list
		if categoryValSize > 0 && !isItemInCategoryList(tt.categoryVal) {
			if gotError != nil {
				t.Errorf("TriviaRequest(%v, %v): error not expected, got %v", tt.categoryVal, limitVal, gotError)
			}

			if gotValsSize != EmptyRecordCount {
				t.Errorf("TriviaRequest(%v, %v): did not return the correct number of records, got %d - expected: %d", tt.categoryVal, limitVal, gotValsSize, EmptyRecordCount)
			}

			if len(gotTimestamp) == 0 {
				t.Errorf("TriviaRequest(%v, %v): did not return a valid time stamp, got %s", tt.categoryVal, limitVal, gotTimestamp)
			}
		}
	}
}

func BenchmarkTriviaRequest(b *testing.B) {
	// benchmark
	category := ""
	limit := 10

	for idx := 0; idx < b.N; idx++ {
		TriviaRequest(category, limit)
	}
}
