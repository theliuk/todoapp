package inmem_test

import (
	"github.com/theliuk/todoapp"
	"github.com/theliuk/todoapp/inmem"
	"reflect"
	"testing"
)

func TestUserService(t *testing.T) {
	t.Run(`Todo`, func(t *testing.T) {
		t.Parallel()

		const existingID = "a"
		const unexistingID = "1"

		todoAssociatedToExistingID := &todoapp.Todo{
			Description: "something to do",
			IsDone:      false,
		}

		userService := &inmem.UserService{
			Todos: map[string]*todoapp.Todo{
				existingID: todoAssociatedToExistingID,
			},
		}

		t.Run(`returns the Todo associated to the provided ID`, func(t *testing.T) {
			want := todoAssociatedToExistingID
			got, err := userService.Todo(existingID)

			if err != nil {
				t.Fatalf(`Todo(%q) failed because: %v`, existingID, err)
			}

			if !reflect.DeepEqual(*got, *want) {
				t.Fatalf(`Todo(%q) = %v | want %v`, existingID, *got, *want)
			}
		})

		t.Run(`returns an ErrTodoNotFound if the ID is not found`, func(t *testing.T) {
			got, err := userService.Todo(unexistingID)

			if got != nil {
				t.Fatalf(`Todo(%q) = %v | want %v`, unexistingID, got, nil)
			}

			if ok, _ := todoapp.IsErrTodoNotFound(err); !ok {
				t.Fatalf(`Todo(%q) does not return an ErrTodoNotFound`, unexistingID)
			}
		})
	})

	t.Run(`CreateTodo`, func(t *testing.T) {
		t.Parallel()
		const newUnique = "A"

		uniqueIDGenerator := func(...interface{}) string {
			return newUnique
		}

		todoToCreate := todoapp.Todo{
			Description: "something to do",
			IsDone:      false,
		}

		userService := &inmem.UserService{
			Todos:             map[string]*todoapp.Todo{},
			UniqueIDGenerator: uniqueIDGenerator,
		}

		t.Run(`saves the Todo passed as paremeter and returns its ID`, func(t *testing.T) {
			want := newUnique
			got, err := userService.CreateTodo(todoToCreate)

			if err != nil {
				t.Fatalf(`CreateTodo(%v) failed because: %v`, todoToCreate, err)
			}

			if got != want {
				t.Fatalf(`CreateTodo(%v) = %q | want %q`, todoToCreate, got, want)
			}
		})

	})

	t.Run(`DeleteTodo`, func(t *testing.T) {
		t.Parallel()
	})

	t.Run(`UpdateTodo`, func(t *testing.T) {
		t.Parallel()
	})
}
