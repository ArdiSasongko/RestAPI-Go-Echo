package app

import (
	usercontroller "first-project/pkg/user/user.controller"

	"github.com/labstack/echo/v4"
)

func UserRouter(e *echo.Echo, userController usercontroller.UserControllerInterface) {
	apiV1 := e.Group("api/v1/user")
	apiV1.POST("/register", userController.Create)
	apiV1.POST("/login", userController.Login)
	apiV1.GET("/history", userController.GetId, JWTProtect())
	apiV1.PUT("/:id/update", userController.Update, JWTProtect(), AccessUser())
	apiV1.DELETE("/:id/delete", userController.Delete, JWTProtect(), AccessUser())
}
