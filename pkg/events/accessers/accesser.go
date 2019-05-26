package accessers

import(
    "github.com/ololko/simple-HTTP-server/pkg/events/models"
)

type DataAccesser interface{
    Read (models.RequestT) (models.AnswerT, error)
    Write (models.EventT) ([]byte, error)
}