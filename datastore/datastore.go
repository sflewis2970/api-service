package datastore

import "sync"

type QuestionDS struct {
	mapMutex    sync.Mutex
	QuestionMap map[string]QuestionAndAnswer
}

type QuestionAndAnswer struct {
	Question string
	Answer   string
}

func (qds *QuestionDS) AddQuestionAndAnswer(questionID string, question string, answer string) {
	qds.mapMutex.Lock()
	defer qds.mapMutex.Unlock()

	qa := QuestionAndAnswer{
		Question: question,
		Answer:   answer,
	}

	qds.QuestionMap[questionID] = qa
}

func (qds *QuestionDS) CheckAnswer(questionID string, answer string) bool {
	qds.mapMutex.Lock()
	defer qds.mapMutex.Unlock()

	qa, itemFound := qds.QuestionMap[questionID]

	if itemFound {
		delete(qds.QuestionMap, questionID)

		if qa.Answer == answer {
			return true
		}
	}

	return false
}

func InitializeDataStore() *QuestionDS {
	ds := new(QuestionDS)

	ds.QuestionMap = make(map[string]QuestionAndAnswer)

	return ds
}
