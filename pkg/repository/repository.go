package repository

import (
	"time"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/jmoiron/sqlx"
)

type IAuthorization interface {
	CreateUser(user cargodelivery.User) (int, error)
	GetUser(username string) (cargodelivery.User, error)
	GetUserById(userId int) (cargodelivery.User, error)
}

type IProject interface {
	GetUserProject(userId, projectId int) (cargodelivery.Project, error)
	CreateProject(userId int, project cargodelivery.Project, access string) (int, error)
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
	CreateOrder(clientId int, order cargodelivery.Order) (int, error)
	GetAllOrders(clientId int) ([]cargodelivery.Order, error)
	GetOrderById(clientId, orderId int) (cargodelivery.Order, error)
	DeleteOrder(clientId, orderId int) error
	UpdateOrder(clientId, orderId int, input cargodelivery.UpdateOrder) error
	GetOrdersAndCargos(clientId int) ([]cargodelivery.OrderCargo, error)
}

type IDepo interface {
	CreateDepo(projectId int, depo cargodelivery.Depo) (int, error)
	GetDepo(projectId int) (cargodelivery.Client, error)
	UpdateDepo(projectId int, input cargodelivery.UpdateDepo) error
}

type IRoute interface {
	CreateRoute(projectId int, solutionToDb cargodelivery.SolutionToDb, date time.Time) (int, error)
	GetAllRoutes(projectId int) ([]cargodelivery.Routedb, error)
	GetRouteById(projectId, routeId int) (cargodelivery.Routedb, error)
	DeleteRoute(projectId, routeId int) error
}

type IOSM interface {
	GetDistanceMatrix(clients []cargodelivery.Client) ([][]float32, error)
	GetRoutePoints(clients []cargodelivery.Client) ([][]float32, error)
}

type Repository struct {
	IAuthorization
	IProject
	ICar
	ICargo
	IClient
	IOrder
	IRoute
	IOSM
	IDepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IAuthorization: NewAuthPostgres(db),
		IProject:       NewProjectPostgres(db),
		ICar:           NewCarPostgres(db),
		ICargo:         NewCargoPostgres(db),
		IClient:        NewClientPostgres(db),
		IOrder:         NewOrderPostgres(db),
		IRoute:         NewRoutePostgres(db),
		IOSM:           NewOSMRepo(),
		IDepo:          NewDepoPostgres(db),
	}
}
