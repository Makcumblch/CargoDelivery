package cargodelivery

import "errors"

type Order struct {
	Id      int  `json:"id" db:"id"`
	IdCargo int  `json:"idCargo" binding:"required" db:"cargo_id"`
	Count   uint `json:"count" binding:"required" db:"count"`
}

type UpdateOrder struct {
	IdCargo *int  `json:"idCargo"`
	Count   *uint `json:"count"`
}

func (i *UpdateOrder) ValidateUpdateOrder() error {
	if i.IdCargo == nil && i.Count == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
