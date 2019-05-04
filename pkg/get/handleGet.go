package get

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

  var(
    from int64
    to int64
    count int64
    searchedEvent string
    )

func init(){
  from = 0
  to = 0
  count = 0
  searchedEvent = ""
  from = math.MinInt64
  to = math.MaxInt64
}

func HandleGet(w http.ResponseWriter, r *http.Request, app *firebase.App){

  client, err := app.Firestore(context.Background())
            if err != nil {
                log.Fatalln(err)
            }
  defer client.Close()

  var requestLine = strings.Split(r.URL.RawQuery,"&")  

  from = math.MinInt64
  to = math.MaxInt64
  searchedEvent = ""
  request = parseRequest(requestLine)

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