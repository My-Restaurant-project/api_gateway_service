// package handler

// import (
// 	rese "api_gateway/genproto/reservation_service"
// )

// type reservationHandlerImpl struct {
// 	reservationService rese.ReservationServiceClient
// }

// func NewReservationHandler(restaurantService rese.ReservationServiceClient) reservationHandler {
// 	return &reservationHandlerImpl{restaurantService: restaurantService}
// }

package handler

import (
	rese "api_gateway/genproto/reservation_service"

	"github.com/gin-gonic/gin"
)

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

func (r *reservationHandlerImpl) CreateRestaurant(ctx *gin.Context) {}
func (r *reservationHandlerImpl) GetRestaurants(ctx *gin.Context)   {}
func (r *reservationHandlerImpl) GetRestaurant(ctx *gin.Context)    {}
func (r *reservationHandlerImpl) UpdateRestaurant(ctx *gin.Context) {}
func (r *reservationHandlerImpl) DeleteRestaurant(ctx *gin.Context) {}

func (r *reservationHandlerImpl) CreateReservation(ctx *gin.Context) {}
func (r *reservationHandlerImpl) GetReservations(ctx *gin.Context)   {}
func (r *reservationHandlerImpl) GetReservation(ctx *gin.Context)    {}
func (r *reservationHandlerImpl) UpdateReservation(ctx *gin.Context) {}
func (r *reservationHandlerImpl) DeleteReservation(ctx *gin.Context) {}
func (r *reservationHandlerImpl) CheckReservation(ctx *gin.Context)  {}

func (r *reservationHandlerImpl) CreateReservationOrder(ctx *gin.Context) {}
func (r *reservationHandlerImpl) PayForReservation(ctx *gin.Context)      {}

func (r *reservationHandlerImpl) AddMenu(ctx *gin.Context)    {}
func (r *reservationHandlerImpl) GetMenus(ctx *gin.Context)   {}
func (r *reservationHandlerImpl) GetMenu(ctx *gin.Context)    {}
func (r *reservationHandlerImpl) UpdateMenu(ctx *gin.Context) {}
func (r *reservationHandlerImpl) DeleteMenu(ctx *gin.Context) {}
