package main

import (
	"log"

	handler "github.com/Projects/Restaurant_Reservation_System/api_gateway/api/handlers"
	"github.com/Projects/Restaurant_Reservation_System/api_gateway/api/routes"
	"github.com/Projects/Restaurant_Reservation_System/api_gateway/genproto/authentication_service"
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
	authClient := authentication_service.NewAuthenticationServiceClient(authServerConn)

	handlers := handler.NewHandlers(authClient)
	log.Println("Starting API Gateway...")

	server := routes.NewServer(handlers)

	server.InitRoutes(mux)

	mux.Run(":8070")
}
