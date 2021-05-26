package model

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v3"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	NickName  string
	Age       int
}

type Users []User

/* バリデーション
- first_name: 必須, 文字列の長さ 1 ~ 10
- last_name: 必須, 文字列の長さ 1 ~ 10
- nick_name: 必須, 文字列の長さ 1 ~ 10
- age      : 必須, 最低値: 0 ~ 最大値: 130
*/
func (u *User) validates() error {
	return validation.ValidateStruct(
		u,
		validation.Field(
			&u.FirstName,
			validation.Required,
			validation.RuneLength(1, 10),
		),
		validation.Field(
			&u.LastName,
			validation.Required,
			validation.RuneLength(1, 10),
		),
		validation.Field(
			&u.Age,
			validation.Required,
			validation.Min(0),
			validation.Max(130),
		),
	)
}

func NewUser(firstname, lastname, nickname string, age int) (*User, error) {
	user := &User{
		FirstName: firstname,
		LastName:  lastname,
		NickName:  nickname,
		Age:       age,
	}

	if err := user.validates(); err != nil {
		err = fmt.Errorf("入力情報が不適切です %w", err)
		return nil, err
	}

	return user, nil
}
