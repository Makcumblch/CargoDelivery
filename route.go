package cargodelivery

type RouteSettings struct {
	Count   uint    `json:"count" binding:"required"`
	EvCount uint    `json:"evCount" binding:"required"`
	TMax    float32 `json:"tMax" binding:"required"`
	TMin    float32 `json:"tMin" binding:"required"`
}

type DeliveryTaskData struct {
	Cars    []Car
	Clients []ClientOrders
}

type Route struct {
	Clients   []Client    `json:"clients"`
	Waypoints [][]float32 `json:"waypoints"`
}

type Item struct {
	Cargo  Cargo  `json:"cargo"`
	Client Client `json:"client"`
}

type CarRoute struct {
	Car
	Route []Route `json:"route"`
	Items []Item  `json:"items"`
}

type RouteSolution struct {
	CarsRouteSolution []CarRoute `json:"carsRouteSolution"`
	RoutingCost       float32    `json:"routeCost"`
	PackingCost       float32    `json:"packingCost"`
}
