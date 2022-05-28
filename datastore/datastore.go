package datastore

import (
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/sflewis2970/trivia-service/common"
)

type QuestionDS struct {
	mapMutex    sync.Mutex
	QuestionMap map[string]QuestionAndAnswer
}

type QuestionAndAnswer struct {
	Question  string
	Category  string
	AnswerPos int
	Answer    string
	Correct   bool
	Message   string
}

func (qds *QuestionDS) AddQuestionAndAnswer(questionID string, qa QuestionAndAnswer) {
	qds.mapMutex.Lock()
	defer qds.mapMutex.Unlock()

	log.Printf("Adding question to map")
	qds.QuestionMap[questionID] = qa
}

func (qds *QuestionDS) CheckAnswer(questionID string, answer string) (string, *QuestionAndAnswer) {
	qds.mapMutex.Lock()
	defer qds.mapMutex.Unlock()

	newQA := new(QuestionAndAnswer)
	log.Printf("Looking up question ID: %v", questionID)
	qa, itemFound := qds.QuestionMap[questionID]

	// Get timestamp right after checking to see if item is in map
	timestamp := common.GetFormattedTime(time.Now(), "Mon Jan 2 15:04:05 2006")

	if itemFound {
		log.Printf("Found question in map: %v", questionID)

		// Update fields for new Question and Answer
		newQA.Question = qa.Question
		newQA.Category = qa.Category
		newQA.Answer = qa.Answer

		// Convert answer string to answer number value
		nAnswer, convErr := strconv.Atoi(answer)
		if convErr != nil {
			log.Print("Error converting string to number")
		}

		// Delete the record from map
		delete(qds.QuestionMap, questionID)
		log.Print("record deleted!")

		// Check to see the client has provided the correct answer
		if qa.AnswerPos == nAnswer {
			newQA.Correct = true
			newQA.Message = "Congrats! That is correct!"
			return timestamp, newQA
		} else {
			newQA.Correct = false
			newQA.Message = "Nice try! That is NOT correct"
		}
	} else {
		newQA.Correct = false
		newQA.Message = "Question not found"
	}

	return timestamp, newQA
}

func InitializeDataStore() *QuestionDS {
	ds := new(QuestionDS)

	ds.QuestionMap = make(map[string]QuestionAndAnswer)

	return ds
}
