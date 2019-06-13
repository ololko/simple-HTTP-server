package access

import (
	"errors"

	"github.com/ololko/simple-HTTP-server/pkg/events/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

type PostgreAccess struct {
	Client *gorm.DB
}


func (d *PostgreAccess) ReadEvent(request models.Request, answer chan<- int32, chanErr chan<- error) {
	var events []models.DatabaseElement
	err := d.Client.Where("type=? AND timestamp>=? AND timestamp<=?", request.Type, request.From, request.To).Find(&events).Error
	if err != nil {
		chanErr <- err
		log.WithFields(log.Fields{
			"type": request.Type,
			"from": request.From,
			"to":   request.To,
		}).Info("Error while reading from database!")
		return
	}

	if len(events) == 0 {
		chanErr <- errors.New("Element is not in database")
		log.WithFields(log.Fields{
			"type": request.Type,
			"from": request.From,
			"to":   request.To,
		}).Info("Requested event does not exist in range!")
		return
	}

	var count int32
	for _, event := range events {
		count += event.Count
	}
	chanErr <- nil
	answer <- count
}

//TIEZ UPRAV

func (d *PostgreAccess) WriteEvent(insert models.Event, chanErr chan<- error) {
	panic("implement me")
	/*d.Client.NewRecord(insert)
	d.Client.Create(&insert)
	chanErr <- nil
	return*/
}
