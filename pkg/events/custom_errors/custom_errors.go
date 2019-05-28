package custom_errors

import "fmt"

type ElementDoesNotExistError struct {
	Path string
	CanContinue bool
	ReallyError bool
}

func (e *ElementDoesNotExistError) Error() string {
	return fmt.Sprintf(e.Path)
}
