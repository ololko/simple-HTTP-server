package apis

import (
	"encoding/json"
	"github.com/ololko/simple-HTTP-server/pkg/events/accessor"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHandleGet(t *testing.T) {
	candidates := []struct{
		events map[string][]models.EventT
		response models.AnswerT
	}{
		{
			events: map[string][]models.EventT{
				"Skuska": {
					{
						Count: 2,
						Type: "Skuska",
						Timestamp: 0,
					},
					{
						Count: 4,
						Type: "Skuska",
						Timestamp: 6,
					},
				},
			},
			response: models.AnswerT{
				Type: "Skuska",
				Count: 6,
			},
		},
		{
			events: map[string][]models.EventT{
				"Skuska": {
					{
						Count: 13,
						Type: "Skuska",
						Timestamp: 0,
					},
				},
			},
			response: models.AnswerT{
				Type: "Skuska",
				Count: 13,
			},
		},
		{
			events: map[string][]models.EventT{
			},
			response: models.AnswerT{
				Type: "Skuska",
				Count: 0,
			},
		},
	}

	for _, c := range candidates {
		req, err := http.NewRequest("GET", "/events?type=Skuska", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		svc := &Service{
			DataAccessor: &accessor.MockAccess{
				Events: c.events,
			},
		}

		handler := http.HandlerFunc(svc.HandleGet)

		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		// Check the response body is what we expect.
		var recieved models.AnswerT
		expected := c.response // expecting structure
		json.Unmarshal(rr.Body.Bytes(), &recieved)

		if recieved != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}



type requestBody struct{
	content map[string]string

}

func (rb *requestBody) Read (p []byte) (n int, err error){
	var i int
	for j, headder := range rb.content {
		for k,c := range headder {
			i = i+k
			p[i] = byte(c)
		}
	}
	return i, nil
}



func TestHandlePost(t *testing.T) {
	candidates := []struct{
		rb requestBody
		response string
	}{
		{
			rb: requestBody. {"type": {"Skuska"}, "count": {"100"}, "timestamp": {"3"}},
			response: "Skuska",
		},
	}

	for _, c := range candidates {
		req, err := http.NewRequest("POST", "/events", c.requestBody)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		svc := &Service{
			DataAccessor: &accessor.MockAccess{
				Events: make(map[string][]models.EventT),
			},
		}

		handler := http.HandlerFunc(svc.HandlePost)

		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		}

		// Check the response requestBody is what we expect.
		expected := c.response // expecting structure
		recieved := rr.Body.String()

		if recieved != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}



}
