package main

import (
	"goServer/adapters"
	"goServer/config"
	"log"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	conf := config.GetConfig()
	adapters.NewDBAdapterRepository(conf)

	log.Println("listening on port:", conf.HttpConfig.HostPort)
	if err := http.ListenAndServe(conf.HttpConfig.HostPort, serveMux); err != nil {
		log.Println("server server", err)
	}
}
