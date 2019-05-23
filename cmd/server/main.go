/*
Main fuinction of server.
Server binds ports here and listens to incomming connection
*/
package main

import (
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/ololko/simple-http-server/pkg/events/apis"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func main() {

	port := ":10000"
	path := "serviceAccountKey.json"

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

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			apis.HandleGet(w, r, client)
		} else if r.Method == "POST" {
			apis.HandlePost(w, r, client)

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
