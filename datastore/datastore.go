package datastore

import (
	"log"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/sflewis2970/trivia-service/common"
)

const (
	DEFAULT_EXPIRATION int = 1  // expirastion time in minutes
	CLEANUP_INTERVAL   int = 10 // expirastion time in minutes
)

type QuestionDS struct {
	memCache *cache.Cache
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

func (qds *QuestionDS) AddQuestionAndAnswer(questionID string, dst DataStoreTable) {
	log.Print("Adding question to map")
	log.Print("Question ID: ", questionID)
	log.Print("Question: ", dst.Question)
	log.Print("Category: ", dst.Category)
	log.Print("Answer: ", dst.Answer)

	qds.memCache.Set(questionID, dst, cache.DefaultExpiration)
}

func (qds *QuestionDS) CheckAnswer(questionID string, response string) (string, *QuestionAndAnswer) {
	newQA := new(QuestionAndAnswer)
	log.Printf("Looking up question ID: %v", questionID)
	item, itemFound := qds.memCache.Get(questionID)
	timestamp := ""

	if itemFound {
		dst, ok := item.(DataStoreTable)
		if !ok {
			log.Print("Error converting interface object: ", item)
		} else {
			// Get timestamp right after checking to see if item is in map
			timestamp = common.GetFormattedTime(time.Now(), "Mon Jan 2 15:04:05 2006")

			log.Print("Found question in map: ", questionID)

			// Update fields for new Question and Answer
			newQA.Question = dst.Question
			newQA.Category = dst.Category
			newQA.Response = response
			newQA.Answer = dst.Answer

			// Delete the record from map
			qds.memCache.Delete(questionID)
			log.Print("record deleted!")

			// Check to see the client has provided the correct answer
			if strings.TrimSpace(dst.Answer) == strings.TrimSpace(response) {
				newQA.Correct = true
				newQA.Message = "Congrats! That is correct!"
				return timestamp, newQA
			} else {
				newQA.Correct = false
				newQA.Message = "Nice try! That is NOT correct"
			}
		}
	} else {
		newQA.Correct = false
		newQA.Message = "Question not found"
	}

	return timestamp, newQA
}

func InitializeDataStore() *QuestionDS {
	ds := new(QuestionDS)

	ds.memCache = cache.New(time.Duration(DEFAULT_EXPIRATION)*time.Minute, time.Duration(CLEANUP_INTERVAL)*time.Minute)

	return ds
}
