package cargodelivery

import (
	"errors"
	"time"
)

type RouteSettings struct {
	EvCount     uint    `json:"evCount" binding:"required"`
	TMax        float32 `json:"tMax" binding:"required"`
	TMin        float32 `json:"tMin" binding:"required"`
	PackingType *bool   `json:"packingType" binding:"required"`
}

type DeliveryTaskData struct {
	Cars    []Car
	Clients []ClientOrders
}

type ClientRoute struct {
	Client
	Index int `json:"index"`
}

type Route struct {
	Clients  []ClientRoute `json:"clients"`
	Polyline [][]float32   `json:"polyline"`
}

type Position struct {
	X float32 `json:"x_pos"`
	Y float32 `json:"y_pos"`
	Z float32 `json:"z_pos"`
}

type Item struct {
	Cargo    Cargo    `json:"cargo"`
	Client   Client   `json:"client"`
	Position Position `json:"position"`
}

type CarRoute struct {
	Car
	Route            Route    `json:"route"`
	Items            [][]Item `json:"items"`
	FreeVolume       float32  `json:"freeVolume"`
	FreeLoadCapacity float32  `json:"freeLoadCapacity"`
}

type RouteSolution struct {
	CarsRouteSolution []CarRoute `json:"carsRouteSolution"`
	Distance          float32    `json:"distance"`
	Fuel              float32    `json:"fuel"`
	PackingCost       float32    `json:"packingCost"`
}

type OSMRTableResponse struct {
	Distances [][]float32 `json:"distances"`
}

type Coords struct {
	Coordinates [][]float32 `json:"coordinates"`
}

type Geometry struct {
	Geometry Coords `json:"geometry"`
}

type OSMRRoteResponse struct {
	Routes []Geometry `json:"routes"`
}

type CarRouteResponse struct {
	Car
	Polyline [][]float32 `json:"polyline"`
}

type RouteResponse struct {
	Id         int                `json:"id"`
	Date       time.Time          `json:"date"`
	Fuel       float32            `json:"fuel"`
	Distance   float32            `json:"distance"`
	Packing    float32            `json:"packing"`
	Clients    []Client           `json:"clients"`
	CarsRoutes []CarRouteResponse `json:"cars_routes"`
}

type Routedb struct {
	Id        int       `json:"id" db:"id"`
	ProjectId int       `json:"project_id" db:"project_id"`
	DT        time.Time `json:"dt" db:"dt"`
	Solution  []byte    `json:"solution" db:"solution"`
}

type SolutionToDb struct {
	Solution RouteSolution `json:"solution"`
	Depo     Client        `json:"depo"`
	Clients  []Client      `json:"clients"`
}

var ErrCreateRouteFewCars = errors.New("недостаточно ТС")
var ErrCreateRouteCars = errors.New("нет транспортных средств")
var ErrCreateRouteClient = errors.New("нет клиентов")
var ErrCreateRouteDepo = errors.New("не установлено депо")
