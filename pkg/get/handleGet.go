//Package which handles GET request.
//Checks conditions and answers to user
package get

import (
	"encoding/json"
	firebase "firebase.google.com/go"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"net/http"
	"strings"
)

func HandleGet(w http.ResponseWriter, r *http.Request, app *firebase.App) {

	client, err := app.Firestore(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	var requestLine = strings.Split(r.URL.RawQuery, "&")

	request, err := parseRequest(r.URL)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	var count int64
	iter := client.Collection("users").Where("Type", "==", request.searchedEvent).Where("Timestamp", ">=", request.from).Where("Timestamp", "<=", request.to).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(502)
			return
		}

		if recData, ok := doc.Data()["Count"].(int64); ok {
			count += recData
		} else {
			w.WriteHeader(500)
			return
		}
	}

	answ := answerT{count, request.searchedEvent}
	var answJson, error = json.Marshal(answ)
	if error != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)
	w.Write(answJson)
}
