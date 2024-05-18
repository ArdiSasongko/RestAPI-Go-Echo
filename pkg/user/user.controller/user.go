package usercontroller

import (
	"first-project/db/model/web"
	"first-project/helper"
	userservice "first-project/pkg/user/user.service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Service userservice.UserServiceInterface
}

func NewUserController(service userservice.UserServiceInterface) *UserController {
	return &UserController{
		Service: service,
	}
}

func (uC *UserController) Create(c echo.Context) error {
	newUser := new(web.UserReq)

	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(newUser); err != nil {
		return err
	}

	saveUser, errSave := uC.Service.Create(*newUser)

	if errSave != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errSave.Error(), nil))
	}

	return c.JSON(http.StatusCreated, helper.ResponseClient(http.StatusCreated, "Success Created User", saveUser))
}
