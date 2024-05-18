package main

import (
	"first-project/db/connection"
	"first-project/helper"
	bicyclecontroller "first-project/pkg/bicycle/bicycle.controller"
	bicyclerepository "first-project/pkg/bicycle/bicycle.repository"
	bicycleservice "first-project/pkg/bicycle/bicycle.service"
	usercontroller "first-project/pkg/user/user.controller"
	userrepository "first-project/pkg/user/user.repository"
	userservice "first-project/pkg/user/user.service"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
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
	token := helper.NewTokenUseCase()

	// user
	userRepo := userrepository.NewUserRepo(DB)
	userService := userservice.NewUserService(userRepo, token)
	userController := usercontroller.NewUserController(userService)

	// bicycle
	bicycleRepo := bicyclerepository.NewBicycleRepo(DB)
	bicycleService := bicycleservice.NewBicycleService(bicycleRepo)
	bicycleController := bicyclecontroller.NewBicycleController(bicycleService)

	server := echo.New()
	server.Validator = &CustomValidator{validator: validator.New()}
	server.HTTPErrorHandler = helper.ValidateBind

	apiV1 := server.Group("api/v1/user")
	// router user
	apiV1.POST("/register", userController.Create)
	apiV1.POST("/login", userController.Login)
	apiV1.GET("/history", userController.GetId, JWTProtect())

	// router bicycle
	server.POST("/bicycle/created", bicycleController.Create)
	// start
	server.Logger.Fatal(server.Start(":8080"))
}

func JWTProtect() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.CustomClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, helper.ResponseClient(http.StatusUnauthorized, "login needed", nil))
		},
	})
}
