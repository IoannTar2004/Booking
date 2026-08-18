package main

import (
	"hotel/internal/handlers"
	"hotel/internal/server"
	"log"
)

func main() {
	handler := handlers.InitHandlers()
	srv := server.Server{}

	err := srv.Run(8080, handler)
	if err != nil {
		log.Fatal(err)
	}
}
