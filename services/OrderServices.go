package services

import (
	"Clean-Architecture/interfaces"
	"Clean-Architecture/models"
)

type OrderService struct {
	interfaces.IOrderService
}

func (orservice OrderService) GetOrderById(id int) (models.Order, error) {
	order, err := orservice.GetOrderById(id)
	if err != nil {
		return models.Order{}, nil
	}
	return order, nil
}

func (orservice OrderService) GetOrders() ([]models.Order, error) {
	orders, err := orservice.GetOrders()
	if err != nil {
		return []models.Order{}, nil
	}
	return orders, nil
}
