package handler

import (
	auth "api_gateway/genproto/authentication_service"
	payment "api_gateway/genproto/payment_service"
	rese "api_gateway/genproto/reservation_service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Auth() AuthHandler
	Reservation() ReservationHandler
	Payment() PaymentHandler
}

type handlers struct {
	authHandler        AuthHandler
	reservationHandler ReservationHandler
	paymentHandler     PaymentHandler
}

func NewHandlers(authSer auth.AuthenticationServiceClient, reser rese.ReservationServiceClient, pay payment.PaymentServiceClient) Handlers {
	return &handlers{
		authHandler:        NewAuthHandler(authSer),
		reservationHandler: NewReservationHandler(reser),
		paymentHandler:     NewPaymentHandler(pay),
	}
}

func (h *handlers) Auth() AuthHandler {
	return h.authHandler
}

func (h *handlers) Reservation() ReservationHandler {
	return h.reservationHandler
}

func (h *handlers) Payment() PaymentHandler {
	return h.paymentHandler
}

// Reservation Handler Implementation

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

	// Reservation
	CreateReservation(c *gin.Context)
	GetReservations(c *gin.Context)
	GetReservation(c *gin.Context)
	UpdateReservation(c *gin.Context)
	DeleteReservation(c *gin.Context)
	CheckReservation(c *gin.Context)

	CreateReservationOrder(c *gin.Context)
	PayForReservation(c *gin.Context)

	// Menu
	AddMenu(c *gin.Context)
	GetMenus(c *gin.Context)
	GetMenu(c *gin.Context)
	UpdateMenu(c *gin.Context)
	DeleteMenu(c *gin.Context)
}

// @Summary Add new order for reservation
// @Description Adding new order for reservation
// @Tags Reservation
// @Accept  json
// @Produce  json
// @Param AddReservationOrderRequest body reservation_service.AddReservationOrderRequest true "Creating new order"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reservation/order [post]
func (r *reservationHandlerImpl) CreateReservationOrder(ctx *gin.Context) {
	var resOrderReq rese.AddReservationOrderRequest

	if ctx.ShouldBindJSON(&resOrderReq) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	resOrderRes, err := r.reservationService.AddReservationOrder(context.TODO(), &resOrderReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Reservation order added successfully", "reservation_order": resOrderRes})

}

func (r *reservationHandlerImpl) PayForReservation(ctx *gin.Context) {}
