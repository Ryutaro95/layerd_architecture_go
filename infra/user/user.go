package infra

import (
	"fmt"

	"github.com/Ryutaro95/domain/model"
	repository "github.com/Ryutaro95/domain/repository/user"
)

type userRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (ur *userRepository) Create(u *model.User) *model.User {
	fmt.Printf("ユーザー名: %v %v\nニックネーム: %v\n年齢: %v\n", u.FirstName, u.LastName, u.NickName, u.Age)
	return u
}
