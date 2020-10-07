package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Name": "Quan",
		"Age": "2",
	})
}
func HandleSignUp(c echo.Context) error {
	type User struct {
		Email string `json:"email"`
		Name string `json:"name"`
	}
	user := User{
		Email: "123",
		Name: "Quan",
	}
	return c.JSON(http.StatusOK, user)
}