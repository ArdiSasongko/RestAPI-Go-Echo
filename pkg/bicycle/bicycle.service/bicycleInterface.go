package bicycleservice

import (
	"first-project/db/model/entity"
	"first-project/db/model/web"
	"first-project/helper"
)

type BicycleServiceInterface interface {
	Create(req web.BicycleReq) (helper.CustomResponse, error)
	GetBicycle(id int) (entity.BicycleEntity, error)
	GetBicycles() ([]entity.BicycleEntity, error)
}
