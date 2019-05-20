package inmem

import (
	"github.com/theliuk/todoapp"
	"strconv"
	"sync"
)

type UniqueIDGenerator interface {
	GenerateUniqueID(...interface{}) string
}

// IncrementalIDGenerator returns an incremental interger as string.
// Must be used in a safe way.
type IncrementalIDGenerator struct {
	counter int
}

//GenerateUniqueID returns an incremental integer as a string at each call.
func (incIDGen *IncrementalIDGenerator) GenerateUniqueID(...interface{}) (ID string) {
	ID = strconv.Itoa(incIDGen.counter)
	incIDGen.counter++
	return
}

type todoService struct {
	mtx    sync.RWMutex
	todos  map[string]todoapp.Todo
	uIDGen UniqueIDGenerator
}

//NewTodoService returns a new instance of an in-memory implementation of todoapp.TodoService
func NewTodoService(udIDGen UniqueIDGenerator) todoapp.TodoService {
	return &todoService{
		todos:  make(map[string]todoapp.Todo),
		uIDGen: udIDGen,
	}
}

//Get takes an ID and returns the corresponding Todo, if any is associated to
//the specific ID; an error is returned otherwise.
func (usrv *todoService) Get(ID string) (todoapp.Todo, error) {
	usrv.mtx.RLock()
	defer usrv.mtx.RUnlock()

	todo, ok := usrv.todos[ID]

	if !ok {
		return todoapp.Todo{}, &errTodoNotFound{ID}
	}

	return todo, nil
}

// Create takes a Todo with all the information available, except for the ID,
// and returns the ID of the created Todo
func (usrv *todoService) Create(t todoapp.Todo) (string, error) {
	usrv.mtx.Lock()
	defer usrv.mtx.Unlock()

	newUniqueID := usrv.uIDGen.GenerateUniqueID()

	usrv.todos[newUniqueID] = t
	return newUniqueID, nil
}

//Delete removes the Todo associated to the ID if any; otherwise it returns an error
func (usrv *todoService) Delete(ID string) error {
	usrv.mtx.Lock()
	defer usrv.mtx.Unlock()

	if _, todoIsPresent := usrv.todos[ID]; !todoIsPresent {
		return &errTodoNotFound{ID}
	}

	delete(usrv.todos, ID)
	return nil
}

// Update updates the Todo associated to an ID; if upsert is set, the association is created
// if does NOT already exist
func (usrv *todoService) Update(ID string, todo todoapp.Todo, upsert bool) error {
	usrv.mtx.Lock()
	defer usrv.mtx.Unlock()

	if upsert {
		usrv.todos[ID] = todo
		return nil
	}

	if _, todoAlreadyPresent := usrv.todos[ID]; !todoAlreadyPresent {
		return &errTodoNotFound{ID}
	}

	usrv.todos[ID] = todo
	return nil
}
