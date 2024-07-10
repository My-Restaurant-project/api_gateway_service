package main

import (
	"log"

	api "api_gateway/api"
	handler "api_gateway/api/handlers"
	"api_gateway/config"
	auth "api_gateway/genproto/authentication_service"
	reser "api_gateway/genproto/reservation_service"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	mux := gin.Default()

	authServerConn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connecting to authenication service: ", err)
	}

	reserServerConn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Failed to connecting to authenication service: ", err)
	}

	authClient := auth.NewAuthenticationServiceClient(authServerConn)
	reservationClinet := reser.NewReservationServiceClient(reserServerConn)

	handlers := handler.NewHandlers(authClient, reservationClinet)
	log.Println("Starting API Gateway...")

	server := api.NewServer(handlers)

	server.InitRoutes(mux)

	mux.Run(":" + config.Load().URL_PORT)
}
