package decode

import (
	"encoding/json"
	"buffers"
    "net/http"
)

func Decode(req *http.Request) buffers.Event {
    decoder := json.NewDecoder(req.Body)
    var t buffers.Event
    err := decoder.Decode(&t)
    if err != nil {
    	t.Timestamp = 0
        t.Count = 0
        t.Type = ""
    }
    return t
}
