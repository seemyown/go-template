package logging

// Пакет для настройки логирования

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

func createLogger(logFile string) zerolog.Logger {
	if err := os.MkdirAll(filepath.Dir(logFile), os.ModePerm); err != nil {
		panic(err)
	}

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	consoleWriter := zerolog.ConsoleWriter{
		Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: true,
	}
	consoleWriter.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	consoleWriter.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	consoleWriter.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	consoleWriter.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	multiWriter := zerolog.MultiLevelWriter(consoleWriter, file)

	return zerolog.New(multiWriter).With().Timestamp().Logger()
}

var (
	AppLogger = createLogger("logs/app.log")
	MiddlewareLogger = createLogger("logs/middleware.log")
	ControllerLogger = createLogger("logs/controller.log")
	BrokerLogger = createLogger("logs/broker.log")
	ServiceLogger = createLogger("logs/service.log")
	KeyValueLogger = createLogger("logs/key_value.log")
	ServerLogger = createLogger("logs/server.log")
	ToolLogger = createLogger("logs/tool.log")
	RelationLogger = createLogger("logs/relation.log")
)
