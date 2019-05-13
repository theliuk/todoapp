package inmem_test

import (
	"github.com/theliuk/todoapp"
	"github.com/theliuk/todoapp/inmem"
	"reflect"
	"testing"
)

func TestUserService(t *testing.T) {
	t.Run(`Get`, func(t *testing.T) {
		t.Parallel()

		const existingID = "a"
		const unexistingID = "1"

		todoAssociatedToExistingID := todoapp.Todo{
			Description: "something to do",
			IsDone:      false,
		}

		userService := &inmem.UserService{
			Todos: map[string]todoapp.Todo{
				existingID: todoAssociatedToExistingID,
			},
		}

		t.Run(`returns the Todo associated to the provided ID`, func(t *testing.T) {
			want := todoAssociatedToExistingID
			got, err := userService.Get(existingID)

			if err != nil {
				t.Fatalf(`Get(%q) failed because: %v`, existingID, err)
			}

			if !reflect.DeepEqual(got, want) {
				t.Fatalf(`Get(%q) = %v | want %v`, existingID, got, want)
			}
		})

		t.Run(`returns an ErrTodoNotFound if the ID is NOT found`, func(t *testing.T) {
			want := todoapp.Todo{}
			got, err := userService.Get(unexistingID)

			if got != want {
				t.Fatalf(`Get(%q) = %v | want %v`, unexistingID, got, want)
			}

			if ok, _ := todoapp.IsErrTodoNotFound(err); !ok {
				t.Fatalf(`Get(%q) does NOT return an ErrTodoNotFound`, unexistingID)
			}
		})
	})

	t.Run(`Create`, func(t *testing.T) {
		t.Parallel()
		const newUniqueID = "A"

		uniqueIDGenerator := func(...interface{}) string {
			return newUniqueID
		}

		todoToCreate := todoapp.Todo{
			Description: "something to do",
			IsDone:      false,
		}

		userService := &inmem.UserService{
			Todos:             map[string]todoapp.Todo{},
			UniqueIDGenerator: uniqueIDGenerator,
		}

		t.Run(`saves the Todo passed as parameter and returns its ID`, func(t *testing.T) {
			want := newUniqueID
			got, err := userService.Create(todoToCreate)

			if err != nil {
				t.Fatalf(`Create(%v) failed because: %v`, todoToCreate, err)
			}

			if got != want {
				t.Fatalf(`Create(%v) = %q | want %q`, todoToCreate, got, want)
			}

			if _, ok := userService.Todos[want]; !ok {
				t.Fatalf(`Create(%v) does NOT create a new Todo`, todoToCreate)
			}
		})
	})

	t.Run(`DeleteTodo`, func(t *testing.T) {
		t.Parallel()

		t.Run(`removes the Todo associated to the ID`, func(t *testing.T) {
			const existingId = "A"

			userService := &inmem.UserService{
				Todos: map[string]todoapp.Todo{
					existingId: todoapp.Todo{},
				},
			}

			err := userService.Delete(existingId)

			if err != nil {
				t.Fatalf(`Delete(%q) fails because: %v`, existingId, err)
			}

			if _, ok := userService.Todos[existingId]; ok {
				t.Fatalf(`Delete(%q) does NOT remove Todo associated with ID %q`, existingId, existingId)
			}
		})

		t.Run(`returns an ErrTodoNotFound if the ID is not found`, func(t *testing.T) {
			const nonExistingID = "B"

			userService := &inmem.UserService{
				Todos: map[string]todoapp.Todo{},
			}

			err := userService.Delete(nonExistingID)

			if ok, _ := todoapp.IsErrTodoNotFound(err); !ok {
				t.Fatalf(`Delete(%q) does NOT return an ErrTodoNotFound`, nonExistingID)
			}
		})

	})

	t.Run(`UpdateTodo`, func(t *testing.T) {
		t.Parallel()
		//No TDD? Oops :)
	})
}
