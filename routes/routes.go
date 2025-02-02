package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go_rest_api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// server.POST("/events", middlewares.Authorization, createEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authorization)
	authenticated.POST("/events", createEvent) //两种 中间件 使用方法
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("login", login)

}
