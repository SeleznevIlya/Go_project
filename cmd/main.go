package main

import (
	"log"

	"github.com/SeleznevIlya/Go_project"
)

func main() {
	srv := new(Go_project.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
