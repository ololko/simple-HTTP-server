//Package which handles POST request.
//Checks conditions and creates file in database
package post

import (
	firebase "firebase.google.com/go"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
)

func HandlePost(w http.ResponseWriter, r *http.Request, app *firebase.App) {

	client, err := app.Firestore(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	var newEvent eventT
	newEvent = decode(r)

	if newEvent.Type == "" {
		w.WriteHeader(400)
		return
	}

	var DocRef, _, error = client.Collection("users").Add(context.Background(), map[string]interface{}{
		"Count":     newEvent.Count,
		"Type":      newEvent.Type,
		"Timestamp": newEvent.Timestamp,
	})
	if error != nil {
		w.WriteHeader(502)
		return
	}

	w.Header().Set("Content-type", "text/plain")
	w.WriteHeader(201)
	w.Write([]byte(DocRef.ID))
	return
}
