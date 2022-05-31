package api

import (
	"testing"
)

func TestTriviaRequest(t *testing.T) {
	// Test cases
	testCases := []struct {
		categoryVal string
		limitVal    int
	}{
		{categoryVal: ""},
		{categoryVal: "", limitVal: 1},
		{categoryVal: "", limitVal: 5},
		{categoryVal: "", limitVal: 10},
		{categoryVal: "", limitVal: 20},
		{categoryVal: "", limitVal: 40},
		{categoryVal: "general"},
		{categoryVal: "general", limitVal: 1},
		{categoryVal: "general", limitVal: 5},
		{categoryVal: "general", limitVal: 10},
		{categoryVal: "general", limitVal: 20},
		{categoryVal: "general", limitVal: 40},
		{categoryVal: "lang"},
		{categoryVal: "lang", limitVal: 1},
		{categoryVal: "lang", limitVal: 5},
		{categoryVal: "lang", limitVal: 10},
		{categoryVal: "lang", limitVal: 20},
		{categoryVal: "lang", limitVal: 40},
	}

	for _, tt := range testCases {
		gotError, gotVals, _ := TriviaRequest(tt.categoryVal, tt.limitVal)

		gotValsSize := len(gotVals)
		categoryValSize := len(tt.categoryVal)

		// Category and limit have empty values
		if categoryValSize == 0 && tt.limitVal == 0 {
			if gotError != nil {
				t.Errorf("TriviaRequest(%v, %v): error not expected, got %v", tt.categoryVal, tt.limitVal, gotError)
			}

			if gotValsSize != TriviaMaxRecordCount {
				t.Errorf("TriviaRequest(%v, %v): did not return the correct number of records, got %d - expected: %d", tt.categoryVal, tt.limitVal, gotValsSize, TriviaMaxRecordCount)
			}
		}

		// Category has a empty value while limit has a value greater than 0
		if categoryValSize == 0 && tt.limitVal > 0 {
			if gotError != nil {
				t.Errorf("TriviaRequest(%v, %v): error not expected, got %v", tt.categoryVal, tt.limitVal, gotError)
			}

			if tt.limitVal <= APIMaxRecordCount && gotValsSize != tt.limitVal {
				t.Errorf("TriviaRequest(%v, %v): did not return the correct number of records, got %d - expected: %d", tt.categoryVal, tt.limitVal, gotValsSize, tt.limitVal)
			}

			if tt.limitVal > APIMaxRecordCount && gotValsSize != APIMaxRecordCount {
				t.Errorf("TriviaRequest(%v, %v): did not return the correct number of records, got %d - expected: %d", tt.categoryVal, tt.limitVal, gotValsSize, APIMaxRecordCount)
			}
		}

		// Category has a non-empty value while limit has an empty
		if categoryValSize > 0 && isItemInCategoryList(tt.categoryVal) && tt.limitVal == 0 {
			if gotError != nil {
				t.Errorf("TriviaRequest(%v, %v): error not expected, got %v", tt.categoryVal, tt.limitVal, gotError)
			}

			if gotValsSize != TriviaMaxRecordCount {
				t.Errorf("TriviaRequest(%v, %v): did not return the correct number of records, got %d - expected: %d", tt.categoryVal, tt.limitVal, gotValsSize, TriviaMaxRecordCount)
			}
		}

		// Category has a non-empty value while limit has a value greater than zero
		if categoryValSize > 0 && isItemInCategoryList(tt.categoryVal) && tt.limitVal > 0 {
			if gotError != nil {
				t.Errorf("TriviaRequest(%v, %v): error not expected, got %v", tt.categoryVal, tt.limitVal, gotError)
			}

			if tt.limitVal <= APIMaxRecordCount && gotValsSize != tt.limitVal {
				t.Errorf("TriviaRequest(%v, %v): did not return the correct number of records, got %d - expected: %d", tt.categoryVal, tt.limitVal, gotValsSize, tt.limitVal)
			}

			if tt.limitVal > APIMaxRecordCount && gotValsSize != APIMaxRecordCount {
				t.Errorf("TriviaRequest(%v, %v): did not return the correct number of records, got %d - expected: %d", tt.categoryVal, tt.limitVal, gotValsSize, APIMaxRecordCount)
			}
		}

		// Category has a non-empty value while and the category is not list in the category list
		if categoryValSize > 0 && !isItemInCategoryList(tt.categoryVal) {
			if gotError != nil {
				t.Errorf("TriviaRequest(%v, %v): error not expected, got %v", tt.categoryVal, tt.limitVal, gotError)
			}

			if gotValsSize != EmptyRecordCount {
				t.Errorf("TriviaRequest(%v, %v): did not return the correct number of records, got %d - expected: %d", tt.categoryVal, tt.limitVal, gotValsSize, EmptyRecordCount)
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
