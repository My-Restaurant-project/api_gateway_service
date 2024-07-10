package handler

import (
	res "api_gateway/genproto/reservation_service"

	"github.com/gin-gonic/gin"
)

type reservationHandler interface {
	//Restaurant
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

type reservationImpl struct {
	restaurantService res.ReservationServiceClient
}

func NewReservationHandler(restaurantService res.ReservationServiceClient) reservationHandler {
	return &reservationImpl{restaurantService: restaurantService}
}

func (r *reservationImpl) CreateRestaurant(ctx *gin.Context) {}
func (r *reservationImpl) GetRestaurants(ctx *gin.Context)   {}
func (r *reservationImpl) GetRestaurant(ctx *gin.Context)    {}
func (r *reservationImpl) UpdateRestaurant(ctx *gin.Context) {}
func (r *reservationImpl) DeleteRestaurant(ctx *gin.Context) {}

func (r *reservationImpl) CreateReservation(ctx *gin.Context) {}
func (r *reservationImpl) GetReservations(ctx *gin.Context)   {}
func (r *reservationImpl) GetReservation(ctx *gin.Context)    {}
func (r *reservationImpl) UpdateReservation(ctx *gin.Context) {}
func (r *reservationImpl) DeleteReservation(ctx *gin.Context) {}
func (r *reservationImpl) CheckReservation(ctx *gin.Context)  {}

func (r *reservationImpl) CreateReservationOrder(ctx *gin.Context) {}
func (r *reservationImpl) PayForReservation(ctx *gin.Context)      {}

func (r *reservationImpl) AddMenu(ctx *gin.Context)    {}
func (r *reservationImpl) GetMenus(ctx *gin.Context)   {}
func (r *reservationImpl) GetMenu(ctx *gin.Context)    {}
func (r *reservationImpl) UpdateMenu(ctx *gin.Context) {}
func (r *reservationImpl) DeleteMenu(ctx *gin.Context) {}
