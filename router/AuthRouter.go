package router

import (
	"authentication/controller"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.RouterGroup) {
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
	router.POST("/verify-token", controller.VerifyToken)
}