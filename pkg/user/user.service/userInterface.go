package userservice

import (
	"first-project/db/model/entity"
	"first-project/db/model/web"
	"first-project/helper"
)

type UserServiceInterface interface {
	Create(req web.UserReq) (helper.CustomResponse, error)
	Login(email, password string) (helper.CustomResponse, error)
	GetID(id int) (entity.UserHistoryEntity, error)
	Update(id int, req web.UserUpdateReq) (helper.CustomResponse, error)
}
