package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/tarkanaciksoz/todo-list/helpers"
	"golang.org/x/exp/slices"
)

type Todo struct {
	Id     int    `json:"id" gorm:"primary_key"`
	Value  string `json:"value"`
	Marked int    `json:"marked"`
}
type Todos []*Todo

var TodoList = Todos{}

func (todo *Todo) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(todo)
}

func (todo Todo) ToJSON(w io.Writer, message string) error {
	response := helpers.SetAndGetResponse(true, message, todo, http.StatusOK)
	_, err := fmt.Fprintln(w, response)
	if err != nil {
		err = fmt.Errorf("print problem: %s\n", err.Error())
		return err
	}
	return nil
}

func (todo Todos) ToJSON(w io.Writer, message string) error {
	response := helpers.SetAndGetResponse(true, message, todo, http.StatusOK)
	_, err := fmt.Fprintln(w, response)
	if err != nil {
		err = fmt.Errorf("print problem: %s\n", err.Error())
		return err
	}
	return nil
}

func GetTodos() Todos {
	return TodoList
}

func AddTodo(todo *Todo) *Todo {
	todoIndex := slices.IndexFunc(TodoList, func(t *Todo) bool {
		return t.Id == todo.Id
	})

	if todoIndex >= 0 || todo.Id == 0 {
		todo.Id = len(TodoList) + 1
	}

	TodoList = append(TodoList, todo)
	return todo
}

func (todo *Todo) UpdateTodo() error {
	todoIndex := slices.IndexFunc(TodoList, func(t *Todo) bool {
		return t.Id == todo.Id
	})

	if todoIndex < 0 {
		return errors.New("Todo with id:" + strconv.Itoa(todo.Id) + " Couldn't find to update")
	}

	todo = TodoList[todoIndex]

	if todo.Marked == 1 {
		todo.Marked = 0
	} else {
		todo.Marked = 1
	}

	TodoList[todoIndex] = todo
	return nil
}

func (todo *Todo) DeleteTodo() error {
	todoIndex := slices.IndexFunc(TodoList, func(t *Todo) bool {
		return t.Id == todo.Id
	})

	if todoIndex < 0 {
		return errors.New("Todo with id:" + strconv.Itoa(todo.Id) + " Couldn't find to delete")
	}

	TodoList[todoIndex] = TodoList[len(TodoList)-1]
	TodoList[len(TodoList)-1] = &Todo{}
	TodoList = TodoList[:len(TodoList)-1]

	return nil
}

func DeleteAllTodos() {
	TodoList = Todos{}
}
