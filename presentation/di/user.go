package di

import (
	usecase "github.com/Ryutaro95/application/usecase/user"
	repository "github.com/Ryutaro95/domain/repository/user"
	infra "github.com/Ryutaro95/infra/user"
	"github.com/Ryutaro95/presentation/handler"
)

func InjectUserRepository() repository.UserRepository {
	userRepo := infra.NewUserRepository()
	return userRepo
}

func InjectUserUsecase() usecase.UserUsecase {
	userRepo := InjectUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)
	return userUsecase
}

func InjectUserHandler() handler.UserHandler {
	userUsecase := InjectUserUsecase()
	userHandler := handler.NewUserHandler(userUsecase)

	return userHandler
}
