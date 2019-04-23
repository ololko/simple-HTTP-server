package main


import (
  "fmt"
  "encoding/json"
  "log"
  "net/http"
  "decode"
  "eventStructure"
  "strings"
  "strconv"
  "google.golang.org/api/iterator"
  "os"
  //"net/url"
  "golang.org/x/net/context"
  firebase "firebase.google.com/go"
  //"firebase.google.com/go/auth"
  "google.golang.org/api/option"
)

type Answer struct {
  Count int64
  Type string
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


//      SERVER SIDE
    maxTime := 0
    minTime := 9999

    http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request){
      if r.Method == "GET"{
          var querries = strings.Split(r.URL.RawQuery,"&")
          
          from := minTime
          to := maxTime
          var count int64
          var searchedEvent string

          for i := 0; i < len(querries); i++ {
            if strings.Contains(querries[i],"from="){
              from,err = strconv.Atoi(querries[i][5:])
              if err != nil {
                log.Fatal("Couldn't convert string to integer!")
              }
              continue
            }
            if strings.Contains(querries[i],"to="){
              to,err = strconv.Atoi(querries[i][3:])
              if err != nil {
                log.Fatal("Couldn't convert string to integer!")
              }
              continue
            }
            if strings.Contains(querries[i],"type="){
              searchedEvent = querries[i][5:]
              continue
            }
          }

          iter := client.Collection("users").Where("Type", "==", searchedEvent).Where("Timestamp", ">=", from).Where("Timestamp", "<=", to).Documents(ctx)
          for {
                doc, err := iter.Next()
                if err == iterator.Done {
                  fmt.Println("iterator done")
                        break
                }
                if err != nil {
                        log.Fatalf("Error reading data %v", err)
                }

                if dataC, ok := doc.Data()["Count"].(int64); ok {
                    count += dataC
                } else {
                    log.Fatal("Couldn't convert int to integer64!")
                }
          }

          fmt.Println(count)
          //send Json
          answ := Answer{count, searchedEvent}
          var answJson,err = json.Marshal(answ)
          if err != nil {
            log.Fatalf("Nepodarilo sa vytvorit Json")
          }
          w.Header().Set("Content-type", "application/json")
          w.WriteHeader(http.StatusOK)
          w.Write(answJson)


        } else if r.Method == "POST"{
            var newEvent eventStructure.Event
            newEvent = decode.Decode(r)

            if newEvent.Type == ""{
              log.Fatalf("Hod err zle poslane parametre")
            }

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

    log.Fatal(http.ListenAndServe(port, nil))

}