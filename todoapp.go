package todoapp

//Todo represents something that has to be done and has a binary state.
type Todo struct {
	ID          string
	Description string
	IsDone      bool
}

//TodoService defines all the operations that are supported by a service which allows
//a client to deal with todos
type TodoService interface {
	Todo(ID string) (Todo, error)
	CreateTodo(t Todo) error
	DeleteTodo(t Todo) error
	UpdateTodo(t Todo) error
}

type ErrTodoNotFound interface {
	ErrTodoNotFound() (bool, string)
}

func IsErrTodoNotFound(err error) (bool, string) {
	if e, ok := err.(ErrTodoNotFound); ok {
		return e.ErrTodoNotFound()
	}

	return false, ""
}
