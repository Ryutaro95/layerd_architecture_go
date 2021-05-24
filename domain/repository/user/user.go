package repository

import "github.com/Ryutaro95/domain/model"

type UserRepository interface {
	Create(u *model.User) *model.User
}
