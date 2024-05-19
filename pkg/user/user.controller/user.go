package usercontroller

import (
	"first-project/db/model/web"
	"first-project/helper"
	userservice "first-project/pkg/user/user.service"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
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

func (uC *UserController) Login(c echo.Context) error {
	loginUser := new(web.UserLoginReq)

	if err := c.Bind(loginUser); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(loginUser); err != nil {
		return err
	}

	userLogin, errLogin := uC.Service.Login(loginUser.Email, loginUser.Password)

	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errLogin.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Login Success", userLogin))
}

func (uC *UserController) GetId(c echo.Context) error {
	// authHeader := c.Request().Header.Get("Authorization")
	// token, errToken := helper.ValidToken(authHeader)

	// if errToken != nil {
	// 	return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errToken.Error(), nil))
	// }

	userToken := c.Get("user").(*jwt.Token)
	claims, _ := userToken.Claims.(*helper.CustomClaims)

	getUser, errUser := uC.Service.GetID(claims.UserID)

	if errUser != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errUser.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "User Detail", getUser))
}

func (uC *UserController) Update(c echo.Context) error {
	updateUser := new(web.UserUpdateReq)
	id, errId := strconv.Atoi(c.Param("id"))

	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errId.Error(), nil))
	}

	if err := c.Bind(updateUser); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(updateUser); err != nil {
		return err
	}

	saveUpdate, errUpdate := uC.Service.Update(id, *updateUser)

	if errUpdate != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success Update", saveUpdate))
}
