package accessor

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"github.com/ololko/simple-HTTP-server/pkg/events/custom_errors"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

type FirestoreAccess struct {
	Client *firestore.Client
}

func (d *FirestoreAccess) ReadEvent(request models.RequestT) (models.AnswerT, custom_errors.ElementDoesNotExistError) {
	var count int64
	iter := d.Client.Collection("users").Where("Type", "==", request.Type).Where("Timestamp", ">=", request.From).Where("Timestamp", "<=", request.To).Documents(context.Background())
	if iter == nil {
		fmt.Println("Iter is nil")
		return models.AnswerT{Type:request.Type, Count:0},custom_errors.ElementDoesNotExistError{"Searched element is not in database", true, true}
	}
	fmt.Println("iter is not nil")
	fmt.Println(iter)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			return models.AnswerT{}, custom_errors.ElementDoesNotExistError{"Searched element is not in database", false, true}
		}

		if recData, ok := doc.Data()["Count"].(int64); ok {
			count += recData
		} else {
			return models.AnswerT{}, custom_errors.ElementDoesNotExistError{"Searched element is not in database", false, true}
		}
	}

	return models.AnswerT{count, request.Type}, custom_errors.ElementDoesNotExistError{"AKO TOMUTO DAT NIL?", true, false}
}

func (d *FirestoreAccess) WriteEvent(insert models.EventT) ([]byte, error) {
	DocRef, _, err := d.Client.Collection("users").Add(context.Background(), insert)
	if err != nil {
		return []byte{}, err
	}
	return []byte(DocRef.ID), nil
}
