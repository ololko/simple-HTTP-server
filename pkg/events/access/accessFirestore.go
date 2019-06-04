package access

import (
	"cloud.google.com/go/firestore"
	"errors"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

type FirestoreAccess struct {
	Client *firestore.Client
}

func (d *FirestoreAccess) ReadEvent(request models.RequestT, answer chan<- models.AnswerT, chanErr chan<- error) {
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
				answer <- models.AnswerT{}
				chanErr <- errors.New("Element is not in database")
				return
			}
		}
		if err != nil {
			log.WithFields(log.Fields{
				"type": request.Type,
			}).Error("Unexpected error with firestore while reading")
			answer <- models.AnswerT{}
			chanErr <- err
			return
		}

		if recData, ok := doc.Data()["Count"].(int64); ok {
			count += recData
		} else {
			log.WithFields(log.Fields{
				"type": request.Type,
			}).Error("Unexpected error. Database is incosistent in COUNT field!")
			answer <- models.AnswerT{}
			chanErr <- err
			return
		}
		elementExists = true
	}

	answer <- models.AnswerT{Count:count, Type:request.Type}
	chanErr <- nil
	return
}

func (d *FirestoreAccess) WriteEvent(insert models.EventT, answer chan<- string, errChan chan<- error) {
	_, _, err := d.Client.Collection("users").Add(context.Background(), insert)
	if err != nil {
		log.WithFields(log.Fields{
			"type": insert.Type,
		}).Error("Unexpected error while creating new event in database")
		errChan <- err
		answer <- ""
		return
	}
	errChan <- nil
	answer <- insert.Type
	return
}
