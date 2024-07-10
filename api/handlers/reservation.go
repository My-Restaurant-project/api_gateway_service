package handler

import (
	rese "api_gateway/genproto/reservation_service"
	"context"
	"net/http"

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

func (r *reservationHandlerImpl) CreateRestaurant(ctx *gin.Context) {
	var addRestReq rese.AddRestaurantRequest

	err := ctx.BindJSON(&addRestReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}

	addRestRes, err := r.reservationService.AddRestaurant(context.TODO(), &addRestReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Restaurant added successfully", "restaurant": addRestRes})
}

func (r *reservationHandlerImpl) GetRestaurants(ctx *gin.Context) {
	getRestReq := rese.GetRestaurantsRequest{}

	getRestRes, err := r.reservationService.GetRestaurants(context.TODO(), &getRestReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Restaurants retrieved successfully", "restaurants": getRestRes})
}
func (r *reservationHandlerImpl) GetRestaurant(ctx *gin.Context) {
	id := ctx.Param("id")

	getRestReq := rese.GetRestaurantRequest{Id: id}

	getRestRes, err := r.reservationService.GetRestaurant(context.TODO(), &getRestReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Restaurant retrieved successfully", "restaurant": getRestRes})
}

func (r *reservationHandlerImpl) UpdateRestaurant(ctx *gin.Context) {
	id := ctx.Param("id")

	var updateRestReq rese.UpdateRestaurantRequest
	err := ctx.BindJSON(&updateRestReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}
	updateRestReq.Id = id

	updateRestRes, err := r.reservationService.UpdateRestaurant(context.TODO(), &updateRestReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Restaurant updated successfully", "restaurant": updateRestRes})
}

func (r *reservationHandlerImpl) DeleteRestaurant(ctx *gin.Context) {
	id := ctx.Param("id")

	deleteRestReq := rese.DeleteRestaurantRequest{Id: id}

	_, err := r.reservationService.DeleteRestaurant(context.TODO(), &deleteRestReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Restaurant deleted successfully"})
}

func (r *reservationHandlerImpl) CreateReservation(ctx *gin.Context) {
	var addResReq rese.AddReservationRequest

	err := ctx.BindJSON(&addResReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}

	addResRes, err := r.reservationService.AddReservation(context.TODO(), &addResReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Reservation added successfully", "reservation": addResRes})
}

func (r *reservationHandlerImpl) GetReservations(ctx *gin.Context) {
	getResReq := rese.GetReservationsRequest{}

	getResRes, err := r.reservationService.GetReservations(context.TODO(), &getResReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Reservations retrieved successfully", "reservations": getResRes})
}

func (r *reservationHandlerImpl) GetReservation(ctx *gin.Context) {
	id := ctx.Param("id")

	getResReq := rese.GetReservationRequest{Id: id}

	getResRes, err := r.reservationService.GetReservation(context.TODO(), &getResReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Reservation retrieved successfully", "reservation": getResRes})
}

func (r *reservationHandlerImpl) UpdateReservation(ctx *gin.Context) {
	id := ctx.Param("id")

	var updateResReq rese.UpdateReservationRequest
	err := ctx.BindJSON(&updateResReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}
	updateResReq.Id = id

	updateResRes, err := r.reservationService.UpdateReservation(context.TODO(), &updateResReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Reservation updated successfully", "reservation": updateResRes})
}

func (r *reservationHandlerImpl) DeleteReservation(ctx *gin.Context) {
	id := ctx.Param("id")

	deleteResReq := rese.DeleteReservationRequest{Id: id}

	_, err := r.reservationService.DeleteReservation(context.TODO(), &deleteResReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Reservation deleted successfully"})
}

func (r *reservationHandlerImpl) CheckReservation(ctx *gin.Context) {
	// need to implement checkReservation
}

func (r *reservationHandlerImpl) CreateReservationOrder(ctx *gin.Context) {}
func (r *reservationHandlerImpl) PayForReservation(ctx *gin.Context)      {}

func (r *reservationHandlerImpl) AddMenu(ctx *gin.Context)    {}
func (r *reservationHandlerImpl) GetMenus(ctx *gin.Context)   {}
func (r *reservationHandlerImpl) GetMenu(ctx *gin.Context)    {}
func (r *reservationHandlerImpl) UpdateMenu(ctx *gin.Context) {}
func (r *reservationHandlerImpl) DeleteMenu(ctx *gin.Context) {}
