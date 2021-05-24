package main

import (
	"github.com/Ryutaro95/presentation/router"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	router.InitRouting(e)
	e.Logger.Fatal(e.Start(":8080"))
}
