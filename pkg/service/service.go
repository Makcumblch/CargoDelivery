package service

import (
	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/Makcumblch/CargoDelivery/pkg/repository"
)

type IAuthorization interface {
	CreateUser(user cargodelivery.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
	GetUserById(userId int) (cargodelivery.User, error)
}

type IProject interface {
	GetUserProject(userId, projectId int) (cargodelivery.Project, error)
	CreateProject(userId int, project cargodelivery.Project) (int, error)
	GetAllProjects(userId int) ([]cargodelivery.Project, error)
	GetProjectById(userId, projectId int) (cargodelivery.Project, error)
	DeleteProject(userId, projectId int) error
	UpdateProject(userId, projectId int, input cargodelivery.UpdateProject) error
}

type ICar interface {
	CreateCar(projectId int, car cargodelivery.Car) (int, error)
	GetAllCars(projectId int) ([]cargodelivery.Car, error)
	GetCarById(projectId, carId int) (cargodelivery.Car, error)
	DeleteCar(projectId, carId int) error
	UpdateCar(projectId, carId int, input cargodelivery.UpdateCar) error
}

type ICargo interface {
	CreateCargo(projectId int, cargo cargodelivery.Cargo) (int, error)
	GetAllCargos(projectId int) ([]cargodelivery.Cargo, error)
	GetCargoById(projectId, cargoId int) (cargodelivery.Cargo, error)
	DeleteCargo(projectId, cargoId int) error
	UpdateCargo(projectId, cargoId int, input cargodelivery.UpdateCargo) error
}

type IClient interface {
	CreateClient(projectId int, client cargodelivery.Client) (int, error)
	GetAllClients(projectId int) ([]cargodelivery.Client, error)
	GetClientById(projectId, clientId int) (cargodelivery.Client, error)
	DeleteClient(projectId, clientId int) error
	UpdateClient(projectId, clientId int, input cargodelivery.UpdateClient) error
}

type IOrder interface {
	CreateOrder(projectId, clientId int, order cargodelivery.Order) (int, error)
	GetAllOrders(clientId int) ([]cargodelivery.Order, error)
	GetOrderById(clientId, orderId int) (cargodelivery.Order, error)
	DeleteOrder(clientId, orderId int) error
	UpdateOrder(clientId, orderId int, input cargodelivery.UpdateOrder) error
}

type IDepo interface {
	CreateDepo(projectId int, depo cargodelivery.Depo) (int, error)
	GetDepo(projectId int) (cargodelivery.Client, error)
	UpdateDepo(projectId int, input cargodelivery.UpdateDepo) error
}

type IRoute interface {
	CreateRoute(projectId int, settingsRoute cargodelivery.RouteSettings) (cargodelivery.RouteSolution, error)
}

type Service struct {
	IAuthorization
	IProject
	ICar
	ICargo
	IClient
	IOrder
	IRoute
	IDepo
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		IAuthorization: NewAuthService(repos.IAuthorization),
		IProject:       NewProjectService(repos.IProject),
		ICar:           NewCarService(repos.ICar),
		ICargo:         NewCargoService(repos.ICargo),
		IClient:        NewClientService(repos.IClient),
		IOrder:         NewOrderService(repos.IOrder, repos.ICargo),
		IDepo:          NewDepoService(repos.IDepo),
		IRoute:         NewRouteService(repos.ICar, repos.IClient, repos.IOrder, repos.IRoute, repos.IDepo, repos.IOSM),
	}
}
