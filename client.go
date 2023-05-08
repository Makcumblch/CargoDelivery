package cargodelivery

import "errors"

type Depo struct {
	Id      int     `json:"id" db:"id"`
	Address string  `json:"address" binding:"required" db:"address"`
	CoordX  float32 `json:"coordX" binding:"required" db:"coord_x"`
	CoordY  float32 `json:"coordY" binding:"required" db:"coord_y"`
}

type Client struct {
	Id      int     `json:"id" db:"id"`
	Address string  `json:"address" binding:"required" db:"address"`
	CoordX  float32 `json:"coordX" binding:"required" db:"coord_x"`
	CoordY  float32 `json:"coordY" binding:"required" db:"coord_y"`
	Name    string  `json:"name" binding:"required" db:"name"`
}

type ClientOrders struct {
	Client
	Orders []OrderCargo
}

type UpdateClient struct {
	Address *string  `json:"address"`
	Name    *string  `json:"name"`
	CoordX  *float32 `json:"coordX"`
	CoordY  *float32 `json:"coordY"`
}

type UpdateDepo struct {
	Address *string  `json:"address"`
	CoordX  *float32 `json:"coordX"`
	CoordY  *float32 `json:"coordY"`
}

func (i *UpdateClient) ValidateUpdateClient() error {
	if i.Address == nil && i.Name == nil && i.CoordX == nil && i.CoordY == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

func (i *UpdateDepo) ValidateUpdateDepo() error {
	if i.Address == nil && i.CoordX == nil && i.CoordY == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
