package service

import (
	"fmt"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/Makcumblch/CargoDelivery/pkg/repository"
)

type OrderService struct {
	repoCargo repository.ICargo
	repoOrder repository.IOrder
}

func NewOrderService(repoOrder repository.IOrder, repoCargo repository.ICargo) *OrderService {
	return &OrderService{repoOrder: repoOrder, repoCargo: repoCargo}
}

func (o *OrderService) CreateOrder(projectId, clientId int, order cargodelivery.Order) (int, error) {
	_, err := o.repoCargo.GetCargoById(projectId, order.IdCargo)
	if err != nil {
		return 0, fmt.Errorf("cargo not found: %s", err.Error())
	}

	return o.repoOrder.CreateOrder(clientId, order)
}
