package main

import (
	"api-gateway/api"
	"api-gateway/api/handler"
	"api-gateway/service"
	"log"
)

func main() {
	service := service.New()

	handler := handler.NewHandler(service.Auth, service.Reservation, service.Payment)

	router := api.NewRouter(handler)

	log.Fatal(router.Run())
}
