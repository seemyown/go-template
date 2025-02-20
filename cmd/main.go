package main

import (
	"go-fiber-template/internal/server"
	"go-fiber-template/pkg/logging"
)

var appLogger = logging.AppLogger

func main() {
	app, err := server.New()
	if err != nil {
		appLogger.Fatal().Err(err).Msg("Ошибка инициализации приложения")
	}
	appLogger.Info().Msg("Приложение инициализировано")
	appLogger.Fatal().Err(app.Listen(":8080")).Msg("Application terminated")
}
