package orderrepository

import (
	"first-project/db/model/domain"

	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{
		DB: db,
	}
}

func (oR *OrderRepo) Create(order domain.Orders) (domain.Orders, error) {
	if err := oR.DB.Create(&order).Error; err != nil {
		return domain.Orders{}, err
	}

	return order, nil
}
