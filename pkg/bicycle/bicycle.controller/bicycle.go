package bicyclecontroller

import (
	"first-project/db/model/web"
	"first-project/helper"
	bicycleservice "first-project/pkg/bicycle/bicycle.service"
	orderservice "first-project/pkg/order/order.service"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type BicycleController struct {
	Service      bicycleservice.BicycleServiceInterface
	orderService orderservice.OrderServiceInterface
}

func NewBicycleController(service bicycleservice.BicycleServiceInterface, order orderservice.OrderServiceInterface) *BicycleController {
	return &BicycleController{
		Service:      service,
		orderService: order,
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

func (bC *BicycleController) GetBicycle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, errData := bC.Service.GetBicycle(id)

	if errData != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errData.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Bicycle Found", data))
}

func (bC *BicycleController) GetBicycles(c echo.Context) error {
	data, errData := bC.Service.GetBicycles()

	if errData != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errData.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Bicycle Found", data))
}

func (bC *BicycleController) BuyBicycle(c echo.Context) error {
	newOrder := new(web.OrderReq)
	id, _ := strconv.Atoi(c.Param("id"))
	// authHeader := c.Request().Header.Get("Authorization")
	// token, errToken := helper.ValidToken(authHeader)

	// if errToken != nil {
	// 	return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errToken.Error(), nil))
	// }
	userId := c.Get("user").(*jwt.Token)
	claims, _ := userId.Claims.(*helper.CustomClaims)

	if err := c.Bind(newOrder); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(newOrder); err != nil {
		return err
	}

	createOrder, errOrder := bC.orderService.Create(claims.UserID, id, *newOrder)

	if errOrder != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errOrder.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success Buy Bicycle", createOrder))
}
