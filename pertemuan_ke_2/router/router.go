package router

import (
	"belajar_golang/pertemuan_ke_2/handler"
	"belajar_golang/pertemuan_ke_2/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	usersPublicEndpoint := r.Group("/users")
	usersPublicEndpoint.GET("/:id", handler.GetUser)
	usersPublicEndpoint.GET("/", handler.GetAllUsers)

	usersPrivateEndpoint := r.Group("/users")
	usersPrivateEndpoint.Use(middleware.AuthMiddleware())
	usersPrivateEndpoint.POST("/", handler.CreateUser)
	usersPrivateEndpoint.PUT("/:id", handler.UpdateUser)
	usersPrivateEndpoint.DELETE("/:id", handler.DeleteUser)
}