package main

import (
	"github.com/labstack/echo"
	"gitlab.com/martha.sutopo/workshop/handler"
)

func main() {
	e := echo.New()
	e.POST("/login", handler.Login())
	e.Logger.Fatal(e.Start(":5000"))
}
