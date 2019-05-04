package main

import (
  "github.com/ololko/simple-http-server/pkg/decode"
  "github.com/ololko/simple-http-server/pkg/eventStructure" 
  "github.com/ololko/simple-http-server/pkg/handleGet"
  "github.com/ololko/simple-http-server/pkg/limit"
  "log"
  "net/http"
  "os"
  "math"
  "golang.org/x/net/context"
  firebase "firebase.google.com/go"
  "google.golang.org/api/option"
)


func main() {

    port := ":" + os.Args[1]
    path := os.Args[2]

    opt := option.WithCredentialsFile(path)
    ctx := context.Background()
    app, err := firebase.NewApp(ctx, nil, opt)
    if err != nil {
      log.Fatalln(err)
    }

//      SERVER SIDE
    var limit limit.Limit
    limit.MaxTime = math.MaxInt64
    limit.MinTime = math.MinInt64

    http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request){
      if r.Method == "GET"{
        handleGet.HandleGET(w, r, app, limit)
        
        } else if r.Method == "POST"{

            client, err := app.Firestore(ctx)
            if err != nil {
                log.Fatalln(err)
            }
            defer client.Close()

            var newEvent eventStructure.Event
            newEvent = decode.Decode(r)

            if newEvent.Type == ""{
              w.WriteHeader(400)
              return
            }

            if newEvent.Timestamp > limit.maxTime {
              limit.maxTime = newEvent.Timestamp
            }
            if newEvent.Timestamp < limit.minTime {
              limit.minTime = newEvent.Timestamp
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

        } else {
          w.WriteHeader(501)
        }
    })

    log.Fatal(http.ListenAndServe(port, nil))
}
