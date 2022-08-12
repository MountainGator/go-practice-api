package services

import (
	"context"
	"example.com/api_practice/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	// pointer to the mongo object
	user_collection *mongo.Collection
	ctx				contxt.Context
}
//func is a pointer so you include that first
func (u *UserServiceImpl) CreateUser(user *models.User) error {
	return nil
}
// after the name of the function you add the return types
func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) GetAll() []*models.User {
	return nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	return nil
}