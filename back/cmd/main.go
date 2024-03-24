package main

import (
	todo "github.com/romAYAYA/golang-todo"
	"github.com/romAYAYA/golang-todo/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("An error occurred: %s", err.Error())
	}
}
