package inmem

import (
	"github.com/theliuk/todoapp"
	"sync"
)

// UserService is an actual implementation of the todoapp.UserService interface
// which works with in memory storage.
type UserService struct {
	mtx   sync.RWMutex
	Todos map[string]*todoapp.Todo
}

func (usrv *UserService) Todo(ID string) (*todoapp.Todo, error) {
	usrv.mtx.RLock()
	defer usrv.mtx.RUnlock()

	todo, ok := usrv.Todos[ID]

	if !ok {
		return nil, &errTodoNotFound{ID}
	}

	return todo, nil
}
