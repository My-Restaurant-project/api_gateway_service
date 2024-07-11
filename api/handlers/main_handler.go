package handler

import (
	auth "api_gateway/genproto/authentication_service"
	rese "api_gateway/genproto/reservation_service"

	"github.com/gin-gonic/gin"
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

type reservationHandlerImpl struct {
	reservationService rese.ReservationServiceClient
}

func NewReservationHandler(reser rese.ReservationServiceClient) ReservationHandler {
	return &reservationHandlerImpl{reservationService: reser}
}

type ReservationHandler interface {
	CreateRestaurant(*gin.Context)
	GetRestaurants(*gin.Context)
	GetRestaurant(*gin.Context)
	UpdateRestaurant(*gin.Context)
	DeleteRestaurant(*gin.Context)

	//Reservation
	CreateReservation(c *gin.Context)
	GetReservations(c *gin.Context)
	GetReservation(c *gin.Context)
	UpdateReservation(c *gin.Context)
	DeleteReservation(c *gin.Context)
	CheckReservation(c *gin.Context)

	CreateReservationOrder(c *gin.Context)
	PayForReservation(c *gin.Context)

	//Menu
	AddMenu(c *gin.Context)
	GetMenus(c *gin.Context)
	GetMenu(c *gin.Context)
	UpdateMenu(c *gin.Context)
	DeleteMenu(c *gin.Context)
}

func (r *reservationHandlerImpl) CheckReservation(ctx *gin.Context) {
	// need to implement checkReservation
	// var gResOrdReq rese.ReservationOrder
	// ctx.BindJSON()
}

func (r *reservationHandlerImpl) CreateReservationOrder(ctx *gin.Context) {}
func (r *reservationHandlerImpl) PayForReservation(ctx *gin.Context)      {}

func (r *reservationHandlerImpl) AddMenu(ctx *gin.Context)    {}
func (r *reservationHandlerImpl) GetMenus(ctx *gin.Context)   {}
func (r *reservationHandlerImpl) GetMenu(ctx *gin.Context)    {}
func (r *reservationHandlerImpl) UpdateMenu(ctx *gin.Context) {}
func (r *reservationHandlerImpl) DeleteMenu(ctx *gin.Context) {}
