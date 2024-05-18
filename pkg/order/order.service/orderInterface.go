package orderservice

import (
	"first-project/db/model/web"
	"first-project/helper"
)

type OrderServiceInterface interface {
	Create(token string, id int, req web.OrderReq) (helper.CustomResponse, error)
}
