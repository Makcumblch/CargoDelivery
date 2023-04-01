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

func (o *OrderService) GetAllOrders(clientId int) ([]cargodelivery.Order, error) {
	return o.repoOrder.GetAllOrders(clientId)
}

func (o *OrderService) GetOrderById(clientId, orderId int) (cargodelivery.Order, error) {
	return o.repoOrder.GetOrderById(clientId, orderId)
}

func (o *OrderService) DeleteOrder(clientId, orderId int) error {
	return o.repoOrder.DeleteOrder(clientId, orderId)
}

func (o *OrderService) UpdateOrder(clientId, orderId int, input cargodelivery.UpdateOrder) error {
	if err := input.ValidateUpdateOrder(); err != nil {
		return err
	}
	return o.repoOrder.UpdateOrder(clientId, orderId, input)
}
