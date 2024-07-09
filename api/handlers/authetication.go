package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Projects/Restaurant_Reservation_System/api_gateway/config"
	auth "github.com/Projects/Restaurant_Reservation_System/api_gateway/genproto/authentication_service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type authHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	GetProfileId(c *gin.Context)
}

type authHandlerImpl struct {
	authService auth.AuthenticationServiceClient
}

type Claims struct {
	Email string `json:"email,omitempty"`
	jwt.StandardClaims
}

func (h *authHandlerImpl) Login(c *gin.Context) {
	// Define a struct for the login request
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	// Bind JSON request body to the struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
	log.Println(req.Email, req.Password)

	// Authenticate user
	user, err := h.authService.Login(context.Background(), &auth.LoginRequest{Password: req.Password, Email: req.Email})
	if err != nil {
		// Handle different types of errors
		switch {
		case errors.Is(err, errors.New("ErrorInvalid credentials")):
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		default:
			log.Println("Login failed", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred during login"})
		}
		return
	}

	// Generate JWT token
	token, err := h.generateJWT()
	if err != nil {
		log.Println("Failed to generate JWT", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate authentication token"})
		return
	}

	// Return success response with token
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"Login": gin.H{
			"Success": user.Success,
			// Add other non-sensitive user details as needed
		},
	})
}

func (h *authHandlerImpl) Register(c *gin.Context) {
	// Use the predefined auth.RegisterRequest type
	var req auth.RegisterRequest
	// Bind JSON request body to the struct
	var profile auth.Profile

	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
	req.Profile = &profile

	// Validate the request (assuming auth.RegisterRequest doesn't have built-in validation)
	if err := h.validateRegisterRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(req.Profile)
	// Register the user
	user, err := h.authService.Register(c, &req)
	if err != nil {
		// Handle different types of errors
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

	// Return success response with token and user details
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

// Validation function (if needed)
func (h *authHandlerImpl) validateRegisterRequest(req *auth.RegisterRequest) error {
	if len(req.Profile.Name) < 3 || len(req.Profile.Name) > 30 {
		return errors.New("username must be between 3 and 30 characters")
	}
	if len(req.Profile.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	// Add more validation as needed
	return nil
}

func (h *authHandlerImpl) GetProfileId(c *gin.Context) {
	id := c.Param("id")
	log.Println(id)
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
