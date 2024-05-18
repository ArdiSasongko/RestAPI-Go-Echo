package bicyclerepository

import "first-project/db/model/domain"

type BicycleControllerInterface interface {
	Create(bicycle domain.Bicycles) (domain.Bicycles, error)
}
