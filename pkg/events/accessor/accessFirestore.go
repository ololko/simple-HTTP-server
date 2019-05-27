package accessor

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

type FirestoreAccess struct {
	Client *firestore.Client
}

func (d *FirestoreAccess) ReadEvent(request models.RequestT) (models.AnswerT, error) {
	var count int64
	iter := d.Client.Collection("users").Where("Type", "==", request.Type).Where("Timestamp", ">=", request.From).Where("Timestamp", "<=", request.To).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			return models.AnswerT{}, err
		}

		if recData, ok := doc.Data()["Count"].(int64); ok {
			count += recData
		} else {
			return models.AnswerT{}, err
		}
	}

	return models.AnswerT{count, request.Type}, nil
}

func (d *FirestoreAccess) WriteEvent(insert models.EventT) ([]byte, error) {
	DocRef, _, err := d.Client.Collection("users").Add(context.Background(), insert)
	if err != nil {
		return []byte{}, err
	}
	return []byte(DocRef.ID), nil
}
