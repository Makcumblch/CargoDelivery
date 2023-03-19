package cargodelivery

type Order struct {
	Id      int  `json:"id" db:"id"`
	IdCargo int  `json:"idCargo" binding:"required" db:"cargo_id"`
	Count   uint `json:"count" binding:"required" db:"count"`
}
