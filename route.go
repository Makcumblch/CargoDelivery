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

type RouteSolution struct {
	Data [][]float32
}
