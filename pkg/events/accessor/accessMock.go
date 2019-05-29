package accessor

import (
	"errors"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
)

type MockAccess struct {
	Events map[string][]models.EventT
}

func (d *MockAccess) ReadEvent(request models.RequestT, answer chan<- models.AnswerT, errChan chan<- error) {
	_, exists := d.Events[request.Type]
	if !exists {
		errChan <- errors.New("err: element does not exist")
		answer <- models.AnswerT{}
		return
	}

	var count int64
	inRange := false
	for _,event := range d.Events[request.Type]{
		if event.Timestamp >= request.From && event.Timestamp <= request.To {
			count = count + event.Count
			inRange = true
		}
	}

	if inRange{
		errChan <- nil
		answer <- models.AnswerT{count,request.Type}
		return
	}else{
		errChan <- errors.New("Searched event does not exist in range")
		answer <- models.AnswerT{}
		return
	}
}

func (d *MockAccess) WriteEvent(newEvent models.EventT, answer chan<- string, errChan chan<- error) {
	d.Events[newEvent.Type] = append(d.Events[newEvent.Type], newEvent)

	errChan <- nil
	answer <- newEvent.Type
	return
}