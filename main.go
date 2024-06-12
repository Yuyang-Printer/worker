package main

import (
	"log"
	"worker/api"
)

func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start()
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
