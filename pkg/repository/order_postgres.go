package repository

import (
	"fmt"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (o *OrderPostgres) CreateOrder(clientId int, order cargodelivery.Order) (int, error) {
	var id int
	createCarQuery := fmt.Sprintf("INSERT INTO %s (client_id, cargo_id, count) VALUES ($1, $2, $3) RETURNING id", ordersTable)
	row := o.db.QueryRow(createCarQuery, clientId, order.IdCargo, order.Count)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
