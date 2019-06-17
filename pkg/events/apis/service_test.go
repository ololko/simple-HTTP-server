package apis

import (
	"database/sql"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ololko/simple-HTTP-server/pkg/events/access"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
	"math"
	"net/http"
	"testing"
)

type ApiSuite struct {
	suite.Suite
	service *Service
	//client  *firestore.Client
	client *sql.DB
}

type errorFromGRPC struct {
	code int
}

type requestInvalidTo struct {
	Type string
	From int
	To   string
}

type requestInvalidFrom struct {
	Type string
	From string
	To   int
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

func (s *ApiSuite) SetupSuite() {
	/*err := myViper.ReadConfig("viperConfig", "../../../configs/")
	if err != nil {
		fmt.Println(err)
		return
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", viper.GetString("host"), viper.GetInt("dbPort"), viper.GetString("user"), viper.GetString("dbname"))
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	//db.DropTableIfExists(&models.DatabaseElement{})
	db.AutoMigrate(&models.DatabaseElement{})*/

	/*	datAcc := &access.PostgreAccess{Client: db}
		s.service = NewService(datAcc)*/

	//mockDB
	s.service = NewService(&access.MockAccess{})

	//firestoreDb
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
		Events: map[string][]models.DatabaseElement{
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

func (s *ApiSuite) TestGet() {
	candidates := []struct {
		request       *models.Request
		expectedBody  *models.Answer
		expectedError error
	}{
		{
			request: &models.Request{
				Type: "Skuska",
				From: 3,
				To:   7,
			},
			expectedBody: &models.Answer{
				Type:  "Skuska",
				Count: 12,
			},
		},
		{
			request: &models.Request{
				Type: "Skuska",
				From: 10,
				To:   math.MaxInt32,
			},
			expectedBody: &models.Answer{
				Type:  "Skuska",
				Count: 1,
			},
		},
		{
			request: &models.Request{
				Type: "Skuska",
				From: 3,
				To:   -9,
			},
			expectedError: status.Error(http.StatusNotFound, "Not found"),
		},
		{
			request: &models.Request{
				From: math.MinInt32,
				To:   math.MaxInt32,
				Type: "NoData",
			},
			expectedError: status.Error(http.StatusNotFound, "Not found"),
		},
		{
			request: &models.Request{
				Type: "NoData",
			},
			expectedError: status.Error(http.StatusNotFound, "Not found"),
		},
		{
			request: &models.Request{
				Type: "Skuska",
				From: -60,
				To:   -12,
			},
			expectedError: status.Error(http.StatusNotFound, "Not found"),
		},
	}

	for _, candidate := range candidates {
		resp, err := s.service.ReadEvent(context.Background(), candidate.request)

		s.Equal(candidate.expectedBody, resp)
		s.Equal(candidate.expectedError, err)
	}
}

func (s *ApiSuite) TestPost() {
	candidates := []struct {
		newEvent      *models.Event
		expectedBody  interface{}
		expectedError error
	}{
		{
			newEvent: &models.Event{
				Type:      "Skuska",
				Timestamp: 100,
				Count:     3,
			},
			expectedBody: &empty.Empty{},
		},
		{
			newEvent: &models.Event{
				Type:      "BryndzoveHalusky",
				Timestamp: 5,
				Count:     36,
			},
			expectedBody: &empty.Empty{},
		},
		/*{
			newEvent: &eventInvalidTimestamp{
				Type:      "Skuska",
				Timestamp: "100",
				Count:     3,
			},
			expectedBody: &empty.Empty{},
		},
		{
			newEvent: eventInvalidTimestamp{
				Type:      "DacoIne",
				Timestamp: "10sd5s",
				Count:     36,
			},
			expectedBody:     "",
			expectedCode: http.StatusBadRequest,
		},
		{
			newEvent: eventInvalidTimestamp{
				Type:      "ZaseNovyEvent",
				Timestamp: "96sd",
				Count:     36,
			},
			expectedBody:     "",
			expectedCode: http.StatusBadRequest,
		},
		{
			newEvent: eventInvalidType{
				Type:      true,
				Timestamp: 100,
				Count:     3,
			},
			expectedBody:     "",
			expectedCode: http.StatusBadRequest,
		},
		{
			newEvent: eventInvalidType{
				Type:      false,
				Timestamp: 10,
				Count:     36,
			},
			expectedBody:     "",
			expectedCode: http.StatusBadRequest,
		},*/
	}

	for _, candidate := range candidates {
		resp, err := s.service.CreateEvent(context.Background(), candidate.newEvent)

		s.Equal(candidate.expectedBody, resp)
		s.Equal(candidate.expectedError, err)
	}
}

func TestApiSuite(t *testing.T) {
	suite.Run(t, new(ApiSuite))
}

func (s *ApiSuite) TearDownSuite() {
	//s.NoError(s.client.Close())
}
