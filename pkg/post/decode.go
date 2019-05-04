package post

import (
	"encoding/json"
    "net/http"
)

func decode(req *http.Request) eventT {
    decoder := json.NewDecoder(req.Body)
    var t eventT
    err := decoder.Decode(&t)
    if err != nil {
    	t.Timestamp = 0
        t.Count = 0
        t.Type = ""
    }
    return t
}
