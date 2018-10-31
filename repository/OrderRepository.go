package repository

import (
	"Clean-Architecture/interfaces"
	"Clean-Architecture/models"
	"fmt"
)

type OrderRepository struct {
	interfaces.IDbHandeler
}

func (repo *OrderRepository) FetchOrder(id int) (models.Order, error) {
	rows, err := repo.Query(fmt.Sprintf("select * from order where orderid=%d", id))
	if err != nil {
		return models.Order{}, nil
	}
	var order models.Order
	rows.Next()
	rows.Scan(&order.Id, &order.Status)
	return order, nil
}

func (repo *OrderRepository) FetchAllOrder() ([]models.Order, error) {
	rows, err := repo.Query(fmt.Sprintf("select * from order"))
	if err != nil {
		return []models.Order{}, nil
	}
	var (
		order  models.Order
		orders []models.Order
	)
	for rows.Next() {
		rows.Scan(&order.Id, &order.Status)
		orders = append(orders, order)
	}
	return orders, nil
}
