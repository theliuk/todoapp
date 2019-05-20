package inmem

import (
	"github.com/theliuk/todoapp"
	"reflect"
	"testing"
)

func TestTodoService(t *testing.T) {
	t.Run(`Get`, func(t *testing.T) {
		t.Parallel()

		const existingID = "a"
		const unexistingID = "1"

		todoAssociatedToExistingID := todoapp.Todo{
			Description: "something to do",
			IsDone:      false,
		}

		todoService := &todoService{
			todos: map[string]todoapp.Todo{
				existingID: todoAssociatedToExistingID,
			},
		}

		t.Run(`returns the Todo associated to the provided ID`, func(t *testing.T) {
			want := todoAssociatedToExistingID
			got, err := todoService.Get(existingID)

			if err != nil {
				t.Fatalf(`Get(%q) failed because: %v`, existingID, err)
			}

			if !reflect.DeepEqual(got, want) {
				t.Fatalf(`Get(%q) = %v | want %v`, existingID, got, want)
			}
		})

		t.Run(`returns an ErrTodoNotFound if the ID is NOT found`, func(t *testing.T) {
			want := todoapp.Todo{}
			got, err := todoService.Get(unexistingID)

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

		todoService := &todoService{
			todos:  map[string]todoapp.Todo{},
			uIDGen: &uniqueIDGeneratorMock{newUniqueID},
		}

		todoToCreate := todoapp.Todo{
			Description: "something to do",
			IsDone:      false,
		}

		t.Run(`saves the Todo passed as parameter and returns its ID`, func(t *testing.T) {
			want := newUniqueID
			got, err := todoService.Create(todoToCreate)

			if err != nil {
				t.Fatalf(`Create(%v) failed because: %v`, todoToCreate, err)
			}

			if got != want {
				t.Fatalf(`Create(%v) = %q | want %q`, todoToCreate, got, want)
			}

			if _, isTodoPresent := todoService.todos[want]; !isTodoPresent {
				t.Fatalf(`Create(%v) does NOT create a new Todo`, todoToCreate)
			}
		})
	})

	t.Run(`Delete`, func(t *testing.T) {
		t.Parallel()

		t.Run(`removes the Todo associated to the ID`, func(t *testing.T) {
			const existingID = "A"

			todoService := &todoService{
				todos: map[string]todoapp.Todo{
					existingID: todoapp.Todo{},
				},
			}

			err := todoService.Delete(existingID)

			if err != nil {
				t.Fatalf(`Delete(%q) fails because: %v`, existingID, err)
			}

			if _, isTodoStillPresent := todoService.todos[existingID]; isTodoStillPresent {
				t.Fatalf(`Delete(%q) does NOT remove Todo associated with ID %q`, existingID, existingID)
			}
		})

		t.Run(`returns an ErrTodoNotFound if the ID is not found`, func(t *testing.T) {
			const nonExistingID = "B"

			todoService := &todoService{
				todos: map[string]todoapp.Todo{},
			}

			err := todoService.Delete(nonExistingID)

			if ok, _ := todoapp.IsErrTodoNotFound(err); !ok {
				t.Fatalf(`Delete(%q) does NOT return an ErrTodoNotFound`, nonExistingID)
			}
		})
	})

	t.Run(`Update`, func(t *testing.T) {
		t.Parallel()
		//No TDD? Oops :)
	})
}

type uniqueIDGeneratorMock struct {
	FakeID string
}

func (uIDGenMock *uniqueIDGeneratorMock) GenerateUniqueID(...interface{}) string {
	return uIDGenMock.FakeID
}
