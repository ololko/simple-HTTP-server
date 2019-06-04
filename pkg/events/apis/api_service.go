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
	DataAccessor access.DataAccesser
}

func NewService(dataAccessor access.DataAccesser) *Service {
	return &Service{DataAccessor: dataAccessor}
}



func (s *Service) HandleGet(c echo.Context) error {
	request, err := fillRequestStruck(c.Request().URL)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Preco to nevidim v testoch? A vlastne ziadny error")
	}

	data := make(chan models.AnswerT, 1)
	errChan := make(chan error)

	go s.DataAccessor.ReadEvent(request, data, errChan)
	if <-errChan != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Preco to nevidim v testoch? A vlastne ziadny error")
	}

	log.WithFields(log.Fields{
		"method": "get",
		"url"	: c.Request().URL.String(),
	}).Info("Sending positive answer to GET request")
	return c.JSON(http.StatusOK, <-data)
}

func (s *Service) HandlePost(c echo.Context) error {
	var newEvent models.EventT
	err := json.NewDecoder(c.Request().Body).Decode(&newEvent)
	if err != nil {
		log.WithFields(log.Fields{
			"method": "post",
			"body" : c.Request().Body,
		}).Error("Body has bad structure")
		return echo.NewHTTPError(http.StatusBadRequest, "Preco to nevidim v testoch? A vlastne ziadny error")
	}

	returnType := make(chan string)
	chanErr := make(chan error)

	go s.DataAccessor.WriteEvent(newEvent, returnType, chanErr)
	if err = <-chanErr; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Preco to nevidim v testoch? A vlastne ziadny error")
	}

	log.WithFields(log.Fields{
		"method": "post",
		"body"	: c.Request().Body,
	}).Info("Sending positive answer to POST request")
	return c.String(http.StatusCreated, <-returnType)
}
