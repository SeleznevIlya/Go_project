package main

import (
	"log"

	"github.com/SeleznevIlya/Go_project"
	"github.com/SeleznevIlya/Go_project/pkg/handler"
	"github.com/SeleznevIlya/Go_project/pkg/repository"
	"github.com/SeleznevIlya/Go_project/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(Go_project.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
