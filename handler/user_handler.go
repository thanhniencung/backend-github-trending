package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user123": "Ryan",
		"email": "ryan@gmail.com",
	})
}

func HandleSignUp(c echo.Context) error {
	type User struct {
		RyanEmail string `json:"email"`
		FullName string `json:"name"`
		Age int `json:"age"`
	}

	user := User{
		RyanEmail:    "ryan@gmail.com",
		FullName: "Ryan Nguyen",
		Age: 90,
	}
	return c.JSON(http.StatusOK, user)
}
