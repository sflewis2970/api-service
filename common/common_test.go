package common

import (
	"fmt"
	"testing"
	"time"
)

func TestGetFormattedTime(t *testing.T) {
	// Get current time
	timeNow := time.Now()
	fmt.Println("current time: ", timeNow)

	// Test cases
	testCases := []struct {
		timeVal       time.Time
		timeFormatVal string
	}{
		{timeVal: timeNow, timeFormatVal: "Mon Jan 2 15:04:05 2006"},
	}

	for _, tt := range testCases {
		gotVal := GetFormattedTime(tt.timeVal, tt.timeFormatVal)

		if len(gotVal) <= 0 {
			t.Errorf("GetFormattedTime(%v, %v): expected results from the API call. Got: %s", tt.timeVal, tt.timeFormatVal, gotVal)
		}
	}
}

func BenchmarkGetFormattedTime(b *testing.B) {
	// benchmark
	timeNow := time.Now()
	timeFormat := "Mon Jan 2 15:04:05 2006"

	for idx := 0; idx < b.N; idx++ {
		GetFormattedTime(timeNow, timeFormat)
	}
}
