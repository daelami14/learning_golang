package main

import (
	"belajar_golang/pertemuan_ke_1/middleware"
	"belajar_golang/pertemuan_ke_1/router"
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
