package router

import (
	"authentication/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	router.POST("/register", controller.Login)
	router.POST("/login", controller.Login)
}