package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/Makcumblch/CargoDelivery/pkg/repository"
	solutiondeliverytask "github.com/Makcumblch/CargoDelivery/pkg/service/solutionDeliveryTask"
	"github.com/spf13/viper"
)

type RouteService struct {
	carRepo    repository.ICar
	clientRepo repository.IClient
	orderRepo  repository.IOrder
	repo       repository.IRoute
	osm        repository.IOSM
}

func NewRouteService(carRepo repository.ICar, clientRepo repository.IClient, orderRepo repository.IOrder, repo repository.IRoute, osm repository.IOSM) *RouteService {
	return &RouteService{carRepo: carRepo, clientRepo: clientRepo, orderRepo: orderRepo, repo: repo, osm: osm}
}

func getTaskDeliveryData(projectId int, carRepo repository.ICar, clientRepo repository.IClient, orderRepo repository.IOrder) (cargodelivery.DeliveryTaskData, error) {
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

	var clientsOrders []cargodelivery.ClientOrders
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

type OSMRTableResponse struct {
	Distances [][]float32 `json:"distances"`
}

func getDistanceMatrix(taskData cargodelivery.DeliveryTaskData) ([][]float32, error) {
	coords := make([]string, 0)
	for _, client := range taskData.Clients {
		coords = append(coords, fmt.Sprintf("%f,%f", client.CoordX, client.CoordY))
	}

	resp, err := http.Get(fmt.Sprintf("http://%s:%s/table/v1/driving/%s?annotations=distance", viper.GetString("osmr.addres"), viper.GetString("osmr.port"), strings.Join(coords, ";")))
	if err != nil {
		return make([][]float32, 0), err
	}

	var result OSMRTableResponse
	json.NewDecoder(resp.Body).Decode(&result)

	return result.Distances, nil
}

func (r *RouteService) CreateRoute(projectId int, settingsRoute cargodelivery.RouteSettings) (cargodelivery.RouteSolution, error) {

	taskData, err := getTaskDeliveryData(projectId, r.carRepo, r.clientRepo, r.orderRepo)
	if err != nil {
		return cargodelivery.RouteSolution{}, err
	}

	distanceMatrix, err := getDistanceMatrix(taskData)
	if err != nil {
		return cargodelivery.RouteSolution{}, err
	}

	deliverySolution, err := solutiondeliverytask.GetDeliverySolution(taskData, distanceMatrix, settingsRoute)
	if err != nil {
		return cargodelivery.RouteSolution{}, err
	}

	return deliverySolution, nil
}
