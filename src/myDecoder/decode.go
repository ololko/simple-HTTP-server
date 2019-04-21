package myDecoder 

import (
	"encoding/json"
	"eventStructure"
    "log"
    "net/http"
)

func Decode(req *http.Request) eventStructure.Event {
    decoder := json.NewDecoder(req.Body)
    var t eventStructure.Event
    err := decoder.Decode(&t)
    if err != nil {
    	log.Println("error")
        panic(err)
    }
    log.Println(t.Count)
    log.Println(t)
    return t
}