package main

import (
	"log"

	api "api_gateway/api"
	handler "api_gateway/api/handlers"
	"api_gateway/config"
	auth "api_gateway/genproto/authentication_service"
	payment "api_gateway/genproto/payment_service"
	reser "api_gateway/genproto/reservation_service"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	mux := gin.Default()

	authServerConn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to authentication service: ", err)
	}
	defer authServerConn.Close()

	reserServerConn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to reservation service: ", err)
	}
	defer reserServerConn.Close()

	paymentServerConn, err := grpc.NewClient("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to payment service: ", err)
	}
	defer paymentServerConn.Close()

	authClient := auth.NewAuthenticationServiceClient(authServerConn)
	reservationClient := reser.NewReservationServiceClient(reserServerConn)
	paymentClient := payment.NewPaymentServiceClient(paymentServerConn)

	handlers := handler.NewHandlers(authClient, reservationClient, paymentClient)
	log.Println("Starting API Gateway...")

	server := api.NewServer(handlers)

	server.InitRoutes(mux)

	mux.Run(":" + config.Load().URL_PORT)
}
