package repository

import (
	"fmt"
	"strings"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/jmoiron/sqlx"
)

type DepoPostgres struct {
	db *sqlx.DB
}

func NewDepoPostgres(db *sqlx.DB) *DepoPostgres {
	return &DepoPostgres{db: db}
}

func (c *DepoPostgres) CreateDepo(projectId int, depo cargodelivery.Depo) (int, error) {
	var id int
	createDepoQuery := fmt.Sprintf("INSERT INTO %s (project_id, address, coord_x, coord_y) VALUES ($1, $2, $3, $4) RETURNING id", depoTable)
	row := c.db.QueryRow(createDepoQuery, projectId, depo.Address, depo.CoordX, depo.CoordY)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (c *DepoPostgres) GetDepo(projectId int) (cargodelivery.Client, error) {
	var depo cargodelivery.Client

	query := fmt.Sprintf("SELECT id, address, coord_x, coord_y FROM %s WHERE project_id = $1", depoTable)
	err := c.db.Get(&depo, query, projectId)

	return depo, err
}

func (c *DepoPostgres) UpdateDepo(projectId int, input cargodelivery.UpdateDepo) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

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
	query := fmt.Sprintf("UPDATE %s SET %s WHERE project_id = $%d", clientsTable, setQuery, argId)
	args = append(args, projectId)

	_, err := c.db.Exec(query, args...)

	return err
}
