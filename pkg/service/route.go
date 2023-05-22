package service

import (
	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/Makcumblch/CargoDelivery/pkg/repository"
	solutiondeliverytask "github.com/Makcumblch/CargoDelivery/pkg/service/solutionDeliveryTask"
)

type RouteService struct {
	carRepo    repository.ICar
	clientRepo repository.IClient
	orderRepo  repository.IOrder
	depoRepo   repository.IDepo
	repo       repository.IRoute
	osm        repository.IOSM
}

func NewRouteService(carRepo repository.ICar, clientRepo repository.IClient, orderRepo repository.IOrder, repo repository.IRoute, depoRepo repository.IDepo, osm repository.IOSM) *RouteService {
	return &RouteService{carRepo: carRepo, clientRepo: clientRepo, orderRepo: orderRepo, repo: repo, depoRepo: depoRepo, osm: osm}
}

func getTaskDeliveryData(projectId int, carRepo repository.ICar, clientRepo repository.IClient, orderRepo repository.IOrder, depoRepo repository.IDepo) (cargodelivery.DeliveryTaskData, error) {
	var taskData cargodelivery.DeliveryTaskData

	cars, err := carRepo.GetAllCars(projectId)
	if err != nil {
		return taskData, err
	}
	taskData.Cars = cars

	clients, err := clientRepo.GetAllClients(projectId)
	if err != nil {
		return taskData, err
	}

	depo, err := depoRepo.GetDepo(projectId)
	if err != nil {
		return taskData, err
	}

	var clientsOrders []cargodelivery.ClientOrders
	clientsOrders = append(clientsOrders, cargodelivery.ClientOrders{Client: depo, Orders: make([]cargodelivery.OrderCargo, 0)})
	for _, client := range clients {
		orderCargos, err := orderRepo.GetOrdersAndCargos(client.Id)
		if err != nil {
			return taskData, err
		}
		clientsOrders = append(clientsOrders, cargodelivery.ClientOrders{Client: client, Orders: orderCargos})
	}
	taskData.Clients = clientsOrders

	return taskData, nil
}

func (r *RouteService) getDistanceMatrix(taskData cargodelivery.DeliveryTaskData) ([][]float32, error) {
	clients := make([]cargodelivery.Client, 0)
	for _, cl := range taskData.Clients {
		clients = append(clients, cargodelivery.Client{Id: cl.Id, Address: cl.Address, CoordX: cl.CoordX, CoordY: cl.CoordY, Name: cl.Name})
	}
	return r.osm.GetDistanceMatrix(clients)
}

func (r *RouteService) setRoutePoints(bestSolution *cargodelivery.RouteSolution) {
	for _, car := range bestSolution.CarsRouteSolution {
		clients := make([]cargodelivery.Client, 0)
		for _, cl := range car.Route.Clients {
			clients = append(clients, cargodelivery.Client{Id: cl.Id, Address: cl.Address, CoordX: cl.CoordX, CoordY: cl.CoordY, Name: cl.Name})
		}
		car.Route.Waypoints = r.osm.GetRoutePoints(clients)
	}
}

func (r *RouteService) CreateRoute(projectId int, settingsRoute cargodelivery.RouteSettings) (cargodelivery.RouteSolution, error) {

	taskData, err := getTaskDeliveryData(projectId, r.carRepo, r.clientRepo, r.orderRepo, r.depoRepo)
	if err != nil {
		return cargodelivery.RouteSolution{}, err
	}

	distanceMatrix, err := r.getDistanceMatrix(taskData)
	if err != nil {
		return cargodelivery.RouteSolution{}, err
	}

	deliverySolution, err := solutiondeliverytask.GetDeliverySolution(taskData, distanceMatrix, settingsRoute)
	if err != nil {
		return cargodelivery.RouteSolution{}, err
	}

	r.setRoutePoints(&deliverySolution)

	return deliverySolution, nil
}
