//Package which handles POST request.
//Checks conditions and creates file in database
package apis

import (
	"encoding/json"
	"net/http"

	//"cloud.google.com/go/firestore"
	"github.com/ololko/simple-http-server/pkg/events/models"
	//"golang.org/x/net/context"
)

func (s *Service) HandlePost(w http.ResponseWriter, r *http.Request) {

	var newEvent models.EventT
	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Id, err := s.DataAccesser.Write(newEvent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	

	w.Header().Set("Content-type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write(Id)
	return
}
