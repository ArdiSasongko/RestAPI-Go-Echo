package bicycleservice

import (
	"first-project/db/model/domain"
	"first-project/db/model/web"
	"first-project/helper"
	bicyclerepository "first-project/pkg/bicycle/bicycle.repository"
)

type BicycleService struct {
	Repo bicyclerepository.BicycleControllerInterface
}

func NewBicycleService(repo bicyclerepository.BicycleControllerInterface) *BicycleService {
	return &BicycleService{
		Repo: repo,
	}
}

func (bS *BicycleService) Create(req web.BicycleReq) (helper.CustomResponse, error) {
	newBicycle := domain.Bicycles{
		Name:     req.Name,
		Price:    req.Price,
		Quantity: req.Quantity,
	}

	saveBicycle, errBicycle := bS.Repo.Create(newBicycle)

	if errBicycle != nil {
		return nil, errBicycle
	}

	data := helper.CustomResponse{
		"bicycle_id": saveBicycle.BicycleID,
		"name":       saveBicycle.Name,
		"price":      saveBicycle.Price,
		"quantity":   saveBicycle.Quantity,
	}

	return data, nil
}
