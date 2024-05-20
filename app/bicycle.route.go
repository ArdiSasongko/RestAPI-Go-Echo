package app

import (
	bicyclecontroller "first-project/pkg/bicycle/bicycle.controller"

	"github.com/labstack/echo/v4"
)

func BicycleRoute(e *echo.Echo, bicycleController bicyclecontroller.BicycleControllerInterface) {
	apiV1 := e.Group("api/v1/bicycle")
	apiV1.POST("/created", bicycleController.Create)
	apiV1.GET("/:id", bicycleController.GetBicycle, JWTProtect())
	apiV1.POST("/:id/buy", bicycleController.BuyBicycle, JWTProtect())
	apiV1.GET("/all", bicycleController.GetBicycles, JWTProtect())
}
