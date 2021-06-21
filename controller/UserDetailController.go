package controller

import (
	"authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserDetail(c * gin.Context){
	user, isExisted := service.UserDetail(c);
	if(!isExisted){
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Can not get user detail!", "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Get user detail successful!", "data" : user})
}
