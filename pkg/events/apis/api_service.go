package apis

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ololko/simple-HTTP-server/pkg/events/access"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	DataAccessor access.DataAccessor
}

func NewService(dataAccessor access.DataAccessor) *Service {
	return &Service{DataAccessor: dataAccessor}
}

func (s *Service) HandleGet(c echo.Context) error {
	request, err := fillRequestStruck(c.Request().URL)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	data := make(chan models.AnswerT, 1)
	errChan := make(chan error)

	go s.DataAccessor.ReadEvent(request, data, errChan)
	if <-errChan != nil {
		return c.NoContent(http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"method": "GET",
		"url":    c.Request().URL.String(),
	}).Info("Sending positive answer")
	return c.JSON(http.StatusOK, <-data)
}

func (s *Service) HandlePost(c echo.Context) error {
	var newEvent models.EventT
	err := json.NewDecoder(c.Request().Body).Decode(&newEvent)
	if err != nil {
		log.WithFields(log.Fields{
			"method": "POST",
			"body":   c.Request().Body,
		}).Error("Body has bad structure")
		return c.NoContent(http.StatusBadRequest)
	}

	chanErr := make(chan error)

	go s.DataAccessor.WriteEvent(newEvent, chanErr)
	if err = <-chanErr; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	log.WithFields(log.Fields{
		"method": "POST",
		"body":   c.Request().Body,
	}).Info("Sending positive answer")
	return c.JSON(http.StatusCreated, newEvent.Type)
}
