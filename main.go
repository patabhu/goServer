package main

import (
	"goServer/adapters"
	"goServer/app/controller"
	"goServer/app/repository"
	"goServer/app/usecase"
	"goServer/config"
	"log"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	conf := config.GetConfig()
	dbconn := adapters.NewDBAdapterRepository(conf)

	toDorp := repository.NewToDoRepo(dbconn)
	toDouc := usecase.NewToDoUsecase(toDorp)
	controller.NewTodo(serveMux, toDouc)

	log.Println("listening on port:", conf.HttpConfig.HostPort)
	if err := http.ListenAndServe(conf.HttpConfig.HostPort, serveMux); err != nil {
		log.Println("server server", err)
	}
}
