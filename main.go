package main

import (
	"backend-github-trending/handler"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", handler.Welcome)

	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)

	e.Logger.Fatal(e.Start(":3000"))
}
