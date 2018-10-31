package interfaces

import (
	"Clean-Architecture/models"
)

type IOrderRepository interface {
	FetchOrder(id int) (models.Order, error)
	FetchAll() ([]models.Order, error)
}
