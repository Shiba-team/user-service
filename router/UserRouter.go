package router

import (
	"authentication/controller"
	"authentication/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	router.GET("/me", middleware.Authentication(nil), controller.UserDetail)
}