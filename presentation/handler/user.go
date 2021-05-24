package handler

import (
	"net/http"

	usecase "github.com/Ryutaro95/application/usecase/user"
	"github.com/labstack/echo"
)

type UserHandler interface {
	Post() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

type requestUser struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	NickName  string `json:"nickname"`
	Age       int    `json:"age"`
}
type responseUser struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	NickName  string `json:"nickname"`
	Age       int    `json:"age"`
}

func (uh *userHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		
		createdUser, err := uh.userUsecase.Create(req.FirstName, req.LastName, req.NickName, req.Age)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		res := &responseUser{
			FirstName: createdUser.FirstName,
			LastName: createdUser.LastName,
			NickName: createdUser.NickName,
			Age: createdUser.Age,
		}
		return c.JSON(http.StatusCreated, res)
	}
}
