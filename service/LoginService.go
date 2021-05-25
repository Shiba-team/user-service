package service

import (
	"authentication/common"
	"authentication/config"
	"authentication/model"
	repo "authentication/repository/repoImpl"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type authCustomClaims struct {
	username string
	jwt.StandardClaims
}

func GenerateToken(username string) string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	claims := &authCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return t
}

func  ValidateToken(encodedToken string) (*jwt.Token, error) {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(secret), nil
	})

}

func Login(dto model.UserLoginDto) (bool,string) {


	userRepo := repo.NewUserRepo(config.Mongo.UserCollection)
	user, err := userRepo.FindByUsername(dto.Username); if err != nil{
		return false, "Username or password incorrect!"
	}
	
	if(!common.CheckPasswordHash(dto.Password, user.Password)){
		return false, "Username or password incorrect!"
	}
	return true, GenerateToken(dto.Username)
}


