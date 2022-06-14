package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func TriviaRequest(category string, limit int) ([]TriviaResponse, string, error) {
	// Build URL string
	url := TriviaURL

	// Add optional parametes string
	// Get category string
	categoryLength := len(category)
	if categoryLength > 0 {
		url = url + "?category=" + category
	}

	// Set limit default value
	if limit == 0 {
		limit = TriviaMaxRecordCount
	}

	// Add limit string to the end of the url
	if categoryLength > 0 {
		url = url + "&limit=" + fmt.Sprint(limit)
	} else {
		url = url + "?limit=" + fmt.Sprint(limit)
	}

	headers := []common.HTTPHeader{
		{Key: RapidAPIHostKey, Value: TriviaAPIHostValue},
		{Key: RapidAPIKey, Value: RapidAPIValue},
	}

	// Create a http request
	method := "GET"
	request, requestErr := common.CreateRequest(method, url, headers, nil)
	if requestErr != nil {
		log.Print("Error creating request...")
		return nil, "", requestErr
	}

	// Execute request
	response, responseErr := common.ExecuteRequest(request)
	if responseErr != nil {
		log.Print("Error executing request...")
		return nil, "", responseErr
	}
	defer response.Body.Close()

	// Get timestamp right after receiving a valid request
	timestamp := common.GetFormattedTime(time.Now(), "Mon Jan 2 15:04:05 2006")

	// Parse request body
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Print("Error reading response...")
		return nil, "", readErr
	}

	// Parse response into JSON format
	responses := make([]TriviaResponse, 0)
	unmarshalErr := json.Unmarshal(body, &responses)
	if unmarshalErr != nil {
		log.Print("Error unmarshaling response...")
		return nil, "", unmarshalErr
	}

	// Return a valid response (in JSON format) as well as a timestamp
	return responses, timestamp, nil
}
