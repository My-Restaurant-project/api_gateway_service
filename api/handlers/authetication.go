package handler

import (
	"api_gateway/config"
	auth "api_gateway/genproto/authentication_service"
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email string `json:"email,omitempty"`
	jwt.StandardClaims
}

type AuthHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	GetProfileId(c *gin.Context)
}

type authHandlerImpl struct {
	authService auth.AuthenticationServiceClient
}

func NewAuthHandler(authSer auth.AuthenticationServiceClient) AuthHandler {
	return &authHandlerImpl{authService: authSer}
}

// @Summary Login and Getting Token
// @Description User inserts their credentials like email and password
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param LoginRequest body authentication_service.LoginRequest true "example@gmail.com"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
func (h *authHandlerImpl) Login(c *gin.Context) {
	var req auth.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	user, err := h.authService.Login(context.Background(), &req)
	if err != nil {
		switch {
		case errors.Is(err, errors.New("ErrorInvalid credentials")):
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		default:
			log.Println("Login failed", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred during login"})
		}
		return
	}

	token, err := h.generateJWT()
	if err != nil {
		log.Println("Failed to generate JWT", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate authentication token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"Login": gin.H{
			"Success": user.Success,
		},
	})
}

// @Summary Register user with credentials
// @Description User inserts their credentials
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param profile body authentication_service.Profile true "registering"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
func (h *authHandlerImpl) Register(c *gin.Context) {
	var req auth.RegisterRequest
	var profile auth.Profile

	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
	req.Profile = &profile

	if err := h.validateRegisterRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.Register(c, &req)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "already exists"):
			c.JSON(http.StatusConflict, gin.H{"error": "User with this email already exists"})
		case strings.Contains(err.Error(), "invalid username"):
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
		case strings.Contains(err.Error(), "invalid password"):
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		default:
			log.Println("Registration failed", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred during registration"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":         &user.Profile.Id,
			"username":   user.Profile.Name,
			"email":      user.Profile.Email,
			"password":   user.Profile.Password,
			"role":       user.Profile.Role,
			"created_at": time.Now(),
			"updated_at": time.Now(),
		},
	})
}

func (h *authHandlerImpl) validateRegisterRequest(req *auth.RegisterRequest) error {
	if len(req.Profile.Name) < 3 || len(req.Profile.Name) > 30 {
		return errors.New("username must be between 3 and 30 characters")
	}
	if len(req.Profile.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	return nil
}

// @Summary Get profile by ID
// @Description Retrieve user profile information using user ID
// @Tags Profile
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/profile/{id} [get]
func (h *authHandlerImpl) GetProfileId(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	req := auth.UserIdRequest{
		Id: id,
	}

	resp, err := h.authService.GetProfileById(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *authHandlerImpl) generateJWT() (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{

		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Load().SECRET_KEY))
}
