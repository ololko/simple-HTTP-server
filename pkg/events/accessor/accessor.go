package accessor

import(
    "github.com/ololko/simple-HTTP-server/pkg/events/models"
)

type DataAccesser interface{
    ReadEvent(models.RequestT) (models.AnswerT, error)
    WriteEvent(models.EventT) ([]byte, error)
}