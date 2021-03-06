package controller

import (
	"authentication/model"
	"authentication/service"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)



func Register(c * gin.Context){
	var input model.UserRegisterDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	  }
	err := service.Register(input)
	if err != ""{
		c.JSON(http.StatusInternalServerError, gin.H{"success" : false, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success" : true, "message": "Register successful!"})
}

func Login(c * gin.Context){
	var input model.UserLoginDto;
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	resutl, msg := service.Login(input)
	if(!resutl){
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": msg})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"success": false, "message": "Login successful!", "data":msg})
}

func VerifyToken(c * gin.Context) {

	tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			log.Println("unAuth")
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Token is required!",
			})
			c.Abort()
			return
		}
	
	token := strings.Split(tokenHeader, " ")[1]
	

	result, err := service.ValidateToken(token);
	if err != nil {
		log.Println("unAuth")
		c.JSON(401, gin.H{
			"error": "No Authentication",
		})
		c.Abort()
		return
	}
	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		username := string(claims["username"].(string))
		c.JSON(http.StatusAccepted, gin.H{"username":username})
	}
}
