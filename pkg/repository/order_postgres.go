package repository

import (
	"fmt"
	"strings"

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

func (o *OrderPostgres) GetAllOrders(clientId int) ([]cargodelivery.Order, error) {
	var orders []cargodelivery.Order

	query := fmt.Sprintf("SELECT ord.id, ord.cargo_id, ord.count, ca.name FROM %s ord INNER JOIN %s ca ON ord.cargo_id = ca.id WHERE ord.client_id = $1", ordersTable, cargosTable)
	err := o.db.Select(&orders, query, clientId)

	return orders, err
}

func (o *OrderPostgres) GetOrderById(clientId, orderId int) (cargodelivery.Order, error) {
	var order cargodelivery.Order

	query := fmt.Sprintf("SELECT ord.id, ord.cargo_id, ord.count, ca.name FROM %s ord INNER JOIN %s ca ON ord.cargo_id = ca.id WHERE ord.client_id = $1 AND ord.id = $2", ordersTable, cargosTable)
	err := o.db.Get(&order, query, clientId, orderId)

	return order, err
}

func (o *OrderPostgres) DeleteOrder(clientId, orderId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE client_id = $1 AND id = $2", ordersTable)
	_, err := o.db.Exec(query, clientId, orderId)

	return err
}

func (o *OrderPostgres) UpdateOrder(clientId, orderId int, input cargodelivery.UpdateOrder) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.IdCargo != nil {
		setValues = append(setValues, fmt.Sprintf("cargo_id=$%d", argId))
		args = append(args, *input.IdCargo)
		argId++
	}

	if input.Count != nil {
		setValues = append(setValues, fmt.Sprintf("count=$%d", argId))
		args = append(args, *input.Count)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE client_id = $%d AND id = $%d", ordersTable, setQuery, argId, argId+1)
	args = append(args, clientId, orderId)

	_, err := o.db.Exec(query, args...)

	return err
}
