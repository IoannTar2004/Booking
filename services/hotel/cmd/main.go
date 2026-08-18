package main

import (
	"hotel/internal/handlers"
	"hotel/internal/infrastructure/config"
	"hotel/internal/infrastructure/database"
	"hotel/internal/infrastructure/server"
	"log"
	"os"

	_ "hotel/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Booking project
// @version         1.0
// @description     Hotel microservice.

// @host      localhost:8080
// @BasePath  /api/v1
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

	db, err := database.NewPostgresDB(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	handler := handlers.Init(db)
	handler.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	srv := server.Server{}
	if err = srv.Run(8080, handler); err != nil {
		log.Fatal(err)
	}
}
