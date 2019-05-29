package apis

import (
	"bytes"
	"encoding/json"
	"github.com/ololko/simple-HTTP-server/pkg/events/accessor"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ApiSuite struct {
	suite.Suite
	service *Service
}

func (s *ApiSuite) SetupSuite() {
	s.service = NewService(&accessor.MockAccess{})
}

func (s *ApiSuite) SetupTest() {
	// add new data to database
	s.service.DataAccessor = &accessor.MockAccess{
		Events: map[string][]models.EventT{
			"Skuska": {
				{
					Count:     2,
					Type:      "Skuska",
					Timestamp: 0,
				},
				{
					Count:     4,
					Type:      "Skuska",
					Timestamp: 6,
				},
				{
					Count:     8,
					Type:      "Skuska",
					Timestamp: 4,
				},
				{
					Count:     1,
					Type:      "Skuska",
					Timestamp: 10,
				},
			},
		},
	}
}

func (s *ApiSuite) TestHandleGet() {
	candidates := []struct {
		url          string
		expected     models.AnswerT
		expectedCode int
	}{
		{
			url: "/events?type=Skuska&from=3&to=7",
			expected: models.AnswerT{
				Type:  "Skuska",
				Count: 12,
			},
			expectedCode: http.StatusOK,
		},
		{
			url:          "/events?type=Skuska&from=3&to=75fdg",
			expected:     models.AnswerT{},
			expectedCode: http.StatusBadRequest,
		},
		{
			url: "/events?type=NoData&from=3",
			expected: models.AnswerT{
				Type:  "",
				Count: 0,
			},
			expectedCode: http.StatusNotFound,
		},
		{
			url: "/events?type=Skuska&from=10",
			expected: models.AnswerT{
				Type:  "Skuska",
				Count: 1,
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, c := range candidates {
		req, err := http.NewRequest("GET", c.url, nil)
		s.NoError(err)

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(s.service.HandleGet)

		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		s.Equal(c.expectedCode, rr.Code)

		received := models.AnswerT{}

		err = json.Unmarshal(rr.Body.Bytes(), &received)
		s.NoError(err)

		// Check the response body is what we expect.
		s.Equal(c.expected, received)
	}
}

func (s *ApiSuite) TestHandlePost() {
	candidates := []struct {
		newEvent models.EventT
		response string
		expectedCode int
	}{
		{
			newEvent: models.EventT{
				Type: "Skuska",
				Timestamp: 100,
				Count: 3,
			},
			response: "Skuska",
			expectedCode:http.StatusCreated,
		},
		{
			newEvent: models.EventT{
				Type: "Skuska",
				Timestamp: 5,
				Count: 36,
			},
			response: "Skuska",
			expectedCode:http.StatusCreated,
		},
	}

	for _, c := range candidates {
		bb, _ := json.Marshal(c.newEvent)
		req, err := http.NewRequest("POST", "/events", bytes.NewBuffer(bb))
		s.NoError(err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.service.HandlePost)
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		s.Equal(c.expectedCode,rr.Code)

		// Check the response requestBody is what we expect.
		s.Equal(c.response,rr.Body.String())
	}
}

func TestApiSuite(t *testing.T) {
	suite.Run(t, new(ApiSuite))
}
