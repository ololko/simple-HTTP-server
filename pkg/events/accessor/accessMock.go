package accessor

import (
	"github.com/ololko/simple-HTTP-server/pkg/events/custom_errors"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
)

type MockAccess struct {
	Events map[string][]models.EventT
}

func (d *MockAccess) ReadEvent(request models.RequestT) (models.AnswerT, custom_errors.ElementDoesNotExistError) {
	var count int64

	_, exists := d.Events[request.Type]
	if !exists {
		return models.AnswerT{count,request.Type}, custom_errors.ElementDoesNotExistError{"Element does not exist", true, true}
	}


	for _,event := range d.Events[request.Type]{
		if event.Timestamp >= request.From && event.Timestamp <= request.To {
			count = count + event.Count
		}
	}

	return models.AnswerT{count,request.Type}, custom_errors.ElementDoesNotExistError{"AKO SEM DAT NIL?", true, false}
}

func (d *MockAccess) WriteEvent(newEvent models.EventT) ([]byte, error) {
	d.Events[newEvent.Type] = append(d.Events[newEvent.Type], newEvent)
	return []byte(newEvent.Type),nil
}