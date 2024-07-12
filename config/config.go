package config

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT        string
	AUTH_PORT        string
	PAYMENT_PORT     string
	RESERVATION_PORT string
	ACCESS_TOKEN     string
}

var Logger *slog.Logger

func InitLogger() {
	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	handler := slog.NewJSONHandler(logFile, nil)
	Logger = slog.New(handler)
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	cfg := Config{}
	cfg.HTTP_PORT = cast.ToString(Coalesce("HTTP_PORT", ":8080"))
	cfg.AUTH_PORT = cast.ToString(Coalesce("AUTH_PORT", ":50050"))
	cfg.RESERVATION_PORT = cast.ToString(Coalesce("RESERVATION_PORT", "50051"))
	cfg.PAYMENT_PORT = cast.ToString(Coalesce("PAYMENT_PORT", ":50052"))
	cfg.ACCESS_TOKEN = cast.ToString(Coalesce("ACCESS_TOKEN", "my_secret_key"))

	return cfg
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
