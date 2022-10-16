package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tarkanaciksoz/todo-list/helpers"
)

func MethodNotFoundHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		resp := helpers.SetAndGetResponse(false, "Method Not Found", nil, http.StatusNotFound)
		http.Error(rw, resp, http.StatusNotFound)
	})
}

func MethodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		resp := helpers.SetAndGetResponse(false, "Method Not Allowed", nil, http.StatusBadRequest)
		http.Error(rw, resp, http.StatusBadRequest)
	})
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-type", "application/json; charset=UTF-8")
		rw.Header().Set("Access-Control-Allow-Origin", "*")

		ctx := r.Context()
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}

func ApplicationRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				_, err := fmt.Fprintln(os.Stderr, "Recovered from application error occurred")
				if err != nil {
					resp := helpers.SetAndGetResponse(false, "Internal Server Error at Application Recovery", nil, http.StatusInternalServerError)
					http.Error(rw, resp, http.StatusInternalServerError)
					return
				}
				_, _ = fmt.Fprintln(os.Stderr, err)

				resp := helpers.SetAndGetResponse(false, "Internal Server Error", nil, http.StatusInternalServerError)
				http.Error(rw, resp, http.StatusInternalServerError)
				return
			}
		}()

		ctx := r.Context()
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
