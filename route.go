package cargodelivery

type RouteSettings struct {
	EvCount uint    `json:"evCount" binding:"required"`
	TMax    float32 `json:"tMax" binding:"required"`
	TMin    float32 `json:"tMin" binding:"required"`
}

type DeliveryTaskData struct {
	Cars    []Car
	Clients []ClientOrders
}

type ClientRoute struct {
	Client
	Index int
}

type Route struct {
	Clients   []ClientRoute `json:"clients"`
	Waypoints [][]float32   `json:"waypoints"`
}

type Position struct{
	X float32 `json:"x_pos"`
	Y float32 `json:"y_pos"`
	Z float32 `json:"z_pos"`
}

type Item struct {
	Cargo  Cargo  	  `json:"cargo"`
	Client Client 	  `json:"client"`
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
	RoutingCost       float32    `json:"routeCost"`
	PackingCost       float32    `json:"packingCost"`
}

type OSMRTableResponse struct {
	Distances [][]float32 `json:"distances"`
}
