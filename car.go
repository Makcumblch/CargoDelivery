package cargodelivery

import "errors"

type Car struct {
	Id              int     `json:"id" db:"id"`
	Name            string  `json:"name" db:"name" binding:"required"`
	LoadCapacity    float32 `json:"loadCapacity" binding:"required" db:"load_capacity"`       // кг
	Width           float32 `json:"width" binding:"required" db:"width"`                      // м
	Height          float32 `json:"height" binding:"required" db:"height"`                    // м
	Length          float32 `json:"length" binding:"required" db:"length"`                    // м
	FuelConsumption float32 `json:"fuelConsumption" binding:"required" db:"fuel_consumption"` // л/100км
}

type UpdateCar struct {
	Name            *string  `json:"name"`
	LoadCapacity    *float32 `json:"loadCapacity"`    // кг
	Width           *float32 `json:"width"`           // м
	Height          *float32 `json:"height"`          // м
	Length          *float32 `json:"length"`          // м
	FuelConsumption *float32 `json:"fuelConsumption"` // л/100км
}

func (i *UpdateCar) ValidateUpdateCar() error {
	if i.Name == nil && i.LoadCapacity == nil && i.FuelConsumption == nil && i.Height == nil && i.Length == nil && i.Width == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
