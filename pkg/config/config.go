package config

import (
	"fmt"
	"go-fiber-template/pkg/logging"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var log = logging.ToolLogger

type Config struct {
	DBHost        string
	DBPort        string
	DBName        string
	DBUser        string
	DBPass        string
	SecretKey     string
	TokenIssuer   string
	NatsServer    string
	NatsPort      string
	RedisHost     string
	RedisPort     string
	RedisCacheDB  int
	RedisStoreDB  int
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msg("Error loading .env file")
	}
	return &Config{
		DBHost:        getEnvString("DB_HOST", "localhost"),
		DBPort:        getEnvString("DB_PORT", "5432"),
		DBName:        getEnvString("DB_NAME", "postgres"),
		DBUser:        getEnvString("DB_USER", "postgres"),
		DBPass:        getEnvString("DB_PASS", "root"),
		NatsServer:    getEnvString("NATS_SERVER", "localhost"),
		NatsPort:      getEnvString("NATS_PORT", "4222"),
		RedisHost:     getEnvString("REDIS_HOST", "localhost"),
		RedisPort:     getEnvString("REDIS_PORT", "6379"),
		RedisCacheDB:  getEnvInt("REDIS_CACHE_DB", 0),
		RedisStoreDB:  getEnvInt("REDIS_STORE_DB", 1),
		SecretKey:     getEnvString("SECRET_KEY", "secretKey"),
		TokenIssuer:   getEnvString("TOKEN_ISSUER", "jwt.issuer"),
	}
}

func getEnvString(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if value, err := strconv.Atoi(value); err == nil {
			return value
		}
	}
	return defaultValue
}

func (c *Config) DBConnString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBName, c.DBPass,
	)
}

func (c *Config) NatsConnString() string {
	return fmt.Sprintf(
		"nats://%s:%s",
		c.NatsServer, c.NatsPort,
	)
}

func (c *Config) RedisAddrString() string {
	return fmt.Sprintf(
		"%s:%s",
		c.RedisHost, c.RedisPort,
	)
}

func (c *Config) RedisCacheDBNumber() int {
	return c.RedisCacheDB
}

func (c *Config) RedisStoreDBNumber() int {
	return c.RedisStoreDB
}