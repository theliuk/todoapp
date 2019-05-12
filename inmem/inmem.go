package inmem

import (
	"github.com/theliuk/todoapp"
	"sync"
)

// UserService is an actual implementation of the todoapp.UserService interface
// which works with in memory storage.

type UniqueIDGenerator func(...interface{}) string
type UserService struct {
	mtx               sync.RWMutex
	UniqueIDGenerator func(...interface{}) string
	Todos             map[string]*todoapp.Todo
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

// CreateTodo takes a Todo with all the information available, except for the ID,
// and returns the ID of the created Todo
func (usrv *UserService) CreateTodo(t todoapp.Todo) (string, error) {
	usrv.mtx.Lock()
	defer usrv.mtx.Unlock()

	newUniqueID := usrv.UniqueIDGenerator()

	usrv.Todos[newUniqueID] = &t
	return newUniqueID, nil
}
