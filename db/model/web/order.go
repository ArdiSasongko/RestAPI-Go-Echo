package web

type OrderReq struct {
	UserIDFK    int `validate:"required" json:"user_id_fK"`
	BicycleIDFK int `validate:"required" json:"bicycle_id_fk"`
	Quantity    int `validate:"required" json:"quantity"`
	TotalPrice  int `validate:"required" json:"total_price"`
}
