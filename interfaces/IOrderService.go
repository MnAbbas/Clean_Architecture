package interfaces

import (
	"Clean-Architecture/models"
)

type IOrderService interface {
	GetOrderById(id int) (models.Order, error)
	GetOrders() ([]models.Order, error)
}
