package repository

import "github.com/Ryutaro95/domain/model"

type UserRepository interface {
	Create(u *model.User) *model.User
	GetByID(userID int) (*model.User, error)
	GetAll() (*model.Users, error)
}
