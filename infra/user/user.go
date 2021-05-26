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

func (ur *userRepository) GetAll() (*model.Users, error) {
	users := &model.Users{
		model.User{
			FirstName: "最初の",
			LastName:  "ユーザーです",
			NickName:  "初代",
			Age:       1,
		},
		model.User{
			FirstName: "２番めの",
			LastName:  "ユーザーです",
			NickName:  "二代目",
			Age:       2,
		},
		model.User{
			FirstName: "３番めの",
			LastName:  "ユーザーです",
			NickName:  "三代目",
			Age:       3,
		},
	}

	return users, nil
}
