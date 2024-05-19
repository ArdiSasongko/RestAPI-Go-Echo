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

	apiV1 := server.Group("api/v1/user")
	// router user
	apiV1.POST("/register", userController.Create)
	apiV1.POST("/login", userController.Login)
	apiV1.GET("/history", userController.GetId, JWTProtect())
	apiV1.PUT("/:id/update", userController.Update, JWTProtect(), AccessUser())
	apiV1.DELETE("/:id/delete", userController.Delete, JWTProtect(), AccessUser())

	// router bicycle
	server.POST("/bicycle/created", bicycleController.Create)
	server.GET("/bicycle/:id", bicycleController.GetBicycle, JWTProtect())
	server.POST("/bicycle/:id/buy", bicycleController.BuyBicycle, JWTProtect())
	server.GET("/bicycles", bicycleController.GetBicycles, JWTProtect())

	return server
}
