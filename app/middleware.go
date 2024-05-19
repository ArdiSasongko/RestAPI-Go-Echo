package app

import (
	"first-project/helper"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

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

func AccessUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims, _ := user.Claims.(*helper.CustomClaims)
			idParam, errParam := strconv.Atoi(c.Param("id"))

			if errParam != nil {
				return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, "id parameter not valid", nil))
			}

			if claims.UserID != idParam {
				return c.JSON(http.StatusUnauthorized, helper.ResponseClient(http.StatusUnauthorized, "Unauthorized", nil))
			}

			return next(c)
		}
	}
}
