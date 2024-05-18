package userservice

import (
	"first-project/db/model/web"
	"first-project/helper"
)

type UserServiceInterface interface {
	Create(req web.UserReq) (helper.CustomResponse, error)
}
