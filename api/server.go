package main


import (
  "fmt"
  "log"
  "net/http"
  "myDecoder"
  "eventStructure"
  "strings"
  "strconv"
  "google.golang.org/api/iterator"
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
    maxTime := 0
    minTime := 0

    http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request){
      if r.Method == "GET"{
          var querries = strings.Split(r.URL.RawQuery,"&")
          
          from := minTime
          to := maxTime
          //count := 0
          var searchedEvent string
          for i := 0; i < len(querries); i++ {
            if strings.Contains(querries[i],"from="){
              from,err = strconv.Atoi(querries[i][5:])
              if err != nil {
                log.Fatal("Couldn't convert string to integer!")
              }
            }
            if strings.Contains(querries[i],"to="){
              to,err = strconv.Atoi(querries[i][3:])
              if err != nil {
                log.Fatal("Couldn't convert string to integer!")
              }
            }
            if strings.Contains(querries[i],"type="){
              searchedEvent = querries[i][5:]
            }
          }

          iter := client.Collection("users").Documents(ctx)
          for {
                  doc, err := iter.Next()
                  if err == iterator.Done {
                    break
                  }
                  if err != nil {
                    log.Fatalf("Error reading data %v", err) 
                  }
                  fmt.Println(doc.Data())
          }

          fmt.Println(from)
          fmt.Println(to)
          fmt.Println(searchedEvent)
          //send Json

        } else if r.Method == "POST"{
            var newEvent eventStructure.Event
            newEvent = myDecoder.Decode(r)

            if newEvent.Timestamp > maxTime {
              maxTime = newEvent.Timestamp
            }
            if newEvent.Timestamp < minTime {
              minTime = newEvent.Timestamp
            }

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