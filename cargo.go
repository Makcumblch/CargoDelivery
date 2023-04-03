package cargodelivery

import "errors"

type Cargo struct {
	Id     int      `json:"id" db:"id"`
	Name   *string  `json:"name" db:"name" binding:"required"`
	Width  *float32 `json:"width" binding:"required" db:"width"`   // м
	Height *float32 `json:"height" binding:"required" db:"height"` // м
	Length *float32 `json:"length" binding:"required" db:"length"` // м
	Weight *float32 `json:"weight" binding:"required" db:"weight"` // кг
}

type UpdateCargo struct {
	Name   *string  `json:"name"`
	Width  *float32 `json:"width"`  // м
	Height *float32 `json:"height"` // м
	Length *float32 `json:"length"` // м
	Weight *float32 `json:"weight"` // кг
}

func (i *UpdateCargo) ValidateUpdateCargo() error {
	if i.Name == nil && i.Width == nil && i.Height == nil && i.Length == nil && i.Weight == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
