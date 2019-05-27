/*
Main fuinction of server.
Server binds ports here and listens to incomming connection
*/
package main

import (
	"fmt"
	"github.com/ololko/simple-HTTP-server/pkg/events/accessor"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/ololko/simple-HTTP-server/pkg/events/apis"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

const(
	port = ":10000"
	path = "serviceAccountKey.json"
)

func main() {
	opt := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	datAcc := &accessor.MockAccess{make(map[string][]models.EventT)}
	svc := apis.NewService(datAcc)

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			svc.HandleGet(w, r)
		} else if r.Method == "POST" {
			svc.HandlePost(w, r)

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
