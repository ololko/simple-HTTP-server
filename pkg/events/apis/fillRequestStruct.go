package apis

import (
	"math"
	"net/http"
	"strconv"

	"github.com/ololko/simple-HTTP-server/pkg/events/models"
)

func fillRequestStruct(r *http.Request) (models.RequestT, error) {
	q := r.URL.Query()
	var to int64 = math.MaxInt64
	var from int64 = math.MinInt64
	var err error
	var request models.RequestT

	if q.Get("from") != "" {
		from, err = strconv.ParseInt(q.Get("from"), 10, 64)
		if err != nil {
			return request, err
		}
	}

	if q.Get("to") != "" {
		to, err = strconv.ParseInt(q.Get("to"), 10, 64)
		if err != nil {
			return request, err
		}
	}

	searchedEvent := q.Get("type")

	request.From = from
	request.To = to
	request.Type = searchedEvent

	return request, nil
}
