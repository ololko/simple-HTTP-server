package access

import (
    "github.com/ololko/simple-HTTP-server/pkg/events/models"
)

type DataAccessor interface{
    ReadEvent(models.Request, chan<- int32, chan<- error)
    WriteEvent(models.Event, chan<- error)
}