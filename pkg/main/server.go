package main

//fmt.Println(reflect.TypeOf(tst))

import (
  "fmt"
  "reflect"
  "encoding/json"
  "log"
  "net/http"
  "decode"
  "eventStructure"
  "strings"
  "strconv"
  "google.golang.org/api/iterator"
  "os"
  "golang.org/x/net/context"
  firebase "firebase.google.com/go"
  "google.golang.org/api/option"
)

type AnswerStruct struct {
  Count int64
  Type string
}
type Limit struct {
  minTime int64
  maxTime int64
}

func handleGET(w http.ResponseWriter, r *http.Request, client firestore.Client ,limit Limit){
  
  var querries = strings.Split(r.URL.RawQuery,"&")  

  from := limit.minTime
  to := limit.maxTime
  var count int64
  var searchedEvent string

  for i := 0; i < len(querries); i++ {
    if strings.Contains(querries[i],"from="){
      from,err := strconv.ParseInt(querries[i][5:], 10, 64)
      if err != nil {
        w.WriteHeader(400)
        return
      }
      continue
    }
    if strings.Contains(querries[i],"to="){
      to,err := strconv.ParseInt(querries[i][3:], 10, 64)
      if err != nil {
        w.WriteHeader(400)
        return
      }
      continue
    }
    if strings.Contains(querries[i],"type="){
      searchedEvent = querries[i][5:]
      continue
    }
  }

  iter := client.Collection("users").Where("Type", "==", searchedEvent).Where("Timestamp", ">=", from).Where("Timestamp", "<=", to).Documents(context.Background())
  for {
        doc, err := iter.Next()
        if err == iterator.Done {
          break
        }
        if err != nil {
          w.WriteHeader(502)
          return
        }

        if dataC, ok := doc.Data()["Count"].(int64); ok {
            count += dataC
        } else {
            w.WriteHeader(500)
            return
        }
  }

  answ := AnswerStruct{count, searchedEvent}
  var answJson,err = json.Marshal(answ)
  if err != nil {
    w.WriteHeader(500)
    return
  }

  w.Header().Set("Content-type", "application/json")
  w.WriteHeader(200)
  w.Write(answJson)
}



func main() {

    port := ":" + os.Args[1]
    path := os.Args[2]

    opt := option.WithCredentialsFile(path)
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
    fmt.Println(reflect.TypeOf(client))

//      SERVER SIDE
    var limit Limit
    limit.maxTime = 0
    limit.minTime = 9999

    http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request){
      if r.Method == "GET"{
        handleGET(w, r, client, limit)
        
        } else if r.Method == "POST"{
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

            var DocRef ,_, err = client.Collection("users").Add(ctx, map[string]interface{}{
                "Count"     : newEvent.Count,
                "Type"      : newEvent.Type,
                "Timestamp" : newEvent.Timestamp,
            })
            if err != nil {
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