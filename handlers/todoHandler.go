package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tarkanaciksoz/todo-list/helpers"
	"github.com/tarkanaciksoz/todo-list/models"
)

type TodoHandler struct {
	l *log.Logger
}

func NewTodoHandler(l *log.Logger) *TodoHandler {
	return &TodoHandler{l}
}

func (uHandler *TodoHandler) GetTodos(rw http.ResponseWriter, _ *http.Request) {
	uHandler.l.Println("Handle GetTodos method")

	todos := models.GetTodos()

	err := todos.ToJSON(rw, "Todos listed successfully")
	if err != nil {
		uHandler.l.Println("Error while handling GetTodos method: " + err.Error())
		resp := helpers.SetAndGetResponse(false, err.Error(), nil, http.StatusBadRequest)
		http.Error(rw, resp, http.StatusOK)
		return
	}
	uHandler.l.Println("GetTodos method successfully handled")
}

func (uHandler *TodoHandler) AddTodo(rw http.ResponseWriter, r *http.Request) {
	uHandler.l.Println("Handle AddTodo method")

	newTodo := models.Todo{}
	err := newTodo.FromJSON(r.Body)
	if err != nil {
		uHandler.l.Println("Error while handling AddTodo method: " + err.Error())
		resp := helpers.SetAndGetResponse(false, "Invalid JSON Data", nil, http.StatusBadRequest)
		http.Error(rw, resp, http.StatusOK)
		return
	}

	todo := models.AddTodo(&newTodo)

	err = todo.ToJSON(rw, "Todo created successfully")
	if err != nil {
		uHandler.l.Println("Error while handling AddTodo method: " + err.Error())
		resp := helpers.SetAndGetResponse(false, err.Error(), nil, http.StatusBadRequest)
		http.Error(rw, resp, http.StatusBadRequest)
		return
	}
	uHandler.l.Println("AddTodo method successfully handled")
}

func (uHandler *TodoHandler) MarkTodo(rw http.ResponseWriter, r *http.Request) {
	uHandler.l.Println("Handle MarkTodo method")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		uHandler.l.Println("Error while handling MarkTodo method: " + err.Error())
		resp := helpers.SetAndGetResponse(false, "Unable to convert id "+vars["id"]+": "+err.Error(), nil, http.StatusBadRequest)
		http.Error(rw, resp, http.StatusOK)
		return
	}

	todo := models.Todo{
		Id: id,
	}

	err = todo.UpdateTodo()
	if err != nil {
		uHandler.l.Println("Error while handling MarkTodo method: " + err.Error())
		resp := helpers.SetAndGetResponse(false, err.Error(), nil, http.StatusBadRequest)
		http.Error(rw, resp, http.StatusOK)
		return
	}

	response := helpers.SetAndGetResponse(true, "Todo with id:"+strconv.Itoa(todo.Id)+" successfully marked", nil, http.StatusOK)
	fmt.Fprintln(rw, response)
	uHandler.l.Println("MarkTodo method successfully handled")
}

func (uHandler *TodoHandler) DeleteAllTodos(rw http.ResponseWriter, _ *http.Request) {
	uHandler.l.Println("Handle DeleteAllTodos method")

	models.DeleteAllTodos()

	response := helpers.SetAndGetResponse(true, "All todos successfully deleted", nil, http.StatusOK)
	fmt.Fprintln(rw, response)
	uHandler.l.Println("DeleteAllTodos method successfully handled")
}

func (uHandler *TodoHandler) DeleteTodo(rw http.ResponseWriter, r *http.Request) {
	uHandler.l.Println("Handle DeleteTodo method")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		uHandler.l.Println("Error while handling DeleteTodo method: " + err.Error())
		resp := helpers.SetAndGetResponse(false, "Unable to convert id "+vars["id"]+": "+err.Error(), nil, http.StatusBadRequest)
		http.Error(rw, resp, http.StatusOK)
		return
	}

	todo := models.Todo{
		Id: id,
	}

	err = todo.DeleteTodo()
	if err != nil {
		uHandler.l.Println("Error while handling DeleteTodo method: " + err.Error())
		resp := helpers.SetAndGetResponse(false, err.Error(), nil, http.StatusBadRequest)
		http.Error(rw, resp, http.StatusOK)
		return
	}

	response := helpers.SetAndGetResponse(true, "Todo with id:"+strconv.Itoa(todo.Id)+" successfully deleted", nil, http.StatusOK)
	fmt.Fprintln(rw, response)

	uHandler.l.Println("DeleteTodo method successfully handled")
}
