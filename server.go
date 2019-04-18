package main


import (
  //"fmt"
  "log"

  "golang.org/x/net/context"

  firebase "firebase.google.com/go"
  //"firebase.google.com/go/auth"

  "google.golang.org/api/option"
)

func main() {

    opt := option.WithCredentialsFile("./serviceAccountKey.json")
    ctx := context.Background()
    app, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
      log.Fatalln(err)
    }

    client, err := app.Firestore(ctx)
    if err != nil {
        log.Fatalln(err)
    }

    _, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
        "first": "Ada",
        "last":  "Lovelace",
        "born":  1815,
    })
    if err != nil {
        log.Fatalf("Failed adding alovelace: %v", err)
    }

    defer client.Close()

//      SERVER SIDE
    /*http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "events")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))*/

}