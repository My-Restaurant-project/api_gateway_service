package handler

import (
	"api_gateway/api/models"
	rese "api_gateway/genproto/reservation_service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Add new reservation
// @Description Adding new reservation
// @Tags Reservation
// @Accept  json
// @Produce  json
// @Param AddReservationRequest body reservation_service.AddReservationRequest true "Creating new reservation"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reservation/ [post]
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

// @Summary Get all reservation
// @Description Geting all reservation
// @Tags Reservation
// @Accept  json
// @Produce  json
// @Param GetReservationsRequest body reservation_service.GetReservationsRequest true "Get all reservation"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reservation/getall [post]
func (r *reservationHandlerImpl) GetReservations(ctx *gin.Context) {
	getResReq := rese.GetReservationsRequest{}

	if err := ctx.ShouldBindJSON(&getResReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}

	getResRes, err := r.reservationService.GetReservations(context.TODO(), &getResReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Reservations retrieved successfully", "reservations": getResRes.Reservations})
}

// @
// @Summary Get reservation by ID
// @Description Get reservation information using reservation ID
// @Tags Reservation
// @Accept  json
// @Produce  json
// @Param id path string true "Reservation ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reservation/{id} [get]
func (r *reservationHandlerImpl) GetReservation(ctx *gin.Context) {
	id := ctx.Param("id")

	getResReq := rese.GetReservationRequest{Id: id}

	getResRes, err := r.reservationService.GetReservation(context.TODO(), &getResReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Reservation retrieved successfully", "reservation": getResRes.Reservation})
}

// @
// @Summary Update reservation by ID
// @Description Update reservation information using reservation ID
// @Tags Reservation
// @Accept  json
// @Produce  json
// @Param id path string true "Reservation ID"
// @Param UpdateReservationRequest body reservation_service.UpdateReservationRequest true "Updating reservation"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reservation/{id} [put]
func (r *reservationHandlerImpl) UpdateReservation(ctx *gin.Context) {
	id := ctx.Param("id")

	var updateResReq rese.UpdateReservationRequest
	updateResReq.Id = id
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
	ctx.JSON(http.StatusOK, gin.H{"message": "Reservation updated successfully", "reservation": updateResRes.Reservation})
}

// @Summary Delete reservation by ID
// @Description Delete reservation information using reservation ID
// @Tags Reservation
// @Accept  json
// @Produce  json
// @Param id path string true "Reservation ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reservation/{id} [delete]
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

// @Summary Checking Reservation Availability
// @Description Can check Reservation availability via ids
// @Tags Reservation
// @Accept  json
// @Produce  json
// @Param check body models.CheckReservationFilter true "Check Reservation"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reservation/check [post]
func (r *reservationHandlerImpl) CheckReservation(ctx *gin.Context) {
	var checkReser models.CheckReservationFilter
	var reserReq rese.GetReservationRequest

	err := ctx.ShouldBindJSON(&checkReser)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "Invalid format input: " + err.Error()})
		return
	}
	reserReq.Id = checkReser.ReservationID

	reserRes, err := r.reservationService.GetReservation(ctx, &reserReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Reservation not found: " + err.Error()})
		return
	}

	if reserRes.Reservation == nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "Reservation not found"})
		return
	}

	if reserRes.Reservation.GetStatus() == "inactive" {
		ctx.JSON(http.StatusOK, gin.H{"error": "Reservation is not active"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Reservation is available", "reservation": reserRes.Reservation})

}
