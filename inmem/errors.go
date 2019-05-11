package inmem

import (
	"fmt"
)

type errTodoNotFound struct {
	ID string
}

func (e *errTodoNotFound) Error() string {
	return fmt.Sprintf("todo with id %s not found", e.ID)
}

func (e *errTodoNotFound) ErrTodoNotFound() (bool, string) {
	return true, e.ID
}
