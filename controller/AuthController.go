package controller

import (
	"authentication/model"
	"authentication/service"
	"fmt"
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
		c.JSON(http.StatusInternalServerError, gin.H{"messageError": err})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Register successful!"})
}

func Login(c * gin.Context){
	var input model.UserLoginDto;
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resutl, msg := service.Login(input)
	if(!resutl){
		c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
		return
	}
	fmt.Println(msg)
	c.JSON(http.StatusAccepted, gin.H{"token":msg})
}

func VerifyToken(c * gin.Context){
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

	c.JSON(401, gin.H{
		"error": "No Authentication",
	})
}