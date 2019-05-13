package todoapp

//Todo represents something that has to be done and has a binary state.
type Todo struct {
	Description string
	IsDone      bool
}

//TodoService defines all the operations that are supported by a service which allows
//a client to deal with todos
type TodoService interface {
	Get(ID string) (Todo, error)
	Create(t Todo) (string, error)
	Delete(ID string) error
	Update(ID string, t Todo, upsert bool) error
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
