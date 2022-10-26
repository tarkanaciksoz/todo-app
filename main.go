package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/tarkanaciksoz/todo-list/router"
)

func main() {
	logger := log.New(os.Stdout, "api-todo-list: ", log.LstdFlags)
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		logger.Printf("You must declare APP_ENV before run")
		os.Exit(1)
	}

	err := godotenv.Load(".env" + "." + appEnv)
	if err != nil {
		logger.Printf("Error while Read .env file: %s\n", err.Error())
		os.Exit(1)
	}
	var bindAddress = os.Getenv("BIND_ADDRESS")

	s := http.Server{
		Addr:         bindAddress,
		Handler:      router.ApplicationRecovery(router.Middleware(router.Init(logger))),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		logger.Printf("Starting server on port %s\n", bindAddress)

		err := s.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = s.Shutdown(ctx)
	if err != nil {
		logger.Printf("Shutdown problem: %s\n", err.Error())
		return
	}
}
