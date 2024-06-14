package router

import (
	"belajar_golang/unit_testing/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", handler.RootHandler)

	r.POST("/post", handler.PostHandler)

}
