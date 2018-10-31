package controllers

import (
	"Clean-Architecture/interfaces"
	"Clean-Architecture/payload"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type OrderContoroller struct {
	interfaces.IOrderService
}

func (controller OrderContoroller) GetOrder(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		//
	}

	order, err := controller.GetOrderById(intId)

	json.NewEncoder(res).Encode(payloads.ScuccessGetOrder{Code: http.StatusOK, Response: order})

}

func (controller OrderContoroller) GetAllOrders(res http.ResponseWriter, req *http.Request) {
	orders, err := controller.GetOrders()
	if err != nil {

	}
	json.NewEncoder(res).Encode(payloads.SuccessGetAllTodo{Code: http.StatusOK, Response: orders})
}
