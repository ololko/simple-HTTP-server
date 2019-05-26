package accessers

import "github.com/ololko/simple-HTTP-server/pkg/events/models"

type MockAccesser struct {
}

func (d *MockAccesser) Read(request models.RequestT) (models.AnswerT, error) {
	return models.AnswerT{}, nil
}

func (d *MockAccesser) Write(request models.EventT) ([]byte, error) {
	return []byte{}, nil
}