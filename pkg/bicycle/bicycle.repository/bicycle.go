package bicyclerepository

import (
	"errors"
	"first-project/db/model/domain"

	"gorm.io/gorm"
)

type BicycleRepo struct {
	DB *gorm.DB
}

func NewBicycleRepo(db *gorm.DB) *BicycleRepo {
	return &BicycleRepo{
		DB: db,
	}
}

func (bR *BicycleRepo) Create(bicycle domain.Bicycles) (domain.Bicycles, error) {
	if err := bR.DB.Create(&bicycle).Error; err != nil {
		return domain.Bicycles{}, errors.New("failed to create bicycle")
	}

	return bicycle, nil
}
