package main

import (
	"belajar_golang/pertemuan_ke_4/entity"
	"belajar_golang/pertemuan_ke_4/handler"
	"belajar_golang/pertemuan_ke_4/repository/slice"
	"belajar_golang/pertemuan_ke_4/router"
	"belajar_golang/pertemuan_ke_4/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// setup service
	var mockUserDBInSlice []entity.User
	userRepo := slice.NewUserRepository(mockUserDBInSlice)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	router.SetupRouter(r, userHandler)

	// Run the server
	log.Println("Running server on port 8080")
	r.Run(":8080")
}
