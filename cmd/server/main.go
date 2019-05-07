/*
Main fuinction of server.
Server binds ports here and listens to incomming connection
*/
package main

import (
	firebase "firebase.google.com/go"
	"github.com/ololko/simple-http-server/pkg/get"
	"github.com/ololko/simple-http-server/pkg/post"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
	"net/http"
)

func main() {

	port := ":10000"
	path := "serviceAccountKey.json"

	opt := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			get.HandleGet(w, r, app)

		} else if r.Method == "POST" {
			post.HandlePost(w, r, app)

		} else {
			w.WriteHeader(501)
		}
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
