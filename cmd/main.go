package main

import (
	"log"

	"github.com/ksilex/todo-app"
	"github.com/ksilex/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured %s", err.Error())
	}

}
