package bicyclecontroller

import "github.com/labstack/echo/v4"

type BicycleControllerInterface interface {
	Create(c echo.Context) error
}
