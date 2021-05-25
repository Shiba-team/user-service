package repository

import (
	"authentication/model"
)

type UserRepo interface {
	FindByUsername(username string) (model.User , error)
	FindAllUser() ([]model.User , error)
	InsertUser(user model.User) (error)
}