package web

type OrderReq struct {
	Quantity   int `validate:"required" json:"quantity"`
	TotalPrice int `validate:"required" json:"total_price"`
}
