package service

import (
	"authentication/model"
	"log"

	"github.com/gin-gonic/gin"
)
func UserDetail(c *gin.Context) (interface{}, bool){
	log.Println(c)
	user, isExisted := c.Get("user");
	if isExisted{
		return model.EntityToUserDto(user.(model.User)), isExisted
	}
	return nil, isExisted

}