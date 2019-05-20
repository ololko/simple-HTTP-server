//function parseRequest parses GET request propted by user
//function checks if integers are well written
package get

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"fmt"
	"net"
	"net/url"
)

func parseRequest(requestLine []string) (requestT, error) {
	request := requestT{
		math.MinInt64,
		math.MaxInt64,
		"",
	}

	u, err := url.Parse(requestLine)
    if err != nil {
        panic(err)
    }

	fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])




   /* 
	for i := 0; i < len(requestLine); i++ {
		if strings.Contains(requestLine[i], "from=") {
			from, err := strconv.ParseInt(requestLine[i][5:], 10, 64)
			if err != nil {
				return request, errors.New("Parsing error")
			}
			request.from = from
			continue
		}
		if strings.Contains(requestLine[i], "to=") {
			to, err := strconv.ParseInt(requestLine[i][3:], 10, 64)
			if err != nil {
				return request, errors.New("Parsing error")
			}
			request.to = to
			continue
		}
		if strings.Contains(requestLine[i], "type=") {
			searchedEvent := requestLine[i][5:]
			request.searchedEvent = searchedEvent
			continue
		}
	}
	return request, nil*/
}
