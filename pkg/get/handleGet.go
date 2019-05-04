package get

import(
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "strings"
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

  var requestLine = strings.Split(r.URL.RawQuery,"&")  

  request,err := parseRequest(requestLine)
  if err != nil{
    w.WriteHeader(400)
    return
  }

  var count int64
  count = 0
  iter := client.Collection("users").Where("Type", "==", request.searchedEvent).Where("Timestamp", ">=", request.from).Where("Timestamp", "<=", request.to).Documents(context.Background())
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

        if recData, ok := doc.Data()["Count"].(int64); ok {
            count += recData
        } else {
            w.WriteHeader(500)
            return
        }
  }

  answ := answerT{count, request.searchedEvent}
  var answJson,error = json.Marshal(answ)
  if error != nil {
    w.WriteHeader(500)
    return
  }

  w.Header().Set("Content-type", "application/json")
  w.WriteHeader(200)
  w.Write(answJson) 
}