package handler

import (
	auth "api_gateway/genproto/authentication_service"
	rese "api_gateway/genproto/reservation_service"
)

func NewAuthHandler(authSer auth.AuthenticationServiceClient) authHandler {
	return &authHandlerImpl{authService: authSer}
}

type Handlers struct {
	authHandler
	reservationHandler
}

func NewHandlers(authSer auth.AuthenticationServiceClient, reser rese.ReservationServiceClient) *Handlers {
	auth := NewAuthHandler(authSer)
	res := NewReservationHandler(reser)
	return &Handlers{authHandler: auth, reservationHandler: res}
}
