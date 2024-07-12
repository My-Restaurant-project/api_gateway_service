package handler

import (
	pb "api_gateway/genproto/payment_service"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentHandler interface {
	CreatePayment(ctx *gin.Context)
	GetPayment(ctx *gin.Context)
	ListPayments(ctx *gin.Context)
	UpdatePayment(ctx *gin.Context)
	DeletePayment(ctx *gin.Context)
}

type paymentHandlerImpl struct {
	paymentService pb.PaymentServiceClient
}

func NewPaymentHandler(paymentService pb.PaymentServiceClient) PaymentHandler {
	return &paymentHandlerImpl{paymentService: paymentService}
}

// @Summary Add new payment
// @Description Adding new payment
// @Tags Payment
// @Accept  json
// @Produce  json
// @Param AddPaymentRequest body payment_service.CreatePaymentRequest true "Creating new payment"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payment/ [post]
func (r *paymentHandlerImpl) CreatePayment(ctx *gin.Context) {
	var createPaymentReq pb.CreatePaymentRequest
	err := ctx.BindJSON(&createPaymentReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}

	createPaymentRes, err := r.paymentService.CreatePayment(context.TODO(), &createPaymentReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Payment added successfully", "payment": createPaymentRes.Payment})
}

// @Summary Get payment by ID
// @Description Get payment information using payment ID
// @Tags Payment
// @Accept  json
// @Produce  json
// @Param id path string true "Payment ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payment/{id} [get]
func (r *paymentHandlerImpl) GetPayment(ctx *gin.Context) {
	id := ctx.Param("id")

	getPaymentReq := pb.GetPaymentRequest{Id: id}

	getPaymentRes, err := r.paymentService.GetPayment(context.TODO(), &getPaymentReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Payment retrieved successfully", "payment": getPaymentRes.Payment})
}

// @Summary List payments
// @Description List payments
// @Tags Payment
// @Accept  json
// @Produce  json
// @Param reservation_id query string false "Reservation ID"
// @Param payment_method query string false "Payment Method"
// @Param payment_status query string false "Payment Status"
// @Param limit query int false "Limit"
// @Param offset query string false "Offset"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payments/ [get]
func (r *paymentHandlerImpl) ListPayments(ctx *gin.Context) {
	listPaymentsReq := pb.ListPaymentsRequest{
		ReservationId: ctx.Query("reservation_id"),
		PaymentMethod: ctx.Query("payment_method"),
		PaymentStatus: ctx.Query("payment_status"),
		Limit:         getLimitFromQuery(ctx),
		Offset:        ctx.Query("offset"),
	}

	listPaymentsRes, err := r.paymentService.ListPayments(context.TODO(), &listPaymentsReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Payments retrieved successfully", "payments": listPaymentsRes.Payments})
}

// @Summary Update payment by ID
// @Description Update payment information using payment ID
// @Tags Payment
// @Accept  json
// @Produce  json
// @Param id path string true "Payment ID"
// @Param UpdatePaymentRequest body payment_service.UpdatePaymentRequest true "Updating payment"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payment/{id} [put]
func (r *paymentHandlerImpl) UpdatePayment(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatePaymentReq pb.UpdatePaymentRequest
	err := ctx.BindJSON(&updatePaymentReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}
	updatePaymentReq.Payment.Id = id

	updatePaymentRes, err := r.paymentService.UpdatePayment(context.TODO(), &updatePaymentReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Payment updated successfully", "payment": updatePaymentRes.Payment})
}

// @Summary Delete payment by ID
// @Description Delete payment information using payment ID
// @Tags Payment
// @Accept  json
// @Produce  json
// @Param id path string true "Payment ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payment/{id} [delete]
func (r *paymentHandlerImpl) DeletePayment(ctx *gin.Context) {
	id := ctx.Param("id")

	deletePaymentReq := pb.DeletePaymentRequest{Id: id}

	_, err := r.paymentService.DeletePayment(context.TODO(), &deletePaymentReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Payment deleted successfully"})
}

func getLimitFromQuery(ctx *gin.Context) int32 {
	limit := int32(10)
	if ctx.Query("limit") != "" {
		if l, err := strconv.ParseInt(ctx.Query("limit"), 10, 32); err == nil {
			limit = int32(l)
		}
	}
	return limit
}
