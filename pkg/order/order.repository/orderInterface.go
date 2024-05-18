package orderrepository

import "first-project/db/model/domain"

type OrderRepositoryInterface interface {
	Create(order domain.Orders) (domain.Orders, error)
}
