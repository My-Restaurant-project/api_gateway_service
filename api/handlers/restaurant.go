package handler

import (
	rese "api_gateway/genproto/reservation_service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Add new restaurant
// @Description Adding new restaurant
// @Tags Restaurant
// @Accept  json
// @Produce  json
// @Param AddRestaurantRequest body reservation_service.AddRestaurantRequest true "Creating new restaurnat"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /restaurant/ [post]
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
	updateRestReq.Id = id

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
