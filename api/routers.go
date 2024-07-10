package api

import (
	_ "api_gateway/api/docs"
	handler "api_gateway/api/handlers"
	"api_gateway/api/middlewares"

	"github.com/gin-gonic/gin"
)

// @title Restaurant Reservation System API
// @version 1.0
// @description This is a sample server for a restaurant reservation system.
// @host localhost:8070

type Server struct {
	Handlers handler.Handlers
}

func NewServer(hands handler.Handlers) *Server {
	return &Server{Handlers: hands}
}

func (s *Server) InitRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", s.Handlers.Auth().Login)
		authGroup.POST("/register", s.Handlers.Auth().Register)
	}
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.Use(middlewares.JWTMiddlewares)
	r.GET("/auth/profile/:id", s.Handlers.Auth().GetProfileId)

	restaurantGroup := r.Group("/restaurant")
	{
		restaurantGroup.POST("/", s.Handlers.Reservation().CreateRestaurant)
		restaurantGroup.GET("/:id", s.Handlers.Reservation().GetRestaurant)
		restaurantGroup.PUT("/:id", s.Handlers.Reservation().UpdateRestaurant)
		restaurantGroup.DELETE("/:id", s.Handlers.Reservation().DeleteRestaurant)
		restaurantGroup.GET("/", s.Handlers.Reservation().GetRestaurants)
	}
	reservationGroup := r.Group("/reservation")
	{
		reservationGroup.POST("/", s.Handlers.Reservation().CreateReservation)
		reservationGroup.GET("/:id", s.Handlers.Reservation().GetReservation)
		reservationGroup.PUT("/:id", s.Handlers.Reservation().UpdateReservation)
		reservationGroup.DELETE("/:id", s.Handlers.Reservation().DeleteReservation)
		reservationGroup.GET("/", s.Handlers.Reservation().GetReservations)
		reservationGroup.POST("/check", s.Handlers.Reservation().CheckReservation)
		reservationGroup.POST("/order", s.Handlers.Reservation().CreateReservationOrder)
	}
}
