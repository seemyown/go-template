package main

import (
	"go-fiber-template/internal/server"
	"go-fiber-template/pkg/logging"
)

var log = logging.New(logging.Config{
	FileName: "servise",
	Name:     "service",
})

func main() {
	app, err := server.New()
	if err != nil {
		log.Fatal(err, "Ошибка инициализации приложения")
	}
	log.Info("Приложение инициализировано")
	log.Fatal(app.Listen(":8080"), "Application terminated")
}
