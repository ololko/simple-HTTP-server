package accessor

import(
    "github.com/ololko/simple-HTTP-server/pkg/events/custom_errors"
    "github.com/ololko/simple-HTTP-server/pkg/events/models"
)

type DataAccesser interface{
    ReadEvent(models.RequestT) (models.AnswerT, custom_errors.ElementDoesNotExistError)
    WriteEvent(models.EventT) ([]byte, error)
}