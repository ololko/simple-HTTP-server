package readers

import(
	"github.com/ololko/simple-http-server/pkg/events/models"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"fmt"
	"golang.org/x/net/context"
)

type FirestoreAccesser struct {
	Client *firestore.Client
}

type MockAccesser struct {
}


func(d *MockAccesser) Read (request models.RequestT) (models.AnswerT,error) {
	return models.AnswerT{},nil
}

func(d *MockAccesser) Write (request models.EventT) ([]byte, error) {
	return []byte{}, nil
}

func(d *FirestoreAccesser) Read (request models.RequestT) (models.AnswerT,error) {
	var count int64
	iter := d.Client.Collection("users").Where("Type", "==", request.Type).Where("Timestamp", ">=", request.From).Where("Timestamp", "<=", request.To).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			return models.AnswerT{},err
		}

		if recData, ok := doc.Data()["Count"].(int64); ok {
			count += recData
		} else {
			return models.AnswerT{},err
		}
	}

	 return models.AnswerT{count, request.Type},nil
}


func(d *FirestoreAccesser) Write (insert models.EventT) ([]byte, error){
	DocRef, _, err := d.Client.Collection("users").Add(context.Background(), insert)
	if err != nil {
		return []byte{}, err
	}
	return []byte(DocRef.ID), nil
}