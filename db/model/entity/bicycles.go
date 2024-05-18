package entity

import (
	"first-project/db/model/domain"
	"time"
)

type BicycleEntity struct {
	BicycleID int       `json:"bicycle_id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}

func ToBicycleEntity(bicycle domain.Bicycles) BicycleEntity {
	return BicycleEntity{
		BicycleID: bicycle.BicycleID,
		Name:      bicycle.Name,
		Price:     bicycle.Price,
		Quantity:  bicycle.Quantity,
		CreatedAt: bicycle.CreatedAt,
	}
}

func ToBicycleListEntity(bicycles []domain.Bicycles) []BicycleEntity {
	dataBicycles := []BicycleEntity{}
	for _, v := range bicycles {
		dataBicycles = append(dataBicycles, ToBicycleEntity(v))
	}
	return dataBicycles
}
