package handler

import (
	"api_gateway/api/models"
	auth "api_gateway/genproto/authentication_service"
	rese "api_gateway/genproto/reservation_service"
	"net/http"

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

	err := ctx.ShouldBindJSON(checkReser)
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

func (r *reservationHandlerImpl) CreateReservationOrder(ctx *gin.Context) {}
func (r *reservationHandlerImpl) PayForReservation(ctx *gin.Context)      {}

func (r *reservationHandlerImpl) AddMenu(ctx *gin.Context)    {}
func (r *reservationHandlerImpl) GetMenus(ctx *gin.Context)   {}
func (r *reservationHandlerImpl) GetMenu(ctx *gin.Context)    {}
func (r *reservationHandlerImpl) UpdateMenu(ctx *gin.Context) {}
func (r *reservationHandlerImpl) DeleteMenu(ctx *gin.Context) {}
