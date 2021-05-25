package di

import (
	usecase "github.com/Ryutaro95/application/usecase/user"
	infra "github.com/Ryutaro95/infra/user"
	"github.com/Ryutaro95/presentation/handler"
)

func InjectUser() (userHandler handler.UserHandler) {
	userRepo := infra.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler = handler.NewUserHandler(userUsecase)

	return
}
