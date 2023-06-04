package service

import (
	"encoding/json"
	"time"

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
		return taskData, cargodelivery.ErrCreateRouteDepo
	}
	depo.Name = "Депо"

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

func (r *RouteService) setRoutePoints(bestSolution cargodelivery.RouteSolution) (cargodelivery.RouteSolution, error) {
	for i := 0; i < len(bestSolution.CarsRouteSolution); i++ {
		clients := make([]cargodelivery.Client, 0)
		for _, cl := range bestSolution.CarsRouteSolution[i].Route.Clients {
			clients = append(clients, cargodelivery.Client{Id: cl.Id, Address: cl.Address, CoordX: cl.CoordX, CoordY: cl.CoordY, Name: cl.Name})
		}
		polyline, err := r.osm.GetRoutePoints(clients)
		if err != nil {
			return cargodelivery.RouteSolution{}, err
		}
		bestSolution.CarsRouteSolution[i].Route.Polyline = polyline
	}
	return bestSolution, nil
}

func (r *RouteService) CreateRoute(projectId int, settingsRoute cargodelivery.RouteSettings) (cargodelivery.RouteResponse, error) {

	taskData, err := getTaskDeliveryData(projectId, r.carRepo, r.clientRepo, r.orderRepo, r.depoRepo)
	if err != nil {
		return cargodelivery.RouteResponse{}, err
	}

	distanceMatrix, err := r.getDistanceMatrix(taskData)
	if err != nil {
		return cargodelivery.RouteResponse{}, err
	}

	deliverySolution, err := solutiondeliverytask.GetDeliverySolution(taskData, distanceMatrix, settingsRoute)
	if err != nil {
		return cargodelivery.RouteResponse{}, err
	}

	deliverySolution, err = r.setRoutePoints(deliverySolution)
	if err != nil {
		return cargodelivery.RouteResponse{}, err
	}

	var clients = make([]cargodelivery.Client, 0)
	for _, client := range taskData.Clients {
		clients = append(clients, client.Client)
	}

	date := time.Now()

	solutionToDb := cargodelivery.SolutionToDb{Solution: deliverySolution, Depo: taskData.Clients[0].Client, Clients: clients}

	id, err := r.repo.CreateRoute(projectId, solutionToDb, date)
	if err != nil {
		return cargodelivery.RouteResponse{}, err
	}

	carsRoutes := make([]cargodelivery.CarRouteResponse, 0)
	for _, r := range deliverySolution.CarsRouteSolution {
		carRoute := cargodelivery.CarRouteResponse{Car: r.Car, Polyline: r.Route.Polyline}
		carsRoutes = append(carsRoutes, carRoute)
	}

	return cargodelivery.RouteResponse{Id: id, Date: date, Distance: deliverySolution.Distance, Fuel: deliverySolution.Fuel, Packing: deliverySolution.PackingCost, Clients: solutionToDb.Clients, CarsRoutes: carsRoutes}, nil
}

func (r *RouteService) GetAllRoutes(projectId int) ([]cargodelivery.RouteResponse, error) {
	routeResponse := make([]cargodelivery.RouteResponse, 0)
	routes, err := r.repo.GetAllRoutes(projectId)
	if err != nil {
		return make([]cargodelivery.RouteResponse, 0), err
	}
	for _, route := range routes {
		var response cargodelivery.RouteResponse

		response.Id = route.Id
		response.Date = route.DT

		var solution cargodelivery.SolutionToDb
		err = json.Unmarshal(route.Solution, &solution)
		if err != nil {
			return make([]cargodelivery.RouteResponse, 0), err
		}
		response.Clients = solution.Clients
		if len(response.Clients) > 0 {
			response.Clients[0].Name = "Депо"
		}
		response.Distance = solution.Solution.Distance
		response.Fuel = solution.Solution.Fuel
		response.Packing = solution.Solution.PackingCost
		response.CarsRoutes = make([]cargodelivery.CarRouteResponse, 0)
		for _, r := range solution.Solution.CarsRouteSolution {
			carRoute := cargodelivery.CarRouteResponse{Car: r.Car, Polyline: r.Route.Polyline}
			response.CarsRoutes = append(response.CarsRoutes, carRoute)
		}

		routeResponse = append(routeResponse, response)
	}

	return routeResponse, nil
}

// func (r *RouteService) GetRouteById(projectId, routeId int) (cargodelivery.RouteSolution, error) {

// }

func (r *RouteService) DeleteRoute(projectId, routeId int) error {
	return r.repo.DeleteRoute(projectId, routeId)
}
