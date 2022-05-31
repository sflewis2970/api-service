package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sflewis2970/trivia-service/common"
)

const (
	RapidAPIHostKey string = "X-RapidAPI-Host"
	RapidAPIKey     string = "X-RapidAPI-Key"
	RapidAPIValue   string = "1f8720c0c7msh43fe783209a6813p1833b2jsnc2300c30b9a9"

	TriviaURL          string = "https://trivia-by-api-ninjas.p.rapidapi.com/v1/trivia"
	TriviaAPIHostValue string = "trivia-by-api-ninjas.p.rapidapi.com"

	TriviaCategoryCount  int = 14
	EmptyRecordCount     int = 0
	TriviaMaxRecordCount int = 5
	APIMaxRecordCount    int = 30
)

var CategoryList = [TriviaCategoryCount]string{"artliterature", "language", "sciencenature", "general", "fooddrink", "peopleplaces",
	"geography", "historyholidays", "entertainment", "toysgames", "music", "mathematics", "religionmythology", "sportsleisure"}

type TriviaResponse struct {
	Category string `json:"category"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func isItemInCategoryList(item string) bool {
	for _, category := range CategoryList {
		if item == category {
			return true
		}
	}

	return false
}

func TriviaRequest(category string, limit int) (error, []TriviaResponse, string) {
	// Build URL string
	url := TriviaURL

	// Add optional parametes string
	// Get category string
	categoryLength := len(category)
	if categoryLength > 0 {
		url = url + "?category=" + category
	}

	// Add limit string to the end of the url
	if limit == 0 {
		limit = TriviaMaxRecordCount
	}

	if categoryLength > 0 {
		url = url + "&limit=" + fmt.Sprint(limit)
	} else {
		url = url + "?limit=" + fmt.Sprint(limit)
	}

	// Create new http request
	request, requestErr := http.NewRequest("GET", url, nil)
	if requestErr != nil {
		return requestErr, nil, ""
	}

	// Setup request headers
	request.Header.Add(RapidAPIHostKey, TriviaAPIHostValue)
	request.Header.Add(RapidAPIKey, RapidAPIValue)

	// Get response from http request
	response, responseErr := http.DefaultClient.Do(request)
	if responseErr != nil {
		return requestErr, nil, ""
	}
	defer response.Body.Close()

	// Get timestamp right after receiving a valid request
	timestamp := common.GetFormattedTime(time.Now(), "Mon Jan 2 15:04:05 2006")

	// Parse request body
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return readErr, nil, ""
	}

	// Parse response into JSON format
	responses := make([]TriviaResponse, 0)
	unmarshalErr := json.Unmarshal(body, &responses)
	if unmarshalErr != nil {
		return unmarshalErr, nil, ""
	}

	// Return a valid response (in JSON format) as well as a timestamp
	return nil, responses, timestamp
}
