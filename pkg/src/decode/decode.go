package decode

import (
	"encoding/json"
	"eventStructure"
    "net/http"
)

func Decode(req *http.Request) eventStructure.Event {
    decoder := json.NewDecoder(req.Body)
    var t eventStructure.Event
    err := decoder.Decode(&t)
    if err != nil {
    	t.Timestamp = 0
        t.Count = 0
        t.Type = ""
    }
    return t
}
