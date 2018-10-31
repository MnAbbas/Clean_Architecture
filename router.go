package main

import (
	"sync"

	"github.com/go-chi/chi"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {
	mycontrolers := ServiceContainer().InjectOrderController()
	r := chi.NewRouter()
	r.HandleFunc("/order/{id}", mycontrolers.GetOrder)
	r.HandleFunc("/orders", mycontrolers.GetOrders)

	return r
}

var (
	instace    *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if instace == nil {
		routerOnce.Do(func() {
			instace = &router{}
		})
	}
	return instace
}
