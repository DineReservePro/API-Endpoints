package handler

import (
	"api-gateway/generated/auth_service"
	"api-gateway/generated/payment_service"
	"api-gateway/generated/reservation_service"
)

type Handler struct {
	Auth auth_service.AuthServiceClient
	Reservation reservation_service.ReservationServiceClient
	Payment payment_service.PaymentServiceClient
}

func NewHandler (a auth_service.AuthServiceClient, r reservation_service.ReservationServiceClient, p payment_service.PaymentServiceClient) *Handler {
	return &Handler{
		Auth:  a,
		Reservation: r,
		Payment: p,
	}
}