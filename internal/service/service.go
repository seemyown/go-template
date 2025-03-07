package service

import (
	"go-fiber-template/pkg/logging"
)

var log = logging.New(logging.Config{
	FileName: "servise",
	Name:     "service",
})
