package access

import (
    "github.com/ololko/simple-HTTP-server/pkg/events/models"
)

type DataAccesser interface{
    ReadEvent(models.RequestT, chan<- models.AnswerT, chan<- error)
    WriteEvent(models.EventT, chan<- string, chan<- error)
}