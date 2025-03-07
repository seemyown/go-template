package controller

import (
	"go-fiber-template/pkg/logging"
)

var log = logging.New(logging.Config{
	FileName: "controller",
	Name:     "controller",
})

// Модуль для интерфесов контроллеров
