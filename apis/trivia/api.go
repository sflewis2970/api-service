package trivia

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sflewis2970/trivia-service/common"
	"github.com/sflewis2970/trivia-service/config"
	"github.com/sflewis2970/trivia-service/messages"
)

//goland:noinspection SpellCheckingInspection,SpellCheckingInspection
const (
	// RapidAPIHostKey  string = "X-RapidAPI-Host"
	// RapidAPIKey      string = "X-RapidAPI-Key"
	// RapidAPIKeyValue string = "b2c5514e45msh28ba59311e819e7p148a88jsn1a173e6b9eab"

	RapidApiURL        string = "https://trivia-by-api-ninjas.p.rapidapi.com/v1/trivia"
	TriviaAPIHostValue string = "trivia-by-api-ninjas.p.rapidapi.com"

	// TriviaCategoryCount  int = 14
	EmptyRecordCount     int = 0
	TriviaMaxRecordCount int = 5
	AnswerCount          int = 4
)

var CategoryList = []string{"artliterature", "language", "sciencenature", "general", "fooddrink", "peopleplaces",
	"geography", "historyholidays", "entertainment", "toysgames", "music", "mathematics", "religionmythology", "sportsleisure"}

type TriviaResponse struct {
	Category string `json:"category"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type API struct {
	cfg     *config.Config
	cfgData *config.CfgData
}

// GetTrivia exported type method
func (a *API) GetTrivia() (messages.Trivia, error) {
	// Initialize data store when needed
	requestComplete := false
	var apiResponseErr error
	var apiResponses []TriviaResponse
	apiResponsesSize := 0
	triviaRequestAttCtr := 0
	timestamp := ""
	nbrRequestAtt := 0

	// Check for duplicates for marking the request as complete
	// Question (Request) Response message
	var trivia messages.Trivia

	for !requestComplete {
		// Send request to API
		log.Print("Sending request to API!")
		apiResponses, timestamp, apiResponseErr = a.triviaRequest()
		nbrRequestAtt++
		triviaRequestAttCtr++
		trivia.NbrRequestAtt = nbrRequestAtt
		apiResponsesSize = len(apiResponses)
		log.Print("number of responses: ", apiResponsesSize)

		// Update API Response timestamp
		trivia.Timestamp = timestamp

		if apiResponseErr == nil {
			// The trivia API used is a basic account, therefore, clients are NOT allowed to supply a limit.
			// basic accounts returns at least 1 array element item returned from the API. After getting a valid response
			// from the API, generate a question ID
			if apiResponsesSize > 0 {
				trivia.QuestionID = uuid.New().String()
				trivia.QuestionID = common.BuildNewUUID(trivia.QuestionID, messages.DASH, messages.FIRST_SET)
				trivia.Category = apiResponses[0].Category
				if len(trivia.Category) == 0 {
					log.Print("no category returned in response")
				}
				trivia.Question = apiResponses[0].Question
				trivia.Answer = strings.ToUpper(apiResponses[0].Answer)

				// Build choices string
				var choiceList []string
				choiceList = append(choiceList, apiResponses[0].Answer)
				additionalAns, multiAnsErr := a.returnMultipleAnswers(AnswerCount)
				if multiAnsErr != nil {
					log.Print("error returned generating additional answers: ", multiAnsErr.Error())
				}
				choiceList = append(choiceList, additionalAns...)

				// Shuttle list
				choiceList = common.ShuffleList(choiceList)

				// Add a message filler to the beginning of the list
				trivia.Choices = append(trivia.Choices, messages.MAKE_SELECTION_MSG)
				trivia.Choices = append(trivia.Choices, choiceList...)

				// The request is complete when there is a valid question
				//
				if (len(trivia.Question) > 0) || (triviaRequestAttCtr > (a.cfgData.TriviaAPI.NbrOfRetries + 1)) {
					requestComplete = true
				}
			}
		} else {
			// If an error occurs let the client know after all the retries hsave been exausted
			if triviaRequestAttCtr > (a.cfgData.TriviaAPI.NbrOfRetries + 1) {
				return trivia, apiResponseErr
			}
		}
	}

	return trivia, nil
}

// unexported type method
// triviaRequest is a function that sends a request to the API to retrieve the trivia question
// trivia response is an array of a structure. The structure contains 3 fields: category (string),
// question (string), answer (string).
func (a *API) triviaRequest() ([]TriviaResponse, string, error) {
	// Build URL string
	url := RapidApiURL

	// Create a http request
	method := "GET"
	var reqBody io.Reader
	log.Print("URL being processed: ", url)
	// Setup request headers
	request, requestErr := http.NewRequest(method, url, reqBody)
	request.Header.Add("x-rapidapi-key", "b2c5514e45msh28ba59311e819e7p148a88jsn1a173e6b9eab")
	request.Header.Add("x-rapidapi-host", "trivia-by-api-ninjas.p.rapidapi.com")
	if requestErr != nil {
		log.Print("A request error has occurred...: ", requestErr)
		return []TriviaResponse{}, "", requestErr
	}

	// Execute request
	// Get response from http request
	log.Print("Performing Default Do command")
	response, responseErr := http.DefaultClient.Do(request)
	statusCodeStr := common.ProcessStatusCode(response.StatusCode)
	if responseErr != nil {
		errMsg := "Doing request command returned a response error, with status code: " + strconv.Itoa(response.StatusCode)
		log.Print(errMsg)
		return []TriviaResponse{}, statusCodeStr, responseErr
	} else {
		log.Print("Response returned with no errors")
	}
	defer response.Body.Close()

	// Get timestamp right after receiving a valid request
	timestamp := common.GetFormattedTime(time.Now(), "Mon Jan 2 15:04:05 2006")

	// Parse request body
	body, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		log.Print("Error reading response...", readErr)
		return []TriviaResponse{}, "", readErr
	}

	// Parse response into JSON format
	var responses []TriviaResponse
	unmarshalErr := json.Unmarshal(body, &responses)
	if unmarshalErr != nil {
		errMsg := "Error unmarshalling response...: " + unmarshalErr.Error()
		log.Print(errMsg)
		return []TriviaResponse{}, "", unmarshalErr
	}

	// Return a valid response (in JSON format) as well as a timestamp
	return responses, timestamp, nil
}

func (a *API) returnMultipleAnswers(answerCnt int) ([]string, error) {
	answers := []string{}

	for idx := 0; idx < answerCnt; idx++ {
		var apiResponseErr error
		var apiResponses []TriviaResponse
		apiResponses, _, apiResponseErr = a.triviaRequest()
		apiResponsesSize := len(apiResponses)
		if apiResponseErr != nil {

		} else {
			if apiResponsesSize > 0 {
				answers = append(answers, apiResponses[0].Answer)
			}

		}

	}

	return answers, nil

}

// Create new API object
func New() *API {
	log.Print("Creating API object...")
	api := new(API)
	api.cfg = config.New()
	var cfgDataErr error
	api.cfgData, cfgDataErr = api.cfg.GetData(config.REFRESH_CONFIG_DATA)
	if cfgDataErr != nil {
		log.Printf("error getting config data: %s\n", cfgDataErr.Error())
	}

	return api
}
