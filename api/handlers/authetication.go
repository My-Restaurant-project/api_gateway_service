package handler

import (
	"context"

	auth "github.com/Projects/Restaurant_Reservation_System/api_gateway/genproto/authentication_service"
	"github.com/gin-gonic/gin"
)

type authHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	GetProfileId(c *gin.Context)
}

type authHandlerImpl struct {
	authService auth.AuthenticationServiceClient
}

func (h *authHandlerImpl) Login(c *gin.Context) {

	res, err := h.authService.Login(context.Background(), &auth.LoginRequest{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": res.Token})
	// ctx := context.Background()
	// h.authService.Login(ctx, )
	// Implement login logic using h.authService
}

func (h *authHandlerImpl) Register(c *gin.Context) {
	// Implement register logic using h.authService
}

func (h *authHandlerImpl) GetProfileId(c *gin.Context) {
	// Implement get profile ID logic using h.authService
}
