package apis

import (
	"encoding/json"
	"fmt"
	"github.com/ololko/simple-HTTP-server/pkg/events/accessor"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
	"net/http"
)

type Service struct {
	DataAccessor accessor.DataAccesser
}

func NewService(dataAccessor accessor.DataAccesser) *Service {
	return &Service{DataAccessor: dataAccessor}
}

func (s *Service) HandleGet(w http.ResponseWriter, r *http.Request) {
	request, err := fillRequestStruck(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	data := make(chan models.AnswerT)
	errChan := make (chan error)

	go s.DataAccessor.ReadEvent(request, data, errChan)
	if <-errChan != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}


	answerJSON, err := json.Marshal(<-data)
	if err != nil {
		fmt.Println("error creating JSON")
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(answerJSON)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *Service) HandlePost(w http.ResponseWriter, r *http.Request) {

	var newEvent models.EventT
	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	returnType := make(chan string)
	chanErr := make(chan error)

	go s.DataAccessor.WriteEvent(newEvent, returnType, chanErr)
	if <-chanErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(<-returnType))
	if err != nil {
		return
	}
	return
}
