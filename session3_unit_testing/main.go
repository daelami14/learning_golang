package main

import (
	"belajar_golang/unit_testing/middleware"
	"belajar_golang/unit_testing/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	router.SetupRouter(r)

	log.Println("Running server on port 8080")
	r.Run(":8080")

}
