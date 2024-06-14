package main

import (
	"learning_golang/session4_unit_test_crud/entity"
	"learning_golang/session4_unit_test_crud/handler"
	"learning_golang/session4_unit_test_crud/repository/slice"
	"learning_golang/session4_unit_test_crud/router"
	"learning_golang/session4_unit_test_crud/service"
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
