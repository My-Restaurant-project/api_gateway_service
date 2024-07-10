package handler

import "github.com/gin-gonic/gin"

type reservationHandler interface {
	CreateReservation(c *gin.Context)
	GetReservations(c *gin.Context)
	GetReservation(c *gin.Context)
	UpdateReservation(c *gin.Context)
	DeleteReservation(c *gin.Context)
}

type reservationImpl struct {
	reservationService ReservationServiceClient
}
