package tests

import (
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/tarkanaciksoz/todo-list/configs/database"
	"github.com/tarkanaciksoz/todo-list/configs/router"
	"github.com/tarkanaciksoz/todo-list/helpers"
	"github.com/tarkanaciksoz/todo-list/models"
)

type Test struct {
	TestName             string
	TestRequest          Request
	ExpectedTestResponse string
	ExpectedTestError    error
}

type Tests []Test

type Request struct {
	Method string
	Url    string
	Body   string
}

func TestTodoHandler(t *testing.T) {
	logger := log.New(os.Stdout, "api-todo-list: ", log.LstdFlags)
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		logger.Printf("You must declare APP_ENV before run")
		os.Exit(1)
	}

	err := godotenv.Load("../.env" + "." + appEnv)
	if err != nil {
		logger.Printf("Error while Read .env file: %s\n", err.Error())
		os.Exit(1)
	}

	DB, err := database.Init()
	if err != nil {
		logger.Printf("Error starting database connection: %s\n", err.Error())
		os.Exit(1)
	}

	for _, test := range getTests() {
		t.Run(test.TestName, func(t *testing.T) {
			responseRecorder := httptest.NewRecorder()

			handler := router.ApplicationRecovery(router.Middleware(router.Init(logger, DB)))

			request, _ := http.NewRequest(test.TestRequest.Method, test.TestRequest.Url, strings.NewReader(test.TestRequest.Body))

			handler.ServeHTTP(responseRecorder, request)

			response := strings.TrimSpace(responseRecorder.Body.String())

			if !reflect.DeepEqual(test.ExpectedTestResponse, response) {
				t.Errorf("%s failed. Expected: '%v' - Got : '%v'", test.TestName, test.ExpectedTestResponse, response)
			}

			if !errors.Is(test.ExpectedTestError, err) {
				t.Errorf("Expected Error failed. Expected: %v - Got : %v", test.ExpectedTestError, err)
			}
		})
	}
}

