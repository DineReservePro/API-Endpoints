package service

import (
	"api-gateway/config"
	"api-gateway/generated/auth_service"
	"api-gateway/generated/payment_service"
	"api-gateway/generated/reservation_service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	Auth auth_service.AuthServiceClient
	Reservation reservation_service.ReservationServiceClient
	Payment payment_service.PaymentServiceClient
}

func New() *Service {
	cfg := config.Load()
	connAuth, err := grpc.NewClient("localhost"+cfg.AUTH_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("ERROR: ", err.Error())
	}

	authService := auth_service.NewAuthServiceClient(connAuth)

	reservationConn, err := grpc.NewClient("localhost"+cfg.RESERVATION_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Error: ", err.Error())
		return nil
	}
	reservationService := reservation_service.NewReservationServiceClient(reservationConn)


	paymentConn, err := grpc.NewClient("localhost"+cfg.PAYMENT_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("ERROR: ", err.Error())
		return nil
	}
	paymentService := payment_service.NewPaymentServiceClient(paymentConn)


	return &Service{
		Auth: authService,
		Reservation: reservationService,
		Payment: paymentService,
	}
}