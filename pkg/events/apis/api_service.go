package apis

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
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

func (s *Service) ReadEvent(ctx context.Context, request *models.Request) (*models.Answer, error) {
	data := make(chan int32, 1)
	errChan := make(chan error, 1)

	go s.DataAccessor.ReadEvent(*request, data, errChan)
	if <-errChan != nil {
		return nil, status.Errorf(http.StatusNotFound, "Not found")
	}

	log.WithFields(log.Fields{
		"method": "GET",
		//"url":    c.Request().URL.String(),		//ako sem kurna hodim url?
	}).Info("Sending positive answer")
	return &models.Answer{
		Type:  request.Type,
		Count: <-data,
	}, nil
}

//Vyries nacitanie z JSON
func (s *Service) CreateEvent(ctx context.Context, insert *models.Event) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

/*
func (s *Service) HandlePost(ctx context.Context) (*empty.Empty, error) {
	//var newEvent models.EventT
	//err := c.Bind(&newEvent)
	//if err != nil {
	//	log.WithFields(log.Fields{
	//		"method": "POST",
	//		"body":   c.Request().Body,
	//	}).Error("Body has bad structure")
	//	return c.NoContent(http.StatusBadRequest)
	//}

	chanErr := make(chan error)

	newEvent := models.databaseElement{
		Type: //ako to brat z JSON
	}

	go s.DataAccessor.WriteEvent(newEvent, chanErr)
	if err = <-chanErr; err != nil {
		return &empty.Empty{}, status.Errorf(http.StatusNotFound, "Error while writing to database!")		//ako toto returnut?
	}

	log.WithFields(log.Fields{
		"method": "POST",
		"body":   //c.Request().Body,
	}).Info("Sending positive answer")

	//set returning code to 201
	//return google.protobuff.empty //dako takto
}

*/
