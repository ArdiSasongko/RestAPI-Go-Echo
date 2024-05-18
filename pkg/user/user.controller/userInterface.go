package usercontroller

import "github.com/labstack/echo/v4"

type UserControllerInterface interface {
	Create(c echo.Context) error
	Login(c echo.Context) error
}
