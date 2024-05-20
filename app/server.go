package app

import (
	"first-project/helper"
	bicyclecontroller "first-project/pkg/bicycle/bicycle.controller"
	usercontroller "first-project/pkg/user/user.controller"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cV *CustomValidator) Validate(i interface{}) error {
	return cV.validator.Struct(i)
}

func InitialServer(userController usercontroller.UserControllerInterface, bicycleController bicyclecontroller.BicycleControllerInterface) *echo.Echo {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	server := echo.New()
	server.Validator = &CustomValidator{validator: validator.New()}
	server.HTTPErrorHandler = helper.ValidateBind

	UserRouter(server, userController)
	BicycleRoute(server, bicycleController)

	return server
}
