package bicyclecontroller

import "github.com/labstack/echo/v4"

type BicycleControllerInterface interface {
	Create(c echo.Context) error
	GetBicycle(c echo.Context) error
	GetBicycles(c echo.Context) error
	BuyBicycle(c echo.Context) error
}
