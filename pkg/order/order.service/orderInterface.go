package orderservice

import (
	"first-project/db/model/web"
	"first-project/helper"
)

type OrderServiceInterface interface {
	Create(userId int, id int, req web.OrderReq) (helper.CustomResponse, error)
}
