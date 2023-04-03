package repository

import (
	"fmt"
	"strings"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/jmoiron/sqlx"
)

type CargoPostgres struct {
	db *sqlx.DB
}

func NewCargoPostgres(db *sqlx.DB) *CargoPostgres {
	return &CargoPostgres{db: db}
}

func (c *CargoPostgres) CreateCargo(projectId int, cargo cargodelivery.Cargo) (int, error) {
	var id int
	createCarQuery := fmt.Sprintf("INSERT INTO %s (project_id, name, width, height, length, weight) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", cargosTable)
	row := c.db.QueryRow(createCarQuery, projectId, *cargo.Name, *cargo.Width, *cargo.Height, *cargo.Length, *cargo.Weight)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (c *CargoPostgres) GetAllCargos(projectId int) ([]cargodelivery.Cargo, error) {
	var cargos []cargodelivery.Cargo

	query := fmt.Sprintf("SELECT id, name, width, height, length, weight FROM %s WHERE project_id = $1", cargosTable)
	err := c.db.Select(&cargos, query, projectId)

	return cargos, err
}

func (c *CargoPostgres) GetCargoById(projectId, cargoId int) (cargodelivery.Cargo, error) {
	var cargo cargodelivery.Cargo

	query := fmt.Sprintf("SELECT id, name, width, height, length, weight FROM %s WHERE project_id = $1 AND id = $2", cargosTable)
	err := c.db.Get(&cargo, query, projectId, cargoId)

	return cargo, err
}

func (c *CargoPostgres) DeleteCargo(projectId, cargoId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE project_id = $1 AND id = $2", cargosTable)
	_, err := c.db.Exec(query, projectId, cargoId)

	return err
}

func (c *CargoPostgres) UpdateCargo(projectId, cargoId int, input cargodelivery.UpdateCargo) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Width != nil {
		setValues = append(setValues, fmt.Sprintf("width=$%d", argId))
		args = append(args, *input.Width)
		argId++
	}

	if input.Height != nil {
		setValues = append(setValues, fmt.Sprintf("height=$%d", argId))
		args = append(args, *input.Height)
		argId++
	}

	if input.Length != nil {
		setValues = append(setValues, fmt.Sprintf("length=$%d", argId))
		args = append(args, *input.Length)
		argId++
	}

	if input.Weight != nil {
		setValues = append(setValues, fmt.Sprintf("weight=$%d", argId))
		args = append(args, *input.Weight)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE project_id = $%d AND id = $%d", cargosTable, setQuery, argId, argId+1)
	args = append(args, projectId, cargoId)

	_, err := c.db.Exec(query, args...)

	return err
}
