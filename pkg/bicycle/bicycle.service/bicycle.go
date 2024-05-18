package bicycleservice

import (
	"first-project/db/model/domain"
	"first-project/db/model/entity"
	"first-project/db/model/web"
	"first-project/helper"
	bicyclerepository "first-project/pkg/bicycle/bicycle.repository"
)

type BicycleService struct {
	Repo bicyclerepository.BicycleRepositoryInterface
}

func NewBicycleService(repo bicyclerepository.BicycleRepositoryInterface) *BicycleService {
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

func (bS *BicycleService) GetBicycle(id int) (entity.BicycleEntity, error) {
	bicycle, errBicyle := bS.Repo.GetBicycle(id)

	if errBicyle != nil {
		return entity.BicycleEntity{}, errBicyle
	}

	// data := entity.BicycleEntity{
	// 	BicycleID: bicycle.BicycleID,
	// 	Name: bicycle.Name,
	// 	Price: bicycle.Price,
	// 	Quantity: bicycle.Quantity,
	// 	CreatedAt: bicycle.CreatedAt,
	// }

	return entity.ToBicycleEntity(bicycle), nil
}

func (bS *BicycleService) GetBicycles() ([]entity.BicycleEntity, error) {
	bicycles, errBicycles := bS.Repo.GetBicycles()

	if errBicycles != nil {
		return nil, errBicycles
	}

	return entity.ToBicycleListEntity(bicycles), nil
}
