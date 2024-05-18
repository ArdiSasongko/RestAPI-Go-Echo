package orderservice

import (
	"errors"
	"first-project/db/model/domain"
	"first-project/db/model/web"
	"first-project/helper"
	bicyclerepository "first-project/pkg/bicycle/bicycle.repository"
	orderrepository "first-project/pkg/order/order.repository"
)

type OrderService struct {
	Repo        orderrepository.OrderRepositoryInterface
	Token       helper.TokenUseCaseInterface
	BicycleRepo bicyclerepository.BicycleRepositoryInterface
}

func NewOrderService(repo orderrepository.OrderRepositoryInterface, token helper.TokenUseCaseInterface, bicycleRepo bicyclerepository.BicycleRepositoryInterface) *OrderService {
	return &OrderService{
		Repo:        repo,
		Token:       token,
		BicycleRepo: bicycleRepo,
	}
}

func (oS *OrderService) Create(token string, id int, req web.OrderReq) (helper.CustomResponse, error) {
	tokenV, errToken := oS.Token.DecodeToken(token)
	if errToken != nil {
		return nil, errToken
	}
	claims, _ := tokenV.Claims.(*helper.CustomClaims)

	bicycle, errBicycle := oS.BicycleRepo.GetBicycle(id)

	if errBicycle != nil {
		return nil, errBicycle
	}

	totalPrice := bicycle.Price * req.Quantity

	if req.TotalPrice < totalPrice {
		return nil, errors.New("not enough money")
	}

	newOrder := domain.Orders{
		UserIDFK:    claims.UserID,
		BicycleIDFK: id,
		Quantity:    req.Quantity,
		TotalPrice:  totalPrice,
	}

	createOrder, errOrder := oS.Repo.Create(newOrder)

	if errOrder != nil {
		return nil, errOrder
	}

	data := helper.CustomResponse{
		"order_id":    createOrder.OrderID,
		"user_id":     createOrder.UserIDFK,
		"bicycle_id":  createOrder.BicycleIDFK,
		"quantity":    createOrder.Quantity,
		"total_price": createOrder.TotalPrice,
		"created_at":  createOrder.CreatedAt,
	}

	return data, nil
}
