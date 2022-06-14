package datastore

import (
	"math"

	"github.com/patrickmn/go-cache"
)

// Datastore component contants
const (
	DEFAULT_EXPIRATION int = 1  // expirastion time in minutes
	CLEANUP_INTERVAL   int = 10 // expirastion time in minutes
)

// Datastore server contants
const (
	// DS_NOT_STARTED -- Datastore server has not been started or initialized
	DS_NOT_STARTED int = iota
	// DS_RUNNING -- Datastore server has been started and is ready for messages
	DS_RUNNING
	// DS_INVALID_SERVER_NAME -- When requesting the Datastore server status the wrong server name was provided
	DS_INVALID_SERVER_NAME
	// DS_REQUEST_ERROR
	DS_REQUEST_ERROR
	// DS_RESPONSE_ERROR
	DS_RESPONSE_ERROR
	// DS_UNAVAILABLE -- When requesting the Datastore server status the server never responded or the connect was refused
	DS_UNAVAILABLE int = math.MaxInt
	// DS_MAX_STATUS_ATTEMPTS -- The maximum number of allowed attempts to get server status
	//                           in a single periodic setting
	// DS_MAX_STATUS_ATTEMPTS -- This is no longer be used be for now we keep the setting in case
	//                           we go back to
	DS_MAX_STATUS_ATTEMPTS int    = 3
	DS_HOST                string = "http://127.0.0.1"
	DS_PORT                string = ":9090"
	DS_STATUS_PATH         string = "/api/v1/ds/status"
	DS_ADD_QUESTION_PATH   string = "/api/v1/ds/addquestion"
	DS_CHECK_ANSWER_PATH   string = "/api/v1/ds/checkanswer"
	DS_SERVERNAME          string = "servername="
)

// Status Response receives a response from a server in the network with the server status information
type StatusCode int
type StatusResponse struct {
	ServerName string     `json:"servername"`
	Timestamp  string     `json:"timestamp"`
	Status     StatusCode `json:"status"`
	Message    string     `json:"message,omitempty"`
	Warning    string     `json:"warning,omitempty"`
	Error      string     `json:"error,omitempty"`
}

// Add Question Request sends a request to the datastore server to add a question to the datastore
type AddQuestionRequest struct {
	QuestionID string `json:"questionid"`
	Question   string `json:"question"`
	Category   string `json:"category"`
	Answer     string `json:"answer"`
}

// Add Question Response sends a request to the datastore server to add a question to the datastore
type AddQuestionResponse struct {
	QuestionID      string `json:"questionid"`
	Question        string `json:"question"`
	Category        string `json:"category"`
	Answer          string `json:"answer"`
	Timestamp       string `json:"timestamp"`
	Action          string `json:"action"`
	RecordsAffected string `json:"recordsaffected"`
	Message         string `json:"message,omitempty"`
	Warning         string `json:"warning,omitempty"`
	Error           string `json:"error,omitempty"`
}

type CheckAnswerRequest struct {
	ServerName string `json:"servername"`
	QuestionID string `json:"questionid"`
	Response   string `json:"response"`
}

type CheckAnswerResponse struct {
	QuestionID string `json:"questionid"`
	Question   string `json:"question"`
	Category   string `json:"category"`
	Answer     string `json:"answer"`
	Response   string `json:"response"`
	Timestamp  string `json:"timestamp"`
	Correct    bool   `json:"correct"`
	Message    string `json:"message,omitempty"`
	Warning    string `json:"warning,omitempty"`
	Error      string `json:"error,omitempty"`
}

type DataStoreTable struct {
	Question string
	Category string
	Answer   string
}

type QuestionAndAnswer struct {
	Question string
	Category string
	Response string
	Answer   string
	Correct  bool
	Message  string
}

type QuestionDS struct {
	useLocalDB bool
	dsStatus   StatusCode
	memCache   *cache.Cache
}
