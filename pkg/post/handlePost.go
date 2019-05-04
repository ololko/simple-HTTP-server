package post

import (
  "log"
  "net/http"
  "os"
  "golang.org/x/net/context"
  firebase "firebase.google.com/go"
  "google.golang.org/api/option"
)

func HandlePost(w http.ResponseWriter, r *http.Request, app *firebase.App) {
	
	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println(err)
	    return
	}
	defer client.Close()

	var newEvent eventT
	newEvent = decode(r)

	if newEvent.Type == ""{
		w.WriteHeader(400)
		return
	}

	var DocRef ,_, error = client.Collection("users").Add(ctx, map[string]interface{}{
	"Count"     : newEvent.Count,
	"Type"      : newEvent.Type,
	"Timestamp" : newEvent.Timestamp,
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