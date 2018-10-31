package payloads

import (
	"Clean-Architecture/models"
)

type ScuccessGetOrder struct {
	Code     int
	Response models.Order
}

type SuccessGetAllTodo struct {
	Code     int
	Response []models.Order
}
