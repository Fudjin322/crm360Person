package main

import (
	_ "github.com/lib/pq"
	"log"
	"sqlBoiler/api"
)

func main() {
	server := api.NewServer()
	err := server.Start()
	if err != nil {
		log.Fatal("Невозможно запустить сервер ", err)
	}

}
