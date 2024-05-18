package entity

import "first-project/db/model/domain"

type UserHistoryEntity struct {
	UserID int         `json:"user_id"`
	Name   string      `json:"name"`
	Email  string      `json:"email"`
	Orders interface{} `json:"orders"`
}

type UserEnity struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

func ToUserHistoryEntity(user domain.User) UserHistoryEntity {
	var orders []OrderEntity
	if len(user.Orders) != 0 {
		for _, v := range user.Orders {
			orders = append(orders, OrderEntity{
				OrderID:     v.OrderID,
				UserIDFK:    v.UserIDFK,
				BicycleIDFK: v.BicycleIDFK,
				Quantity:    v.Quantity,
				TotalPrice:  v.TotalPrice,
				CreatedAt:   v.CreatedAt,
			})
		}

		return UserHistoryEntity{
			UserID: user.UserID,
			Name:   user.Name,
			Email:  user.Email,
			Orders: orders,
		}
	}

	return UserHistoryEntity{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
		Orders: "Never Made an Order",
	}
}

func ToUserEntity(user domain.User) UserEnity {
	return UserEnity{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
	}
}

func ToUsersEntity(users []domain.User) []UserEnity {
	usersData := []UserEnity{}
	for _, v := range users {
		usersData = append(usersData, ToUserEntity(v))
	}
	return usersData
}
