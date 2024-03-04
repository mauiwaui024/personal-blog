package main

import (
	"day_06"
	"day_06/internal/credentials"
	"day_06/internal/handler"
	"day_06/internal/repository"
	"day_06/internal/service"
	"log"
)

func main() {
	//генерация логотипа
	// logogenerate.LogoGenerator()
	//сначала создаем репозиторий, потом сервис, который зависит от репозитория,
	//а потом хэндлер, который зависит от сервисов
	dbConfig, err := credentials.GetDBConfig("./internal/credentials/credentials.txt")
	if err != nil {
		log.Fatal(err)
	}
	db, err := repository.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handler := handler.NewHandler(services)
	srv := new(day_06.Server)

	if err := srv.Run("8888", handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
