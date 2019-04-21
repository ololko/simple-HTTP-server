package main


import (
  "fmt"
  "log"
  "net/http"
  "myDecoder"
  "eventStructure"
  //"net/url"
  "golang.org/x/net/context"
  firebase "firebase.google.com/go"
  //"firebase.google.com/go/auth"
  "google.golang.org/api/option"
)

func main() {

    opt := option.WithCredentialsFile("../configs/serviceAccountKey.json")
    ctx := context.Background()
    app, err := firebase.NewApp(ctx, nil, opt)
    if err != nil {
      log.Fatalln(err)
    }

    client, err := app.Firestore(ctx)
    if err != nil {
        log.Fatalln(err)
    }
    defer client.Close()

//      SERVER SIDE
    http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request){
      if r.Method == "GET"{
          fmt.Fprintf(w, "NOT implemented")
          fmt.Println(r.URL.RawQuery)

        } else if r.Method == "POST"{

          var newEvent eventStructure.Event
          newEvent = myDecoder.Decode(r)
          fmt.Fprintf(w, newEvent.Type)

          _, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
              "Count"     : newEvent.Count,
              "Type"      : newEvent.Type,
              "Timestamp" : newEvent.Timestamp,
          })
          if err != nil {
              log.Fatalf("Failed adding alovelace: %v", err)
          }


        } else {
          fmt.Fprintf(w, "hod err not supported method")

        }
    })

    log.Fatal(http.ListenAndServe(":10000", nil))

}