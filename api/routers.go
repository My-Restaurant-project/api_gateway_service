package api

import (
	_ "api_gateway/api/docs"
	handler "api_gateway/api/handlers"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Restaurant Reservation System API
// @version 1.0
// @description This is a sample server for a restaurant reservation system.
// @host localhost:8080
type Server struct {
	Handlers handler.Handlers
}

func NewServer(hands handler.Handlers) *Server {
	return &Server{Handlers: hands}
}

func (s *Server) InitRoutes(r *gin.Engine) {
	r.GET("swagger/*any", ginSwagger.WrapHandler(files.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	auth := s.Handlers.Auth()

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", auth.Login)
		authGroup.POST("/register", auth.Register)
	}

	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	// r.Use(middlewares.JWTMiddlewares)

	r.GET("/auth/profile/:id", auth.GetProfileId)

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
		reservationGroup.POST("/getall", reservation.GetReservations)
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
