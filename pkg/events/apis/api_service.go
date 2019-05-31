package apis

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/ololko/simple-HTTP-server/pkg/events/access"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
	"net/http"
)

type Service struct {
	DataAccessor access.DataAccesser
}

func marshalAndSend(w http.ResponseWriter, data models.AnswerT, statusCode int) error {
	answerJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Content-type", "application/json")
	_, err = w.Write(answerJSON)
	if err != nil {
		return err
	}
	return nil
}

func NewService(dataAccessor access.DataAccesser) *Service {
	return &Service{DataAccessor: dataAccessor}
}

func (s *Service) HandleGet(c echo.Context) error {
	request, err := fillRequestStruck(c.Request().URL)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	data := make(chan models.AnswerT, 1)
	errChan := make(chan error)

	go s.DataAccessor.ReadEvent(request, data, errChan)
	if <-errChan != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, <-data)
}

func (s *Service) HandlePost(c echo.Context) error {

	var newEvent models.EventT
	err := json.NewDecoder(c.Request().Body).Decode(&newEvent)
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest)
	}

	returnType := make(chan string)
	chanErr := make(chan error)

	go s.DataAccessor.WriteEvent(newEvent, returnType, chanErr)
	if err = <-chanErr; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, <-returnType)
}
