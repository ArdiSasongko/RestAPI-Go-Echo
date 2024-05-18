package domain

import "time"

type Orders struct {
	OrderID     int `gorm:"column:order_id;primaryKey;autoIncrement"`
	UserIDFK    int `gorm:"column:user_id_fk"`
	BicycleIDFK int `gorm:"column:bicycle_id_fk"`
	Quantity    int `gorm:"column:quantity"`
	TotalPrice  int `gorm:"column:total_price"`
	CreatedAt   time.Time
}
