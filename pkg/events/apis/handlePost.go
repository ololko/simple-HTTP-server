//Package which handles POST request.
//Checks conditions and creates file in database
package apis

import (
	"encoding/json"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/ololko/simple-http-server/pkg/events/models"
	"golang.org/x/net/context"
)

func HandlePost(w http.ResponseWriter, r *http.Request, client *firestore.Client) {

	var newEvent models.EventT
	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var DocRef, _, error = client.Collection("users").Add(context.Background(), map[string]interface{}{
		"Count":     newEvent.Count,
		"Type":      newEvent.Type,
		"Timestamp": newEvent.Timestamp,
	})
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(DocRef.ID))
	return
}
