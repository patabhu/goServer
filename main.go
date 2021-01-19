package main

import (
	"goServer/config"
	"log"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	conf := config.GetConfig()

	log.Println("listening on port:", conf.HttpConfig.HostPort)
	if err := http.ListenAndServe(conf.HttpConfig.HostPort, serveMux); err != nil {
		log.Println("server server", err)
	}
}
