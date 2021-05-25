package router

import (
	"authentication/constant"
	"authentication/controller"
	"authentication/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.RouterGroup) {
	router.GET("/get-all", middleware.Authentication(constant.ADMIN), controller.GetAllUser)
}