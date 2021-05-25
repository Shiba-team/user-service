package repoimpl

import (
	"authentication/model"
	repo "authentication/repository"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepoImpl struct{
	UserCollection *mongo.Collection
}

func NewUserRepo(userCollection *mongo.Collection)  repo.UserRepo{
	return &UserRepoImpl{
		UserCollection: userCollection,
	}
}

func (u *UserRepoImpl) FindByUsername(username string) (model.User , error){
	var user model.User
	if err := u.UserCollection.FindOne(context.Background(), bson.M{"username" : username}).Decode(&user); err != nil{
		return user, err
	}
	return user, nil
}

func (u *UserRepoImpl) FindAllUser() ([]model.User , error){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel();
	users := make([]model.User, 0)
	cursor, err := u.UserCollection.Find(ctx, bson.M{})
	if err != nil{
		log.Fatal(err)
		return users, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx){
		var user model.User
		if err = cursor.Decode(&user); err != nil {
			log.Fatal(err)
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepoImpl) InsertUser(user model.User) (error){
	_, err := u.UserCollection.InsertOne(context.Background(), user)
	 if err != nil{
		log.Fatal(err)
		return err
	}
	return nil
}