package main

import (
	"api-gateway/api"
	"api-gateway/api/handler"
	"api-gateway/config"
	"api-gateway/service"
	"log"
)

func main() {
	cfg := config.Load()
	config.InitLogger()
	logger := config.Logger
	logger.Info("Starting the application...")

	srv := service.New()
	h := handler.NewHandler(srv.Auth, srv.Reservation, srv.Payment, logger)

	r := api.NewRouter(h)

	logger.Info("Server is running", "PORT", cfg.HTTP_PORT)
	log.Fatal(r.Run(cfg.HTTP_PORT))
}
