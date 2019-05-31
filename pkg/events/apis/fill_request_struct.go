package apis

import (
	"math"
	"net/url"
	"strconv"

	"github.com/ololko/simple-HTTP-server/pkg/events/models"
)

func fillRequestStruck(u *url.URL) (models.RequestT, error) {
	q := u.Query()
	var to int64 = math.MaxInt64
	var from int64 = math.MinInt64
	var err error

	if q.Get("from") != "" {
		from, err = strconv.ParseInt(q.Get("from"), 10, 64)
		if err != nil {
			return models.RequestT{}, err
		}
	}

	if q.Get("to") != "" {
		to, err = strconv.ParseInt(q.Get("to"), 10, 64)
		if err != nil {
			return models.RequestT{}, err
		}
	}

	request := models.RequestT{}

	request.From = from
	request.To = to
	request.Type = q.Get("type")

	return request, nil
}
