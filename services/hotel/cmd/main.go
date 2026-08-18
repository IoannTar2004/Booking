package main

import (
	"hotel/internal/handlers"
	"hotel/internal/infrastructure/config"
	"hotel/internal/infrastructure/database"
	"hotel/internal/infrastructure/server"
	"log"
	"os"
)

func main() {
	profile := ""
	if len(os.Args) > 1 {
		profile = os.Args[1]
	}

	conf, err := config.Load(profile)
	if err != nil {
		log.Fatal(err)
	}

	if err = database.MigrateUp("file://migrations", conf); err != nil {
		log.Fatal(err)
	}

	handler := handlers.InitHandlers()

	srv := server.Server{}
	if err = srv.Run(8080, handler); err != nil {
		log.Fatal(err)
	}
}
