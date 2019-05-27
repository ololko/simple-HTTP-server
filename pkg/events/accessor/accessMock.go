package accessor

import (
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
)

type MockAccess struct {
	Events map[string][]models.EventT
}

func (d *MockAccess) ReadEvent(request models.RequestT) (models.AnswerT, error) {
	var count int64

	for _,event := range d.Events[request.Type]{
		if event.Timestamp >= request.From && event.Timestamp <= request.To {
			count = count + event.Count
		}
	}

	var retVal = models.AnswerT{count,request.Type}
	return retVal, nil
}

func (d *MockAccess) WriteEvent(newEvent models.EventT) ([]byte, error) {
	d.Events[newEvent.Type] = append(d.Events[newEvent.Type], newEvent)
	return []byte(newEvent.Type),nil
}