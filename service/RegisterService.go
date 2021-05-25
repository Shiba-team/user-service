package service

import (
	"authentication/config"
	"authentication/model"
	repo "authentication/repository/repoImpl"
)
func Register(dto model.UserRegisterDto) (string){
	userRepo := repo.NewUserRepo(config.Mongo.UserCollection)
	if _, findErr := userRepo.FindByUsername(dto.Username); findErr != nil {
			return "Username has already existed!"
		}
	user := model.UserRegisterDtoToEntity(dto)
	insertErr := userRepo.InsertUser(user.(model.User))
		if insertErr != nil {
		 return "Can not create user!"
		}
	return ""
}