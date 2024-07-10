package handler

import (
	auth "api_gateway/genproto/authentication_service"
	rese "api_gateway/genproto/reservation_service"
)

type Handlers interface {
	Auth() AuthHandler
	Reservation() ReservationHandler
}

type handlers struct {
	authHandler        AuthHandler
	reservationHandler ReservationHandler
}

func NewHandlers(authSer auth.AuthenticationServiceClient, reser rese.ReservationServiceClient) Handlers {
	return &handlers{
		authHandler:        NewAuthHandler(authSer),
		reservationHandler: NewReservationHandler(reser),
	}
}

func (h *handlers) Auth() AuthHandler {
	return h.authHandler
}

func (h *handlers) Reservation() ReservationHandler {
	return h.reservationHandler
}