func getTests() Tests {
	tests := Tests{
		//GIVEN, WHEN, THEN - 1 *** START
		Test{
			TestName:             "GET TODOS-1",
			TestRequest:          Request{"GET", "/todo/getTodos", ""},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "Todos listed successfully", &models.Todos{}, http.StatusOK),
			ExpectedTestError:    nil,
		},
		Test{
			TestName:    "ADD TODO-1",
			TestRequest: Request{"POST", "/todo/createTodo", `{"value":"buy some milk"}`},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "Todo created successfully", &models.Todo{
				Id:    1,
				Value: "buy some milk",
			}, http.StatusOK),
			ExpectedTestError: nil,
		},
		Test{
			TestName:    "GET TODOS-2",
			TestRequest: Request{"GET", "/todo/getTodos", ""},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "Todos listed successfully", &models.Todos{
				{
					Id:     1,
					Value:  "buy some milk",
					Marked: 0,
				},
			}, http.StatusOK),
			ExpectedTestError: nil,
		},
		//		GIVEN, WHEN, THEN - 1 *** END

		//		GIVEN, WHEN, THEN - 2 *** START
		Test{
			TestName:    "ADD TODO-2",
			TestRequest: Request{"POST", "/todo/createTodo", `{"value":"enjoy the assignment"}`},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "Todo created successfully", &models.Todo{
				Id:    2,
				Value: "enjoy the assignment",
			}, http.StatusOK),
			ExpectedTestError: nil,
		},
		Test{
			TestName:    "GET TODOS-3",
			TestRequest: Request{"GET", "/todo/getTodos", ""},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "Todos listed successfully", &models.Todos{
				{
					Id:     1,
					Value:  "buy some milk",
					Marked: 0,
				},
				{
					Id:     2,
					Value:  "enjoy the assignment",
					Marked: 0,
				},
			}, http.StatusOK),
			ExpectedTestError: nil,
		},
		//		GIVEN, WHEN, THEN - 2 *** END

		//		GIVEN, WHEN, THEN - 3 *** START
		Test{
			TestName:             "MARK TODO-1",
			TestRequest:          Request{"POST", "/todo/markTodo/1", ""},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "Todo with id:1 successfully marked", nil, http.StatusOK),
			ExpectedTestError:    nil,
		},
		Test{
			TestName:    "GET TODOS-4",
			TestRequest: Request{"GET", "/todo/getTodos", ""},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "Todos listed successfully", &models.Todos{
				{
					Id:     1,
					Value:  "buy some milk",
					Marked: 1,
				},
				{
					Id:     2,
					Value:  "enjoy the assignment",
					Marked: 0,
				},
			}, http.StatusOK),
			ExpectedTestError: nil,
		},
		//		GIVEN, WHEN, THEN - 3 *** END

		//		GIVEN, WHEN, THEN - 4 *** START
		Test{
			TestName:             "MARK TODO-2",
			TestRequest:          Request{"POST", "/todo/markTodo/1", ""},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "Todo with id:1 successfully marked", nil, http.StatusOK),
			ExpectedTestError:    nil,
		},
		Test{
			TestName:    "GET TODOS-5",
			TestRequest: Request{"GET", "/todo/getTodos", ""},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "Todos listed successfully", &models.Todos{
				{
					Id:     1,
					Value:  "buy some milk",
					Marked: 0,
				},
				{
					Id:     2,
					Value:  "enjoy the assignment",
					Marked: 0,
				},
			}, http.StatusOK),
			ExpectedTestError: nil,
		},
		//		GIVEN, WHEN, THEN - 4 *** END

		//		EMPTY LIST
		Test{
			TestName:             "DELETE ALL TODOS-1",
			TestRequest:          Request{"GET", "/todo/deleteAllTodos", ""},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "All todos successfully deleted", nil, http.StatusOK),
			ExpectedTestError:    nil,
		},
		//		EMPTY LIST

		//		GIVEN, WHEN, THEN - 5 *** START
		Test{
			TestName:    "ADD TODO-3",
			TestRequest: Request{"POST", "/todo/createTodo", `{"value":"rest for a while"}`},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "Todo created successfully", &models.Todo{
				Id:    1,
				Value: "rest for a while",
			}, http.StatusOK),
			ExpectedTestError: nil,
		},
		Test{
			TestName:    "GET TODOS-6",
			TestRequest: Request{"GET", "/todo/getTodos", ""},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "Todos listed successfully", &models.Todos{
				{
					Id:     1,
					Value:  "rest for a while",
					Marked: 0,
				},
			}, http.StatusOK),
			ExpectedTestError: nil,
		},
		Test{
			TestName:             "DELETE TODO-1",
			TestRequest:          Request{"GET", "/todo/deleteTodo/1", ""},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "Todo with id:1 successfully deleted", nil, http.StatusOK),
			ExpectedTestError:    nil,
		},
		Test{
			TestName:             "GET TODOS-7",
			TestRequest:          Request{"GET", "/todo/getTodos", ""},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "Todos listed successfully", &models.Todos{}, http.StatusOK),
			ExpectedTestError:    nil,
		},
		//		GIVEN, WHEN, THEN - 5 *** END

		//		EMPTY LIST
		Test{
			TestName:             "DELETE ALL TODOS-2",
			TestRequest:          Request{"GET", "/todo/deleteAllTodos", ""},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "All todos successfully deleted", nil, http.StatusOK),
			ExpectedTestError:    nil,
		},
		//		EMPTY LIST

		//		GIVEN, WHEN, THEN - 6 *** START
		Test{
			TestName:    "ADD TODO-4",
			TestRequest: Request{"POST", "/todo/createTodo", `{"value":"rest for a while"}`},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "Todo created successfully", &models.Todo{
				Id:    1,
				Value: "rest for a while",
			}, http.StatusOK),
			ExpectedTestError: nil,
		},
		Test{
			TestName:    "ADD TODO-5",
			TestRequest: Request{"POST", "/todo/createTodo", `{"value":"drink water"}`},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "Todo created successfully", &models.Todo{
				Id:    2,
				Value: "drink water",
			}, http.StatusOK),
			ExpectedTestError: nil,
		},
		Test{
			TestName:    "GET TODOS-8",
			TestRequest: Request{"GET", "/todo/getTodos", ""},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "Todos listed successfully", &models.Todos{
				{
					Id:     1,
					Value:  "rest for a while",
					Marked: 0,
				},
				{
					Id:     2,
					Value:  "drink water",
					Marked: 0,
				},
			}, http.StatusOK),
			ExpectedTestError: nil,
		},
		Test{
			TestName:             "DELETE TODO-2",
			TestRequest:          Request{"GET", "/todo/deleteTodo/1", ""},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "Todo with id:1 successfully deleted", nil, http.StatusOK),
			ExpectedTestError:    nil,
		},
		Test{
			TestName:    "GET TODOS-9",
			TestRequest: Request{"GET", "/todo/getTodos", ""},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "Todos listed successfully", &models.Todos{
				{
					Id:     2,
					Value:  "drink water",
					Marked: 0,
				},
			}, http.StatusOK),
			ExpectedTestError: nil,
		},
		//		GIVEN, WHEN, THEN - 6 *** END

		//		EMPTY LIST
		Test{
			TestName:             "DELETE ALL TODOS-2",
			TestRequest:          Request{"GET", "/todo/deleteAllTodos", ""},
			ExpectedTestResponse: helpers.SetAndGetResponse(true, "All todos successfully deleted", nil, http.StatusOK),
			ExpectedTestError:    nil,
		},
		//		EMPTY LIST
	}

	return tests
}
