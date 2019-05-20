//Function decodes JSON and fills database structure
package post

import (
	"encoding/json"
	"net/http"
)

func decode(req *http.Request) (eventT, error) {
	decoder := json.NewDecoder(req.Body)
	var t eventT
	err := decoder.Decode(&t)
	return t, err
}
