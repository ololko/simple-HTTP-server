package access

import (
	"cloud.google.com/go/firestore"
	"errors"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"math"
)

type FirestoreAccess struct {
	Client *firestore.Client
}

func (d *FirestoreAccess) ReadEvent(request models.Request, answer chan<- int32, chanErr chan<- error) {
	var count int64
	elementExists := false
	iter := d.Client.Collection("users").Where("Type", "==", request.Type).Where("Timestamp", ">=", request.From).Where("Timestamp", "<=", request.To).Documents(context.Background())
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			if elementExists {
				break
			} else {
				log.WithFields(log.Fields{
					"type": request.Type,
					"from": request.From,
					"to":	request.To,
				}).Info("Requested event does not exist in range!")
				answer <- 0
				chanErr <- errors.New("Element is not in database")
				return
			}
		}
		if err != nil {
			log.WithFields(log.Fields{
				"type": request.Type,
			}).Error("Unexpected error with firestore while reading")
			answer <- 0
			chanErr <- err
			return
		}

		if recData, ok := doc.Data()["Count"].(int64); ok {
			count += recData
		} else {
			log.WithFields(log.Fields{
				"type": request.Type,
			}).Error("Unexpected error. Database is incosistent in COUNT field!")
			answer <- 0
			chanErr <- err
			return
		}
		elementExists = true
	}

	if count > math.MaxInt32{
		count = math.MaxInt32
	}
	answer <- int32(count)
	chanErr <- nil
	return
}

func (d *FirestoreAccess) WriteEvent(insert models.Event, errChan chan<- error) {
	panic("implement me")
/*		_, _, err := d.Client.Collection("users").Add(context.Background(), insert)
	if err != nil {
		log.WithFields(log.Fields{
			"type": insert.Type,
		}).Error("Unexpected error while creating new event in database")
		errChan <- err
		return
	}

	errChan <- nil
	return*/
}
