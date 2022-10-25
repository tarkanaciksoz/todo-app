package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tarkanaciksoz/todo-list/handlers"
	"gorm.io/gorm"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func Init(logger *log.Logger, DB *gorm.DB) *mux.Router {
	uHandler := handlers.NewTodoHandler(logger, DB)

	mappedRoutes := make(map[string]Routes)
	mappedRoutes[http.MethodGet] = Routes{
		Route{
			Name:        "GET TODO LIST",
			Method:      http.MethodGet,
			Pattern:     "/todo/getTodos",
			HandlerFunc: uHandler.GetTodos,
		},
	}

	mappedRoutes[http.MethodPost] = Routes{
		Route{
			Name:        "CREATE TODO",
			Method:      http.MethodPost,
			Pattern:     "/todo/createTodo",
			HandlerFunc: uHandler.AddTodo,
		},
		Route{
			Name:        "MARK TODO",
			Method:      http.MethodPost,
			Pattern:     "/todo/markTodo/{id:[0-9]+}",
			HandlerFunc: uHandler.MarkTodo,
		},
		Route{
			Name:        "DELETE ALL TODOS",
			Method:      http.MethodPost,
			Pattern:     "/todo/deleteAllTodos",
			HandlerFunc: uHandler.DeleteAllTodos,
		},
		Route{
			Name:        "DELETE TODO",
			Method:      http.MethodPost,
			Pattern:     "/todo/deleteTodo/{id:[0-9]+}",
			HandlerFunc: uHandler.DeleteTodo,
		},
	}

	/*mappedRoutes[http.MethodPut] = Routes{

	}

	mappedRoutes[http.MethodDelete] = Routes{

	}*/

	router := mux.NewRouter()
	router.NotFoundHandler = MethodNotFoundHandler()
	router.MethodNotAllowedHandler = MethodNotAllowedHandler()

	for method, routes := range mappedRoutes {
		methodRout := router.Methods(method).Subrouter()
		for _, route := range routes {
			methodRout.HandleFunc(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}
