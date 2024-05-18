package web

type BicycleReq struct {
	Name     string `validate:"required" json:"name"`
	Price    int    `validate:"required" json:"price"`
	Quantity int    `validate:"required" json:"quantity"`
}

type BicycleUpdateReq struct {
	Price    int `json:"price"`
	Quantity int `json:"quantity"`
}
