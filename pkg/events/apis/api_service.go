package apis

import (
	"encoding/json"
	"fmt"
	"github.com/ololko/simple-HTTP-server/pkg/events/accessor"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
	"github.com/ololko/simple-HTTP-server/pkg/events/custom_errors"
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
		data,_ := json.Marshal(models.AnswerT{})
		_,err := w.Write(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(err)
		return
	}

	var customErr custom_errors.ElementDoesNotExistError
	data,customErr := s.DataAccessor.ReadEvent(request)
	if customErr.ReallyError == true{
		if customErr.CanContinue == true {
			w.WriteHeader(http.StatusNotFound)
			answerJSON, _ := json.Marshal(data)
			w.Write(answerJSON)
			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}



	answerJSON, err := json.Marshal(data)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_,err = w.Write(answerJSON)
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

	Id, err := s.DataAccessor.WriteEvent(newEvent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	_,err = w.Write(Id)
	if err != nil {
		return
	}
	return
}


