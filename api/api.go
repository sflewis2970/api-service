package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sflewis2970/trivia-service/common"
	"github.com/sflewis2970/trivia-service/datastores"
	"github.com/sflewis2970/trivia-service/messages"
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

type API struct {
}

// exported type method
func (a *API) GetQuestion() (messages.QuestionResponse, datastores.DataStoreTable) {
	// Initialize data store when needed
	categoryStr := ""
	limit := 0
	requestComplete := false
	var apiResponseErr error
	var apiResponses []TriviaResponse
	apiResponsesSize := 0
	timestamp := ""

	// Check for duplicates for marking the request as complete
	for !requestComplete {
		// Send request to API
		apiResponses, timestamp, apiResponseErr = a.triviaRequest(categoryStr, limit)

		// Get API Response size
		apiResponsesSize = len(apiResponses)

		if apiResponsesSize > 0 {
			// When results are returned, make sure there are no duplicate answers
			if !a.containsDuplicates(apiResponses) {
				log.Print("No duplicates found...")
				requestComplete = true
			} else {
				log.Print("Found duplicates...")
			}
		} else {
			// An error occurred or no results found
			requestComplete = true
		}
	}

	// Question (Request) Response message
	var qResponse messages.QuestionResponse

	// Build API Response
	qResponse.Timestamp = timestamp

	var dsTable datastores.DataStoreTable
	if apiResponseErr != nil {
		// If an error occurs let the client know
		qResponse.Error = apiResponseErr.Error()
	} else if apiResponsesSize == 0 {
		// If no error occurs but no response is returned
		// let the client know, the most likely case where this happens
		// is when the client supplies a category that does not exist
		qResponse.Warning = "No question was returned, invalid category selected"
	} else {
		// Since the client is no longer allowed to supply a limit
		// there should be five items returned from the API
		// After getting a valid response from the API, generate a question ID
		qResponse.QuestionID = uuid.New().String()
		qResponse.QuestionID = common.BuildUUID(qResponse.QuestionID, messages.DASH, messages.ONE_SET)
		qResponse.Category = apiResponses[0].Category
		qResponse.Question = apiResponses[0].Question

		// Build choices string
		choiceList := []string{}
		for idx := 0; idx < apiResponsesSize; idx++ {
			choiceList = append(choiceList, apiResponses[idx].Answer)
		}

		// Shuttle list
		choiceList = common.ShuffleList(choiceList)

		// Add a message filler to the beginning of the list
		qResponse.Choices = append(qResponse.Choices, messages.MAKE_SELECTION_MSG)
		qResponse.Choices = append(qResponse.Choices, choiceList...)

		// Create data store table struct
		dsTable.Question = qResponse.Question
		dsTable.Category = qResponse.Category
		dsTable.Answer = apiResponses[0].Answer
	}

	return qResponse, dsTable
}

// unexported type method
func (a *API) triviaRequest(category string, limit int) ([]TriviaResponse, string, error) {
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

// containsDuplicates checks the slice for any duplicate items
func (a *API) containsDuplicates(items []TriviaResponse) bool {
	// Initialize the map for usage
	itemsMap := make(map[string]int)

	// Since maps uses unique keys, use the string value of answer to be the key
	for idx, item := range items {
		itemsMap[item.Answer] = idx + 1
	}

	// If the size of the map is the same size of the slice, then there are no duplicates
	if len(itemsMap) != len(items) {
		return true
	}

	// Otherwise return false
	return false
}

func New() *API {
	log.Print("Creating API object...")
	api := new(API)

	return api
}

// unexported functions
func isItemInCategoryList(item string) bool {
	for _, category := range CategoryList {
		if item == category {
			return true
		}
	}

	return false
}
