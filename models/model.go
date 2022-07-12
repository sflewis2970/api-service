package models

import (
	"log"

	"github.com/sflewis2970/trivia-service/config"
	"github.com/sflewis2970/trivia-service/datastores"
	"github.com/sflewis2970/trivia-service/messages"
)

type Model struct {
	cfgData *config.ConfigData
	ds      *datastores.DataStore
}

func (m *Model) InsertQuestion(questionID string, dsTable datastores.DataStoreTable) error {
	return m.ds.Insert(questionID, dsTable)
}

func (m *Model) AnswerQuestion(aRequest messages.AnswerRequest) messages.AnswerResponse {
	// Initialize data store when needed
	var aResponse messages.AnswerResponse
	timestamp, newQA, getErr := m.ds.Get(aRequest.QuestionID)
	if getErr != nil {
		errMsg := "Datastore: Get error..."
		log.Print(errMsg, ": ", getErr)
		aResponse.Error = errMsg
	} else {
		// Build Response mesasge
		if len(newQA.Question) > 0 {
			aResponse.Question = newQA.Question
			aResponse.Category = newQA.Category
			aResponse.Answer = newQA.Answer
			aResponse.Response = aRequest.Response
			aResponse.Timestamp = timestamp

			if aRequest.Response == newQA.Answer {
				aResponse.Correct = true
				aResponse.Message = m.cfgData.Messages.CongratsMsg
			} else {
				aResponse.Correct = false
				aResponse.Message = m.cfgData.Messages.TryAgainMsg
			}
		} else {
			aResponse.Message = newQA.Message
			aResponse.Warning = newQA.Warning
			aResponse.Error = newQA.Error
		}
	}

	return aResponse
}

func New() *Model {
	log.Print("Creating model object...")
	model := new(Model)

	// Get config data
	var cfgDataErr error
	model.cfgData, cfgDataErr = config.Get().GetData()
	if cfgDataErr != nil {
		log.Print("Error getting config data: ", cfgDataErr)
		return nil
	}

	// Get new datastore
	model.ds = datastores.New()

	return model
}
