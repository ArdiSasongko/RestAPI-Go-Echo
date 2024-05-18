package entity

import "time"

type OrderEntity struct {
	OrderID     int       `json:"order_id"`
	UserIDFK    int       `json:"user_id_fk"`
	BicycleIDFK int       `json:"bicycle_id_fk"`
	Quantity    int       `json:"quantity"`
	TotalPrice  int       `json:"total_price"`
	CreatedAt   time.Time `json:"created_at"`
}
