package apis

import(
	"github.com/ololko/simple-http-server/pkg/events/readers"
)

type Service struct {
	DataAccesser readers.DataAccesser
}

type Mock struct {}