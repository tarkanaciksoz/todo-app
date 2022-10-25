package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/tarkanaciksoz/todo-list/helpers"
	"gorm.io/gorm"
)

type Todo struct {
	Id     int    `json:"id" gorm:"primary_key"`
	Value  string `json:"value"`
	Marked int    `json:"marked"`
}
type Todos []Todo

func (todo *Todo) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(todo)
}

func (todo *Todos) FromJSON(r io.Reader) error {
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

func (todos Todos) ToJSON(w io.Writer, message string) error {
	response := helpers.SetAndGetResponse(true, message, todos, http.StatusOK)
	_, err := fmt.Fprintln(w, response)
	if err != nil {
		err = fmt.Errorf("print problem: %s\n", err.Error())
		return err
	}
	return nil
}

func GetTodos(db *gorm.DB) (todos Todos, err error) {
	if err = db.Find(&todos).Error; err != nil {
		return todos, err
	}

	return todos, err
}

func AddTodo(todo *Todo, db *gorm.DB) (*Todo, error) {
	result := db.Create(&todo)
	if result.Error != nil {
		err := fmt.Errorf("%s : %w", "Todo couldn't created", result.Error)
		return nil, err
	}

	return todo, nil
}

func DeleteAllTodos(db *gorm.DB) error {
	result := db.Exec("TRUNCATE TABLE todos")
	if result.Error != nil {
		err := fmt.Errorf("%s : %w", "Todos couldn't deleted", result.Error)
		return err
	}

	return nil
}

func (todo *Todo) UpdateTodo(db *gorm.DB) error {
	result := db.First(&todo)
	if result.Error != nil {
		err := fmt.Errorf("%s : %w", "Todo with id:"+strconv.Itoa(todo.Id)+" couldn't found to mark", result.Error)
		return err
	}

	if todo.Marked == 1 {
		todo.Marked = 0
	} else {
		todo.Marked = 1
	}

	result = db.Save(&todo)
	if result.Error != nil {
		err := fmt.Errorf("%s : %w", "Todo with id:"+strconv.Itoa(todo.Id)+" couldn't updated", result.Error)
		return err
	}

	return nil
}

func (todo *Todo) DeleteTodo(db *gorm.DB) error {
	result := db.Delete(&todo)
	if result.Error != nil {
		err := fmt.Errorf("%s : %w", "Todo with id : "+strconv.Itoa(todo.Id)+" couldn't deleted", result.Error)
		return err
	}

	return nil
}
