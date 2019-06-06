package apis

import (
	"bytes"
	"cloud.google.com/go/firestore"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/ololko/simple-HTTP-server/pkg/events/access"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ApiSuite struct {
	suite.Suite
	service *Service
	client  *firestore.Client
}

type eventInvalidTimestamp struct {
	Type      string
	Count     int
	Timestamp string
}

type eventInvalidType struct {
	Type      bool
	Count     int
	Timestamp int
}


const (
	path = "../../../configs/serviceAccountKey.json"
)

func (s *ApiSuite) SetupSuite() {
	s.service = NewService(&access.MockAccess{})

	/*opt := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	datAcc := &access.FirestoreAccess{Client: client}
	s.service = NewService(datAcc)
	s.client = client*/
}

func (s *ApiSuite) SetupTest() {
	// add new data to database
	s.service.DataAccessor = &access.MockAccess{
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



func (s *ApiSuite) TestGetValidInputs() {
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
			url: "/events?type=Skuska&from=10",
			expected: models.AnswerT{
				Type:  "Skuska",
				Count: 1,
			},
			expectedCode: http.StatusOK,
		},
		{
			url: "/events?type=Skuska&to=7&from=3",
			expected: models.AnswerT{
				Type:  "Skuska",
				Count: 12,
			},
			expectedCode: http.StatusOK,
		},
	}

	e := echo.New()

	for _, candidate := range candidates {
		req := httptest.NewRequest(http.MethodGet, candidate.url, nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		h := s.service.HandleGet

		s.NoError(h(ctx))
		var received models.AnswerT
		s.NoError(json.Unmarshal(rec.Body.Bytes(), &received))

		assert.Equal(s.T(), candidate.expectedCode, rec.Code)
		assert.Equal(s.T(), candidate.expected, received)
	}
}

func (s *ApiSuite) TestGetBadRequest() {
	candidates := []struct {
		url          string
		expected     string
		expectedCode int
	}{
		{
			url:          "/events?type=Skuska&from=3&to=75fdg",
			expectedCode: http.StatusBadRequest,
		},
		{
			url:          "/events?type=Skuska&from=t&to=",
			expectedCode: http.StatusBadRequest,
		},
	}

	e := echo.New()

	for _, candidate := range candidates {
		req := httptest.NewRequest(http.MethodGet, candidate.url, nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		h := s.service.HandleGet

		s.NoError(h(ctx))

		assert.Equal(s.T(), candidate.expectedCode, rec.Code)
		assert.Equal(s.T(), candidate.expected, rec.Body.String())
	}
}

func (s *ApiSuite) TestGetNotFound() {
	candidates := []struct {
		url          string
		expected     string
		expectedCode int
	}{
		{
			url:          "/events?type=NoData&from=3",
			expectedCode: http.StatusNotFound,
		},
		{
			url:          "/events?type=NoData",
			expectedCode: http.StatusNotFound,
		},
		{
			url: "/events?type=Skuska&from=7&to=3",
			expectedCode: http.StatusNotFound,
		},
	}

	e := echo.New()

	for i, candidate := range candidates {
		req := httptest.NewRequest(http.MethodGet, candidate.url, nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		h := s.service.HandleGet

		s.NoError(h(ctx))

		assert.Equal(s.T(), candidate.expectedCode, rec.Code, "Error in expected code in test number %d", i)
		assert.Equal(s.T(), candidate.expected, rec.Body.String())
	}
}

func (s *ApiSuite) TestPostValidBody() {
	candidates := []struct {
		newEvent     models.EventT
		response     string
		expectedCode int
	}{
		{
			newEvent: models.EventT{
				Type:      "Skuska",
				Timestamp: 100,
				Count:     3,
			},
			response:     "Skuska",
			expectedCode: http.StatusCreated,
		},
		{
			newEvent: models.EventT{
				Type:      "BryndzoveHalusky",
				Timestamp: 5,
				Count:     36,
			},
			response:     "BryndzoveHalusky",
			expectedCode: http.StatusCreated,
		},
	}

	e := echo.New()
	h := s.service.HandlePost

	for _, c := range candidates {
		bb, _ := json.Marshal(c.newEvent)
		req := httptest.NewRequest(http.MethodPost, "/events", bytes.NewBuffer(bb))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		s.NoError(h(ctx))
		var received string
		s.NoError(json.Unmarshal(rec.Body.Bytes(), &received))

		assert.Equal(s.T(), c.expectedCode, rec.Code)
		assert.Equal(s.T(), c.response, received)
	}
}

func (s *ApiSuite) TestPostInvalidTimestamp() {
	candidates := []struct {
		newEvent     eventInvalidTimestamp
		response     string
		expectedCode int
	}{
		{
			newEvent: eventInvalidTimestamp{
				Type:      "Skuska",
				Timestamp: "100",
				Count:     3,
			},
			response:     "",
			expectedCode: http.StatusBadRequest,
		},
		{
			newEvent: eventInvalidTimestamp{
				Type:      "DacoIne",
				Timestamp: "10sd5s",
				Count:     36,
			},
			response:     "",
			expectedCode: http.StatusBadRequest,
		},
		{
			newEvent: eventInvalidTimestamp{
				Type:      "ZaseNovyEvent",
				Timestamp: "96sd",
				Count:     36,
			},
			response:     "",
			expectedCode: http.StatusBadRequest,
		},
	}

	e := echo.New()
	h := s.service.HandlePost

	for _, c := range candidates {
		bb, _ := json.Marshal(c.newEvent)
		req := httptest.NewRequest(http.MethodPost, "/events", bytes.NewBuffer(bb))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		s.NoError(h(ctx))
		assert.Equal(s.T(), http.StatusBadRequest, rec.Code)
		assert.Equal(s.T(), c.response, rec.Body.String())
	}
}

func (s *ApiSuite) TestPostInValidTypeBool() {
	candidates := []struct {
		newEvent     eventInvalidType
		response     string
		expectedCode int
	}{
		{
			newEvent: eventInvalidType{
				Type:      true,
				Timestamp: 100,
				Count:     3,
			},
			response:     "",
			expectedCode: http.StatusBadRequest,
		},
		{
			newEvent: eventInvalidType{
				Type:      false,
				Timestamp: 10,
				Count:     36,
			},
			response:     "",
			expectedCode: http.StatusBadRequest,
		},
	}

	e := echo.New()
	h := s.service.HandlePost

	for _, c := range candidates {
		bb, _ := json.Marshal(c.newEvent)
		req := httptest.NewRequest(http.MethodPost, "/events", bytes.NewBuffer(bb))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		s.NoError(h(ctx))
		assert.Equal(s.T(), c.expectedCode, rec.Code)
		assert.Equal(s.T(), c.response, rec.Body.String())
		fmt.Println(c.newEvent)
	}
}

func TestApiSuite(t *testing.T) {
	suite.Run(t, new(ApiSuite))
}

func (s *ApiSuite) TearDownSuite() {
	//s.NoError(s.client.Close())
}
