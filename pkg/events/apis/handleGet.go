//Package apis  handles GET request.
//Checks conditions and answers to user
package apis

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"github.com/ololko/simple-http-server/pkg/events/models"
)

func (s *Service) HandleGet(w http.ResponseWriter, r *http.Request) {

	request, err := fillRequestStruct(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	data,err := s.DataAccesser.Read(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	answJSON, err := json.Marshal(data)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(answJSON)
}
