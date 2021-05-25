package middleware

import (
	"authentication/config"
	"authentication/constant"
	"authentication/service"
	"log"
	"net/http"
	"strings"

	repo "authentication/repository/repoImpl"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication(role interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		userRepo := repo.NewUserRepo(config.Mongo.UserCollection)

		if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
			username := string(claims["username"].(string))
			user, err := userRepo.FindByUsername(username); if err != nil{
				c.JSON(http.StatusForbidden, gin.H{
					"error": "User not exists!",
				})
				c.Abort()
				return
			}
			if role != nil && user.Role != role.(constant.Role){
				c.JSON(http.StatusForbidden, gin.H{
					"error": "Do not have permission!",
				})
				c.Abort()
				return
			}
			c.Set("user", user)
		}
		c.Next()		
	}
}
