package main

import (
  "github.com/ololko/simple-http-server/pkg/get"
  "github.com/ololko/simple-http-server/pkg/post"
  "log"
  "net/http"
  "os"
  "golang.org/x/net/context"
  firebase "firebase.google.com/go"
  "google.golang.org/api/option"
)


func main() {

  port := ":" + os.Args[1]
  path := os.Args[2]

  opt := option.WithCredentialsFile(path)
  app, err := firebase.NewApp(context.Background(), nil, opt)
  if err != nil {
    log.Fatalln(err)
  }

  http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request){
    if r.Method == "GET"{
      get.HandleGet(w, r, app)

    } else if r.Method == "POST"{
      post.HandlePost(w,r,app)

      } else {
        w.WriteHeader(501)
      }
  })

  log.Fatal(http.ListenAndServe(port, nil))
}
