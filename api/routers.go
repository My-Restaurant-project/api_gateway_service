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
	Handlers *handler.Handlers
}

func NewServer(hands *handler.Handlers) *Server {
	return &Server{Handlers: hands}
}

func (s *Server) InitRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", s.Handlers.Login)
		authGroup.POST("/register", s.Handlers.Register)
	}
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.Use(middlewares.JWTMiddlewares)
	r.GET("auth/profile/:id", s.Handlers.GetProfileId)

	r.GET("/auth/profile/:id", s.Handlers.GetProfileId)

	restaurantGroup := r.Group("/restaurant")
	{
		restaurantGroup.GET("/:id")
		restaurantGroup.POST()
	}
}
