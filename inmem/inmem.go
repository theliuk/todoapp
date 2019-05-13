package inmem

import (
	"github.com/theliuk/todoapp"
	"sync"
)

// UserService is an actual implementation of the todoapp.UserService interface
// which works with in memory storage.
type UserService struct {
	mtx               sync.RWMutex
	UniqueIDGenerator func(...interface{}) string
	Todos             map[string]todoapp.Todo
}

//Get takes an ID and returns the corresponding Todo, if any is associated to
//the specific ID; an error is returned otherwise.
func (usrv *UserService) Get(ID string) (todoapp.Todo, error) {
	usrv.mtx.RLock()
	defer usrv.mtx.RUnlock()

	todo, ok := usrv.Todos[ID]

	if !ok {
		return todoapp.Todo{}, &errTodoNotFound{ID}
	}

	return todo, nil
}

// Create takes a Todo with all the information available, except for the ID,
// and returns the ID of the created Todo
func (usrv *UserService) Create(t todoapp.Todo) (string, error) {
	usrv.mtx.Lock()
	defer usrv.mtx.Unlock()

	newUniqueID := usrv.UniqueIDGenerator()

	usrv.Todos[newUniqueID] = t
	return newUniqueID, nil
}

//Delete removes the Todo associated to the ID if any; otherwise it returns an error
func (usrv *UserService) Delete(ID string) error {
	usrv.mtx.Lock()
	defer usrv.mtx.Unlock()

	if _, todoIsPresent := usrv.Todos[ID]; !todoIsPresent {
		return &errTodoNotFound{ID}
	}

	delete(usrv.Todos, ID)
	return nil
}

// Update updates the Todo associated to an ID; if upsert is set, the association is created
// if does NOT already exist
func (usrv *UserService) Update(ID string, todo todoapp.Todo, upsert bool) error {
	usrv.mtx.Lock()
	defer usrv.mtx.Unlock()

	if upsert {
		usrv.Todos[ID] = todo
		return nil
	}

	if _, todoAlreadyPresent := usrv.Todos[ID]; !todoAlreadyPresent {
		return &errTodoNotFound{ID}
	}

	usrv.Todos[ID] = todo
	return nil
}
