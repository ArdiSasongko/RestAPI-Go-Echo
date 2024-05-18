package main

import (
	"first-project/db/connection"
	"first-project/helper"
	usercontroller "first-project/pkg/user/user.controller"
	userrepository "first-project/pkg/user/user.repository"
	userservice "first-project/pkg/user/user.service"
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

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}
	DB := connection.DBConn()

	userRepo := userrepository.NewUserRepo(DB)
	userService := userservice.NewUserService(userRepo)
	userController := usercontroller.NewUserController(userService)

	server := echo.New()
	server.Validator = &CustomValidator{validator: validator.New()}
	server.HTTPErrorHandler = helper.ValidateBind

	// router
	server.POST("/register", userController.Create)

	// start
	server.Logger.Fatal(server.Start(":8080"))
}
