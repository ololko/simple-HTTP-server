//Package which handles GET request.
//Checks conditions and answers to user
package apis

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/ololko/simple-http-server/pkg/events/models"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

func HandleGet(w http.ResponseWriter, r *http.Request, client *firestore.Client) {

	var request models.RequestT
	request, err := fillRequestStruct(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	var count int64
	iter := client.Collection("users").Where("Type", "==", request.Type).Where("Timestamp", ">=", request.From).Where("Timestamp", "<=", request.To).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if recData, ok := doc.Data()["Count"].(int64); ok {
			count += recData
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	answ := models.AnswerT{count, request.Type}
	answJson, err := json.Marshal(answ)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(answJson)
}
