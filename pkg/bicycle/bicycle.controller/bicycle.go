package bicyclecontroller

import (
	"first-project/db/model/web"
	"first-project/helper"
	bicycleservice "first-project/pkg/bicycle/bicycle.service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BicycleController struct {
	Service bicycleservice.BicycleServiceInterface
}

func NewBicycleController(service bicycleservice.BicycleServiceInterface) *BicycleController {
	return &BicycleController{
		Service: service,
	}
}

func (bC *BicycleController) Create(c echo.Context) error {
	newBicycle := new(web.BicycleReq)

	if err := c.Bind(newBicycle); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(newBicycle); err != nil {
		return err
	}

	saveBicycle, errBicycle := bC.Service.Create(*newBicycle)

	if errBicycle != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errBicycle.Error(), nil))
	}

	return c.JSON(http.StatusCreated, helper.ResponseClient(http.StatusCreated, "Success Created Bicycle", saveBicycle))
}
