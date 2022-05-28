package api

import "testing"

func TestTriviaRequest(t *testing.T) {
	// Test cases
	testCases := []struct {
		categoryVal      string
		limitVal         string
		expectedErrorVal error
	}{
		{categoryVal: "", limitVal: "", expectedErrorVal: nil},
		{categoryVal: "", limitVal: "1", expectedErrorVal: nil},
		{categoryVal: "", limitVal: "5", expectedErrorVal: nil},
		{categoryVal: "", limitVal: "10", expectedErrorVal: nil},
	}

	for _, tt := range testCases {
		gotError, gotVals, gotTimestamp := TriviaRequest(tt.categoryVal, tt.limitVal)

		if gotError != tt.expectedErrorVal {
			t.Errorf("TriviaRequest(%v, %v): expected %v, got %v", tt.categoryVal, tt.limitVal, tt.expectedErrorVal, gotError)
		}

		if len(gotVals) <= 0 {
			t.Errorf("TriviaRequest(%v, %v): expected results from the API call.", tt.categoryVal, tt.limitVal)
		}

		if len(gotTimestamp) <= 0 {
			t.Errorf("TriviaRequest(%v, %v): expected a valid generated timestamp.", tt.categoryVal, tt.limitVal)
		}
	}
}

func BenchmarkTriviaRequest(b *testing.B) {
	// benchmark
	category := ""
	limit := "10"

	for idx := 0; idx < b.N; idx++ {
		TriviaRequest(category, limit)
	}
}
