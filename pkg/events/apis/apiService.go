package apis

import "github.com/ololko/simple-HTTP-server/pkg/events/accessers"

func NewFirestoreService(dataAccesser accessers.DataAccesser) *Service {
	return &Service{DataAccesser: dataAccesser}
}

func NewMockService() *Mock {
	return &Mock{}
}

type Service struct {
	DataAccesser accessers.DataAccesser
}

type Mock struct{}
