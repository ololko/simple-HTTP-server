package handleGet

import(
  "encoding/json"
  "fmt"
  "log"
  "math"
  "net/http"
  "strconv"
  "strings"
  "github.com/ololko/simple-http-server/pkg/answerStructure"
  "google.golang.org/api/iterator" 
  "golang.org/x/net/context"
  firebase "firebase.google.com/go"
)

func HandleGet(w http.ResponseWriter, r *http.Request, app *firebase.App){

  client, err := app.Firestore(context.Background())
            if err != nil {
                log.Fatalln(err)
            }
            defer client.Close()


  var querries = strings.Split(r.URL.RawQuery,"&")  

        from := math.MinInt64
        to := math.MaxInt64
        var count int64
        var searchedEvent string

        for i := 0; i < len(querries); i++ {
          if strings.Contains(querries[i],"from="){
            from,err = strconv.ParseInt(querries[i][5:], 10, 64)
            if err != nil {
              w.WriteHeader(400)
              return
            }
            continue
          }
          if strings.Contains(querries[i],"to="){
            to,err = strconv.ParseInt(querries[i][3:], 10, 64)
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
                fmt.Println(err)
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

        answ := answerStruct.AnswerStruct{count, searchedEvent}
        var answJson,error = json.Marshal(answ)
        if error != nil {
          w.WriteHeader(500)
          return
        }

        w.Header().Set("Content-type", "application/json")
        w.WriteHeader(200)
        w.Write(answJson)
  
}