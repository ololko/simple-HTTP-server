package apis

import (
	"bytes"
	"cloud.google.com/go/firestore"
	"encoding/json"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/labstack/echo"
	"github.com/ololko/simple-HTTP-server/pkg/events/access"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ApiSuite struct {
	suite.Suite
	service *Service
	client *firestore.Client
}

const(
	path = "../../../configs/serviceAccountKey.json"
)

func (s *ApiSuite) SetupSuite() {
	//s.service = NewService(&access.MockAccess{})

	opt := option.WithCredentialsFile(path)
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
	s.client = client
}

func (s *ApiSuite) SetupTest() {
	// add new data to database
	/*s.service.DataAccessor = &access.MockAccess{
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
	}*/
}

func (s *ApiSuite) TestHandleGet(t *testing.T) {
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
			expected: models.AnswerT{},
			expectedCode: http.StatusNotFound,
		},
		{
			url: "/events?type=NoData",
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

		if assert.NoError(t, h(ctx)) {
			assert.Equal(t, candidate.expectedCode, rec.Code)
			assert.Equal(t, candidate.expected, rec.Body)
		}
	}
}

func (s *ApiSuite) TestHandlePost(t *testing.T) {
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



	e := echo.New()
	h := s.service.HandlePost

	for _, c := range candidates {
		bb, _ := json.Marshal(c.newEvent)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(bb))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, h(ctx)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, c.response, rec.Body.String())
		}
	}
}

func TestApiSuite(t *testing.T) {
	suite.Run(t, new(ApiSuite))
}


func (s *ApiSuite) TearDownSuite() {
	s.client.Close()
}