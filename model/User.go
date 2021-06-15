package model

import (
	"authentication/common"
	"authentication/constant"
	"time"

	"github.com/google/uuid"
)

type User struct {
    Username   string           `json:"username" validate:"required"`
    Email   string           `json:"email" validate:"required"`
    Password   string            `json:"password" validate:"required"`
    Name       string            `json:"name"`
    Role   constant.Role            `json:"role" validate:"required"`
    Created_at time.Time          `json:"created_at"`
    Updated_at time.Time          `json:"updated_at"`
    SecretKey  string             `json:"secretKey"`
}

type UserRegisterDto struct {
    Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
    Name string `json:"name" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

type UserLoginDto struct {
    Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
    Name       string            `json:"name"`
}

type UserDto struct{
    Username   string           `json:"username"`
    Email   string           `json:"email" `
    Name       string            `json:"name"`
    SecretKey  string             `json:"secretKey"`
}

func UserRegisterDtoToEntity (dto UserRegisterDto) interface{}{
    var user User;
    user.Username = dto.Username;
    user.Email = dto.Email;
    password, err := common.HashPassword(dto.Password)
        if(err != nil){
            return err
        }
    user.Password = string(password)
    user.Role = constant.USER;
    user.SecretKey = uuid.NewString()
    user.Updated_at = time.Now()
    if(user.Created_at == time.Time{}){
        user.Created_at = time.Now()
    }
    return user;
}


func EntityToUserDto (user User) UserDto{
    var userDto UserDto;
    userDto.SecretKey = user.SecretKey
    userDto.Email = user.Email
    userDto.Username = user.Username
    userDto.Name = user.Name
    return userDto;
}