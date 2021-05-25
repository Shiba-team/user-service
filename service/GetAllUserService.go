package service

import (
	"authentication/config"
	"authentication/model"
	repo "authentication/repository/repoImpl"
)

func GetAllUser() ([]model.User, error){
	userRepo := repo.NewUserRepo(config.Mongo.UserCollection)
	return userRepo.FindAllUser()
}