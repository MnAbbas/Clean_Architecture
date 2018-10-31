package main

import (
	"Clean-Architecture/controllers"
	"Clean-Architecture/services"
	"database/sql"
	"gitclean-architecture/infrastructures"
	"gitclean-architecture/repositories"
	"sync"
)

type IServiceContainer interface {
	InjectOrderController() controllers.OrderContoroller
}

type kernal struct{}

func (k *kernal) InjectOrderController() controllers.OrderContoroller {
	sqlConn, err := sql.Open("mysql", "root:local@/abbas")
	MySQLHandler := &infrastructures.MySQLHandler{}
	MySQLHandler.Conn = sqlConn

	orderRepository := &repositories.orderRepository{MySQLHandler}
	orderService := &services.OrderService{orderRepository}
	orderController := controllers.OrderContoroller{orderService}

	return orderController

}

var (
	k           *kernal
	containOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containOnce.Do(func() {
			k = &kernal
		})
	}
	return k
}
