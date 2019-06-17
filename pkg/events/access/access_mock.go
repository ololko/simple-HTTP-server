package access

import (
	"errors"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
	log "github.com/sirupsen/logrus"
)

type MockAccess struct {
	Events map[string][]models.DatabaseElement
}

func (d *MockAccess) ReadEvent(request models.Request, answer chan<- int32, errChan chan<- error) {
	_, exists := d.Events[request.Type]
	if !exists {
		log.WithFields(log.Fields{
			"type": request.Type,
		}).Info("Requested event does not exist!")
		errChan <- errors.New("Searched event does not exist")
		answer <- 0
		return
	}

	var count int32
	inRange := false
	for _, event := range d.Events[request.Type] {
		if event.Timestamp >= request.From && event.Timestamp <= request.To {
			count = count + event.Count
			inRange = true
		}
	}

	if inRange {
		errChan <- nil
		answer <- count
		return
	} else {
		log.WithFields(log.Fields{
			"type": request.Type,
			"from": request.From,
			"to":   request.To,
		}).Info("Requested event does not exist in range!")
		errChan <- errors.New("Searched event does not exist in range")
		answer <- 0
		return
	}
}

func (d *MockAccess) WriteEvent(inserting models.Event, errChan chan<- error) {
	newEvent := models.DatabaseElement{
		Type: inserting.Type,
		Count:inserting.Count,
		Timestamp:inserting.Timestamp,
	}
	d.Events[newEvent.Type] = append(d.Events[newEvent.Type], newEvent)

	errChan <- nil
	return
}
