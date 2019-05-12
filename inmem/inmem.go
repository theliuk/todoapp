package inmem

import (
	"github.com/theliuk/todoapp"
	"sync"
)

// UserService is an actual implementation of the todoapp.UserService interface
// which works with in memory storage.

type IDGenerator func(...interface{}) string
type UserService struct {
	mtx   sync.RWMutex
	IDGen IDGenerator
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

// CreateTodo takes a Todo with all the information available, except for the ID,
// and returns the ID of the created Todo
func (usrv *UserService) CreateTodo(t todoapp.Todo) (string, error) {
	usrv.mtx.Lock()
	defer usrv.mtx.Unlock()

	newID := usrv.IDGen()

	for !usrv.isIDUnique(newID) {
		newID = usrv.IDGen()
	}

	usrv.Todos[newID] = &t
	return newID, nil
}

func (usrv *UserService) isIDUnique(ID string) bool {
	_, isIDAlreadyAssociatedToSomeTodo := usrv.Todos[ID]
	return !isIDAlreadyAssociatedToSomeTodo
}
