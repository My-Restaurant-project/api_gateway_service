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

	reservation := s.Handlers.Reservation()

	restaurantGroup := r.Group("/restaurant")
	{
		restaurantGroup.POST("/", reservation.CreateRestaurant)
		restaurantGroup.GET("/:id", reservation.GetRestaurant)
		restaurantGroup.PUT("/:id", reservation.UpdateRestaurant)
		restaurantGroup.DELETE("/:id", reservation.DeleteRestaurant)
		restaurantGroup.GET("/", reservation.GetRestaurants)
	}
	reservationGroup := r.Group("/reservation")
	{
		reservationGroup.POST("/", reservation.CreateReservation)
		reservationGroup.GET("/:id", reservation.GetReservation)
		reservationGroup.PUT("/:id", reservation.UpdateReservation)
		reservationGroup.DELETE("/:id", reservation.DeleteReservation)
		reservationGroup.GET("/", reservation.GetReservations)
		reservationGroup.POST("/check", reservation.CheckReservation)
		reservationGroup.POST("/order", reservation.CreateReservationOrder)
	}
	menuGroup := r.Group("/menu")
	{
		menuGroup.POST("/", reservation.AddMenu)
		menuGroup.GET("/:id", reservation.GetMenu)
		menuGroup.PUT("/:id", reservation.UpdateMenu)
		menuGroup.DELETE("/:id", reservation.DeleteMenu)
		menuGroup.GET("/", reservation.GetMenus)
	}

}
