package handler

import (
	auth "github.com/Projects/Restaurant_Reservation_System/api_gateway/genproto/authentication_service"
)

func NewAuthHandler(authSer auth.AuthenticationServiceClient, jwtSecret string) authHandler {
	return &authHandlerImpl{
		authService: authSer,
		jwtSecret:   []byte(jwtSecret),
	}
}

type Handlers struct {
	authHandler
}

func NewHandlers(authSer auth.AuthenticationServiceClient,jwtSecret string) *Handlers {
	auth := NewAuthHandler(authSer,jwtSecret)
	return &Handlers{authHandler: auth}
}
