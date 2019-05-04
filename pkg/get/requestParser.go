package get

import()

func parseRequest request() {
	request := request {
		0,
		0,
		""
	}
	for i := 0; i < len(requestLine); i++ {
	    if strings.Contains(requestLine[i],"from="){
	      from,err = strconv.ParseInt(requestLine[i][5:], 10, 64)
	      if err != nil {
	        w.WriteHeader(400)
	        return
	      }
	      request.from = from
	      continue
	    }
	    if strings.Contains(requestLine[i],"to="){
	      to,err = strconv.ParseInt(requestLine[i][3:], 10, 64)
	      if err != nil {
	        w.WriteHeader(400)
	        return
	      }
	      request.to = to
	      continue
	    }
	    if strings.Contains(requestLine[i],"type="){
	      searchedEvent = requestLine[i][5:]
	      request.searchedEvent = searchedEvent
	      continue
	    }
  	}
  	return request
}

