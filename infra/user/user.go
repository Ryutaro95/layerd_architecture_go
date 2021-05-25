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
	// 本来はデータを永続化するが、簡略化のため一旦標準出力
	fmt.Printf("ユーザー名: %v %v\nニックネーム: %v\n年齢: %v\n", u.FirstName, u.LastName, u.NickName, u.Age)
	return u
}

func (ur *userRepository) GetByID(userID int) (*model.User, error) {
	// モックのユーザーオブジェクトを返す
	user := &model.User{
		FirstName: "テスト",
		LastName:  "ユーザー",
		NickName:  "test",
		Age:       33,
	}
	return user, nil
}
