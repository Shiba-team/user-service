package controller

import (
	"net/http"

	"authentication/service"

	"github.com/gin-gonic/gin"
)


func GetAllUser(c * gin.Context){
	result, err := service.GetAllUser();
	if(err != nil){
		c.JSON(http.StatusInternalServerError, gin.H{"success" : false,"message": "Can not get all user!"})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data" : result, "message": "Get all user successful!"})
}