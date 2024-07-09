package routes

import (
	handler "github.com/Projects/Restaurant_Reservation_System/api_gateway/api/handlers"
	"github.com/gin-gonic/gin"
)

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
		authGroup.GET("/profile", s.Handlers.GetProfileId)
	}

	restaurantGroup := r.Group("/restaurant")
	{
		restaurantGroup.GET("/:id")

	}
}
