package router

import (
	"github.com/Ryutaro95/presentation/di"
	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo) {

	userHandler := di.InjectUser()
	e.POST("/users", userHandler.Post())
	e.GET("/user/:id", userHandler.GetByID())
}
