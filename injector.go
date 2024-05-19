//go:build wireinject
// +build wireinject

package main

import (
	"first-project/app"
	"first-project/db/connection"
	"first-project/helper"
	bicyclecontroller "first-project/pkg/bicycle/bicycle.controller"
	bicyclerepository "first-project/pkg/bicycle/bicycle.repository"
	bicycleservice "first-project/pkg/bicycle/bicycle.service"
	orderrepository "first-project/pkg/order/order.repository"
	orderservice "first-project/pkg/order/order.service"
	usercontroller "first-project/pkg/user/user.controller"
	userrepository "first-project/pkg/user/user.repository"
	userservice "first-project/pkg/user/user.service"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var userSet = wire.NewSet(
	userrepository.NewUserRepo,
	wire.Bind(new(userrepository.UserRepoInterface), new(*userrepository.UserRepo)),
	helper.NewTokenUseCase,
	wire.Bind(new(helper.TokenUseCaseInterface), new(*helper.TokenUseCaseImpl)),
	userservice.NewUserService,
	wire.Bind(new(userservice.UserServiceInterface), new(*userservice.UserService)),
	usercontroller.NewUserController,
	wire.Bind(new(usercontroller.UserControllerInterface), new(*usercontroller.UserController)),
)

var orderSet = wire.NewSet(
	orderrepository.NewOrderRepo,
	wire.Bind(new(orderrepository.OrderRepositoryInterface), new(*orderrepository.OrderRepo)),
	orderservice.NewOrderService,
	wire.Bind(new(orderservice.OrderServiceInterface), new(*orderservice.OrderService)),
)

var bicycleSet = wire.NewSet(
	bicyclerepository.NewBicycleRepo,
	wire.Bind(new(bicyclerepository.BicycleRepositoryInterface), new(*bicyclerepository.BicycleRepo)),
	bicycleservice.NewBicycleService,
	wire.Bind(new(bicycleservice.BicycleServiceInterface), new(*bicycleservice.BicycleService)),
	bicyclecontroller.NewBicycleController,
	wire.Bind(new(bicyclecontroller.BicycleControllerInterface), new(*bicyclecontroller.BicycleController)),
)

func StartServer() *echo.Echo {
	wire.Build(
		connection.DBConn,
		userSet,
		orderSet,
		bicycleSet,
		app.InitialServer,
	)
	return nil
}
