package readers

import(
    "github.com/ololko/simple-http-server/pkg/events/models"
)

type DataAccesser interface{
    Read (models.RequestT) (models.AnswerT, error)
    Write (models.EventT) ([]byte, error)
}