package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"go-fiber-template/internal/infrastructure/rest/middleware"
	"go-fiber-template/pkg/config"
	"go-fiber-template/pkg/logging"
)

// Настройка сервера

var log = logging.New(logging.Config{
	FileName: "server",
	Name:     "server",
})

func newApp(cfg *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		Views: html.New("./templates", ".html"),
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))
	app.Use(compress.New())
	app.Use(middleware.LoggingMiddleware)
	app.Use(middleware.JWTMiddleware(cfg))

	app.Static("/docs", "./docs")

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Render("redoc", fiber.Map{
			"Title": "APP",
		})
	})

	app.Get("/docs/sandbox", func(c *fiber.Ctx) error {
		return c.Render("swagger-ui", fiber.Map{
			"Title": "APP",
		})
	})

	return app
}

func New() (*fiber.App, error) {
	cfg := config.NewConfig()
	db, err := sqlx.Connect("postgres", cfg.DBConnString())
	if err != nil {
		log.Error(err, "Ошибка подключения к БД")
		return nil, err
	}
	log.Info().Msg("Подключение к БД установлено")
	natsConn, err := nats.Connect(cfg.NatsConnString())
	if err != nil {
		log.Error(err, "Ошибка подключения к NATS")
		return nil, err
	}
	log.Info("Подключение к NATS установлено")
	redisStoreClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddrString(),
		Password: "",
		DB:       cfg.RedisStoreDB,
	})
	log.Info("Клиент для Redis создан")
	app := newApp(cfg)
	// Инициализация адапторов
	return app, nil
}
