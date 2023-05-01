package repository

import (
	"fmt"
	"strings"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/jmoiron/sqlx"
)

type ClientPostgres struct {
	db *sqlx.DB
}

func NewClientPostgres(db *sqlx.DB) *ClientPostgres {
	return &ClientPostgres{db: db}
}

func (c *ClientPostgres) CreateClient(projectId int, client cargodelivery.Client) (int, error) {
	var id int
	createClientQuery := fmt.Sprintf("INSERT INTO %s (project_id, name, address, coord_x, coord_y) VALUES ($1, $2, $3, $4, $5) RETURNING id", clientsTable)
	row := c.db.QueryRow(createClientQuery, projectId, client.Name, client.Address, client.CoordX, client.CoordY)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (c *ClientPostgres) GetAllClients(projectId int) ([]cargodelivery.Client, error) {
	var clients []cargodelivery.Client

	query := fmt.Sprintf("SELECT id, name, address, coord_x, coord_y FROM %s WHERE project_id = $1", clientsTable)
	err := c.db.Select(&clients, query, projectId)

	return clients, err
}

func (c *ClientPostgres) GetClientById(projectId, clientId int) (cargodelivery.Client, error) {
	var clients cargodelivery.Client

	query := fmt.Sprintf("SELECT id, name, address, coord_x, coord_y FROM %s WHERE project_id = $1 AND id = $2", clientsTable)
	err := c.db.Get(&clients, query, projectId, clientId)

	return clients, err
}

func (c *ClientPostgres) DeleteClient(projectId, clientId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE project_id = $1 AND id = $2", clientsTable)
	_, err := c.db.Exec(query, projectId, clientId)

	return err
}

func (c *ClientPostgres) UpdateClient(projectId, clientId int, input cargodelivery.UpdateClient) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Address != nil {
		setValues = append(setValues, fmt.Sprintf("address=$%d", argId))
		args = append(args, *input.Address)
		argId++
	}

	if input.CoordX != nil {
		setValues = append(setValues, fmt.Sprintf("coord_x=$%d", argId))
		args = append(args, *input.CoordX)
		argId++
	}

	if input.CoordY != nil {
		setValues = append(setValues, fmt.Sprintf("coord_y=$%d", argId))
		args = append(args, *input.CoordY)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE project_id = $%d AND id = $%d", clientsTable, setQuery, argId, argId+1)
	args = append(args, projectId, clientId)

	_, err := c.db.Exec(query, args...)

	return err
}
