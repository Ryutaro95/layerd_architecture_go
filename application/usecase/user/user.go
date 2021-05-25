package usecase

import (
	"github.com/Ryutaro95/domain/model"
	repository "github.com/Ryutaro95/domain/repository/user"
)

type UserUsecase interface {
	Create(firstname, lastname, nickname string, age int) (*model.User, error)
	GetByID(userID int) (*model.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (uu *userUsecase) Create(firstname, lastname, nickname string, age int) (*model.User, error) {
	user, err := model.NewUser(firstname, lastname, nickname, age)
	if err != nil {
		return nil, err
	}

	createdUser := uu.userRepo.Create(user)
	return createdUser, nil
}

func (uu *userUsecase) GetByID(userID int) (*model.User, error) {
	user, err := uu.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
