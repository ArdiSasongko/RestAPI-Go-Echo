package bicyclerepository

import "first-project/db/model/domain"

type BicycleRepositoryInterface interface {
	Create(bicycle domain.Bicycles) (domain.Bicycles, error)
	GetBicycle(id int) (domain.Bicycles, error)
	GetBicycles() ([]domain.Bicycles, error)
}
