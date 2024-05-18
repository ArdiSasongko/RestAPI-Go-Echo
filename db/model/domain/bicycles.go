package domain

import "time"

type Bicycles struct {
	BicycleID int       `gorm:"column:bicycle_id;primaryKey;autoIncrement"`
	Name      string    `gorm:"column:name"`
	Price     int       `gorm:"column:price"`
	Quantity  int       `gorm:"column:quantity"`
	Orders    []*Orders `gorm:"foreignKey:bicycle_id_fk;references:bicycle_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
